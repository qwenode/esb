package esb

import (
    "context"
    "encoding/json"
    "github.com/elastic/go-elasticsearch/v8"
    "github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
    "github.com/elastic/go-elasticsearch/v8/typedapi/types"
    "github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/conflicts"
    "github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
)

type ActiveRecordAlias interface {
    GetIndexAlias() string
}
type ActiveRecord[T ActiveRecordAlias] struct {
    client  *elasticsearch.TypedClient
    refresh bool
    model   T
}

func NewActiveRecord[T ActiveRecordAlias](client *elasticsearch.TypedClient, aliasModel T) *ActiveRecord[T] {
    return &ActiveRecord[T]{
        model:  aliasModel,
        client: client,
    }
}

// 是否立即刷新 20250514
func (r *ActiveRecord[T]) Refresh(refresh bool) *ActiveRecord[T] {
    r.refresh = refresh
    return r
}

// 根据ID查找 20250514
func (r *ActiveRecord[T]) FindOneByID(c context.Context, id string) (T, error) {
    var result T
    response, err := r.client.Get(r.model.GetIndexAlias(), id).Do(c)
    if err != nil {
        return result, err
    }
    
    if !response.Found {
        return result, cc.NewNoData()
    }
    err = json.Unmarshal(response.Source_, &result)
    if err != nil {
        return result, err
    }
    return result, nil
}

// 查找一条信息,要确保数据库中的字段是唯一的,否则数据可能随机返回 20250514
func (r *ActiveRecord[T]) FindOneByField(c context.Context, field string, value types.FieldValue, caseInsensitive bool) (T, error) {
    var result T
    response, err := r.client.Search().Index(r.model.GetIndexAlias()).Query(QFilter(QTerm(field, value, rr.AsPointer(caseInsensitive)))).
        Size(1).Do(c)
    if err != nil {
        return result, err
    }
    if len(response.Hits.Hits) <= 0 {
        return result, cc.NewNoData()
    }
    err = json.Unmarshal(response.Hits.Hits[0].Source_, &result)
    if err != nil {
        return result, err
    }
    return result, nil
}

// 查询id是否存在 20250514
func (r *ActiveRecord[T]) Exist(c context.Context, id string) bool {
    hasID, _ := r.HasID(c, id)
    return hasID
}

// 查询id是否存在 20250514
func (r *ActiveRecord[T]) HasID(c context.Context, id string) (bool, error) {
    return r.client.Exists(r.model.GetIndexAlias(), id).Do(c)
}

// 索引数据,完全覆盖
func (r *ActiveRecord[T]) Index(c context.Context, entity T, id string) (_id string, _err error) {
    h := r.client.Index(r.model.GetIndexAlias()).Id(id).Document(entity)
    if r.refresh {
        h.Refresh(refresh.True)
    }
    response, err := h.Do(c)
    if err != nil {
        return "", err
    }
    return response.Id_, err
}

//  更新已有数据
func (r *ActiveRecord[T]) UpdateByID(c context.Context, entity T, id string) error {
    h := r.client.Update(r.model.GetIndexAlias(), id).Doc(entity)
    if r.refresh {
        h.Refresh(refresh.True)
    }
    _, err := h.Do(c)
    return err
}

// 更新多个字段
func (r *ActiveRecord[T]) UpdateFieldsByID(c context.Context, id string, fields map[string]any) error {
    h := r.client.Update(r.model.GetIndexAlias(), id).Doc(fields)
    if r.refresh {
        h.Refresh(refresh.True)
    }
    _, err := h.Do(c)
    return err
}

// 更新字段
func (r *ActiveRecord[T]) UpdateFieldByID(c context.Context, id, field string, value any) error {
    h := r.client.Update(r.model.GetIndexAlias(), id).Doc(map[string]any{field: value})
    if r.refresh {
        h.Refresh(refresh.True)
    }
    _, err := h.Do(c)
    return err
}

//  更新或创建数据
func (r *ActiveRecord[T]) Upsert(c context.Context, entity T, id string) error {
    h := r.client.Update(r.model.GetIndexAlias(), id).Doc(entity).DocAsUpsert(true)
    if r.refresh {
        h.Refresh(refresh.True)
    }
    _, err := h.Do(c)
    return err
}

//  删除数据
func (r *ActiveRecord[T]) Delete(c context.Context, id string) error {
    h := r.client.Delete(r.model.GetIndexAlias(), id)
    if r.refresh {
        h.Refresh(refresh.True)
    }
    _, err := h.Do(c)
    return err
}

// 表名 20250514
func (r *ActiveRecord[T]) GetIndex() string {
    alias := r.model.GetIndexAlias()
    if alias == "" {
        panic("严重错误,没有索引?")
    }
    return alias
}

// 统计总数量 20250506
func (r *ActiveRecord[T]) Count(c context.Context) (int64, error) {
    h := r.client.Count().Index(r.GetIndex())
    res, err := h.Do(c)
    if err != nil {
        return 0, cc.NewInternalWith(err)
    }
    return res.Count, nil
}

func (r *ActiveRecord[T]) DeleteByField(c context.Context, field string, value []any) error {
    h := r.client.DeleteByQuery(r.GetIndex()).Query(
        &types.Query{
            Bool: &types.BoolQuery{Filter: []types.Query{
                {
                    Terms: &types.TermsQuery{
                        TermsQuery: map[string]types.TermsQueryField{
                            field: value,
                        },
                    },
                },
            }},
        },
    )
    if r.refresh {
        h.Refresh(r.refresh)
    }
    _, err := h.Conflicts(conflicts.Proceed).Do(c)
    
    return err
}

// 批量删除 20250514
func (r *ActiveRecord[T]) BatchDelete(c context.Context, id ...string) error {
    h := r.client.DeleteByQuery(r.GetIndex()).Query(
        &types.Query{
            Bool: &types.BoolQuery{Filter: []types.Query{
                {
                    Terms: &types.TermsQuery{
                        TermsQuery: map[string]types.TermsQueryField{
                            "id": id,
                        },
                    },
                },
            }},
        },
    )
    if r.refresh {
        h.Refresh(r.refresh)
    }
    _, err := h.Conflicts(conflicts.Proceed).Do(c)
    
    return err
}

type ARSearchRequest func(ssi *search.Search)
type ARSearchResponse func(response *search.Response) error

// Search 搜索多条记录
func (r *ActiveRecord[T]) Search(c context.Context, onRequest ARSearchRequest, onResponse ARSearchResponse) error {
    s := r.client.Search().Index(r.model.GetIndexAlias())
    onRequest(s)
    resp, err := s.TrackTotalHits(true).Do(c)
    if err != nil {
        return err
    }
    return onResponse(resp)
}

// GetModel 返回当前模型
func (r *ActiveRecord[T]) GetModel() T {
    return r.model
}

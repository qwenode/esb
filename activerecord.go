package esb

import (
    "context"
    "encoding/json"
    "errors"

    "github.com/elastic/go-elasticsearch/v8"
    "github.com/elastic/go-elasticsearch/v8/typedapi/core/get"
    "github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
    "github.com/elastic/go-elasticsearch/v8/typedapi/types"
    "github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/conflicts"
    "github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
)

var (
    ErrNoData = errors.New("no data")
)

func IsNoData(err error) bool {
    return errors.Is(err, ErrNoData)
}

type ActiveRecordAlias interface {
    GetIndexAlias() string
}
type ActiveRecord[T ActiveRecordAlias] struct {
    client  *elasticsearch.TypedClient
    refresh bool
    entity  T
}

func NewActiveRecord[T ActiveRecordAlias](client *elasticsearch.TypedClient, entity T) *ActiveRecord[T] {
    return &ActiveRecord[T]{
        entity: entity,
        client: client,
    }
}

// 是否立即刷新 20250514
func (r *ActiveRecord[T]) Refresh(refresh bool) *ActiveRecord[T] {
    r.refresh = refresh
    return r
}

// 查询 _id 20250514
func (r *ActiveRecord[T]) FindPK(c context.Context, id string) (T, error) {
    var result T
    response, err := r.client.Get(r.GetAlias(), id).Do(c)
    if err != nil {
        return result, err
    }
    if !response.Found {
        return result, ErrNoData
    }
    err = json.Unmarshal(response.Source_, &result)
    if err != nil {
        return result, err
    }
    return result, nil
}

func (r *ActiveRecord[T]) FindOne(c context.Context, field string, value types.FieldValue) (T, error) {
    return r.FindOneByField(c, field, value, false)
}

// 查找一条信息,要确保数据库中的字段是唯一的,否则数据可能随机返回 20250514
func (r *ActiveRecord[T]) FindOneByField(c context.Context, field string, value types.FieldValue, caseInsensitive bool) (T, error) {
    var result T
    response, err := r.client.Search().Index(r.GetAlias()).Query(&types.Query{
        Bool: &types.BoolQuery{
            Filter: []types.Query{
                {
                    Term: map[string]types.TermQuery{
                        field: {
                            Value:           value,
                            CaseInsensitive: &caseInsensitive,
                        },
                    },
                },
            },
        },
    }).Size(1).Do(c)
    if err != nil {
        return result, err
    }
    if len(response.Hits.Hits) <= 0 {
        return result, ErrNoData
    }
    err = json.Unmarshal(response.Hits.Hits[0].Source_, &result)
    if err != nil {
        return result, err
    }
    return result, nil
}

// 查询id是否存在 20250514
func (r *ActiveRecord[T]) Exist(c context.Context, id string) (bool, error) {
    return r.client.Exists(r.GetAlias(), id).Do(c)
}

// 索引数据,完全覆盖
func (r *ActiveRecord[T]) Index(c context.Context, entity T, id string) (_id string, _err error) {
    h := r.client.Index(r.GetAlias()).Id(id).Document(entity)
    if r.refresh {
        h.Refresh(refresh.True)
    }
    response, err := h.Do(c)
    if err != nil {
        return "", err
    }
    return response.Id_, err
}

//  更新已有数据,更新字段根据结构体定制,但建议优先使用 UpdatePartial
func (r *ActiveRecord[T]) UpdateEntity(c context.Context, entity T, id string) error {
    h := r.client.Update(r.GetAlias(), id).Doc(entity)
    if r.refresh {
        h.Refresh(refresh.True)
    }
    _, err := h.Do(c)
    return err
}

// 局部更新文档
func (r *ActiveRecord[T]) UpdatePartial(c context.Context, id string, fields map[string]any) error {
    marshal, err := json.Marshal(fields)
    if err != nil {
        return err
    }
    h := r.client.Update(r.GetAlias(), id).Doc(marshal)
    if r.refresh {
        h.Refresh(refresh.True)
    }
    _, err = h.Do(c)
    return err
}

// 更新单个字段
func (r *ActiveRecord[T]) UpdateField(c context.Context, id, field string, value any) error {
    marshal, err := json.Marshal(map[string]any{field: value})
    if err != nil {
        return err
    }
    h := r.client.Update(r.GetAlias(), id).Doc(marshal)
    if r.refresh {
        h.Refresh(refresh.True)
    }
    _, err = h.Do(c)
    return err
}

//  更新或创建数据
func (r *ActiveRecord[T]) Upsert(c context.Context, id string, entities T) error {
    h := r.client.Update(r.GetAlias(), id).Doc(entities).DocAsUpsert(true)
    if r.refresh {
        h.Refresh(refresh.True)
    }
    _, err := h.Do(c)
    return err
}

//  删除数据
func (r *ActiveRecord[T]) Delete(c context.Context, id string) error {
    h := r.client.Delete(r.GetAlias(), id)
    if r.refresh {
        h.Refresh(refresh.True)
    }
    _, err := h.Do(c)
    return err
}

// 表名 20250514
func (r *ActiveRecord[T]) GetAlias() string {
    alias := r.GetModel().GetIndexAlias()
    if alias == "" {
        panic("严重错误,没有索引?")
    }
    return alias
}

// 统计总数量 20250506
func (r *ActiveRecord[T]) Count(c context.Context) (int64, error) {
    h := r.client.Count().Index(r.GetAlias())
    res, err := h.Do(c)
    if err != nil {
        return 0, err
    }
    return res.Count, nil
}

// 批量删除 20250722
func (r *ActiveRecord[T]) BatchDeleteByField(c context.Context, field string, value []types.TermsQueryField) error {
    h := r.client.DeleteByQuery(r.GetAlias()).Query(
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

type ActiveRecordSearchRequest func(ssi *search.Search)
type ActiveRecordSearchResponse func(response *search.Response) error

// Search 搜索多条记录
func (r *ActiveRecord[T]) Search(c context.Context, onRequest ActiveRecordSearchRequest, onResponse ActiveRecordSearchResponse) error {
    s := r.client.Search().Index(r.GetAlias())
    onRequest(s)
    resp, err := s.TrackTotalHits(true).Do(c)
    if err != nil {
        return err
    }
    return onResponse(resp)
}

// GetModel 返回当前模型
func (r *ActiveRecord[T]) GetModel() T {
    return r.entity
}

// 解析单条数据 20250722
func FormatOne[T any](response *get.Response, err error) (T, error) {
    var v T
    if err != nil {
        return v, err
    }
    if !response.Found {
        return v, ErrNoData
    }
    err = json.Unmarshal(response.Source_, &v)
    return v, err
}

// 解析后的回调处理,如果_append=false,则数据不会返回 20250722
type FormatHitsPostProcessor[T any] func(src T) (_append bool)

// 解析多条数据返回值 20250722
func FormatSearch[T any](response *search.Response, postprocessor FormatHitsPostProcessor[T]) []T {
    if response == nil || response.Hits.Hits == nil {
        return nil
    }
    hits := response.Hits.Hits
    hLen := len(hits)
    if hLen <= 0 {
        return nil
    }
    if postprocessor == nil {
        postprocessor = func(src T) (_append bool) {
            return true
        }
    }
    ts := make([]T, 0, hLen)
    for _, hit := range hits {
        var v T
        err := json.Unmarshal(hit.Source_, &v)
        if err != nil {
            continue
        }
        if postprocessor(v) {
            ts = append(ts, v)
        }
    }
    return ts
}

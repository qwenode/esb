package multisearch

import (
    "context"
    "errors"
    "strings"

    "github.com/elastic/go-elasticsearch/v8"
    "github.com/elastic/go-elasticsearch/v8/typedapi/core/msearch"
    "github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

//  仅包含指定字段 
func WithIncludeSourceFields(fields ...string) PreProcessor {
    return func(header *types.MultisearchHeader, body *types.MultisearchBody) {
        body.Source_ = types.SourceFilter{
            Includes: fields,
        }
    }
}

//  排除指定字段 
func WithExcludeSourceFields(fields ...string) PreProcessor {
    return func(header *types.MultisearchHeader, body *types.MultisearchBody) {
        body.Source_ = types.SourceFilter{
            Excludes: fields,
        }
    }
}
func WithFunc(optionFunc PreProcessor) PreProcessor {
    return func(header *types.MultisearchHeader, body *types.MultisearchBody) {
        optionFunc(header, body)
    }
}

func WithSort(options ...types.SortCombinations) PreProcessor {
    return func(header *types.MultisearchHeader, body *types.MultisearchBody) {
        body.Sort = options
    }
}
func WithSize10000() PreProcessor {
    return WithSize(10000, 0)
}
func WithSize(size int, from int) PreProcessor {
    return func(header *types.MultisearchHeader, body *types.MultisearchBody) {
        // 如果size=0就执行不到hitLen>0,也就无法获取index,导致aggs没内容,所以size必须大于0 20250523
        if size <= 0 {
            size = 1
        }
        body.Size = &size
        if from > 0 {
            body.From = &from
        }
    }
}

type (
    // 查询前的参数数量 20250722
    PreProcessor func(header *types.MultisearchHeader, body *types.MultisearchBody)
    // 查询成功后的数据处理 20250722
    PostProcessor func(msi *types.MultiSearchItem, resultLength int, index string)
)

// 并发查询 20250722
type MultiSearch struct {
    client         *msearch.Msearch
    postProcessors map[string]PostProcessor
}

func NewBuilder(client *elasticsearch.TypedClient) *MultiSearch {
    return &MultiSearch{client: client.Msearch(), postProcessors: make(map[string]PostProcessor)}
}

func (r *MultiSearch) AddSearch(index string, query *types.Query, postProcessor PostProcessor, size int, preProcessor ...PreProcessor) *MultiSearch {
    header := &types.MultisearchHeader{
        Index: []string{index},
    }
    body := &types.MultisearchBody{
        Query:          query,
        TrackTotalHits: true,
    }
    // 防止aggs没数据 20250722
    if size < 1 {
        size = 1
    }
    body.Size = &size
    for _, option := range preProcessor {
        option(header, body)
    }
    _ = r.client.AddSearch(*header, *body)
    r.postProcessors[index] = postProcessor
    return r
}

type (
    // 自定义表名处理器 20250722
    AliasProcessor func(index string) string
)

// 默认index与alias对应处理器,index必须以_日期结尾,如(prefix_table_20250808)=(prefix_table) 20250722
func DefaultAliasProcessor() AliasProcessor {
    return func(index string) string {
        sub := "_20"
        if !strings.Contains(index, sub) {
            return index
        }
        split := strings.Split(index, sub)
        return strings.Join(split[:len(split)-1], sub)
    }
}

func (r *MultiSearch) Do(c context.Context, aliasProcessors ...AliasProcessor) error {
    response, err := r.client.Do(c)
    if err != nil {
        return err
    }
    aliasProcess := DefaultAliasProcessor()
    if len(aliasProcessors) > 0 {
        aliasProcess = aliasProcessors[0]
    }
    for _, responseItem := range response.Responses {
        item, ok := responseItem.(*types.MultiSearchItem)
        if !ok {
            if responseItem == nil {
                continue
            }
            base := responseItem.(*types.ErrorResponseBase)
            return errors.New(*base.Error.Reason)
        }
        hitLen := len(item.Hits.Hits)
        index := ""
        if hitLen > 0 {
            hit := item.Hits.Hits[0]
            index = aliasProcess(hit.Index_)
        }
        for s, processor := range r.postProcessors {
            if index == s {
                processor(item, hitLen, index)
            }
        }
    }
    return nil
}

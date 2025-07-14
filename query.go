// Package esb 提供了一个基于函数式选项模式的 Elasticsearch 流式查询构建器。
// 它简化了复杂 Elasticsearch 查询的构建过程，同时保持与
// github.com/elastic/go-elasticsearch/v8/typedapi/types 的完全兼容性。
package esb

import (
    "github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// QueryOption 表示一个修改 types.Query 的函数。
// 它遵循函数式选项模式来构建 Elasticsearch 查询。
type QueryOption func(*types.Query)

// NewQuery 通过应用提供的选项创建一个新的 Elasticsearch 查询。
// 它返回一个可以直接用于 go-elasticsearch 客户端的 *types.Query。
//
// 示例：
//   query := esb.NewQuery(
//       esb.Term("status", "published"),
//   )
//   client.Search().Index("articles").Query(query)
func NewQuery(opts ...QueryOption) *types.Query {
    query := &types.Query{}
    for _, opt := range opts {
        opt(query)
    }
    return query
}

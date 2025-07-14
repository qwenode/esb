package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// BoolOption 表示一个修改 types.BoolQuery 的函数。
// 它用于构建包含 Must、Should、Filter 和 MustNot 子句的复杂布尔查询。
type BoolOption func(*types.BoolQuery)

// Bool 使用指定的选项创建一个布尔查询。
// 布尔查询用于使用布尔逻辑组合多个查询。
//
// 示例：
//   query, err := esb.NewQuery(
//       esb.Bool(
//           esb.Must(
//               esb.Term("status", "published"),
//               esb.Range("date").Gte("2023-01-01").Build(),
//           ),
//           esb.Filter(
//               esb.Term("category", "tech"),
//           ),
//           esb.Should(
//               esb.Match("title", "elasticsearch"),
//               esb.Match("content", "search"),
//           ),
//       ),
//   )
func Bool(opts ...BoolOption) QueryOption {
	return func(q *types.Query) {
		boolQuery := &types.BoolQuery{}
		for _, opt := range opts {
			opt(boolQuery)
		}
		q.Bool = boolQuery
	}
}

// Must 指定文档必须匹配的查询条件。
// Must 子句中的所有查询都必须匹配（AND 逻辑）。
//
// 示例：
//   esb.Must(
//       esb.Term("status", "published"),
//       esb.Range("date").Gte("2023-01-01").Build(),
//   )
func Must(opts ...QueryOption) BoolOption {
	return func(bq *types.BoolQuery) {
		for _, opt := range opts {
			subQuery := &types.Query{}
			opt(subQuery)
			bq.Must = append(bq.Must, *subQuery)
		}
	}
}

// Should 指定文档应该匹配的查询条件。
// Should 子句中至少应该匹配一个查询（OR 逻辑）。
// 匹配的 Should 子句越多，文档的得分越高。
//
// 示例：
//   esb.Should(
//       esb.Match("title", "elasticsearch"),
//       esb.Match("content", "search engine"),
//   )
func Should(opts ...QueryOption) BoolOption {
	return func(bq *types.BoolQuery) {
		for _, opt := range opts {
			subQuery := &types.Query{}
			opt(subQuery)
			bq.Should = append(bq.Should, *subQuery)
		}
	}
}

// Filter 指定文档必须匹配的查询条件，
// 但与 Must 不同，Filter 查询不会影响文档得分。
// Filter 查询会被缓存，执行速度比 Must 查询更快。
//
// 示例：
//   esb.Filter(
//       esb.Term("status", "published"),
//       esb.Range("publish_date").Gte("2023-01-01").Build(),
//   )
func Filter(opts ...QueryOption) BoolOption {
	return func(bq *types.BoolQuery) {
		for _, opt := range opts {
			subQuery := &types.Query{}
			opt(subQuery)
			bq.Filter = append(bq.Filter, *subQuery)
		}
	}
}

// MustNot 指定文档不能匹配的查询条件。
// 匹配任何 MustNot 查询的文档都会被排除在结果之外。
//
// 示例：
//   esb.MustNot(
//       esb.Term("status", "deleted"),
//       esb.Term("hidden", true),
//   )
func MustNot(opts ...QueryOption) BoolOption {
	return func(bq *types.BoolQuery) {
		for _, opt := range opts {
			subQuery := &types.Query{}
			opt(subQuery)
			bq.MustNot = append(bq.MustNot, *subQuery)
		}
	}
} 
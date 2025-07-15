package esb

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

// MatchNone 创建一个不匹配任何文档的查询。
// 这个查询在某些复合查询场景中很有用，比如在布尔查询的 must_not 子句中。
//
// 示例：
//   esb.MatchNone()                   // 不匹配任何文档
func MatchNone() QueryOption {
	return func(q *types.Query) {
		q.MatchNone = &types.MatchNoneQuery{}
	}
}

// MatchNoneWithOptions 提供回调函数式的 MatchNone 查询配置。
// 允许设置 boost 等高级选项。
//
// 示例：
//   esb.MatchNoneWithOptions(func(opts *types.MatchNoneQuery) {
//       boost := float32(0.0)
//       opts.Boost = &boost
//   })
func MatchNoneWithOptions(setOpts func(opts *types.MatchNoneQuery)) QueryOption {
	return func(q *types.Query) {
		matchNoneQuery := &types.MatchNoneQuery{}
		if setOpts != nil {
			setOpts(matchNoneQuery)
		}
		q.MatchNone = matchNoneQuery
	}
}
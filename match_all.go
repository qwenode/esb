package esb

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

// MatchAll 创建一个匹配所有文档的查询，给所有文档一个 _score 为 1.0。
// 这是最简单的查询，通常用作默认查询或在复合查询中作为基础。
//
// 示例：
//   esb.MatchAll()                    // 匹配所有文档，score 为 1.0
func MatchAll() QueryOption {
	return func(q *types.Query) {
		q.MatchAll = &types.MatchAllQuery{}
	}
}

// MatchAllWithOptions 提供回调函数式的 MatchAll 查询配置。
// 允许设置 boost 等高级选项。
//
// 示例：
//   esb.MatchAllWithOptions(func(opts *types.MatchAllQuery) {
//       boost := 2.0
//       opts.Boost = &boost
//   })
func MatchAllWithOptions(setOpts func(opts *types.MatchAllQuery)) QueryOption {
	return func(q *types.Query) {
		matchAllQuery := &types.MatchAllQuery{}
		if setOpts != nil {
			setOpts(matchAllQuery)
		}
		q.MatchAll = matchAllQuery
	}
}
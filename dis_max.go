package esb

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

// DisMax 创建一个 Disjunction Max 查询，返回匹配一个或多个包装查询的文档。
// 如果返回的文档匹配多个查询子句，DisMax 查询会为文档分配来自任何匹配子句的最高相关性分数，
// 加上任何其他匹配子查询的平局打破增量。
//
// 这对于搜索多个字段时很有用，你希望最佳匹配字段的分数占主导地位。
//
// 示例：
//   esb.DisMax(
//       esb.Term("title", "quick"),
//       esb.Term("body", "quick"),
//   )
func DisMax(queries ...QueryOption) QueryOption {
	return func(q *types.Query) {
		var queryList []types.Query
		
		for _, queryOpt := range queries {
			if queryOpt != nil {
				subQuery := &types.Query{}
				queryOpt(subQuery)
				queryList = append(queryList, *subQuery)
			}
		}
		
		q.DisMax = &types.DisMaxQuery{
			Queries: queryList,
		}
	}
}

// DisMaxWithOptions 提供回调函数式的 DisMax 查询配置。
// 允许设置 tie_breaker、boost 等高级选项。
//
// 示例：
//   esb.DisMaxWithOptions(
//       []QueryOption{
//           esb.Match("title", "quick brown fox"),
//           esb.Match("body", "quick brown fox"),
//       },
//       func(opts *types.DisMaxQuery) {
//           tieBreaker := 0.3
//           opts.TieBreaker = &tieBreaker
//           boost := float32(1.2)
//           opts.Boost = &boost
//       },
//   )
func DisMaxWithOptions(queries []QueryOption, setOpts func(opts *types.DisMaxQuery)) QueryOption {
	return func(q *types.Query) {
		var queryList []types.Query
		
		for _, queryOpt := range queries {
			if queryOpt != nil {
				subQuery := &types.Query{}
				queryOpt(subQuery)
				queryList = append(queryList, *subQuery)
			}
		}
		
		disMaxQuery := &types.DisMaxQuery{
			Queries: queryList,
		}
		
		if setOpts != nil {
			setOpts(disMaxQuery)
		}
		
		q.DisMax = disMaxQuery
	}
}
package esb

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

// ConstantScore 创建一个常量分数查询，包装一个过滤器查询并为每个匹配的文档
// 返回一个等于 boost 参数值的相关性分数。
//
// 这个查询在你想要过滤文档但不关心相关性分数计算时很有用，
// 可以提高查询性能，因为它跳过了分数计算。
//
// 示例：
//   esb.ConstantScore(
//       esb.Term("status", "published"),
//   )
func ConstantScore(filter QueryOption) QueryOption {
	return func(q *types.Query) {
		filterQuery := &types.Query{}
		if filter != nil {
			filter(filterQuery)
		}
		
		q.ConstantScore = &types.ConstantScoreQuery{
			Filter: filterQuery,
		}
	}
}

// ConstantScoreWithOptions 提供回调函数式的 ConstantScore 查询配置。
// 允许设置 boost 等高级选项。
//
// 示例：
//   esb.ConstantScoreWithOptions(
//       esb.Term("status", "published"),
//       func(opts *types.ConstantScoreQuery) {
//           boost := float32(2.0)
//           opts.Boost = &boost
//       },
//   )
func ConstantScoreWithOptions(filter QueryOption, setOpts func(opts *types.ConstantScoreQuery)) QueryOption {
	return func(q *types.Query) {
		filterQuery := &types.Query{}
		if filter != nil {
			filter(filterQuery)
		}
		
		constantScoreQuery := &types.ConstantScoreQuery{
			Filter: filterQuery,
		}
		
		if setOpts != nil {
			setOpts(constantScoreQuery)
		}
		
		q.ConstantScore = constantScoreQuery
	}
}
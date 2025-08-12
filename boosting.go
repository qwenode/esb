package esb

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

// Boosting 创建一个 Boosting 查询，返回匹配 positive 查询的文档，
// 同时降低也匹配 negative 查询的文档的相关性分数。
//
// 这个查询对于需要降低某些文档相关性而不是完全排除它们的场景很有用。
// 例如，搜索"apple"时，你想要提升关于水果的文档，但降低关于公司的文档。
//
// 示例：
//   esb.Boosting(
//       esb.Term("title", "apple"),           // positive 查询
//       esb.Term("category", "technology"),   // negative 查询
//       0.2,                                  // negative_boost
//   )
func Boosting(positive, negative QueryOption, negativeBoost float64) QueryOption {
    return func(q *types.Query) {
        positiveQuery := &types.Query{}
        negativeQuery := &types.Query{}

        if positive != nil {
            positive(positiveQuery)
        }

        if negative != nil {
            negative(negativeQuery)
        }

        negativeBoostFloat := types.Float64(negativeBoost)

        q.Boosting = &types.BoostingQuery{
            Positive:      *positiveQuery,
            Negative:      *negativeQuery,
            NegativeBoost: negativeBoostFloat,
        }
    }
}

// BoostingWithOptions 提供回调函数式的 Boosting 查询配置。
// 允许设置 boost 等高级选项。
//
// 示例：
//   esb.BoostingWithOptions(
//       esb.Match("title", "apple fruit"),
//       esb.Match("content", "iPhone iPad"),
//       0.3,
//       func(opts *types.BoostingQuery) {
//           boost := float32(1.5)
//           opts.Boost = &boost
//           queryName := "apple_boosting"
//           opts.QueryName_ = &queryName
//       },
//   )
func BoostingWithOptions(positive, negative QueryOption, negativeBoost float64, setOpts func(opts *types.BoostingQuery)) QueryOption {
    return func(q *types.Query) {
        positiveQuery := &types.Query{}
        negativeQuery := &types.Query{}

        if positive != nil {
            positive(positiveQuery)
        }

        if negative != nil {
            negative(negativeQuery)
        }

        negativeBoostFloat := types.Float64(negativeBoost)

        boostingQuery := &types.BoostingQuery{
            Positive:      *positiveQuery,
            Negative:      *negativeQuery,
            NegativeBoost: negativeBoostFloat,
        }

        if setOpts != nil {
            setOpts(boostingQuery)
        }

        q.Boosting = boostingQuery
    }
}

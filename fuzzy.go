package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Fuzzy 创建模糊查询，返回包含与搜索词相似的词条的文档。
// 模糊查询基于 Levenshtein 编辑距离来衡量词条的相似性。
//
// 示例：
//   esb.Fuzzy("username", "john")    // 匹配 "john", "jhon", "joh" 等相似词条
//   esb.Fuzzy("title", "elasticsearch") // 匹配 "elasticsearch", "elasticsearh" 等
//   esb.Fuzzy("product", "iphone")   // 匹配 "iphone", "iphon", "ipone" 等
func Fuzzy(field, value string) QueryOption {
	return func(q *types.Query) {
		q.Fuzzy = map[string]types.FuzzyQuery{
			field: {
				Value: value,
			},
		}
	}
}

// FuzzyOptions 表示模糊查询的高级选项配置。
type FuzzyOptions struct {
	// Boost 用于提高或降低查询的相关性分数的浮点数。
	// 相对于默认值 1.0，0-1.0 之间的值会降低分数，大于 1.0 的值会提高分数。
	Boost *float32
	
	// Fuzziness 允许的最大编辑距离。
	// 可以是数字（0, 1, 2）或字符串（"AUTO"）。
	Fuzziness types.Fuzziness
	
	// MaxExpansions 查询将扩展到的最大词条数。
	// 默认值为 50，较小的值可以提高性能但可能降低召回率。
	MaxExpansions *int
	
	// PrefixLength 不会被"模糊化"的起始字符数。
	// 这有助于减少扩展的词条数量，提高性能。
	PrefixLength *int
	
	// Transpositions 是否包括两个相邻字符的换位操作。
	// 例如，"ab" 到 "ba" 的换位。默认为 true。
	Transpositions *bool
	
	// Rewrite 用于重写查询的方法。
	// 影响查询的执行性能和结果。
	Rewrite *string
	
	// QueryName 为查询设置名称，用于在结果中标识该查询。
	QueryName *string
}

// FuzzyWithOptions 创建带有高级选项的模糊查询。
// 允许对模糊查询进行更精细的控制。
//
// 示例：
//   maxExpansions := 100
//   prefixLength := 2
//   boost := float32(1.5)
//   esb.FuzzyWithOptions("title", "elasticsearch", esb.FuzzyOptions{
//       Fuzziness: types.Fuzziness("AUTO"),
//       MaxExpansions: &maxExpansions,
//       PrefixLength: &prefixLength,
//       Boost: &boost,
//   })
func FuzzyWithOptions(field, value string, options FuzzyOptions) QueryOption {
	return func(q *types.Query) {
		fuzzyQuery := types.FuzzyQuery{
			Value: value,
		}
		
		// 应用选项配置
		if options.Boost != nil {
			fuzzyQuery.Boost = options.Boost
		}
		if options.Fuzziness != nil {
			fuzzyQuery.Fuzziness = options.Fuzziness
		}
		if options.MaxExpansions != nil {
			fuzzyQuery.MaxExpansions = options.MaxExpansions
		}
		if options.PrefixLength != nil {
			fuzzyQuery.PrefixLength = options.PrefixLength
		}
		if options.Transpositions != nil {
			fuzzyQuery.Transpositions = options.Transpositions
		}
		if options.Rewrite != nil {
			fuzzyQuery.Rewrite = options.Rewrite
		}
		if options.QueryName != nil {
			fuzzyQuery.QueryName_ = options.QueryName
		}
		
		q.Fuzzy = map[string]types.FuzzyQuery{
			field: fuzzyQuery,
		}
	}
} 
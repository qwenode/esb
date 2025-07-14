package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/textquerytype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"
)

// MultiMatch 创建多字段匹配查询，允许在多个字段中搜索文本。
// 该查询会分析提供的文本，然后在指定的字段中进行匹配。
//
// 示例：
//   esb.MultiMatch("elasticsearch", "title", "content")          // 在标题和内容字段中搜索
//   esb.MultiMatch("john doe", "first_name", "last_name")        // 在姓名字段中搜索
//   esb.MultiMatch("java programming", "title^2", "content")     // 标题字段权重为2
func MultiMatch(query string, fields ...string) QueryOption {
	return func(q *types.Query) {
		q.MultiMatch = &types.MultiMatchQuery{
			Query:  query,
			Fields: fields,
		}
	}
}

// MultiMatchOptions 表示多字段匹配查询的高级选项配置。
type MultiMatchOptions struct {
	// Analyzer 用于将查询字符串中的文本转换为词条的分析器。
	Analyzer *string
	
	// AutoGenerateSynonymsPhraseQuery 如果为 true，则为多词同义词自动创建短语查询。
	AutoGenerateSynonymsPhraseQuery *bool
	
	// Boost 用于提高或降低查询的相关性分数的浮点数。
	Boost *float32
	
	// CutoffFrequency 允许指定一个绝对或相对的文档频率，高于该频率的词条将被视为通用词条。
	CutoffFrequency *types.Float64
	
	// Fuzziness 允许的最大编辑距离。
	Fuzziness types.Fuzziness
	
	// FuzzyRewrite 用于重写模糊查询的方法。
	FuzzyRewrite *string
	
	// FuzzyTranspositions 是否包括两个相邻字符的换位操作。
	FuzzyTranspositions *bool
	
	// Lenient 如果为 true，则忽略基于格式的错误。
	Lenient *bool
	
	// MaxExpansions 查询将扩展到的最大词条数。
	MaxExpansions *int
	
	// MinimumShouldMatch 必须匹配的最小子句数或百分比。
	MinimumShouldMatch *types.MinimumShouldMatch
	
	// Operator 用于解释查询值中文本的布尔逻辑。
	Operator *operator.Operator
	
	// PrefixLength 不会被"模糊化"的起始字符数。
	PrefixLength *int
	
	// QueryName 为查询设置名称，用于在结果中标识该查询。
	QueryName *string
	
	// Slop 短语查询中词条之间允许的最大间隔数。
	Slop *int
	
	// TieBreaker 用于增加其他匹配子查询分数的系数。
	TieBreaker *types.Float64
	
	// Type 多字段匹配查询的类型。
	Type *textquerytype.TextQueryType
	
	// ZeroTermsQuery 如果分析器移除了所有词条，该查询应该返回什么。
	ZeroTermsQuery *zerotermsquery.ZeroTermsQuery
}

// MultiMatchWithOptions 创建带有高级选项的多字段匹配查询。
// 允许对多字段匹配查询进行更精细的控制。
//
// 示例：
//   boost := float32(1.5)
//   minimumShouldMatch := types.MinimumShouldMatch("75%")
//   esb.MultiMatchWithOptions("elasticsearch java", []string{"title^2", "content"}, esb.MultiMatchOptions{
//       Type: &textquerytype.Bestfields,
//       Operator: &operator.And,
//       MinimumShouldMatch: &minimumShouldMatch,
//       Boost: &boost,
//   })
func MultiMatchWithOptions(query string, fields []string, options MultiMatchOptions) QueryOption {
	return func(q *types.Query) {
		multiMatchQuery := &types.MultiMatchQuery{
			Query:  query,
			Fields: fields,
		}
		
		// 应用选项配置
		if options.Analyzer != nil {
			multiMatchQuery.Analyzer = options.Analyzer
		}
		if options.AutoGenerateSynonymsPhraseQuery != nil {
			multiMatchQuery.AutoGenerateSynonymsPhraseQuery = options.AutoGenerateSynonymsPhraseQuery
		}
		if options.Boost != nil {
			multiMatchQuery.Boost = options.Boost
		}
		if options.CutoffFrequency != nil {
			multiMatchQuery.CutoffFrequency = options.CutoffFrequency
		}
		if options.Fuzziness != nil {
			multiMatchQuery.Fuzziness = options.Fuzziness
		}
		if options.FuzzyRewrite != nil {
			multiMatchQuery.FuzzyRewrite = options.FuzzyRewrite
		}
		if options.FuzzyTranspositions != nil {
			multiMatchQuery.FuzzyTranspositions = options.FuzzyTranspositions
		}
		if options.Lenient != nil {
			multiMatchQuery.Lenient = options.Lenient
		}
		if options.MaxExpansions != nil {
			multiMatchQuery.MaxExpansions = options.MaxExpansions
		}
		if options.MinimumShouldMatch != nil {
			multiMatchQuery.MinimumShouldMatch = *options.MinimumShouldMatch
		}
		if options.Operator != nil {
			multiMatchQuery.Operator = options.Operator
		}
		if options.PrefixLength != nil {
			multiMatchQuery.PrefixLength = options.PrefixLength
		}
		if options.QueryName != nil {
			multiMatchQuery.QueryName_ = options.QueryName
		}
		if options.Slop != nil {
			multiMatchQuery.Slop = options.Slop
		}
		if options.TieBreaker != nil {
			multiMatchQuery.TieBreaker = options.TieBreaker
		}
		if options.Type != nil {
			multiMatchQuery.Type = options.Type
		}
		if options.ZeroTermsQuery != nil {
			multiMatchQuery.ZeroTermsQuery = options.ZeroTermsQuery
		}
		
		q.MultiMatch = multiMatchQuery
	}
}

// MultiMatchBestFields 创建 best_fields 类型的多字段匹配查询。
// 这是默认类型，查找匹配任何字段的文档，但使用最佳字段的分数。
//
// 示例：
//   esb.MultiMatchBestFields("java programming", "title", "content")
func MultiMatchBestFields(query string, fields ...string) QueryOption {
	return MultiMatchWithOptions(query, fields, MultiMatchOptions{
		Type: &textquerytype.Bestfields,
	})
}

// MultiMatchMostFields 创建 most_fields 类型的多字段匹配查询。
// 查找匹配任何字段的文档，并结合每个字段的分数。
//
// 示例：
//   esb.MultiMatchMostFields("java programming", "title", "content", "tags")
func MultiMatchMostFields(query string, fields ...string) QueryOption {
	return MultiMatchWithOptions(query, fields, MultiMatchOptions{
		Type: &textquerytype.Mostfields,
	})
}

// MultiMatchCrossFields 创建 cross_fields 类型的多字段匹配查询。
// 将字段视为一个大字段，查找每个词条在任何字段中的匹配。
//
// 示例：
//   esb.MultiMatchCrossFields("john doe", "first_name", "last_name")
func MultiMatchCrossFields(query string, fields ...string) QueryOption {
	return MultiMatchWithOptions(query, fields, MultiMatchOptions{
		Type: &textquerytype.Crossfields,
	})
}

// MultiMatchPhrase 创建 phrase 类型的多字段匹配查询。
// 对每个字段运行 match_phrase 查询，使用最佳字段的分数。
//
// 示例：
//   esb.MultiMatchPhrase("elasticsearch guide", "title", "content")
func MultiMatchPhrase(query string, fields ...string) QueryOption {
	return MultiMatchWithOptions(query, fields, MultiMatchOptions{
		Type: &textquerytype.Phrase,
	})
}

// MultiMatchPhrasePrefix 创建 phrase_prefix 类型的多字段匹配查询。
// 对每个字段运行 match_phrase_prefix 查询，使用最佳字段的分数。
//
// 示例：
//   esb.MultiMatchPhrasePrefix("elasticsearch sea", "title", "content")
func MultiMatchPhrasePrefix(query string, fields ...string) QueryOption {
	return MultiMatchWithOptions(query, fields, MultiMatchOptions{
		Type: &textquerytype.Phraseprefix,
	})
} 
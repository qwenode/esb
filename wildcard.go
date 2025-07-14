package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Wildcard 创建通配符查询，支持使用通配符模式匹配文档。
// 通配符查询支持 * 和 ? 通配符：
// - * 匹配任意数量的字符（包括零个）
// - ? 匹配任意单个字符
//
// 示例：
//   esb.Wildcard("username", "john*")     // 匹配以 "john" 开头的用户名
//   esb.Wildcard("email", "*@gmail.com") // 匹配所有 Gmail 邮箱
//   esb.Wildcard("product", "iphone-?")  // 匹配 iphone-5, iphone-6 等
func Wildcard(field, value string) QueryOption {
	return func(q *types.Query) {
		q.Wildcard = map[string]types.WildcardQuery{
			field: {
				Value: value,
			},
		}
	}
}

// WildcardOptions 表示通配符查询的高级选项配置。
type WildcardOptions struct {
	// Boost 用于提高或降低查询的相关性分数的浮点数。
	// 相对于默认值 1.0，0-1.0 之间的值会降低分数，大于 1.0 的值会提高分数。
	Boost *float32
	
	// Rewrite 用于重写查询的方法。
	// 影响查询的执行性能和结果。
	Rewrite *string
	
	// CaseInsensitive 设置是否忽略大小写。
	// 如果为 true，则匹配时忽略大小写。
	CaseInsensitive *bool
	
	// QueryName 为查询设置名称，用于在结果中标识该查询。
	QueryName *string
}

// WildcardWithOptions 创建带有高级选项的通配符查询。
// 允许对通配符查询进行更精细的控制。
//
// 示例：
//   boost := float32(2.0)
//   caseInsensitive := true
//   esb.WildcardWithOptions("title", "java*", esb.WildcardOptions{
//       Boost: &boost,
//       CaseInsensitive: &caseInsensitive,
//   })
func WildcardWithOptions(field, value string, options WildcardOptions) QueryOption {
	return func(q *types.Query) {
		wildcardQuery := types.WildcardQuery{
			Value: value,
		}
		
		// 应用选项配置
		if options.Boost != nil {
			wildcardQuery.Boost = options.Boost
		}
		if options.Rewrite != nil {
			wildcardQuery.Rewrite = options.Rewrite
		}
		if options.CaseInsensitive != nil {
			wildcardQuery.CaseInsensitive = options.CaseInsensitive
		}
		if options.QueryName != nil {
			wildcardQuery.QueryName_ = options.QueryName
		}
		
		q.Wildcard = map[string]types.WildcardQuery{
			field: wildcardQuery,
		}
	}
} 
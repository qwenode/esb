package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Prefix 创建前缀查询，返回包含指定前缀的文档。
// 前缀查询用于匹配字段值以特定前缀开头的文档。
//
// 示例：
//   esb.Prefix("username", "john")    // 匹配以 "john" 开头的用户名
//   esb.Prefix("email", "admin")      // 匹配以 "admin" 开头的邮箱
//   esb.Prefix("product", "iphone")   // 匹配以 "iphone" 开头的产品名
func Prefix(field, value string) QueryOption {
	return func(q *types.Query) {
		q.Prefix = map[string]types.PrefixQuery{
			field: {
				Value: value,
			},
		}
	}
}

// PrefixOptions 表示前缀查询的高级选项配置。
type PrefixOptions struct {
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

// PrefixWithOptions 创建带有高级选项的前缀查询。
// 允许对前缀查询进行更精细的控制。
//
// 示例：
//   boost := float32(1.5)
//   caseInsensitive := true
//   esb.PrefixWithOptions("title", "java", esb.PrefixOptions{
//       Boost: &boost,
//       CaseInsensitive: &caseInsensitive,
//   })
func PrefixWithOptions(field, value string, options PrefixOptions) QueryOption {
	return func(q *types.Query) {
		prefixQuery := types.PrefixQuery{
			Value: value,
		}
		
		// 应用选项配置
		if options.Boost != nil {
			prefixQuery.Boost = options.Boost
		}
		if options.Rewrite != nil {
			prefixQuery.Rewrite = options.Rewrite
		}
		if options.CaseInsensitive != nil {
			prefixQuery.CaseInsensitive = options.CaseInsensitive
		}
		if options.QueryName != nil {
			prefixQuery.QueryName_ = options.QueryName
		}
		
		q.Prefix = map[string]types.PrefixQuery{
			field: prefixQuery,
		}
	}
} 
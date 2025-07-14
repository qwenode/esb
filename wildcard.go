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
				Value: &value,
			},
		}
	}
}



 

// WildcardWithOptions 提供回调函数式的 Wildcard 查询配置。
func WildcardWithOptions(field, value string, setOpts func(opts *types.WildcardQuery)) QueryOption {
    return func(q *types.Query) {
        wildcardQuery := types.WildcardQuery{
            Value: &value,
        }
        if setOpts != nil {
            setOpts(&wildcardQuery)
        }
        q.Wildcard = map[string]types.WildcardQuery{
            field: wildcardQuery,
        }
    }
} 
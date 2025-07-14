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



 

// PrefixWithOptions 提供回调函数式的 Prefix 查询配置。
func PrefixWithOptions(field, value string, setOpts func(opts *types.PrefixQuery)) QueryOption {
    return func(q *types.Query) {
        prefixQuery := types.PrefixQuery{
            Value: value,
        }
        if setOpts != nil {
            setOpts(&prefixQuery)
        }
        q.Prefix = map[string]types.PrefixQuery{
            field: prefixQuery,
        }
    }
} 
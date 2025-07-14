package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Exists 创建一个存在性查询，用于匹配包含指定字段的文档。
// 存在性查询返回包含指定字段的已索引值的文档。
// 不管字段的值是什么，只要该字段存在即可。
//
// 示例：
//   esb.Exists("user.name")
//   esb.Exists("metadata.timestamp")
func Exists(field string) QueryOption {
	return func(q *types.Query) {
		q.Exists = &types.ExistsQuery{
			Field: field,
		}
	}
} 
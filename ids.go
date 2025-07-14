package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// IDs 创建 IDs 查询，返回基于文档 ID 的文档。
// 该查询使用存储在 _id 字段中的文档 ID 来查找文档。
//
// 示例：
//   esb.IDs("1")                      // 查找ID为 "1" 的文档
//   esb.IDs("1", "2", "3")           // 查找ID为 "1", "2", "3" 的文档
//   esb.IDs("user-123", "order-456") // 查找多个特定ID的文档
func IDs(ids ...string) QueryOption {
	return func(q *types.Query) {
		q.Ids = &types.IdsQuery{
			Values: ids,
		}
	}
}

// IDsOptions 表示 IDs 查询的高级选项配置。
type IDsOptions struct {
	// Boost 用于提高或降低查询的相关性分数的浮点数。
	// 相对于默认值 1.0，0-1.0 之间的值会降低分数，大于 1.0 的值会提高分数。
	Boost *float32
	
	// QueryName 为查询设置名称，用于在结果中标识该查询。
	QueryName *string
}

// IDsWithOptions 创建带有高级选项的 IDs 查询。
// 允许对 IDs 查询进行更精细的控制。
//
// 示例：
//   boost := float32(2.0)
//   queryName := "user-ids-query"
//   esb.IDsWithOptions([]string{"1", "2", "3"}, esb.IDsOptions{
//       Boost: &boost,
//       QueryName: &queryName,
//   })
func IDsWithOptions(ids []string, options IDsOptions) QueryOption {
	return func(q *types.Query) {
		idsQuery := &types.IdsQuery{
			Values: ids,
		}
		
		// 应用选项配置
		if options.Boost != nil {
			idsQuery.Boost = options.Boost
		}
		if options.QueryName != nil {
			idsQuery.QueryName_ = options.QueryName
		}
		
		q.Ids = idsQuery
	}
}

// IDsFromSlice 从字符串切片创建 IDs 查询。
// 这是一个便利函数，用于从已有的字符串切片创建查询。
//
// 示例：
//   userIds := []string{"user-1", "user-2", "user-3"}
//   esb.IDsFromSlice(userIds)
func IDsFromSlice(ids []string) QueryOption {
	return func(q *types.Query) {
		q.Ids = &types.IdsQuery{
			Values: ids,
		}
	}
} 
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

// IDsSlice 创建一个IDs查询，功能与IDs相同，但接收切片参数而不是可变参数。
// IDsSlice查询用于匹配具有指定ID的文档。
//
// 示例：
//   esb.IDsSlice([]string{"1", "2", "3"})
func IDsSlice(ids []string) QueryOption {
	return func(q *types.Query) {
		q.Ids = &types.IdsQuery{
			Values: ids,
		}
	}
}


// IDsWithOptions 创建带有高级选项的 IDs 查询。
// 允许对 IDs 查询进行更精细的控制。
//
// 示例：
//   boost := float32(2.0)
//   queryName := "user-ids-query"
//   esb.IDsWithOptions([]string{"1", "2", "3"}, func(opts *types.IdsQuery) {
//       opts.Boost = &boost
//       opts.QueryName_ = &queryName
//   })
func IDsWithOptions(ids []string, setOpts func(opts *types.IdsQuery)) QueryOption {
	return func(q *types.Query) {
		idsQuery := &types.IdsQuery{
			Values: ids,
		}
		
		// 应用选项配置
		if setOpts != nil {
			setOpts(idsQuery)
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
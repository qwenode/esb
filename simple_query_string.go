package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// SimpleQueryString 创建一个简单查询字符串查询。
// 此查询使用比标准查询字符串查询更简单的语法。
// 它更适合直接暴露给用户，因为它永远不会抛出语法错误。
//
// 示例：
//
//	query := esb.NewQuery(
//		esb.SimpleQueryString("elasticsearch + search"),
//	)
func SimpleQueryString(query string) QueryOption {
	return func(q *types.Query) {
		q.SimpleQueryString = &types.SimpleQueryStringQuery{
			Query: query,
		}
	}
}

// SimpleQueryStringWithOptions 创建一个带有附加选项的简单查询字符串查询。
//
// 示例：
//
//	query := esb.NewQuery(
//		esb.SimpleQueryStringWithOptions("elasticsearch + search", func(opts *types.SimpleQueryStringQuery) {
//			fields := []string{"title^2", "description"}
//			opts.Fields = fields
//			boost := float32(2.0)
//			opts.Boost = &boost
//		}),
//	)
func SimpleQueryStringWithOptions(query string, setOpts func(opts *types.SimpleQueryStringQuery)) QueryOption {
	return func(q *types.Query) {
		simpleQueryString := &types.SimpleQueryStringQuery{
			Query: query,
		}
		if setOpts != nil {
			setOpts(simpleQueryString)
		}
		q.SimpleQueryString = simpleQueryString
	}
} 
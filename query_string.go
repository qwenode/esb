package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// QueryString 创建一个查询字符串查询。
// 查询字符串语法支持完整的 Lucene 查询字符串语法。
//
// 示例：
//
//	query := esb.NewQuery(
//		esb.QueryString("title:elasticsearch AND (tags:search OR tags:database)"),
//	)
func QueryString(query string) QueryOption {
	return func(q *types.Query) {
		q.QueryString = &types.QueryStringQuery{
			Query: query,
		}
	}
}

// QueryStringWithOptions 创建一个带有附加选项的查询字符串查询。
//
// 示例：
//
//	query := esb.NewQuery(
//		esb.QueryStringWithOptions("title:elasticsearch", func(opts *types.QueryStringQuery) {
//			defaultField := "title"
//			defaultOp := operator.And
//			opts.DefaultField = &defaultField
//			opts.DefaultOperator = &defaultOp
//		}),
//	)
func QueryStringWithOptions(query string, setOpts func(opts *types.QueryStringQuery)) QueryOption {
	return func(q *types.Query) {
		queryString := &types.QueryStringQuery{
			Query: query,
		}
		if setOpts != nil {
			setOpts(queryString)
		}
		q.QueryString = queryString
	}
} 
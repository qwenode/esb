package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Term 创建一个精确词项查询，用于匹配包含指定词项的文档。
// Term 查询用于精确匹配且不会对查询词项进行分析。
//
// 示例：
//   esb.Term("status", "published")
func Term(field, value string) QueryOption {
	return func(q *types.Query) {
		q.Term = map[string]types.TermQuery{
			field: {
				Value: value,
			},
		}
	}
}

// Terms 创建一个多词项查询，用于匹配包含一个或多个精确词项的文档。
// Terms 查询用于匹配任意提供的值。
//
// 示例：
//   esb.Terms("category", "tech", "science", "programming")
func Terms(field string, values ...string) QueryOption {
	return func(q *types.Query) {
		q.Terms = &types.TermsQuery{
			TermsQuery: map[string]types.TermsQueryField{
				field: values,
			},
		}
	}
} 

// TermsSlice 创建一个多词项查询，功能与Terms相同，但接收切片参数而不是可变参数。
// TermsSlice查询用于匹配任意提供的值。
//
// 示例：
//   esb.TermsSlice("category", []string{"tech", "science", "programming"})
func TermsSlice(field string, values []string) QueryOption {
	return func(q *types.Query) {
		q.Terms = &types.TermsQuery{
			TermsQuery: map[string]types.TermsQueryField{
				field: values,
			},
		}
	}
} 
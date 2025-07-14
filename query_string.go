package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// QueryString creates a query string query.
// Query string syntax supports the full Lucene query string syntax.
//
// Example:
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

// QueryStringWithOptions creates a query string query with additional options.
//
// Example:
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
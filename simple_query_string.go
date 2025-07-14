package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// SimpleQueryString creates a simple query string query.
// This query uses a simpler syntax than the standard query string query.
// It's more suitable for exposing directly to users as it will never throw syntax errors.
//
// Example:
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

// SimpleQueryStringWithOptions creates a simple query string query with additional options.
//
// Example:
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
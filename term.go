package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Term creates a term query that matches documents containing an exact term.
// Term queries are used for exact matches and are not analyzed.
//
// Example:
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

// Terms creates a terms query that matches documents containing one or more exact terms.
// Terms queries are used for matching any of the provided values.
//
// Example:
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
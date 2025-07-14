package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Nested creates a nested query.
// This query is used to search nested fields that are arrays of objects.
//
// Example:
//
//	query := esb.NewQuery(
//		esb.Nested("comments", // path to nested field
//			esb.Match("comments.author", "john"), // query on nested field
//		),
//	)
func Nested(path string, query QueryOption) QueryOption {
	return func(q *types.Query) {
		nestedQuery := &types.Query{}
		query(nestedQuery)
		
		q.Nested = &types.NestedQuery{
			Path:  path,
			Query: nestedQuery,
		}
	}
}

// NestedWithOptions creates a nested query with additional options.
//
// Example:
//
//	query := esb.NewQuery(
//		esb.NestedWithOptions("comments",
//			esb.Match("comments.author", "john"),
//			func(opts *types.NestedQuery) {
//				scoreMode := types.NestedScoremodeAvg
//				ignoreUnmapped := true
//				opts.ScoreMode = &scoreMode
//				opts.IgnoreUnmapped = &ignoreUnmapped
//			},
//		),
//	)
func NestedWithOptions(path string, query QueryOption, setOpts func(opts *types.NestedQuery)) QueryOption {
	return func(q *types.Query) {
		nestedQuery := &types.Query{}
		query(nestedQuery)
		
		nested := &types.NestedQuery{
			Path:  path,
			Query: nestedQuery,
		}
		
		if setOpts != nil {
			setOpts(nested)
		}
		
		q.Nested = nested
	}
} 
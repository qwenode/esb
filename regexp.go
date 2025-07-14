package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Regexp creates a regexp query.
// This query allows you to use regular expressions to match field values.
//
// Example:
//
//	query := esb.NewQuery(
//		esb.Regexp("username", "j.*n"),
//	)
func Regexp(field string, value string) QueryOption {
	return func(q *types.Query) {
		if q.Regexp == nil {
			q.Regexp = make(map[string]types.RegexpQuery)
		}
		q.Regexp[field] = types.RegexpQuery{
			Value: value,
		}
	}
}

// RegexpWithOptions creates a regexp query with additional options.
//
// Example:
//
//	query := esb.NewQuery(
//		esb.RegexpWithOptions("username", "j.*n", func(opts *types.RegexpQuery) {
//			flags := "ALL"
//			maxDeterminizedStates := 10000
//			opts.Flags = &flags
//			opts.MaxDeterminizedStates = &maxDeterminizedStates
//		}),
//	)
func RegexpWithOptions(field string, value string, setOpts func(opts *types.RegexpQuery)) QueryOption {
	return func(q *types.Query) {
		if q.Regexp == nil {
			q.Regexp = make(map[string]types.RegexpQuery)
		}
		
		regexpQuery := types.RegexpQuery{
			Value: value,
		}
		if setOpts != nil {
			setOpts(&regexpQuery)
		}
		q.Regexp[field] = regexpQuery
	}
} 
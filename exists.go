package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Exists creates an exists query that matches documents containing the specified field.
// The exists query returns documents that contain an indexed value for a field.
// It doesn't matter what the value is, as long as the field exists.
//
// Example:
//   esb.Exists("user.name")
//   esb.Exists("metadata.timestamp")
func Exists(field string) QueryOption {
	return func(q *types.Query) {
		q.Exists = &types.ExistsQuery{
			Field: field,
		}
	}
} 
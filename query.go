// Package esb provides a fluent query builder for Elasticsearch using functional options pattern.
// It simplifies the construction of complex Elasticsearch queries while maintaining full compatibility
// with github.com/elastic/go-elasticsearch/v8/typedapi/types.
package esb

import (
	"errors"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// QueryOption represents a function that modifies a types.Query.
// It follows the functional options pattern for building Elasticsearch queries.
type QueryOption func(*types.Query)

// Common errors returned by the query builder.
var (
	ErrInvalidQuery = errors.New("invalid query")
	ErrNoOptions    = errors.New("no query options provided")
)

// NewQuery creates a new Elasticsearch query by applying the provided options.
// It returns a *types.Query that can be used directly with the go-elasticsearch client.
//
// Example:
//   query := esb.NewQuery(
//       esb.Term("status", "published"),
//   )
//   client.Search().Index("articles").Query(query)
func NewQuery(opts ...QueryOption) *types.Query {
	query := &types.Query{}
	for _, opt := range opts {
		opt(query)
	}
	return query
} 
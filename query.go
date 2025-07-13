// Package esb provides a fluent query builder for Elasticsearch using functional options pattern.
// It simplifies the construction of complex Elasticsearch queries while maintaining full compatibility
// with github.com/elastic/go-elasticsearch/v8/typedapi/types.
package esb

import (
	"errors"
	"strings"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// QueryOption represents a function that modifies a types.Query.
// It follows the functional options pattern for building Elasticsearch queries.
type QueryOption func(*types.Query) error

// Common errors returned by the query builder.
var (
	ErrInvalidQuery = errors.New("invalid query")
	ErrEmptyField   = errors.New("field name cannot be empty")
	ErrEmptyValue   = errors.New("value cannot be empty")
	ErrNoOptions    = errors.New("no query options provided")
)

// NewQuery creates a new Elasticsearch query by applying the provided options.
// It returns a *types.Query that can be used directly with the go-elasticsearch client.
//
// Example:
//   query, err := esb.NewQuery(
//       esb.Term("status", "published"),
//   )
//   if err != nil {
//       log.Fatal(err)
//   }
//   
//   client.Search().Index("articles").Query(query)
func NewQuery(opts ...QueryOption) (*types.Query, error) {
	if len(opts) == 0 {
		return nil, ErrNoOptions
	}

	query := &types.Query{}
	
	for _, opt := range opts {
		if opt == nil {
			continue // Skip nil options
		}
		
		if err := opt(query); err != nil {
			return nil, err
		}
	}
	
	return query, nil
}

// validateField checks if a field name is valid (non-empty).
func validateField(field string) error {
	if field == "" || strings.TrimSpace(field) == "" {
		return ErrEmptyField
	}
	return nil
}

// validateValue checks if a value is valid (non-empty for strings).
func validateValue(value interface{}) error {
	if str, ok := value.(string); ok && str == "" {
		return ErrEmptyValue
	}
	return nil
} 
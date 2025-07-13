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
	return func(q *types.Query) error {
		if err := validateField(field); err != nil {
			return err
		}
		if err := validateValue(value); err != nil {
			return err
		}
		
		if q.Term == nil {
			q.Term = make(map[string]types.TermQuery)
		}
		
		q.Term[field] = types.TermQuery{
			Value: value,
		}
		
		return nil
	}
}

// Terms creates a terms query that matches documents containing one or more exact terms.
// Terms queries are used for matching any of the provided values.
//
// Example:
//   esb.Terms("category", "tech", "science", "programming")
func Terms(field string, values ...string) QueryOption {
	return func(q *types.Query) error {
		if err := validateField(field); err != nil {
			return err
		}
		if len(values) == 0 {
			return ErrEmptyValue
		}
		
		// Convert string values to FieldValue
		fieldValues := make([]types.FieldValue, len(values))
		for i, v := range values {
			if err := validateValue(v); err != nil {
				return err
			}
			fieldValues[i] = types.FieldValue(v)
		}
		
		q.Terms = &types.TermsQuery{
			TermsQuery: map[string]types.TermsQueryField{
				field: fieldValues,
			},
		}
		
		return nil
	}
} 
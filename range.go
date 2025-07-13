package esb

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// RangeBuilder provides a fluent interface for building range queries.
// It supports chaining methods for setting range conditions like Gte, Gt, Lte, Lt.
type RangeBuilder struct {
	field string
	query types.UntypedRangeQuery
}

// Range creates a new RangeBuilder for the specified field.
// Use the builder methods to set range conditions, then call Build() to get the QueryOption.
//
// Example:
//   esb.Range("age").Gte(18).Lt(65).Build()
//   esb.Range("price").Gt(10.0).Lte(100.0).Build()
func Range(field string) *RangeBuilder {
	return &RangeBuilder{
		field: field,
		query: types.UntypedRangeQuery{},
	}
}

// Gte sets the "greater than or equal to" condition.
// Supports numbers, strings, and dates.
func (rb *RangeBuilder) Gte(value interface{}) *RangeBuilder {
	rb.query.Gte = marshalValue(value)
	return rb
}

// Gt sets the "greater than" condition.
// Supports numbers, strings, and dates.
func (rb *RangeBuilder) Gt(value interface{}) *RangeBuilder {
	rb.query.Gt = marshalValue(value)
	return rb
}

// Lte sets the "less than or equal to" condition.
// Supports numbers, strings, and dates.
func (rb *RangeBuilder) Lte(value interface{}) *RangeBuilder {
	rb.query.Lte = marshalValue(value)
	return rb
}

// Lt sets the "less than" condition.
// Supports numbers, strings, and dates.
func (rb *RangeBuilder) Lt(value interface{}) *RangeBuilder {
	rb.query.Lt = marshalValue(value)
	return rb
}

// From sets the "from" condition (inclusive).
// Supports numbers, strings, and dates.
func (rb *RangeBuilder) From(value interface{}) *RangeBuilder {
	val := marshalValue(value)
	rb.query.From = &val
	return rb
}

// To sets the "to" condition (exclusive).
// Supports numbers, strings, and dates.
func (rb *RangeBuilder) To(value interface{}) *RangeBuilder {
	val := marshalValue(value)
	rb.query.To = &val
	return rb
}

// Boost sets the boost value for the query.
func (rb *RangeBuilder) Boost(boost float32) *RangeBuilder {
	rb.query.Boost = &boost
	return rb
}

// Format sets the date format for date range queries.
func (rb *RangeBuilder) Format(format string) *RangeBuilder {
	rb.query.Format = &format
	return rb
}

// TimeZone sets the time zone for date range queries.
func (rb *RangeBuilder) TimeZone(timeZone string) *RangeBuilder {
	rb.query.TimeZone = &timeZone
	return rb
}

// Build creates the QueryOption from the configured range builder.
// This method validates the field and returns the final QueryOption.
func (rb *RangeBuilder) Build() QueryOption {
	return func(q *types.Query) error {
		if err := validateField(rb.field); err != nil {
			return err
		}
		
		if q.Range == nil {
			q.Range = make(map[string]types.RangeQuery)
		}
		
		q.Range[rb.field] = rb.query
		return nil
	}
}

// marshalValue converts a value to json.RawMessage for the UntypedRangeQuery.
func marshalValue(value interface{}) json.RawMessage {
	data, err := json.Marshal(value)
	if err != nil {
		// Fallback to string representation if marshaling fails
		return json.RawMessage(`"` + "invalid_value" + `"`)
	}
	return json.RawMessage(data)
} 
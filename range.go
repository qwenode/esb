package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/rangerelation"
)

// =============================================================================
// Type-Safe Range Query Builders
// =============================================================================

// NumberRangeBuilder provides a fluent interface for building number range queries.
// It supports chaining methods for setting numeric range conditions.
type NumberRangeBuilder struct {
	field string
	query types.NumberRangeQuery
}

// DateRangeBuilder provides a fluent interface for building date range queries.
// It supports chaining methods for setting date range conditions with format and timezone support.
type DateRangeBuilder struct {
	field string
	query types.DateRangeQuery
}

// TermRangeBuilder provides a fluent interface for building term (string) range queries.
// It supports chaining methods for setting string range conditions.
type TermRangeBuilder struct {
	field string
	query types.TermRangeQuery
}

// =============================================================================
// Factory Functions
// =============================================================================

// NumberRange creates a new NumberRangeBuilder for numeric range queries.
// Use this for integer, float, and other numeric field types.
//
// Example:
//   esb.NumberRange("age").Gte(18.0).Lt(65.0).Build()
//   esb.NumberRange("price").Gt(10.0).Lte(100.0).Build()
func NumberRange(field string) *NumberRangeBuilder {
	return &NumberRangeBuilder{
		field: field,
		query: types.NumberRangeQuery{},
	}
}

// DateRange creates a new DateRangeBuilder for date range queries.
// Use this for date field types with support for format and timezone.
//
// Example:
//   esb.DateRange("created_at").Gte("2023-01-01").Format("yyyy-MM-dd").Build()
//   esb.DateRange("timestamp").Gt("now-1d").TimeZone("UTC").Build()
func DateRange(field string) *DateRangeBuilder {
	return &DateRangeBuilder{
		field: field,
		query: types.DateRangeQuery{},
	}
}

// TermRange creates a new TermRangeBuilder for string/term range queries.
// Use this for keyword, text, and other string field types.
//
// Example:
//   esb.TermRange("username").Gte("alice").Lt("bob").Build()
//   esb.TermRange("category").From("electronics").To("home").Build()
func TermRange(field string) *TermRangeBuilder {
	return &TermRangeBuilder{
		field: field,
		query: types.TermRangeQuery{},
	}
}

// =============================================================================
// NumberRangeBuilder Methods
// =============================================================================

// Gte sets the "greater than or equal to" condition for numeric values.
func (b *NumberRangeBuilder) Gte(value float64) *NumberRangeBuilder {
	b.query.Gte = (*types.Float64)(&value)
	return b
}

// Gt sets the "greater than" condition for numeric values.
func (b *NumberRangeBuilder) Gt(value float64) *NumberRangeBuilder {
	b.query.Gt = (*types.Float64)(&value)
	return b
}

// Lte sets the "less than or equal to" condition for numeric values.
func (b *NumberRangeBuilder) Lte(value float64) *NumberRangeBuilder {
	b.query.Lte = (*types.Float64)(&value)
	return b
}

// Lt sets the "less than" condition for numeric values.
func (b *NumberRangeBuilder) Lt(value float64) *NumberRangeBuilder {
	b.query.Lt = (*types.Float64)(&value)
	return b
}

// From sets the "from" condition (inclusive) for numeric values.
func (b *NumberRangeBuilder) From(value float64) *NumberRangeBuilder {
	b.query.From = (*types.Float64)(&value)
	return b
}

// To sets the "to" condition (exclusive) for numeric values.
func (b *NumberRangeBuilder) To(value float64) *NumberRangeBuilder {
	b.query.To = (*types.Float64)(&value)
	return b
}

// Boost sets the boost value for the numeric range query.
func (b *NumberRangeBuilder) Boost(boost float32) *NumberRangeBuilder {
	b.query.Boost = &boost
	return b
}

// QueryName sets the query name for the numeric range query.
func (b *NumberRangeBuilder) QueryName(name string) *NumberRangeBuilder {
	b.query.QueryName_ = &name
	return b
}

// Relation sets the relation for the numeric range query.
func (b *NumberRangeBuilder) Relation(relation *rangerelation.RangeRelation) *NumberRangeBuilder {
	b.query.Relation = relation
	return b
}

// Build creates the QueryOption from the configured numeric range builder.
func (b *NumberRangeBuilder) Build() QueryOption {
	return func(q *types.Query) {
		q.Range = map[string]types.RangeQuery{
			b.field: b.query,
		}
	}
}

// =============================================================================
// DateRangeBuilder Methods
// =============================================================================

// Gte sets the "greater than or equal to" condition for date values.
func (b *DateRangeBuilder) Gte(value string) *DateRangeBuilder {
	b.query.Gte = &value
	return b
}

// Gt sets the "greater than" condition for date values.
func (b *DateRangeBuilder) Gt(value string) *DateRangeBuilder {
	b.query.Gt = &value
	return b
}

// Lte sets the "less than or equal to" condition for date values.
func (b *DateRangeBuilder) Lte(value string) *DateRangeBuilder {
	b.query.Lte = &value
	return b
}

// Lt sets the "less than" condition for date values.
func (b *DateRangeBuilder) Lt(value string) *DateRangeBuilder {
	b.query.Lt = &value
	return b
}

// From sets the "from" condition (inclusive) for date values.
func (b *DateRangeBuilder) From(value string) *DateRangeBuilder {
	b.query.From = &value
	return b
}

// To sets the "to" condition (exclusive) for date values.
func (b *DateRangeBuilder) To(value string) *DateRangeBuilder {
	b.query.To = &value
	return b
}

// Format sets the date format for date range queries.
func (b *DateRangeBuilder) Format(format string) *DateRangeBuilder {
	b.query.Format = &format
	return b
}

// TimeZone sets the time zone for date range queries.
func (b *DateRangeBuilder) TimeZone(timeZone string) *DateRangeBuilder {
	b.query.TimeZone = &timeZone
	return b
}

// Boost sets the boost value for the date range query.
func (b *DateRangeBuilder) Boost(boost float32) *DateRangeBuilder {
	b.query.Boost = &boost
	return b
}

// QueryName sets the query name for the date range query.
func (b *DateRangeBuilder) QueryName(name string) *DateRangeBuilder {
	b.query.QueryName_ = &name
	return b
}

// Relation sets the relation for the date range query.
func (b *DateRangeBuilder) Relation(relation *rangerelation.RangeRelation) *DateRangeBuilder {
	b.query.Relation = relation
	return b
}

// Build creates the QueryOption from the configured date range builder.
func (b *DateRangeBuilder) Build() QueryOption {
	return func(q *types.Query) {
		q.Range = map[string]types.RangeQuery{
			b.field: b.query,
		}
	}
}

// =============================================================================
// TermRangeBuilder Methods
// =============================================================================

// Gte sets the "greater than or equal to" condition for string values.
func (b *TermRangeBuilder) Gte(value string) *TermRangeBuilder {
	b.query.Gte = &value
	return b
}

// Gt sets the "greater than" condition for string values.
func (b *TermRangeBuilder) Gt(value string) *TermRangeBuilder {
	b.query.Gt = &value
	return b
}

// Lte sets the "less than or equal to" condition for string values.
func (b *TermRangeBuilder) Lte(value string) *TermRangeBuilder {
	b.query.Lte = &value
	return b
}

// Lt sets the "less than" condition for string values.
func (b *TermRangeBuilder) Lt(value string) *TermRangeBuilder {
	b.query.Lt = &value
	return b
}

// From sets the "from" condition (inclusive) for string values.
func (b *TermRangeBuilder) From(value string) *TermRangeBuilder {
	b.query.From = &value
	return b
}

// To sets the "to" condition (exclusive) for string values.
func (b *TermRangeBuilder) To(value string) *TermRangeBuilder {
	b.query.To = &value
	return b
}

// Boost sets the boost value for the term range query.
func (b *TermRangeBuilder) Boost(boost float32) *TermRangeBuilder {
	b.query.Boost = &boost
	return b
}

// QueryName sets the query name for the term range query.
func (b *TermRangeBuilder) QueryName(name string) *TermRangeBuilder {
	b.query.QueryName_ = &name
	return b
}

// Relation sets the relation for the term range query.
func (b *TermRangeBuilder) Relation(relation *rangerelation.RangeRelation) *TermRangeBuilder {
	b.query.Relation = relation
	return b
}

// Build creates the QueryOption from the configured term range builder.
func (b *TermRangeBuilder) Build() QueryOption {
	return func(q *types.Query) {
		q.Range = map[string]types.RangeQuery{
			b.field: b.query,
		}
	}
} 
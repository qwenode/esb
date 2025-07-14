package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/rangerelation"
)

// =============================================================================
// 类型安全的范围查询构建器
// =============================================================================

// NumberRangeBuilder 提供了一个用于构建数值范围查询的流式接口。
// 它支持链式方法来设置数值范围条件。
type NumberRangeBuilder struct {
	field string
	query types.NumberRangeQuery
}

// DateRangeBuilder 提供了一个用于构建日期范围查询的流式接口。
// 它支持链式方法来设置日期范围条件，并支持格式和时区设置。
type DateRangeBuilder struct {
	field string
	query types.DateRangeQuery
}

// TermRangeBuilder 提供了一个用于构建词项（字符串）范围查询的流式接口。
// 它支持链式方法来设置字符串范围条件。
type TermRangeBuilder struct {
	field string
	query types.TermRangeQuery
}

// =============================================================================
// 工厂函数
// =============================================================================

// NumberRange 创建一个用于数值范围查询的 NumberRangeBuilder。
// 用于整数、浮点数和其他数值类型字段。
//
// 示例：
//   esb.NumberRange("age").Gte(18.0).Lt(65.0).Build()
//   esb.NumberRange("price").Gt(10.0).Lte(100.0).Build()
func NumberRange(field string) *NumberRangeBuilder {
	return &NumberRangeBuilder{
		field: field,
		query: types.NumberRangeQuery{},
	}
}

// DateRange 创建一个用于日期范围查询的 DateRangeBuilder。
// 用于日期类型字段，支持格式和时区设置。
//
// 示例：
//   esb.DateRange("created_at").Gte("2023-01-01").Format("yyyy-MM-dd").Build()
//   esb.DateRange("timestamp").Gt("now-1d").TimeZone("UTC").Build()
func DateRange(field string) *DateRangeBuilder {
	return &DateRangeBuilder{
		field: field,
		query: types.DateRangeQuery{},
	}
}

// TermRange 创建一个用于字符串/词项范围查询的 TermRangeBuilder。
// 用于 keyword、text 和其他字符串类型字段。
//
// 示例：
//   esb.TermRange("username").Gte("alice").Lt("bob").Build()
//   esb.TermRange("category").From("electronics").To("home").Build()
func TermRange(field string) *TermRangeBuilder {
	return &TermRangeBuilder{
		field: field,
		query: types.TermRangeQuery{},
	}
}

// =============================================================================
// NumberRangeBuilder 方法
// =============================================================================

// Gte 设置数值的"大于等于"条件。
func (b *NumberRangeBuilder) Gte(value float64) *NumberRangeBuilder {
	b.query.Gte = (*types.Float64)(&value)
	return b
}

// Gt 设置数值的"大于"条件。
func (b *NumberRangeBuilder) Gt(value float64) *NumberRangeBuilder {
	b.query.Gt = (*types.Float64)(&value)
	return b
}

// Lte 设置数值的"小于等于"条件。
func (b *NumberRangeBuilder) Lte(value float64) *NumberRangeBuilder {
	b.query.Lte = (*types.Float64)(&value)
	return b
}

// Lt 设置数值的"小于"条件。
func (b *NumberRangeBuilder) Lt(value float64) *NumberRangeBuilder {
	b.query.Lt = (*types.Float64)(&value)
	return b
}

// From 设置数值的"起始"条件（包含）。
func (b *NumberRangeBuilder) From(value float64) *NumberRangeBuilder {
	b.query.From = (*types.Float64)(&value)
	return b
}

// To 设置数值的"结束"条件（不包含）。
func (b *NumberRangeBuilder) To(value float64) *NumberRangeBuilder {
	b.query.To = (*types.Float64)(&value)
	return b
}

// Boost 设置数值范围查询的权重值。
func (b *NumberRangeBuilder) Boost(boost float32) *NumberRangeBuilder {
	b.query.Boost = &boost
	return b
}

// QueryName 设置数值范围查询的查询名称。
func (b *NumberRangeBuilder) QueryName(name string) *NumberRangeBuilder {
	b.query.QueryName_ = &name
	return b
}

// Relation 设置数值范围查询的关系。
func (b *NumberRangeBuilder) Relation(relation *rangerelation.RangeRelation) *NumberRangeBuilder {
	b.query.Relation = relation
	return b
}

// Build 从配置的数值范围构建器创建 QueryOption。
func (b *NumberRangeBuilder) Build() QueryOption {
	return func(q *types.Query) {
		q.Range = map[string]types.RangeQuery{
			b.field: b.query,
		}
	}
}

// =============================================================================
// DateRangeBuilder 方法
// =============================================================================

// Gte 设置日期的"大于等于"条件。
func (b *DateRangeBuilder) Gte(value string) *DateRangeBuilder {
	b.query.Gte = &value
	return b
}

// Gt 设置日期的"大于"条件。
func (b *DateRangeBuilder) Gt(value string) *DateRangeBuilder {
	b.query.Gt = &value
	return b
}

// Lte 设置日期的"小于等于"条件。
func (b *DateRangeBuilder) Lte(value string) *DateRangeBuilder {
	b.query.Lte = &value
	return b
}

// Lt 设置日期的"小于"条件。
func (b *DateRangeBuilder) Lt(value string) *DateRangeBuilder {
	b.query.Lt = &value
	return b
}

// From 设置日期的"起始"条件（包含）。
func (b *DateRangeBuilder) From(value string) *DateRangeBuilder {
	b.query.From = &value
	return b
}

// To 设置日期的"结束"条件（不包含）。
func (b *DateRangeBuilder) To(value string) *DateRangeBuilder {
	b.query.To = &value
	return b
}

// Format 设置日期范围查询的日期格式。
func (b *DateRangeBuilder) Format(format string) *DateRangeBuilder {
	b.query.Format = &format
	return b
}

// TimeZone 设置日期范围查询的时区。
func (b *DateRangeBuilder) TimeZone(timeZone string) *DateRangeBuilder {
	b.query.TimeZone = &timeZone
	return b
}

// Boost 设置日期范围查询的权重值。
func (b *DateRangeBuilder) Boost(boost float32) *DateRangeBuilder {
	b.query.Boost = &boost
	return b
}

// QueryName 设置日期范围查询的查询名称。
func (b *DateRangeBuilder) QueryName(name string) *DateRangeBuilder {
	b.query.QueryName_ = &name
	return b
}

// Relation 设置日期范围查询的关系。
func (b *DateRangeBuilder) Relation(relation *rangerelation.RangeRelation) *DateRangeBuilder {
	b.query.Relation = relation
	return b
}

// Build 从配置的日期范围构建器创建 QueryOption。
func (b *DateRangeBuilder) Build() QueryOption {
	return func(q *types.Query) {
		q.Range = map[string]types.RangeQuery{
			b.field: b.query,
		}
	}
}

// =============================================================================
// TermRangeBuilder 方法
// =============================================================================

// Gte 设置字符串的"大于等于"条件。
func (b *TermRangeBuilder) Gte(value string) *TermRangeBuilder {
	b.query.Gte = &value
	return b
}

// Gt 设置字符串的"大于"条件。
func (b *TermRangeBuilder) Gt(value string) *TermRangeBuilder {
	b.query.Gt = &value
	return b
}

// Lte 设置字符串的"小于等于"条件。
func (b *TermRangeBuilder) Lte(value string) *TermRangeBuilder {
	b.query.Lte = &value
	return b
}

// Lt 设置字符串的"小于"条件。
func (b *TermRangeBuilder) Lt(value string) *TermRangeBuilder {
	b.query.Lt = &value
	return b
}

// From 设置字符串的"起始"条件（包含）。
func (b *TermRangeBuilder) From(value string) *TermRangeBuilder {
	b.query.From = &value
	return b
}

// To 设置字符串的"结束"条件（不包含）。
func (b *TermRangeBuilder) To(value string) *TermRangeBuilder {
	b.query.To = &value
	return b
}

// Boost 设置词项范围查询的权重值。
func (b *TermRangeBuilder) Boost(boost float32) *TermRangeBuilder {
	b.query.Boost = &boost
	return b
}

// QueryName 设置词项范围查询的查询名称。
func (b *TermRangeBuilder) QueryName(name string) *TermRangeBuilder {
	b.query.QueryName_ = &name
	return b
}

// Relation 设置词项范围查询的关系。
func (b *TermRangeBuilder) Relation(relation *rangerelation.RangeRelation) *TermRangeBuilder {
	b.query.Relation = relation
	return b
}

// Build 从配置的词项范围构建器创建 QueryOption。
func (b *TermRangeBuilder) Build() QueryOption {
	return func(q *types.Query) {
		q.Range = map[string]types.RangeQuery{
			b.field: b.query,
		}
	}
} 
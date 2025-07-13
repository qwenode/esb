package esb

import (
	"sync"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// =============================================================================
// 简化的性能优化示例
// =============================================================================

// OptimizedNumberRangeBuilder 优化版本的数字范围查询构建器
// 使用对象池减少内存分配
type OptimizedNumberRangeBuilder struct {
	field string
	query types.NumberRangeQuery
}

var numberRangeBuilderPool = sync.Pool{
	New: func() interface{} {
		return &OptimizedNumberRangeBuilder{}
	},
}

// OptimizedNumberRange 优化版本的数字范围查询工厂函数
func OptimizedNumberRange(field string) *OptimizedNumberRangeBuilder {
	builder := numberRangeBuilderPool.Get().(*OptimizedNumberRangeBuilder)
	builder.field = field
	builder.query = types.NumberRangeQuery{} // 重置查询状态
	return builder
}

// Gte 优化版本的 Gte 方法，避免不必要的类型转换
func (b *OptimizedNumberRangeBuilder) Gte(value float64) *OptimizedNumberRangeBuilder {
	b.query.Gte = (*types.Float64)(&value)
	return b
}

// Lt 优化版本的 Lt 方法
func (b *OptimizedNumberRangeBuilder) Lt(value float64) *OptimizedNumberRangeBuilder {
	b.query.Lt = (*types.Float64)(&value)
	return b
}

// Build 构建查询选项，使用后会自动回收到对象池
func (b *OptimizedNumberRangeBuilder) Build() QueryOption {
	field := b.field
	query := b.query
	numberRangeBuilderPool.Put(b)
	return func(q *types.Query) {
		if q.Range == nil {
			q.Range = make(map[string]types.RangeQuery)
		}
		q.Range[field] = query
	}
}

// =============================================================================
// 预编译查询模板优化
// =============================================================================

// 常用查询模板缓存
var commonQueryTemplates = map[string]*types.Query{
	"term_status_published": {
		Term: map[string]types.TermQuery{
			"status": {Value: "published"},
		},
	},
	"term_active_true": {
		Term: map[string]types.TermQuery{
			"active": {Value: "true"},
		},
	},
	"exists_author": {
		Exists: &types.ExistsQuery{
			Field: "author",
		},
	},
}

// FastTerm 快速 Term 查询，对常用查询使用预编译模板
func FastTerm(field, value string) QueryOption {
	templateKey := "term_" + field + "_" + value
	if template, exists := commonQueryTemplates[templateKey]; exists {
		return func(q *types.Query) {
			*q = *template // 直接复制预编译的查询
		}
	}
	return Term(field, value) // 回退到标准实现
}

// FastExists 快速 Exists 查询
func FastExists(field string) QueryOption {
	templateKey := "exists_" + field
	if template, exists := commonQueryTemplates[templateKey]; exists {
		return func(q *types.Query) {
			*q = *template
		}
	}
	return Exists(field)
}

// =============================================================================
// 批量查询优化
// =============================================================================

// BatchQueryBuilder 批量查询构建器，减少函数调用开销
type BatchQueryBuilder struct {
	options []QueryOption
}

// NewBatchQueryBuilder 创建批量查询构建器
func NewBatchQueryBuilder() *BatchQueryBuilder {
	return &BatchQueryBuilder{
		options: make([]QueryOption, 0, 8), // 预分配常用容量
	}
}

// Add 添加查询选项
func (b *BatchQueryBuilder) Add(option QueryOption) *BatchQueryBuilder {
	b.options = append(b.options, option)
	return b
}

// Build 批量构建查询
func (b *BatchQueryBuilder) Build() (*types.Query, error) {
	q := &types.Query{}
	for _, opt := range b.options {
		opt(q)
	}
	return q, nil
}

// =============================================================================
// 编译时优化代码生成示例
// =============================================================================

// GeneratedTermQuery 编译时生成的 Term 查询
func GeneratedTermQuery(field, value string) *types.Query {
	return &types.Query{
		Term: map[string]types.TermQuery{
			field: {Value: value},
		},
	}
}

// GeneratedMatchQuery 编译时生成的 Match 查询
func GeneratedMatchQuery(field string, value string) *types.Query {
	return &types.Query{
		Match: map[string]types.MatchQuery{
			field: {Query: value},
		},
	}
}

// =============================================================================
// 性能监控工具
// =============================================================================

// QueryPerformanceMetrics 查询性能指标
type QueryPerformanceMetrics struct {
	QueryType     string
	ExecutionTime int64 // 纳秒
	MemoryUsage   int64 // 字节
	AllocCount    int64 // 分配次数
}

// PerformanceMonitor 性能监控器
type PerformanceMonitor struct {
	metrics []QueryPerformanceMetrics
	mutex   sync.RWMutex
}

// NewPerformanceMonitor 创建性能监控器
func NewPerformanceMonitor() *PerformanceMonitor {
	return &PerformanceMonitor{
		metrics: make([]QueryPerformanceMetrics, 0, 1000),
	}
}

// RecordMetric 记录性能指标
func (pm *PerformanceMonitor) RecordMetric(metric QueryPerformanceMetrics) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	pm.metrics = append(pm.metrics, metric)
}

// GetAverageExecutionTime 获取平均执行时间
func (pm *PerformanceMonitor) GetAverageExecutionTime(queryType string) float64 {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()
	
	var total int64
	var count int64
	
	for _, metric := range pm.metrics {
		if metric.QueryType == queryType {
			total += metric.ExecutionTime
			count++
		}
	}
	
	if count == 0 {
		return 0
	}
	
	return float64(total) / float64(count)
}

// =============================================================================
// 使用示例
// =============================================================================

// ExampleOptimizedUsage 展示如何使用优化后的 API
func ExampleOptimizedUsage() {
	// 使用对象池优化的范围查询
	query1 := NewQuery(
		OptimizedNumberRange("age").Gte(18.0).Lt(65.0).Build(),
	)
	_ = query1

	// 使用预编译模板的快速查询
	query2 := NewQuery(
		FastTerm("status", "published"),
		FastExists("author"),
	)
	_ = query2

	// 使用批量查询构建器
	batchBuilder := NewBatchQueryBuilder()
	query3, _ := batchBuilder.
		Add(FastTerm("active", "true")).
		Add(OptimizedNumberRange("score").Gte(4.0).Build()).
		Build()
	_ = query3
} 
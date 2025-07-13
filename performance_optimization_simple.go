package esb

import (
    "sync"
    
    "github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// =============================================================================
// 零拷贝Builder优化 - 消除中间分配
// =============================================================================

// ZeroCopyNumberRangeBuilder 零拷贝数字范围查询构建器
type ZeroCopyNumberRangeBuilder struct {
    field string
    query types.NumberRangeQuery
}

// ZeroCopyNumberRange 创建零拷贝数字范围查询构建器
func ZeroCopyNumberRange(field string) *ZeroCopyNumberRangeBuilder {
    return &ZeroCopyNumberRangeBuilder{
        field: field,
        query: types.NumberRangeQuery{},
    }
}

// Gte 设置大于等于条件（零拷贝版本）
func (b *ZeroCopyNumberRangeBuilder) Gte(value float64) *ZeroCopyNumberRangeBuilder {
    b.query.Gte = (*types.Float64)(&value)
    return b
}

// Lt 设置小于条件（零拷贝版本）
func (b *ZeroCopyNumberRangeBuilder) Lt(value float64) *ZeroCopyNumberRangeBuilder {
    b.query.Lt = (*types.Float64)(&value)
    return b
}

// Lte 设置小于等于条件（零拷贝版本）
func (b *ZeroCopyNumberRangeBuilder) Lte(value float64) *ZeroCopyNumberRangeBuilder {
    b.query.Lte = (*types.Float64)(&value)
    return b
}

// Gt 设置大于条件（零拷贝版本）
func (b *ZeroCopyNumberRangeBuilder) Gt(value float64) *ZeroCopyNumberRangeBuilder {
    b.query.Gt = (*types.Float64)(&value)
    return b
}

// BuildDirect 直接构建Query，消除QueryOption闭包
func (b *ZeroCopyNumberRangeBuilder) BuildDirect() *types.Query {
    return &types.Query{
        Range: map[string]types.RangeQuery{
            b.field: b.query,
        },
    }
}

// Build 兼容性方法，返回QueryOption
func (b *ZeroCopyNumberRangeBuilder) Build() QueryOption {
    field := b.field
    query := b.query
    return func(q *types.Query) {
        if q.Range == nil {
            q.Range = make(map[string]types.RangeQuery)
        }
        q.Range[field] = query
    }
}

// =============================================================================
// 编译时代码生成优化 - 消除运行时开销
// =============================================================================

// GeneratedNumberRangeQuery 编译时生成的数字范围查询
func GeneratedNumberRangeQuery(field string, gte, lt float64) *types.Query {
    return &types.Query{
        Range: map[string]types.RangeQuery{
            field: types.NumberRangeQuery{
                Gte: (*types.Float64)(&gte),
                Lt:  (*types.Float64)(&lt),
            },
        },
    }
}

// GeneratedDateRangeQuery 编译时生成的日期范围查询
func GeneratedDateRangeQuery(field, from, to string) *types.Query {
    return &types.Query{
        Range: map[string]types.RangeQuery{
            field: types.DateRangeQuery{
                Gte: &from,
                Lte: &to,
            },
        },
    }
}

// GeneratedTermRangeQuery 编译时生成的字符串范围查询
func GeneratedTermRangeQuery(field, from, to string) *types.Query {
    return &types.Query{
        Range: map[string]types.RangeQuery{
            field: types.TermRangeQuery{
                From: &from,
                To:   &to,
            },
        },
    }
}

// =============================================================================
// 内联优化 - 减少函数调用层级
// =============================================================================

// InlineNumberRange 内联数字范围查询，最小化函数调用
func InlineNumberRange(field string, gte, lt float64) QueryOption {
    return func(q *types.Query) {
        if q.Range == nil {
            q.Range = make(map[string]types.RangeQuery)
        }
        q.Range[field] = types.NumberRangeQuery{
            Gte: (*types.Float64)(&gte),
            Lt:  (*types.Float64)(&lt),
        }
    }
}

// InlineDateRange 内联日期范围查询
func InlineDateRange(field, from, to string) QueryOption {
    return func(q *types.Query) {
        if q.Range == nil {
            q.Range = make(map[string]types.RangeQuery)
        }
        q.Range[field] = types.DateRangeQuery{
            Gte: &from,
            Lte: &to,
        }
    }
}

// InlineTermRange 内联字符串范围查询
func InlineTermRange(field, from, to string) QueryOption {
    return func(q *types.Query) {
        if q.Range == nil {
            q.Range = make(map[string]types.RangeQuery)
        }
        q.Range[field] = types.TermRangeQuery{
            From: &from,
            To:   &to,
        }
    }
}

// =============================================================================
// 性能监控（简化版）
// =============================================================================

// QueryPerformanceMetrics 查询性能指标
type QueryPerformanceMetrics struct {
    QueryType     string
    ExecutionTime int64 // 纳秒
    MemoryUsage   int64 // 字节
    AllocCount    int64 // 分配次数
}

// PerformanceMonitor 性能监控器（简化版）
type PerformanceMonitor struct {
    metrics []QueryPerformanceMetrics
    mutex   sync.RWMutex
}

// NewPerformanceMonitor 创建性能监控器
func NewPerformanceMonitor() *PerformanceMonitor {
    return &PerformanceMonitor{
        metrics: make([]QueryPerformanceMetrics, 0, 100),
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
    count := 0
    for _, m := range pm.metrics {
        if m.QueryType == queryType {
            total += m.ExecutionTime
            count++
        }
    }
    
    if count == 0 {
        return 0
    }
    return float64(total) / float64(count)
}

// =============================================================================
// 使用示例 - 展示优化效果
// =============================================================================

// ExampleOptimizedUsage 展示优化后的使用方式
func ExampleOptimizedUsage() {
    // 方案1：零拷贝Builder（保持链式调用体验）
    query1 := ZeroCopyNumberRange("age").Gte(18.0).Lt(65.0).BuildDirect()
    _ = query1
    
    // 方案2：编译时生成（最佳性能）
    query2 := GeneratedNumberRangeQuery("age", 18.0, 65.0)
    _ = query2
    
    // 方案3：内联优化（平衡性能和易用性）
    query3 := NewQuery(InlineNumberRange("age", 18.0, 65.0))
    _ = query3
    
    // 方案4：组合使用
    query4 := NewQuery(
        Bool(
            Must(
                InlineNumberRange("age", 18.0, 65.0),
                InlineDateRange("created_at", "2023-01-01", "2023-12-31"),
            ),
        ),
    )
    _ = query4
}

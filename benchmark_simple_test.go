package esb

import (
	"testing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// BenchmarkOptimizedVsOriginal 对比优化版本与原版本的性能
func BenchmarkOptimizedVsOriginal(b *testing.B) {
	b.Run("Original_NumberRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(NumberRange("age").Gte(18.0).Lt(65.0).Build())
			if err != nil {
				b.Errorf("Original NumberRange failed: %v", err)
			}
		}
	})

	b.Run("Optimized_NumberRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(OptimizedNumberRange("age").Gte(18.0).Lt(65.0).Build())
			if err != nil {
				b.Errorf("Optimized NumberRange failed: %v", err)
			}
		}
	})

	b.Run("Native_NumberRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gte := types.Float64(18.0)
			lt := types.Float64(65.0)
			query := &types.Query{
				Range: map[string]types.RangeQuery{
					"age": types.NumberRangeQuery{
						Gte: &gte,
						Lt:  &lt,
					},
				},
			}
			_ = query
		}
	})
}

// BenchmarkFastTermVsOriginal 对比快速Term查询与原版本
func BenchmarkFastTermVsOriginal(b *testing.B) {
	b.Run("Original_Term", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(Term("status", "published"))
			if err != nil {
				b.Errorf("Original Term failed: %v", err)
			}
		}
	})

	b.Run("Fast_Term_Cached", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(FastTerm("status", "published"))
			if err != nil {
				b.Errorf("Fast Term failed: %v", err)
			}
		}
	})

	b.Run("Fast_Term_NotCached", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(FastTerm("category", "tech"))
			if err != nil {
				b.Errorf("Fast Term failed: %v", err)
			}
		}
	})

	b.Run("Native_Term", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			query := &types.Query{
				Term: map[string]types.TermQuery{
					"status": {Value: "published"},
				},
			}
			_ = query
		}
	})
}

// BenchmarkBatchQueryBuilder 测试批量查询构建器性能
func BenchmarkBatchQueryBuilder(b *testing.B) {
	b.Run("Individual_Queries", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(
				Term("status", "published"),
				Match("title", "elasticsearch"),
				NumberRange("age").Gte(18.0).Build(),
			)
			if err != nil {
				b.Errorf("Individual queries failed: %v", err)
			}
		}
	})

	b.Run("Batch_Query_Builder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			builder := NewBatchQueryBuilder()
			_, err := builder.
				Add(Term("status", "published")).
				Add(Match("title", "elasticsearch")).
				Add(NumberRange("age").Gte(18.0).Build()).
				Build()
			if err != nil {
				b.Errorf("Batch query builder failed: %v", err)
			}
		}
	})
}

// BenchmarkGeneratedQueries 测试编译时生成的查询性能
func BenchmarkGeneratedQueries(b *testing.B) {
	b.Run("ESB_Term", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(Term("status", "published"))
			if err != nil {
				b.Errorf("ESB Term failed: %v", err)
			}
		}
	})

	b.Run("Generated_Term", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			query := GeneratedTermQuery("status", "published")
			_ = query
		}
	})

	b.Run("Generated_Match", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			query := GeneratedMatchQuery("title", "elasticsearch")
			_ = query
		}
	})

	b.Run("Native_Term", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			query := &types.Query{
				Term: map[string]types.TermQuery{
					"status": {Value: "published"},
				},
			}
			_ = query
		}
	})
}

// BenchmarkComplexOptimizedQuery 测试复杂查询的优化效果
func BenchmarkComplexOptimizedQuery(b *testing.B) {
	b.Run("Original_Complex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(
				Bool(
					Must(
						Match("title", "elasticsearch"),
						Term("status", "published"),
						NumberRange("age").Gte(18.0).Lt(65.0).Build(),
					),
					Should(
						Term("category", "tech"),
						Exists("author"),
					),
				),
			)
			if err != nil {
				b.Errorf("Original complex query failed: %v", err)
			}
		}
	})

	b.Run("Optimized_Complex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(
				Bool(
					Must(
						Match("title", "elasticsearch"),
						FastTerm("status", "published"),
						OptimizedNumberRange("age").Gte(18.0).Lt(65.0).Build(),
					),
					Should(
						Term("category", "tech"),
						FastExists("author"),
					),
				),
			)
			if err != nil {
				b.Errorf("Optimized complex query failed: %v", err)
			}
		}
	})
}

// BenchmarkMemoryAllocationOptimized 测试内存分配优化效果
func BenchmarkMemoryAllocationOptimized(b *testing.B) {
	b.Run("Original_Memory", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(
				Bool(
					Must(
						Term("status", "published"),
						NumberRange("age").Gte(18.0).Build(),
					),
				),
			)
			if err != nil {
				b.Errorf("Original memory test failed: %v", err)
			}
		}
	})

	b.Run("Optimized_Memory", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(
				Bool(
					Must(
						FastTerm("status", "published"),
						OptimizedNumberRange("age").Gte(18.0).Build(),
					),
				),
			)
			if err != nil {
				b.Errorf("Optimized memory test failed: %v", err)
			}
		}
	})
} 
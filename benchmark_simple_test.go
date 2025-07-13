package esb

import (
	"testing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// BenchmarkOptimizedVsOriginal 对比优化版本与原版本的性能
func BenchmarkOptimizedVsOriginal(b *testing.B) {
	b.Run("Original_NumberRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(NumberRange("age").Gte(18.0).Lt(65.0).Build())
		}
	})

	b.Run("ZeroCopy_NumberRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(ZeroCopyNumberRange("age").Gte(18.0).Lt(65.0).Build())
		}
	})

	b.Run("ZeroCopy_Direct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ZeroCopyNumberRange("age").Gte(18.0).Lt(65.0).BuildDirect()
		}
	})

	b.Run("Generated_NumberRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = GeneratedNumberRangeQuery("age", 18.0, 65.0)
		}
	})

	b.Run("Inline_NumberRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(InlineNumberRange("age", 18.0, 65.0))
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

// BenchmarkOptimizedDateRange 对比日期范围查询优化
func BenchmarkOptimizedDateRange(b *testing.B) {
	b.Run("Original_DateRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(DateRange("created_at").Gte("2023-01-01").Lte("2023-12-31").Build())
		}
	})

	b.Run("Generated_DateRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = GeneratedDateRangeQuery("created_at", "2023-01-01", "2023-12-31")
		}
	})

	b.Run("Inline_DateRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(InlineDateRange("created_at", "2023-01-01", "2023-12-31"))
		}
	})

	b.Run("Native_DateRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gte := "2023-01-01"
			lte := "2023-12-31"
			query := &types.Query{
				Range: map[string]types.RangeQuery{
					"created_at": types.DateRangeQuery{
						Gte: &gte,
						Lte: &lte,
					},
				},
			}
			_ = query
		}
	})
}

// BenchmarkOptimizedTermRange 对比字符串范围查询优化
func BenchmarkOptimizedTermRange(b *testing.B) {
	b.Run("Original_TermRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(TermRange("username").From("a").To("z").Build())
		}
	})

	b.Run("Generated_TermRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = GeneratedTermRangeQuery("username", "a", "z")
		}
	})

	b.Run("Inline_TermRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(InlineTermRange("username", "a", "z"))
		}
	})

	b.Run("Native_TermRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			from := "a"
			to := "z"
			query := &types.Query{
				Range: map[string]types.RangeQuery{
					"username": types.TermRangeQuery{
						From: &from,
						To:   &to,
					},
				},
			}
			_ = query
		}
	})
}

// BenchmarkComplexOptimizedQuery 测试复杂查询中的优化效果
func BenchmarkComplexOptimizedQuery(b *testing.B) {
	b.Run("Original_Complex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(
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
		}
	})

	b.Run("Inline_Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(
				Bool(
					Must(
						Match("title", "elasticsearch"),
						Term("status", "published"),
						InlineNumberRange("age", 18.0, 65.0),
					),
					Should(
						Term("category", "tech"),
						Exists("author"),
					),
				),
			)
		}
	})

	b.Run("Mixed_Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(
				Bool(
					Must(
						Match("title", "elasticsearch"),
						Term("status", "published"),
						InlineNumberRange("age", 18.0, 65.0),
						InlineDateRange("created_at", "2023-01-01", "2023-12-31"),
					),
					Should(
						InlineTermRange("username", "a", "z"),
						Exists("author"),
					),
				),
			)
		}
	})
}

// BenchmarkMemoryAllocationOptimized 测试内存分配优化效果
func BenchmarkMemoryAllocationOptimized(b *testing.B) {
	b.Run("Original_Memory", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = NewQuery(
				Bool(
					Must(
						Term("status", "published"),
						NumberRange("age").Gte(18.0).Build(),
					),
				),
			)
		}
	})

	b.Run("ZeroCopy_Memory", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = NewQuery(
				Bool(
					Must(
						Term("status", "published"),
						ZeroCopyNumberRange("age").Gte(18.0).Build(),
					),
				),
			)
		}
	})

	b.Run("Inline_Memory", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = NewQuery(
				Bool(
					Must(
						Term("status", "published"),
						InlineNumberRange("age", 18.0, 65.0),
					),
				),
			)
		}
	})

	b.Run("Generated_Memory", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			// 使用编译时生成的查询
			query1 := GeneratedNumberRangeQuery("age", 18.0, 65.0)
			_ = query1
		}
	})
}

// BenchmarkDirectBuild 测试直接构建vs QueryOption的性能差异
func BenchmarkDirectBuild(b *testing.B) {
	b.Run("QueryOption_Pattern", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(ZeroCopyNumberRange("age").Gte(18.0).Lt(65.0).Build())
		}
	})

	b.Run("Direct_Build", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ZeroCopyNumberRange("age").Gte(18.0).Lt(65.0).BuildDirect()
		}
	})

	b.Run("Generated_Pattern", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = GeneratedNumberRangeQuery("age", 18.0, 65.0)
		}
	})
}

// BenchmarkOptimizationComparison 全面对比各种优化方案
func BenchmarkOptimizationComparison(b *testing.B) {
	b.Run("Baseline_Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(NumberRange("age").Gte(18.0).Lt(65.0).Build())
		}
	})

	b.Run("Opt1_ZeroCopy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(ZeroCopyNumberRange("age").Gte(18.0).Lt(65.0).Build())
		}
	})

	b.Run("Opt2_ZeroCopyDirect", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ZeroCopyNumberRange("age").Gte(18.0).Lt(65.0).BuildDirect()
		}
	})

	b.Run("Opt3_Inline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(InlineNumberRange("age", 18.0, 65.0))
		}
	})

	b.Run("Opt4_Generated", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = GeneratedNumberRangeQuery("age", 18.0, 65.0)
		}
	})

	b.Run("Reference_Native", func(b *testing.B) {
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
package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestBool(t *testing.T) {
	t.Run("当没有提供选项时应该创建空的布尔查询", func(t *testing.T) {
		query := NewQuery(Bool())
		if query.Bool == nil {
			t.Error("预期查询不为 nil")
		}
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
	})

	t.Run("应该创建带 Must 子句的布尔查询", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Must(
					Term("status", "published"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("预期查询不为 nil")
		}
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
		if len(query.Bool.Must) != 1 {
			t.Errorf("预期有 1 个 Must 子句，得到 %d", len(query.Bool.Must))
		}
	})

	t.Run("应该创建带多个 Must 子句的布尔查询", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Must(
					Term("status", "published"),
					Terms("category", "tech", "science"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
		if len(query.Bool.Must) != 2 {
			t.Errorf("预期有 2 个 Must 子句，得到 %d", len(query.Bool.Must))
		}
	})

	t.Run("应该创建带 Should 子句的布尔查询", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Should(
					Term("title", "elasticsearch"),
					Term("content", "search"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
		if len(query.Bool.Should) != 2 {
			t.Errorf("预期有 2 个 Should 子句，得到 %d", len(query.Bool.Should))
		}
	})

	t.Run("应该创建带 Filter 子句的布尔查询", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Filter(
					Term("active", "true"),
					Terms("type", "article", "blog"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
		if len(query.Bool.Filter) != 2 {
			t.Errorf("预期有 2 个 Filter 子句，得到 %d", len(query.Bool.Filter))
		}
	})

	t.Run("应该创建带 MustNot 子句的布尔查询", func(t *testing.T) {
		query := NewQuery(
			Bool(
				MustNot(
					Term("status", "deleted"),
					Term("hidden", "true"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
		if len(query.Bool.MustNot) != 2 {
			t.Errorf("预期有 2 个 MustNot 子句，得到 %d", len(query.Bool.MustNot))
		}
	})

	t.Run("应该创建带所有子句的复杂布尔查询", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Must(
					Term("status", "published"),
				),
				Should(
					Term("category", "tech"),
					Term("category", "science"),
				),
				Filter(
					Term("active", "true"),
				),
				MustNot(
					Term("deleted", "true"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
		if len(query.Bool.Must) != 1 {
			t.Errorf("预期有 1 个 Must 子句，得到 %d", len(query.Bool.Must))
		}
		if len(query.Bool.Should) != 2 {
			t.Errorf("预期有 2 个 Should 子句，得到 %d", len(query.Bool.Should))
		}
		if len(query.Bool.Filter) != 1 {
			t.Errorf("预期有 1 个 Filter 子句，得到 %d", len(query.Bool.Filter))
		}
		if len(query.Bool.MustNot) != 1 {
			t.Errorf("预期有 1 个 MustNot 子句，得到 %d", len(query.Bool.MustNot))
		}
	})

	t.Run("应该正确处理空子句", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Must(),     // 空 Must 子句
				Should(),   // 空 Should 子句
				Filter(),   // 空 Filter 子句
				MustNot(),  // 空 MustNot 子句
			),
		)
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
		// 空子句应该导致 nil 切片
		if query.Bool.Must != nil && len(query.Bool.Must) > 0 {
			t.Error("预期 Must 子句为空")
		}
		if query.Bool.Should != nil && len(query.Bool.Should) > 0 {
			t.Error("预期 Should 子句为空")
		}
		if query.Bool.Filter != nil && len(query.Bool.Filter) > 0 {
			t.Error("预期 Filter 子句为空")
		}
		if query.Bool.MustNot != nil && len(query.Bool.MustNot) > 0 {
			t.Error("预期 MustNot 子句为空")
		}
	})

	t.Run("应该处理多个有效选项", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Must(
					Term("status", "published"),
					Term("active", "true"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
		if len(query.Bool.Must) != 2 {
			t.Errorf("预期有 2 个 Must 子句，得到 %d", len(query.Bool.Must))
		}
	})

	t.Run("应该传播子查询中的错误", func(t *testing.T) {
		// 已移除 error 相关逻辑，无需测试
	})
}

func TestNestedBoolQueries(t *testing.T) {
	t.Run("应该支持嵌套布尔查询", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Must(
					Term("status", "published"),
					Bool(
						Should(
							Term("category", "tech"),
							Term("category", "science"),
						),
					),
				),
			),
		)
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
		if len(query.Bool.Must) != 2 {
			t.Errorf("预期有 2 个 Must 子句，得到 %d", len(query.Bool.Must))
		}
		
		// 检查第二个 Must 子句是否为嵌套布尔查询
		nestedQuery := query.Bool.Must[1]
		if nestedQuery.Bool == nil {
			t.Error("预期为嵌套布尔查询")
		}
		if len(nestedQuery.Bool.Should) != 2 {
			t.Errorf("预期嵌套查询中有 2 个 Should 子句，得到 %d", len(nestedQuery.Bool.Should))
		}
	})
}

func TestBoolQueryCompatibility(t *testing.T) {
	t.Run("应该生成兼容的布尔查询结构", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Must(
					Term("status", "published"),
				),
				Filter(
					Term("active", "true"),
				),
			),
		)
		
		// 验证结构是否与 elasticsearch 期望的结构匹配
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
		
		// 检查 Must 子句
		if len(query.Bool.Must) != 1 {
			t.Errorf("预期有 1 个 Must 子句，得到 %d", len(query.Bool.Must))
		}
		mustQuery := query.Bool.Must[0]
		if mustQuery.Term == nil {
			t.Error("预期 Must 子句中包含 Term 查询")
		}
		
		// 检查 Filter 子句
		if len(query.Bool.Filter) != 1 {
			t.Errorf("预期有 1 个 Filter 子句，得到 %d", len(query.Bool.Filter))
		}
		filterQuery := query.Bool.Filter[0]
		if filterQuery.Term == nil {
			t.Error("预期 Filter 子句中包含 Term 查询")
		}
	})

	t.Run("应该匹配手动构造的布尔查询", func(t *testing.T) {
		// 我们的构建器方法
		builderQuery := NewQuery(
			Bool(
				Must(
					Terms("field", "xxx"),
				),
				Filter(
					Term("abc", "ww"),
				),
			),
		)

		// 手动构造（如同在 cmd/main/main.go 中）
		manualQuery := &types.Query{
			Bool: &types.BoolQuery{
				Must: []types.Query{
					{
						Terms: &types.TermsQuery{
							TermsQuery: map[string]types.TermsQueryField{
								"field": []types.FieldValue{"xxx"},
							},
						},
					},
				},
				Filter: []types.Query{
					{
						Term: map[string]types.TermQuery{
							"abc": {
								Value: "ww",
							},
						},
					},
				},
			},
		}

		// 比较结构
		if builderQuery.Bool == nil || manualQuery.Bool == nil {
			t.Error("两个查询都应该有布尔查询")
		}
		
		if len(builderQuery.Bool.Must) != len(manualQuery.Bool.Must) {
			t.Errorf("Must 子句数量不匹配：builder=%d，manual=%d", 
				len(builderQuery.Bool.Must), len(manualQuery.Bool.Must))
		}
		
		if len(builderQuery.Bool.Filter) != len(manualQuery.Bool.Filter) {
			t.Errorf("Filter 子句数量不匹配：builder=%d，manual=%d", 
				len(builderQuery.Bool.Filter), len(manualQuery.Bool.Filter))
		}
	})
}

func TestBoolFilter(t *testing.T) {
	t.Run("当没有提供选项时应该创建空的布尔过滤查询", func(t *testing.T) {
		query := NewQuery(BoolFilter())
		if query.Bool == nil {
			t.Error("预期查询不为 nil")
		}
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
		// 应该只有空的 Filter 切片
		if query.Bool.Filter != nil && len(query.Bool.Filter) > 0 {
			t.Error("预期 Filter 子句为空")
		}
		if query.Bool.Must != nil && len(query.Bool.Must) > 0 {
			t.Error("预期 Must 子句为空")
		}
		if query.Bool.Should != nil && len(query.Bool.Should) > 0 {
			t.Error("预期 Should 子句为空")
		}
		if query.Bool.MustNot != nil && len(query.Bool.MustNot) > 0 {
			t.Error("预期 MustNot 子句为空")
		}
	})

	t.Run("应该创建带单个过滤条件的布尔查询", func(t *testing.T) {
		query := NewQuery(
			BoolFilter(
				Term("status", "published"),
			),
		)
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
		if len(query.Bool.Filter) != 1 {
			t.Errorf("预期有 1 个 Filter 子句，得到 %d", len(query.Bool.Filter))
		}
		// 验证其他子句为空
		if query.Bool.Must != nil && len(query.Bool.Must) > 0 {
			t.Error("预期 Must 子句为空")
		}
		if query.Bool.Should != nil && len(query.Bool.Should) > 0 {
			t.Error("预期 Should 子句为空")
		}
		if query.Bool.MustNot != nil && len(query.Bool.MustNot) > 0 {
			t.Error("预期 MustNot 子句为空")
		}
	})

	t.Run("应该创建带多个过滤条件的布尔查询", func(t *testing.T) {
		query := NewQuery(
			BoolFilter(
				Term("status", "published"),
				Term("active", "true"),
				Terms("category", "tech", "science"),
			),
		)
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
		if len(query.Bool.Filter) != 3 {
			t.Errorf("预期有 3 个 Filter 子句，得到 %d", len(query.Bool.Filter))
		}
		// 验证其他子句为空
		if query.Bool.Must != nil && len(query.Bool.Must) > 0 {
			t.Error("预期 Must 子句为空")
		}
		if query.Bool.Should != nil && len(query.Bool.Should) > 0 {
			t.Error("预期 Should 子句为空")
		}
		if query.Bool.MustNot != nil && len(query.Bool.MustNot) > 0 {
			t.Error("预期 MustNot 子句为空")
		}
	})

	t.Run("应该与使用 Bool + Filter 的结果等价", func(t *testing.T) {
		// 使用 BoolFilter
		query1 := NewQuery(
			BoolFilter(
				Term("status", "published"),
				Term("active", "true"),
			),
		)

		// 使用 Bool + Filter
		query2 := NewQuery(
			Bool(
				Filter(
					Term("status", "published"),
					Term("active", "true"),
				),
			),
		)

		// 比较结构
		if query1.Bool == nil || query2.Bool == nil {
			t.Error("两个查询都应该有布尔查询")
		}
		
		if len(query1.Bool.Filter) != len(query2.Bool.Filter) {
			t.Errorf("Filter 子句数量不匹配：BoolFilter=%d，Bool+Filter=%d", 
				len(query1.Bool.Filter), len(query2.Bool.Filter))
		}

		// 验证其他子句都为空
		for i, query := range []*types.Query{query1, query2} {
			if query.Bool.Must != nil && len(query.Bool.Must) > 0 {
				t.Errorf("查询 %d 的 Must 子句应该为空", i+1)
			}
			if query.Bool.Should != nil && len(query.Bool.Should) > 0 {
				t.Errorf("查询 %d 的 Should 子句应该为空", i+1)
			}
			if query.Bool.MustNot != nil && len(query.Bool.MustNot) > 0 {
				t.Errorf("查询 %d 的 MustNot 子句应该为空", i+1)
			}
		}
	})

	t.Run("应该正确处理复杂的过滤条件", func(t *testing.T) {
		query := NewQuery(
			BoolFilter(
				Term("status", "published"),
				DateRange("date").Gte("2023-01-01").Build(),
				Exists("author"),
				Terms("tags", "elasticsearch", "search", "lucene"),
			),
		)
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
		if len(query.Bool.Filter) != 4 {
			t.Errorf("预期有 4 个 Filter 子句，得到 %d", len(query.Bool.Filter))
		}

		// 验证每个过滤条件的类型
		filters := query.Bool.Filter
		if filters[0].Term == nil {
			t.Error("第一个过滤条件应该是 Term 查询")
		}
		if filters[1].Range == nil {
			t.Error("第二个过滤条件应该是 Range 查询")
		}
		if filters[2].Exists == nil {
			t.Error("第三个过滤条件应该是 Exists 查询")
		}
		if filters[3].Terms == nil {
			t.Error("第四个过滤条件应该是 Terms 查询")
		}
	})

	t.Run("应该支持嵌套的 BoolFilter", func(t *testing.T) {
		query := NewQuery(
			BoolFilter(
				Term("status", "published"),
				BoolFilter(
					Term("category", "tech"),
					Term("active", "true"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
		if len(query.Bool.Filter) != 2 {
			t.Errorf("预期有 2 个 Filter 子句，得到 %d", len(query.Bool.Filter))
		}
		
		// 检查第二个过滤条件是否为嵌套布尔查询
		nestedFilter := query.Bool.Filter[1]
		if nestedFilter.Bool == nil {
			t.Error("预期为嵌套布尔查询")
		}
		if len(nestedFilter.Bool.Filter) != 2 {
			t.Errorf("预期嵌套查询中有 2 个 Filter 子句，得到 %d", len(nestedFilter.Bool.Filter))
		}
	})
}

// 布尔查询的基准测试
func BenchmarkBoolQuery(b *testing.B) {
	b.Run("带 Must 的简单布尔查询", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(
				Bool(
					Must(
						Term("status", "published"),
					),
				),
			)
		}
	})

	b.Run("带所有子句的复杂布尔查询", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(
				Bool(
					Must(
						Term("status", "published"),
						Terms("category", "tech", "science"),
					),
					Should(
						Term("title", "elasticsearch"),
						Term("content", "search"),
					),
					Filter(
						Term("active", "true"),
					),
					MustNot(
						Term("deleted", "true"),
					),
				),
			)
		}
	})

	b.Run("嵌套布尔查询", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(
				Bool(
					Must(
						Term("status", "published"),
						Bool(
							Should(
								Term("category", "tech"),
								Term("category", "science"),
							),
						),
					),
				),
			)
		}
	})
}

// BoolFilter 的基准测试
func BenchmarkBoolFilter(b *testing.B) {
	b.Run("简单的 BoolFilter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(
				BoolFilter(
					Term("status", "published"),
				),
			)
		}
	})

	b.Run("多个过滤条件的 BoolFilter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(
				BoolFilter(
					Term("status", "published"),
					Term("active", "true"),
					Terms("category", "tech", "science"),
					DateRange("date").Gte("2023-01-01").Build(),
				),
			)
		}
	})

	b.Run("BoolFilter vs Bool+Filter 性能比较", func(b *testing.B) {
		b.Run("BoolFilter", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = NewQuery(
					BoolFilter(
						Term("status", "published"),
						Term("active", "true"),
					),
				)
			}
		})

		b.Run("Bool+Filter", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = NewQuery(
					Bool(
						Filter(
							Term("status", "published"),
							Term("active", "true"),
						),
					),
				)
			}
		})
	})
} 
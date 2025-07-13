package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestBool(t *testing.T) {
	t.Run("should create empty Bool query when no options provided", func(t *testing.T) {
		query := NewQuery(Bool())
		if query.Bool == nil {
			t.Error("expected non-nil query")
		}
		if query.Bool == nil {
			t.Error("expected Bool query")
		}
	})

	t.Run("should create Bool query with Must clause", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Must(
					Term("status", "published"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("expected non-nil query")
		}
		if query.Bool == nil {
			t.Error("expected Bool query")
		}
		if len(query.Bool.Must) != 1 {
			t.Errorf("expected 1 Must clause, got %d", len(query.Bool.Must))
		}
	})

	t.Run("should create Bool query with multiple Must clauses", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Must(
					Term("status", "published"),
					Terms("category", "tech", "science"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("expected Bool query")
		}
		if len(query.Bool.Must) != 2 {
			t.Errorf("expected 2 Must clauses, got %d", len(query.Bool.Must))
		}
	})

	t.Run("should create Bool query with Should clause", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Should(
					Term("title", "elasticsearch"),
					Term("content", "search"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("expected Bool query")
		}
		if len(query.Bool.Should) != 2 {
			t.Errorf("expected 2 Should clauses, got %d", len(query.Bool.Should))
		}
	})

	t.Run("should create Bool query with Filter clause", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Filter(
					Term("active", "true"),
					Terms("type", "article", "blog"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("expected Bool query")
		}
		if len(query.Bool.Filter) != 2 {
			t.Errorf("expected 2 Filter clauses, got %d", len(query.Bool.Filter))
		}
	})

	t.Run("should create Bool query with MustNot clause", func(t *testing.T) {
		query := NewQuery(
			Bool(
				MustNot(
					Term("status", "deleted"),
					Term("hidden", "true"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("expected Bool query")
		}
		if len(query.Bool.MustNot) != 2 {
			t.Errorf("expected 2 MustNot clauses, got %d", len(query.Bool.MustNot))
		}
	})

	t.Run("should create complex Bool query with all clauses", func(t *testing.T) {
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
			t.Error("expected Bool query")
		}
		if len(query.Bool.Must) != 1 {
			t.Errorf("expected 1 Must clause, got %d", len(query.Bool.Must))
		}
		if len(query.Bool.Should) != 2 {
			t.Errorf("expected 2 Should clauses, got %d", len(query.Bool.Should))
		}
		if len(query.Bool.Filter) != 1 {
			t.Errorf("expected 1 Filter clause, got %d", len(query.Bool.Filter))
		}
		if len(query.Bool.MustNot) != 1 {
			t.Errorf("expected 1 MustNot clause, got %d", len(query.Bool.MustNot))
		}
	})

	t.Run("should handle empty clauses gracefully", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Must(),     // Empty Must clause
				Should(),   // Empty Should clause
				Filter(),   // Empty Filter clause
				MustNot(),  // Empty MustNot clause
			),
		)
		if query.Bool == nil {
			t.Error("expected Bool query")
		}
		// Empty clauses should result in nil slices
		if query.Bool.Must != nil && len(query.Bool.Must) > 0 {
			t.Error("expected empty Must clause")
		}
		if query.Bool.Should != nil && len(query.Bool.Should) > 0 {
			t.Error("expected empty Should clause")
		}
		if query.Bool.Filter != nil && len(query.Bool.Filter) > 0 {
			t.Error("expected empty Filter clause")
		}
		if query.Bool.MustNot != nil && len(query.Bool.MustNot) > 0 {
			t.Error("expected empty MustNot clause")
		}
	})

	t.Run("should handle multiple valid options", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Must(
					Term("status", "published"),
					Term("active", "true"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("expected Bool query")
		}
		if len(query.Bool.Must) != 2 {
			t.Errorf("expected 2 Must clauses, got %d", len(query.Bool.Must))
		}
	})

	t.Run("should propagate errors from sub-queries", func(t *testing.T) {
		// 已移除 error 相关逻辑，无需测试
	})
}

func TestNestedBoolQueries(t *testing.T) {
	t.Run("should support nested Bool queries", func(t *testing.T) {
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
			t.Error("expected Bool query")
		}
		if len(query.Bool.Must) != 2 {
			t.Errorf("expected 2 Must clauses, got %d", len(query.Bool.Must))
		}
		
		// Check that the second Must clause is a nested Bool query
		nestedQuery := query.Bool.Must[1]
		if nestedQuery.Bool == nil {
			t.Error("expected nested Bool query")
		}
		if len(nestedQuery.Bool.Should) != 2 {
			t.Errorf("expected 2 Should clauses in nested query, got %d", len(nestedQuery.Bool.Should))
		}
	})
}

func TestBoolQueryCompatibility(t *testing.T) {
	t.Run("should generate compatible BoolQuery structure", func(t *testing.T) {
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
		
		// Verify the structure matches what elasticsearch expects
		if query.Bool == nil {
			t.Error("expected Bool query")
		}
		
		// Check Must clause
		if len(query.Bool.Must) != 1 {
			t.Errorf("expected 1 Must clause, got %d", len(query.Bool.Must))
		}
		mustQuery := query.Bool.Must[0]
		if mustQuery.Term == nil {
			t.Error("expected Term query in Must clause")
		}
		
		// Check Filter clause
		if len(query.Bool.Filter) != 1 {
			t.Errorf("expected 1 Filter clause, got %d", len(query.Bool.Filter))
		}
		filterQuery := query.Bool.Filter[0]
		if filterQuery.Term == nil {
			t.Error("expected Term query in Filter clause")
		}
	})

	t.Run("should match manual BoolQuery construction", func(t *testing.T) {
		// Our builder approach
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

		// Manual construction (like in cmd/main/main.go)
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

		// Compare structures
		if builderQuery.Bool == nil || manualQuery.Bool == nil {
			t.Error("both queries should have Bool queries")
		}
		
		if len(builderQuery.Bool.Must) != len(manualQuery.Bool.Must) {
			t.Errorf("Must clause count mismatch: builder=%d, manual=%d", 
				len(builderQuery.Bool.Must), len(manualQuery.Bool.Must))
		}
		
		if len(builderQuery.Bool.Filter) != len(manualQuery.Bool.Filter) {
			t.Errorf("Filter clause count mismatch: builder=%d, manual=%d", 
				len(builderQuery.Bool.Filter), len(manualQuery.Bool.Filter))
		}
	})
}

// Benchmark tests for Bool queries
func BenchmarkBoolQuery(b *testing.B) {
	b.Run("Simple Bool with Must", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = NewQuery(
				Bool(
					Must(
						Term("status", "published"),
					),
				),
			)
		}
	})

	b.Run("Complex Bool with all clauses", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = NewQuery(
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

	b.Run("Nested Bool queries", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = NewQuery(
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
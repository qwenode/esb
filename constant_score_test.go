package esb

import (
	"testing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestConstantScore(t *testing.T) {
	query := NewQuery(
		ConstantScore(
			Term("status", "published"),
		),
	)

	if query.ConstantScore == nil {
		t.Fatal("Expected ConstantScore query to be set")
	}

	if query.ConstantScore.Filter == nil {
		t.Fatal("Expected Filter to be set")
	}

	// 验证内部的 Term 查询
	if query.ConstantScore.Filter.Term == nil {
		t.Fatal("Expected Term query in filter")
	}

	termQuery, exists := query.ConstantScore.Filter.Term["status"]
	if !exists {
		t.Fatal("Expected 'status' field in Term query")
	}

	if termQuery.Value != "published" {
		t.Errorf("Expected term value to be 'published', got %v", termQuery.Value)
	}
}

func TestConstantScoreWithOptions(t *testing.T) {
	boost := float32(2.5)
	queryName := "test_constant_score"

	query := NewQuery(
		ConstantScoreWithOptions(
			Term("category", "tech"),
			func(opts *types.ConstantScoreQuery) {
				opts.Boost = &boost
				opts.QueryName_ = &queryName
			},
		),
	)

	if query.ConstantScore == nil {
		t.Fatal("Expected ConstantScore query to be set")
	}

	if query.ConstantScore.Boost == nil || *query.ConstantScore.Boost != boost {
		t.Errorf("Expected Boost to be %v, got %v", boost, query.ConstantScore.Boost)
	}

	if query.ConstantScore.QueryName_ == nil || *query.ConstantScore.QueryName_ != queryName {
		t.Errorf("Expected QueryName_ to be %v, got %v", queryName, query.ConstantScore.QueryName_)
	}

	// 验证过滤器查询
	if query.ConstantScore.Filter == nil {
		t.Fatal("Expected Filter to be set")
	}

	termQuery, exists := query.ConstantScore.Filter.Term["category"]
	if !exists {
		t.Fatal("Expected 'category' field in Term query")
	}

	if termQuery.Value != "tech" {
		t.Errorf("Expected term value to be 'tech', got %v", termQuery.Value)
	}
}

func TestConstantScoreWithNilFilter(t *testing.T) {
	query := NewQuery(ConstantScore(nil))

	if query.ConstantScore == nil {
		t.Fatal("Expected ConstantScore query to be set")
	}

	if query.ConstantScore.Filter == nil {
		t.Fatal("Expected Filter to be set even with nil filter")
	}
}

func TestConstantScoreWithComplexFilter(t *testing.T) {
	query := NewQuery(
		ConstantScore(
			Bool(
				Must(
					Term("status", "published"),
					Term("type", "article"),
				),
				Filter(
					Term("category", "tech"),
				),
			),
		),
	)

	if query.ConstantScore == nil {
		t.Fatal("Expected ConstantScore query to be set")
	}

	if query.ConstantScore.Filter == nil {
		t.Fatal("Expected Filter to be set")
	}

	// 验证内部的 Bool 查询
	if query.ConstantScore.Filter.Bool == nil {
		t.Fatal("Expected Bool query in filter")
	}

	boolQuery := query.ConstantScore.Filter.Bool

	// 验证 Must 子句
	if len(boolQuery.Must) != 2 {
		t.Fatalf("Expected 2 Must clauses, got %d", len(boolQuery.Must))
	}

	// 验证 Filter 子句
	if len(boolQuery.Filter) != 1 {
		t.Fatalf("Expected 1 Filter clause, got %d", len(boolQuery.Filter))
	}
}

func TestConstantScoreWithNilOptions(t *testing.T) {
	query := NewQuery(
		ConstantScoreWithOptions(
			Term("status", "active"),
			nil,
		),
	)

	if query.ConstantScore == nil {
		t.Fatal("Expected ConstantScore query to be set")
	}

	// 验证 nil 回调不会导致 panic
	if query.ConstantScore.Boost != nil {
		t.Errorf("Expected Boost to be nil, got %v", query.ConstantScore.Boost)
	}
}

// 基准测试
func BenchmarkConstantScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewQuery(
			ConstantScore(
				Term("status", "published"),
			),
		)
	}
}

func BenchmarkConstantScoreWithOptions(b *testing.B) {
	boost := float32(2.0)
	for i := 0; i < b.N; i++ {
		_ = NewQuery(
			ConstantScoreWithOptions(
				Term("status", "published"),
				func(opts *types.ConstantScoreQuery) {
					opts.Boost = &boost
				},
			),
		)
	}
}
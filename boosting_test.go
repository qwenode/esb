package esb

import (
	"testing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestBoosting(t *testing.T) {
	negativeBoost := 0.2

	query := NewQuery(
		Boosting(
			Term("title", "apple"),
			Term("category", "technology"),
			negativeBoost,
		),
	)

	if query.Boosting == nil {
		t.Fatal("Expected Boosting query to be set")
	}

	// 验证 positive 查询
	if query.Boosting.Positive == nil {
		t.Fatal("Expected Positive query to be set")
	}

	if query.Boosting.Positive.Term == nil {
		t.Fatal("Expected Positive query to be Term query")
	}

	titleTerm, exists := query.Boosting.Positive.Term["title"]
	if !exists {
		t.Fatal("Expected 'title' field in Positive Term query")
	}

	if titleTerm.Value != "apple" {
		t.Errorf("Expected positive term value to be 'apple', got %v", titleTerm.Value)
	}

	// 验证 negative 查询
	if query.Boosting.Negative == nil {
		t.Fatal("Expected Negative query to be set")
	}

	if query.Boosting.Negative.Term == nil {
		t.Fatal("Expected Negative query to be Term query")
	}

	categoryTerm, exists := query.Boosting.Negative.Term["category"]
	if !exists {
		t.Fatal("Expected 'category' field in Negative Term query")
	}

	if categoryTerm.Value != "technology" {
		t.Errorf("Expected negative term value to be 'technology', got %v", categoryTerm.Value)
	}

	// 验证 negative_boost
	expectedBoost := types.Float64(negativeBoost)
	if query.Boosting.NegativeBoost != expectedBoost {
		t.Errorf("Expected NegativeBoost to be %v, got %v", expectedBoost, query.Boosting.NegativeBoost)
	}
}

func TestBoostingWithOptions(t *testing.T) {
	negativeBoost := 0.3
	boost := float32(1.5)
	queryName := "test_boosting"

	query := NewQuery(
		BoostingWithOptions(
			Match("title", "apple fruit"),
			Match("content", "iPhone iPad"),
			negativeBoost,
			func(opts *types.BoostingQuery) {
				opts.Boost = &boost
				opts.QueryName_ = &queryName
			},
		),
	)

	if query.Boosting == nil {
		t.Fatal("Expected Boosting query to be set")
	}

	if query.Boosting.Boost == nil || *query.Boosting.Boost != boost {
		t.Errorf("Expected Boost to be %v, got %v", boost, query.Boosting.Boost)
	}

	if query.Boosting.QueryName_ == nil || *query.Boosting.QueryName_ != queryName {
		t.Errorf("Expected QueryName_ to be %v, got %v", queryName, query.Boosting.QueryName_)
	}

	// 验证 negative_boost
	expectedNegativeBoost := types.Float64(negativeBoost)
	if query.Boosting.NegativeBoost != expectedNegativeBoost {
		t.Errorf("Expected NegativeBoost to be %v, got %v", expectedNegativeBoost, query.Boosting.NegativeBoost)
	}

	// 验证 positive 查询是 Match 类型
	if query.Boosting.Positive == nil || query.Boosting.Positive.Match == nil {
		t.Fatal("Expected Positive query to be Match query")
	}

	// 验证 negative 查询是 Match 类型
	if query.Boosting.Negative == nil || query.Boosting.Negative.Match == nil {
		t.Fatal("Expected Negative query to be Match query")
	}
}

func TestBoostingWithNilQueries(t *testing.T) {
	query := NewQuery(
		Boosting(
			nil, // nil positive query
			nil, // nil negative query
			0.5,
		),
	)

	if query.Boosting == nil {
		t.Fatal("Expected Boosting query to be set")
	}

	// nil 查询应该创建空的 Query 对象
	if query.Boosting.Positive == nil {
		t.Fatal("Expected Positive query to be set (even if empty)")
	}

	if query.Boosting.Negative == nil {
		t.Fatal("Expected Negative query to be set (even if empty)")
	}

	expectedBoost := types.Float64(0.5)
	if query.Boosting.NegativeBoost != expectedBoost {
		t.Errorf("Expected NegativeBoost to be %v, got %v", expectedBoost, query.Boosting.NegativeBoost)
	}
}

func TestBoostingWithComplexQueries(t *testing.T) {
	query := NewQuery(
		Boosting(
			// positive: 复合查询
			Bool(
				Must(
					Match("title", "apple"),
					Term("status", "published"),
				),
			),
			// negative: 另一个复合查询
			Bool(
				Should(
					Term("category", "technology"),
					Term("brand", "apple"),
				),
			),
			0.1,
		),
	)

	if query.Boosting == nil {
		t.Fatal("Expected Boosting query to be set")
	}

	// 验证 positive 查询是 Bool 类型
	if query.Boosting.Positive == nil || query.Boosting.Positive.Bool == nil {
		t.Fatal("Expected Positive query to be Bool query")
	}

	// 验证 negative 查询是 Bool 类型
	if query.Boosting.Negative == nil || query.Boosting.Negative.Bool == nil {
		t.Fatal("Expected Negative query to be Bool query")
	}

	// 验证 positive Bool 查询的 Must 子句
	positiveBool := query.Boosting.Positive.Bool
	if len(positiveBool.Must) != 2 {
		t.Fatalf("Expected 2 Must clauses in positive query, got %d", len(positiveBool.Must))
	}

	// 验证 negative Bool 查询的 Should 子句
	negativeBool := query.Boosting.Negative.Bool
	if len(negativeBool.Should) != 2 {
		t.Fatalf("Expected 2 Should clauses in negative query, got %d", len(negativeBool.Should))
	}
}

func TestBoostingWithNilOptions(t *testing.T) {
	query := NewQuery(
		BoostingWithOptions(
			Term("title", "test"),
			Term("category", "spam"),
			0.2,
			nil,
		),
	)

	if query.Boosting == nil {
		t.Fatal("Expected Boosting query to be set")
	}

	// 验证 nil 回调不会导致 panic
	if query.Boosting.Boost != nil {
		t.Errorf("Expected Boost to be nil, got %v", query.Boosting.Boost)
	}

	expectedNegativeBoost := types.Float64(0.2)
	if query.Boosting.NegativeBoost != expectedNegativeBoost {
		t.Errorf("Expected NegativeBoost to be %v, got %v", expectedNegativeBoost, query.Boosting.NegativeBoost)
	}
}

// 基准测试
func BenchmarkBoosting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewQuery(
			Boosting(
				Term("title", "apple"),
				Term("category", "technology"),
				0.2,
			),
		)
	}
}

func BenchmarkBoostingWithOptions(b *testing.B) {
	boost := float32(1.2)
	
	for i := 0; i < b.N; i++ {
		_ = NewQuery(
			BoostingWithOptions(
				Match("title", "apple fruit"),
				Match("content", "iPhone iPad"),
				0.3,
				func(opts *types.BoostingQuery) {
					opts.Boost = &boost
				},
			),
		)
	}
}
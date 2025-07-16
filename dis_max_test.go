package esb

import (
	"testing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestDisMax(t *testing.T) {
	query := NewQuery(
		DisMax(
			Term("title", "quick"),
			Term("body", "quick"),
		),
	)

	if query.DisMax == nil {
		t.Fatal("Expected DisMax query to be set")
	}

	if len(query.DisMax.Queries) != 2 {
		t.Fatalf("Expected 2 queries, got %d", len(query.DisMax.Queries))
	}

	// 验证第一个查询
	firstQuery := query.DisMax.Queries[0]
	if firstQuery.Term == nil {
		t.Fatal("Expected first query to be Term query")
	}

	titleTerm, exists := firstQuery.Term["title"]
	if !exists {
		t.Fatal("Expected 'title' field in first Term query")
	}

	if titleTerm.Value != "quick" {
		t.Errorf("Expected first term value to be 'quick', got %v", titleTerm.Value)
	}

	// 验证第二个查询
	secondQuery := query.DisMax.Queries[1]
	if secondQuery.Term == nil {
		t.Fatal("Expected second query to be Term query")
	}

	bodyTerm, exists := secondQuery.Term["body"]
	if !exists {
		t.Fatal("Expected 'body' field in second Term query")
	}

	if bodyTerm.Value != "quick" {
		t.Errorf("Expected second term value to be 'quick', got %v", bodyTerm.Value)
	}
}

func TestDisMaxWithOptions(t *testing.T) {
	tieBreaker := types.Float64(0.3)
	boost := float32(1.5)
	queryName := "test_dis_max"

	query := NewQuery(
		DisMaxWithOptions(
			[]QueryOption{
				Match("title", "quick brown fox"),
				Match("body", "quick brown fox"),
			},
			func(opts *types.DisMaxQuery) {
				opts.TieBreaker = &tieBreaker
				opts.Boost = &boost
				opts.QueryName_ = &queryName
			},
		),
	)

	if query.DisMax == nil {
		t.Fatal("Expected DisMax query to be set")
	}

	if query.DisMax.TieBreaker == nil || *query.DisMax.TieBreaker != tieBreaker {
		t.Errorf("Expected TieBreaker to be %v, got %v", tieBreaker, query.DisMax.TieBreaker)
	}

	if query.DisMax.Boost == nil || *query.DisMax.Boost != boost {
		t.Errorf("Expected Boost to be %v, got %v", boost, query.DisMax.Boost)
	}

	if query.DisMax.QueryName_ == nil || *query.DisMax.QueryName_ != queryName {
		t.Errorf("Expected QueryName_ to be %v, got %v", queryName, query.DisMax.QueryName_)
	}

	if len(query.DisMax.Queries) != 2 {
		t.Fatalf("Expected 2 queries, got %d", len(query.DisMax.Queries))
	}

	// 验证查询类型
	for i, subQuery := range query.DisMax.Queries {
		if subQuery.Match == nil {
			t.Fatalf("Expected query %d to be Match query", i)
		}
	}
}

func TestDisMaxWithEmptyQueries(t *testing.T) {
	query := NewQuery(DisMax())

	if query.DisMax == nil {
		t.Fatal("Expected DisMax query to be set")
	}

	if len(query.DisMax.Queries) != 0 {
		t.Errorf("Expected 0 queries, got %d", len(query.DisMax.Queries))
	}
}

func TestDisMaxWithNilQueries(t *testing.T) {
	query := NewQuery(
		DisMax(
			Term("title", "test"),
			nil, // nil query should be skipped
			Term("body", "test"),
		),
	)

	if query.DisMax == nil {
		t.Fatal("Expected DisMax query to be set")
	}

	// nil 查询应该被跳过，所以只有 2 个查询
	if len(query.DisMax.Queries) != 2 {
		t.Fatalf("Expected 2 queries (nil should be skipped), got %d", len(query.DisMax.Queries))
	}
}

func TestDisMaxWithMixedQueryTypes(t *testing.T) {
	query := NewQuery(
		DisMax(
			Term("status", "published"),
			Match("title", "elasticsearch guide"),
			Exists("author"),
		),
	)

	if query.DisMax == nil {
		t.Fatal("Expected DisMax query to be set")
	}

	if len(query.DisMax.Queries) != 3 {
		t.Fatalf("Expected 3 queries, got %d", len(query.DisMax.Queries))
	}

	// 验证第一个查询是 Term
	if query.DisMax.Queries[0].Term == nil {
		t.Error("Expected first query to be Term query")
	}

	// 验证第二个查询是 Match
	if query.DisMax.Queries[1].Match == nil {
		t.Error("Expected second query to be Match query")
	}

	// 验证第三个查询是 Exists
	if query.DisMax.Queries[2].Exists == nil {
		t.Error("Expected third query to be Exists query")
	}
}

func TestDisMaxWithNilOptions(t *testing.T) {
	query := NewQuery(
		DisMaxWithOptions(
			[]QueryOption{
				Term("title", "test"),
			},
			nil,
		),
	)

	if query.DisMax == nil {
		t.Fatal("Expected DisMax query to be set")
	}

	// 验证 nil 回调不会导致 panic
	if query.DisMax.Boost != nil {
		t.Errorf("Expected Boost to be nil, got %v", query.DisMax.Boost)
	}

	if len(query.DisMax.Queries) != 1 {
		t.Fatalf("Expected 1 query, got %d", len(query.DisMax.Queries))
	}
}

// 基准测试
func BenchmarkDisMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewQuery(
			DisMax(
				Term("title", "quick"),
				Term("body", "quick"),
			),
		)
	}
}

func BenchmarkDisMaxWithOptions(b *testing.B) {
	tieBreaker := types.Float64(0.3)
	boost := float32(1.2)
	
	for i := 0; i < b.N; i++ {
		_ = NewQuery(
			DisMaxWithOptions(
				[]QueryOption{
					Match("title", "quick brown fox"),
					Match("body", "quick brown fox"),
				},
				func(opts *types.DisMaxQuery) {
					opts.TieBreaker = &tieBreaker
					opts.Boost = &boost
				},
			),
		)
	}
}
package esb

import (
	"testing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestMatchNone(t *testing.T) {
	query := NewQuery(MatchNone())

	if query.MatchNone == nil {
		t.Fatal("Expected MatchNone query to be set")
	}

	// 验证基本结构
	expected := &types.MatchNoneQuery{}
	if query.MatchNone.Boost != expected.Boost {
		t.Errorf("Expected Boost to be %v, got %v", expected.Boost, query.MatchNone.Boost)
	}
}

func TestMatchNoneWithOptions(t *testing.T) {
	boost := float32(0.5)
	queryName := "test_match_none"

	query := NewQuery(
		MatchNoneWithOptions(func(opts *types.MatchNoneQuery) {
			opts.Boost = &boost
			opts.QueryName_ = &queryName
		}),
	)

	if query.MatchNone == nil {
		t.Fatal("Expected MatchNone query to be set")
	}

	if query.MatchNone.Boost == nil || *query.MatchNone.Boost != boost {
		t.Errorf("Expected Boost to be %v, got %v", boost, query.MatchNone.Boost)
	}

	if query.MatchNone.QueryName_ == nil || *query.MatchNone.QueryName_ != queryName {
		t.Errorf("Expected QueryName_ to be %v, got %v", queryName, query.MatchNone.QueryName_)
	}
}

func TestMatchNoneWithNilOptions(t *testing.T) {
	query := NewQuery(MatchNoneWithOptions(nil))

	if query.MatchNone == nil {
		t.Fatal("Expected MatchNone query to be set")
	}

	// 验证 nil 回调不会导致 panic
	if query.MatchNone.Boost != nil {
		t.Errorf("Expected Boost to be nil, got %v", query.MatchNone.Boost)
	}
}

// 测试在布尔查询中使用 MatchNone
func TestMatchNoneInBoolQuery(t *testing.T) {
	query := NewQuery(
		Bool(
			MustNot(
				MatchNone(), // 在 must_not 中使用 match_none 实际上会匹配所有文档
			),
		),
	)

	if query.Bool == nil {
		t.Fatal("Expected Bool query to be set")
	}

	if len(query.Bool.MustNot) != 1 {
		t.Fatalf("Expected 1 MustNot clause, got %d", len(query.Bool.MustNot))
	}

	if query.Bool.MustNot[0].MatchNone == nil {
		t.Error("Expected MustNot clause to contain MatchNone query")
	}
}

// 基准测试
func BenchmarkMatchNone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewQuery(MatchNone())
	}
}

func BenchmarkMatchNoneWithOptions(b *testing.B) {
	boost := float32(0.0)
	for i := 0; i < b.N; i++ {
		_ = NewQuery(
			MatchNoneWithOptions(func(opts *types.MatchNoneQuery) {
				opts.Boost = &boost
			}),
		)
	}
}
package esb

import (
	"testing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestMatchAll(t *testing.T) {
	query := NewQuery(MatchAll())

	if query.MatchAll == nil {
		t.Fatal("Expected MatchAll query to be set")
	}

	// 验证基本结构
	expected := &types.MatchAllQuery{}
	if query.MatchAll.Boost != expected.Boost {
		t.Errorf("Expected Boost to be %v, got %v", expected.Boost, query.MatchAll.Boost)
	}
}

func TestMatchAllWithOptions(t *testing.T) {
	boost := float32(2.5)
	queryName := "test_match_all"

	query := NewQuery(
		MatchAllWithOptions(func(opts *types.MatchAllQuery) {
			opts.Boost = &boost
			opts.QueryName_ = &queryName
		}),
	)

	if query.MatchAll == nil {
		t.Fatal("Expected MatchAll query to be set")
	}

	if query.MatchAll.Boost == nil || *query.MatchAll.Boost != boost {
		t.Errorf("Expected Boost to be %v, got %v", boost, query.MatchAll.Boost)
	}

	if query.MatchAll.QueryName_ == nil || *query.MatchAll.QueryName_ != queryName {
		t.Errorf("Expected QueryName_ to be %v, got %v", queryName, query.MatchAll.QueryName_)
	}
}

func TestMatchAllWithNilOptions(t *testing.T) {
	query := NewQuery(MatchAllWithOptions(nil))

	if query.MatchAll == nil {
		t.Fatal("Expected MatchAll query to be set")
	}

	// 验证 nil 回调不会导致 panic
	if query.MatchAll.Boost != nil {
		t.Errorf("Expected Boost to be nil, got %v", query.MatchAll.Boost)
	}
}

// 基准测试
func BenchmarkMatchAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewQuery(MatchAll())
	}
}

func BenchmarkMatchAllWithOptions(b *testing.B) {
	boost := float32(2.0)
	for i := 0; i < b.N; i++ {
		_ = NewQuery(
			MatchAllWithOptions(func(opts *types.MatchAllQuery) {
				opts.Boost = &boost
			}),
		)
	}
}
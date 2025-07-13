package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestNewQuery(t *testing.T) {
	t.Run("should create empty query when no options provided", func(t *testing.T) {
		query := NewQuery()
		if query == nil {
			t.Error("expected non-nil query")
		}
	})

	t.Run("should create query with valid option", func(t *testing.T) {
		query := NewQuery(Term("status", "published"))
		if query == nil {
			t.Error("expected non-nil query")
		}
	})

	t.Run("should create query with custom option", func(t *testing.T) {
		validOption := func(q *types.Query) {
			// Simple option that sets a field (we'll implement actual options later)
		}

		query := NewQuery(validOption)
		if query == nil {
			t.Error("expected non-nil query")
		}
	})

	t.Run("should apply multiple options in order", func(t *testing.T) {
		callOrder := []int{}
		
		option1 := func(q *types.Query) {
			callOrder = append(callOrder, 1)
		}
		
		option2 := func(q *types.Query) {
			callOrder = append(callOrder, 2)
		}

		query := NewQuery(option1, option2)
		if query == nil {
			t.Error("expected non-nil query")
		}
		
		if len(callOrder) != 2 || callOrder[0] != 1 || callOrder[1] != 2 {
			t.Errorf("expected call order [1, 2], got %v", callOrder)
		}
	})
}



// Benchmark tests to ensure performance is acceptable
func BenchmarkNewQuery(b *testing.B) {
	option := func(q *types.Query) {}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewQuery(option)
	}
}

func BenchmarkNewQueryMultipleOptions(b *testing.B) {
	option1 := func(q *types.Query) {}
	option2 := func(q *types.Query) {}
	option3 := func(q *types.Query) {}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewQuery(option1, option2, option3)
	}
} 
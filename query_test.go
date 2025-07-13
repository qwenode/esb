package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestNewQuery(t *testing.T) {
	t.Run("should return error when no options provided", func(t *testing.T) {
		query, err := NewQuery()
		if err == nil {
			t.Error("expected error when no options provided")
		}
		if err != ErrNoOptions {
			t.Errorf("expected ErrNoOptions, got %v", err)
		}
		if query != nil {
			t.Error("expected nil query when error occurs")
		}
	})

	t.Run("should create empty query with nil option", func(t *testing.T) {
		query, err := NewQuery(nil)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if query == nil {
			t.Error("expected non-nil query")
		}
	})

	t.Run("should create query with valid option", func(t *testing.T) {
		validOption := func(q *types.Query) error {
			// Simple option that sets a field (we'll implement actual options later)
			return nil
		}

		query, err := NewQuery(validOption)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if query == nil {
			t.Error("expected non-nil query")
		}
	})

	t.Run("should return error when option fails", func(t *testing.T) {
		failingOption := func(q *types.Query) error {
			return ErrInvalidQuery
		}

		query, err := NewQuery(failingOption)
		if err == nil {
			t.Error("expected error when option fails")
		}
		if err != ErrInvalidQuery {
			t.Errorf("expected ErrInvalidQuery, got %v", err)
		}
		if query != nil {
			t.Error("expected nil query when error occurs")
		}
	})

	t.Run("should apply multiple options in order", func(t *testing.T) {
		callOrder := []int{}
		
		option1 := func(q *types.Query) error {
			callOrder = append(callOrder, 1)
			return nil
		}
		
		option2 := func(q *types.Query) error {
			callOrder = append(callOrder, 2)
			return nil
		}

		query, err := NewQuery(option1, option2)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if query == nil {
			t.Error("expected non-nil query")
		}
		
		if len(callOrder) != 2 || callOrder[0] != 1 || callOrder[1] != 2 {
			t.Errorf("expected call order [1, 2], got %v", callOrder)
		}
	})

	t.Run("should stop on first error", func(t *testing.T) {
		callOrder := []int{}
		
		option1 := func(q *types.Query) error {
			callOrder = append(callOrder, 1)
			return nil
		}
		
		failingOption := func(q *types.Query) error {
			callOrder = append(callOrder, 2)
			return ErrInvalidQuery
		}
		
		option3 := func(q *types.Query) error {
			callOrder = append(callOrder, 3)
			return nil
		}

		query, err := NewQuery(option1, failingOption, option3)
		if err == nil {
			t.Error("expected error")
		}
		if err != ErrInvalidQuery {
			t.Errorf("expected ErrInvalidQuery, got %v", err)
		}
		if query != nil {
			t.Error("expected nil query when error occurs")
		}
		
		// Should only call first two options
		if len(callOrder) != 2 || callOrder[0] != 1 || callOrder[1] != 2 {
			t.Errorf("expected call order [1, 2], got %v", callOrder)
		}
	})
}



// Benchmark tests to ensure performance is acceptable
func BenchmarkNewQuery(b *testing.B) {
	option := func(q *types.Query) error {
		return nil
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NewQuery(option)
	}
}

func BenchmarkNewQueryMultipleOptions(b *testing.B) {
	option1 := func(q *types.Query) error { return nil }
	option2 := func(q *types.Query) error { return nil }
	option3 := func(q *types.Query) error { return nil }
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NewQuery(option1, option2, option3)
	}
} 
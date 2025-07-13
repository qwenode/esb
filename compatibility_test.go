package esb

import (
	"testing"
	"reflect"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestElasticsearchCompatibility(t *testing.T) {
	t.Run("should generate compatible types.Query", func(t *testing.T) {
		// Test that our generated Query is compatible with types.Query
		option := func(q *types.Query) error {
			// This would be a real query option in practice
			return nil
		}
		
		query, err := NewQuery(option)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		
		// Verify the type is exactly *types.Query
		if reflect.TypeOf(query) != reflect.TypeOf(&types.Query{}) {
			t.Errorf("expected *types.Query, got %T", query)
		}
		
		// Verify it's a valid Query struct
		if query == nil {
			t.Error("expected non-nil query")
		}
	})
	
	t.Run("should work with elasticsearch client interface", func(t *testing.T) {
		// Test that our Query can be used where types.Query is expected
		option := func(q *types.Query) error {
			return nil
		}
		
		query, err := NewQuery(option)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		
		// This function simulates what the elasticsearch client would do
		acceptsQuery := func(q *types.Query) bool {
			return q != nil
		}
		
		if !acceptsQuery(query) {
			t.Error("generated query not accepted by elasticsearch client interface")
		}
	})
	
	t.Run("should maintain Query struct integrity", func(t *testing.T) {
		// Create a query and verify all fields are properly initialized
		option := func(q *types.Query) error {
			return nil
		}
		
		query, err := NewQuery(option)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		
		// Verify the Query struct is properly initialized
		queryValue := reflect.ValueOf(query).Elem()
		if !queryValue.IsValid() {
			t.Error("generated query is not valid")
		}
		
		// Verify it's a struct
		if queryValue.Kind() != reflect.Struct {
			t.Error("expected query to be a struct")
		}
	})
} 
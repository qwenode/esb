package esb

import (
	"testing"
)

// TestHelperFunctions tests the helper functions to improve coverage
func TestHelperFunctions(t *testing.T) {
	// Test intPtr helper
	val := intPtr(42)
	if val == nil {
		t.Error("intPtr should not return nil")
	}
	if *val != 42 {
		t.Errorf("intPtr(*val) = %d, want 42", *val)
	}
	
	// Test stringPtr helper
	str := stringPtr("test")
	if str == nil {
		t.Error("stringPtr should not return nil")
	}
	if *str != "test" {
		t.Errorf("stringPtr(*str) = %s, want test", *str)
	}
	
	// Test float32Ptr helper
	f32 := float32Ptr(3.14)
	if f32 == nil {
		t.Error("float32Ptr should not return nil")
	}
	if *f32 != 3.14 {
		t.Errorf("float32Ptr(*f32) = %f, want 3.14", *f32)
	}
	
	// Test boolPtr helper
	b := boolPtr(true)
	if b == nil {
		t.Error("boolPtr should not return nil")
	}
	if *b != true {
		t.Errorf("boolPtr(*b) = %t, want true", *b)
	}
}

// TestMatchWithOptionsErrorHandling tests error handling in MatchWithOptions
func TestMatchWithOptionsErrorHandling(t *testing.T) {
	// Test with empty options
	query, err := NewQuery(MatchWithOptions("title", "test", MatchOptions{}))
	if err != nil {
		t.Errorf("MatchWithOptions with empty options should not error: %v", err)
	}
	if query.Match == nil {
		t.Error("Match query should not be nil")
	}
	
	// Test with empty field
	_, err = NewQuery(MatchWithOptions("", "test", MatchOptions{
		Boost: float32Ptr(1.5),
	}))
	if err == nil {
		t.Error("Expected error for empty field")
	}
	
	// Test with empty value
	_, err = NewQuery(MatchWithOptions("title", "", MatchOptions{
		Boost: float32Ptr(1.5),
	}))
	if err == nil {
		t.Error("Expected error for empty value")
	}
}

// TestMatchPhraseWithOptionsErrorHandling tests error handling in MatchPhraseWithOptions
func TestMatchPhraseWithOptionsErrorHandling(t *testing.T) {
	// Test with empty options
	query, err := NewQuery(MatchPhraseWithOptions("content", "test phrase", MatchPhraseOptions{}))
	if err != nil {
		t.Errorf("MatchPhraseWithOptions with empty options should not error: %v", err)
	}
	if query.MatchPhrase == nil {
		t.Error("MatchPhrase query should not be nil")
	}
	
	// Test with empty field
	_, err = NewQuery(MatchPhraseWithOptions("", "test phrase", MatchPhraseOptions{
		Slop: intPtr(2),
	}))
	if err == nil {
		t.Error("Expected error for empty field")
	}
	
	// Test with empty phrase
	_, err = NewQuery(MatchPhraseWithOptions("content", "", MatchPhraseOptions{
		Slop: intPtr(2),
	}))
	if err == nil {
		t.Error("Expected error for empty phrase")
	}
}

// TestTermsQueryErrorHandling tests error handling in Terms query
func TestTermsQueryErrorHandling(t *testing.T) {
	// Test with empty field
	_, err := NewQuery(Terms("", "value1", "value2"))
	if err == nil {
		t.Error("Expected error for empty field")
	}
	
	// Test with empty values
	_, err = NewQuery(Terms("field"))
	if err == nil {
		t.Error("Expected error for no values")
	}
	
	// Test with valid values
	query, err := NewQuery(Terms("category", "tech", "science"))
	if err != nil {
		t.Errorf("Terms query with valid values should not error: %v", err)
	}
	if query.Terms == nil {
		t.Error("Terms query should not be nil")
	}
}

// TestTermQueryErrorHandling tests error handling in Term query
func TestTermQueryErrorHandling(t *testing.T) {
	// Test with empty field
	_, err := NewQuery(Term("", "value"))
	if err == nil {
		t.Error("Expected error for empty field")
	}
	
	// Test with empty value
	_, err = NewQuery(Term("field", ""))
	if err == nil {
		t.Error("Expected error for empty value")
	}
	
	// Test with valid field and value
	query, err := NewQuery(Term("status", "published"))
	if err != nil {
		t.Errorf("Term query with valid field and value should not error: %v", err)
	}
	if query.Term == nil {
		t.Error("Term query should not be nil")
	}
}

// TestBoolQueryErrorHandling tests error handling in Bool query
func TestBoolQueryErrorHandling(t *testing.T) {
	// Test with no options
	_, err := NewQuery(Bool())
	if err == nil {
		t.Error("Expected error for Bool query with no options")
	}
	
	// Test with empty Must clause
	query, err := NewQuery(Bool(Must()))
	if err != nil {
		t.Errorf("Bool query with empty Must should not error: %v", err)
	}
	if query.Bool == nil {
		t.Error("Bool query should not be nil")
	}
	
	// Test with empty Should clause
	query, err = NewQuery(Bool(Should()))
	if err != nil {
		t.Errorf("Bool query with empty Should should not error: %v", err)
	}
	if query.Bool == nil {
		t.Error("Bool query should not be nil")
	}
	
	// Test with empty Filter clause
	query, err = NewQuery(Bool(Filter()))
	if err != nil {
		t.Errorf("Bool query with empty Filter should not error: %v", err)
	}
	if query.Bool == nil {
		t.Error("Bool query should not be nil")
	}
	
	// Test with empty MustNot clause
	query, err = NewQuery(Bool(MustNot()))
	if err != nil {
		t.Errorf("Bool query with empty MustNot should not error: %v", err)
	}
	if query.Bool == nil {
		t.Error("Bool query should not be nil")
	}
}

// TestMarshalValueErrorHandling tests error handling in marshalValue
func TestMarshalValueErrorHandling(t *testing.T) {
	// Test with various types
	tests := []struct {
		name  string
		value interface{}
	}{
		{"string", "test"},
		{"int", 42},
		{"int64", int64(42)},
		{"float32", float32(3.14)},
		{"float64", 3.14},
		{"bool", true},
		{"nil", nil},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := marshalValue(tt.value)
			if result == nil {
				t.Error("marshalValue should not return nil")
			}
			if len(result) == 0 {
				t.Error("marshalValue should not return empty slice")
			}
		})
	}
}

// TestRangeBuilderErrorHandling tests error handling in RangeBuilder
func TestRangeBuilderErrorHandling(t *testing.T) {
	// Test with empty field
	_, err := NewQuery(Range("").Gte(10).Build())
	if err == nil {
		t.Error("Expected error for empty field")
	}
	
	// Test with whitespace field
	_, err = NewQuery(Range("   ").Gte(10).Build())
	if err == nil {
		t.Error("Expected error for whitespace field")
	}
	
	// Test chaining methods
	rb := Range("test")
	rb = rb.Gte(10)
	rb = rb.Gt(5)
	rb = rb.Lte(100)
	rb = rb.Lt(95)
	rb = rb.From(10)
	rb = rb.To(90)
	rb = rb.Boost(1.5)
	rb = rb.Format("yyyy-MM-dd")
	rb = rb.TimeZone("UTC")
	
	query, err := NewQuery(rb.Build())
	if err != nil {
		t.Errorf("Chained RangeBuilder should not error: %v", err)
	}
	if query.Range == nil {
		t.Error("Range query should not be nil")
	}
}

// TestComplexErrorPropagation tests error propagation in complex queries
func TestComplexErrorPropagation(t *testing.T) {
	// Test error in deeply nested Bool query
	_, err := NewQuery(
		Bool(
			Must(
				Bool(
					Should(
						Bool(
							Must(
								Term("", "value"), // Error here
							),
						),
					),
				),
			),
		),
	)
	if err == nil {
		t.Error("Expected error to propagate from deeply nested query")
	}
	
	// Test error in Range query within Bool
	_, err = NewQuery(
		Bool(
			Must(
				Match("title", "test"),
				Range("").Gte(10).Build(), // Error here
			),
		),
	)
	if err == nil {
		t.Error("Expected error to propagate from Range query in Bool")
	}
	
	// Test error in MatchWithOptions within Bool
	_, err = NewQuery(
		Bool(
			Should(
				MatchWithOptions("", "test", MatchOptions{ // Error here
					Boost: float32Ptr(1.5),
				}),
			),
		),
	)
	if err == nil {
		t.Error("Expected error to propagate from MatchWithOptions in Bool")
	}
}

// TestEdgeCases tests edge cases and boundary conditions
func TestEdgeCases(t *testing.T) {
	// Test with only nil QueryOptions (should return ErrNoOptions)
	_, err := NewQuery()
	if err == nil {
		t.Error("Expected error for no QueryOptions")
	}
	
	// Test with mixed nil and valid options
	query, err := NewQuery(
		nil,
		Term("status", "published"),
		nil,
	)
	if err != nil {
		t.Errorf("Mixed nil and valid options should not error: %v", err)
	}
	if query.Term == nil {
		t.Error("Term query should not be nil")
	}
	
	// Test Range with zero values
	query, err = NewQuery(Range("count").Gte(0).Lt(0).Build())
	if err != nil {
		t.Errorf("Range with zero values should not error: %v", err)
	}
	if query.Range == nil {
		t.Error("Range query should not be nil")
	}
	
	// Test Terms with single value
	query, err = NewQuery(Terms("category", "single"))
	if err != nil {
		t.Errorf("Terms with single value should not error: %v", err)
	}
	if query.Terms == nil {
		t.Error("Terms query should not be nil")
	}
} 
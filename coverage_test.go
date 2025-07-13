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
	
	// Test with empty field - should work now
	_, err = NewQuery(MatchWithOptions("", "test", MatchOptions{
		Boost: float32Ptr(1.5),
	}))
	if err != nil {
		t.Errorf("Unexpected error for empty field: %v", err)
	}
	
	// Test with empty value - should work now
	_, err = NewQuery(MatchWithOptions("title", "", MatchOptions{
		Boost: float32Ptr(1.5),
	}))
	if err != nil {
		t.Errorf("Unexpected error for empty value: %v", err)
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
	
	// Test with empty field - should work now
	_, err = NewQuery(MatchPhraseWithOptions("", "test phrase", MatchPhraseOptions{
		Slop: intPtr(2),
	}))
	if err != nil {
		t.Errorf("Unexpected error for empty field: %v", err)
	}
	
	// Test with empty phrase - should work now
	_, err = NewQuery(MatchPhraseWithOptions("content", "", MatchPhraseOptions{
		Slop: intPtr(2),
	}))
	if err != nil {
		t.Errorf("Unexpected error for empty phrase: %v", err)
	}
}

// TestTermsQueryErrorHandling tests error handling in Terms query
func TestTermsQueryErrorHandling(t *testing.T) {
	// Test with empty field - should work now
	_, err := NewQuery(Terms("", "value1", "value2"))
	if err != nil {
		t.Errorf("Unexpected error for empty field: %v", err)
	}
	
	// Test with empty values - should work now
	_, err = NewQuery(Terms("field"))
	if err != nil {
		t.Errorf("Unexpected error for no values: %v", err)
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
	// Test with empty field - should work now
	_, err := NewQuery(Term("", "value"))
	if err != nil {
		t.Errorf("Unexpected error for empty field: %v", err)
	}
	
	// Test with empty value - should work now
	_, err = NewQuery(Term("field", ""))
	if err != nil {
		t.Errorf("Unexpected error for empty value: %v", err)
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
// TestMarshalValueErrorHandling has been removed as marshalValue is no longer used
// in the type-safe range query implementation

// TestRangeBuilderErrorHandling tests error handling in type-safe range builders
func TestRangeBuilderErrorHandling(t *testing.T) {
	// Test NumberRange with empty field - should work now
	_, err := NewQuery(NumberRange("").Gte(10.0).Build())
	if err != nil {
		t.Errorf("Unexpected error for empty field: %v", err)
	}
	
	// Test DateRange with whitespace field - should work now
	_, err = NewQuery(DateRange("   ").Gte("2023-01-01").Build())
	if err != nil {
		t.Errorf("Unexpected error for whitespace field: %v", err)
	}
	
	// Test chaining methods on NumberRange
	nrb := NumberRange("test")
	nrb = nrb.Gte(10.0)
	nrb = nrb.Gt(5.0)
	nrb = nrb.Lte(100.0)
	nrb = nrb.Lt(95.0)
	nrb = nrb.From(10.0)
	nrb = nrb.To(90.0)
	nrb = nrb.Boost(1.5)
	
	query, err := NewQuery(nrb.Build())
	if err != nil {
		t.Errorf("Chained NumberRangeBuilder should not error: %v", err)
	}
	if query.Range == nil {
		t.Error("Range query should not be nil")
	}
	
	// Test chaining methods on DateRange
	drb := DateRange("timestamp")
	drb = drb.Gte("2023-01-01")
	drb = drb.Lte("2023-12-31")
	drb = drb.Format("yyyy-MM-dd")
	drb = drb.TimeZone("UTC")
	drb = drb.Boost(2.0)
	
	query, err = NewQuery(drb.Build())
	if err != nil {
		t.Errorf("Chained DateRangeBuilder should not error: %v", err)
	}
	if query.Range == nil {
		t.Error("Range query should not be nil")
	}
}

// TestComplexErrorPropagation tests error propagation in complex queries
func TestComplexErrorPropagation(t *testing.T) {
	// Test deeply nested Bool query - should work now
	_, err := NewQuery(
		Bool(
			Must(
				Bool(
					Should(
						Bool(
							Must(
								Term("", "value"), // No error now
							),
						),
					),
				),
			),
		),
	)
	if err != nil {
		t.Errorf("Unexpected error in deeply nested query: %v", err)
	}
	
	// Test Range query within Bool - should work now
	_, err = NewQuery(
		Bool(
			Must(
				Match("title", "test"),
				NumberRange("").Gte(10.0).Build(), // No error now
			),
		),
	)
	if err != nil {
		t.Errorf("Unexpected error in Range query in Bool: %v", err)
	}
	
	// Test MatchWithOptions within Bool - should work now
	_, err = NewQuery(
		Bool(
			Should(
				MatchWithOptions("", "test", MatchOptions{ // No error now
					Boost: float32Ptr(1.5),
				}),
			),
		),
	)
	if err != nil {
		t.Errorf("Unexpected error in MatchWithOptions in Bool: %v", err)
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
	query, err = NewQuery(NumberRange("count").Gte(0.0).Lt(0.0).Build())
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
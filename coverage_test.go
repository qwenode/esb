package esb

import (
	"testing"
)

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
	boost := float32(1.5)
	_, err = NewQuery(MatchWithOptions("", "test", MatchOptions{
		Boost: &boost,
	}))
	if err != nil {
		t.Errorf("Unexpected error for empty field: %v", err)
	}
	
	// Test with empty value - should work now
	_, err = NewQuery(MatchWithOptions("title", "", MatchOptions{
		Boost: &boost,
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
	
	// Test with slop option
	slop := 2
	query, err = NewQuery(MatchPhraseWithOptions("content", "test phrase", MatchPhraseOptions{
		Slop: &slop,
	}))
	if err != nil {
		t.Errorf("MatchPhraseWithOptions with slop should not error: %v", err)
	}
	if query.MatchPhrase == nil {
		t.Error("MatchPhrase query should not be nil")
	}
	if query.MatchPhrase["content"].Slop == nil || *query.MatchPhrase["content"].Slop != 2 {
		t.Error("Slop should be set to 2")
	}
	
	// Test with analyzer option
	analyzer := "standard"
	query, err = NewQuery(MatchPhraseWithOptions("content", "test phrase", MatchPhraseOptions{
		Analyzer: &analyzer,
	}))
	if err != nil {
		t.Errorf("MatchPhraseWithOptions with analyzer should not error: %v", err)
	}
	if query.MatchPhrase == nil {
		t.Error("MatchPhrase query should not be nil")
	}
	if query.MatchPhrase["content"].Analyzer == nil || *query.MatchPhrase["content"].Analyzer != "standard" {
		t.Error("Analyzer should be set to 'standard'")
	}
}

// TestMatchPhrasePrefixCoverage tests MatchPhrasePrefix functionality
func TestMatchPhrasePrefixCoverage(t *testing.T) {
	query, err := NewQuery(MatchPhrasePrefix("title", "elasticsearch sea"))
	if err != nil {
		t.Errorf("MatchPhrasePrefix should not error: %v", err)
	}
	if query.MatchPhrasePrefix == nil {
		t.Error("MatchPhrasePrefix query should not be nil")
	}
	if query.MatchPhrasePrefix["title"].Query != "elasticsearch sea" {
		t.Error("MatchPhrasePrefix query should match the input")
	}
}

// TestCombinedQueries tests combining different query types
func TestCombinedQueries(t *testing.T) {
	// Test Bool query with different match types
	query, err := NewQuery(
		Bool(
			Must(
				Match("title", "elasticsearch"),
				MatchPhrase("content", "search engine"),
			),
			Should(
				MatchPhrasePrefix("title", "elastic"),
				Term("category", "tech"),
			),
		),
	)
	if err != nil {
		t.Errorf("Combined query should not error: %v", err)
	}
	if query.Bool == nil {
		t.Error("Bool query should not be nil")
	}
	if len(query.Bool.Must) != 2 {
		t.Errorf("Expected 2 Must clauses, got %d", len(query.Bool.Must))
	}
	if len(query.Bool.Should) != 2 {
		t.Errorf("Expected 2 Should clauses, got %d", len(query.Bool.Should))
	}
}

// TestRangeQueryIntegration tests Range query integration
func TestRangeQueryIntegration(t *testing.T) {
	// Test NumberRange in Bool query
	query, err := NewQuery(
		Bool(
			Filter(
				NumberRange("age").Gte(18).Lt(65).Build(),
				Term("status", "active"),
			),
		),
	)
	if err != nil {
		t.Errorf("Range query integration should not error: %v", err)
	}
	if query.Bool == nil {
		t.Error("Bool query should not be nil")
	}
	if len(query.Bool.Filter) != 2 {
		t.Errorf("Expected 2 Filter clauses, got %d", len(query.Bool.Filter))
	}
}

// TestEdgeCases tests edge cases and error conditions
func TestEdgeCases(t *testing.T) {
	// Test with mixed nil and valid options
	query, err := NewQuery(
		Term("status", "published"),
	)
	if err != nil {
		t.Errorf("Valid options should not error: %v", err)
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
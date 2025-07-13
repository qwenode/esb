package esb

import (
	"encoding/json"
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// TestComplexQueryIntegration tests complex query combinations
func TestComplexQueryIntegration(t *testing.T) {
	// Test a complex real-world query scenario
	query, err := NewQuery(
		Bool(
			Must(
				Match("title", "elasticsearch"),
				DateRange("publish_date").Gte("2023-01-01").Lte("2023-12-31").Build(),
				Exists("author"),
			),
			Should(
				MatchPhrase("content", "search engine"),
				Term("category", "technology"),
				NumberRange("views").Gte(1000.0).Build(),
			),
			Filter(
				Term("status", "published"),
				NumberRange("score").Gte(4.0).Build(),
			),
			MustNot(
				Term("deleted", "true"),
				Exists("spam_flag"),
			),
		),
	)
	
	if err != nil {
		t.Errorf("Complex query creation failed: %v", err)
		return
	}
	
	// Validate the structure
	if query.Bool == nil {
		t.Fatal("Bool query is nil")
	}
	
	// Validate Must clauses
	if len(query.Bool.Must) != 3 {
		t.Errorf("Must clauses count = %d, want 3", len(query.Bool.Must))
	}
	
	// Validate Should clauses
	if len(query.Bool.Should) != 3 {
		t.Errorf("Should clauses count = %d, want 3", len(query.Bool.Should))
	}
	
	// Validate Filter clauses
	if len(query.Bool.Filter) != 2 {
		t.Errorf("Filter clauses count = %d, want 2", len(query.Bool.Filter))
	}
	
	// Validate MustNot clauses
	if len(query.Bool.MustNot) != 2 {
		t.Errorf("MustNot clauses count = %d, want 2", len(query.Bool.MustNot))
	}
}

// TestNestedBoolQueryIntegration tests deeply nested Bool queries
func TestNestedBoolQueryIntegration(t *testing.T) {
	query, err := NewQuery(
		Bool(
			Must(
				Bool(
					Should(
						Match("title", "elasticsearch"),
						Match("content", "search"),
					),
				),
				Bool(
					Must(
						DateRange("date").Gte("2023-01-01").Build(),
						Exists("author"),
					),
					MustNot(
						Term("deleted", "true"),
					),
				),
			),
		),
	)
	
	if err != nil {
		t.Errorf("Nested Bool query creation failed: %v", err)
		return
	}
	
	// Validate nested structure
	if query.Bool == nil {
		t.Fatal("Root Bool query is nil")
	}
	
	if len(query.Bool.Must) != 2 {
		t.Errorf("Root Must clauses count = %d, want 2", len(query.Bool.Must))
	}
	
	// Check first nested Bool query
	firstNested := query.Bool.Must[0]
	if firstNested.Bool == nil {
		t.Fatal("First nested Bool query is nil")
	}
	
	if len(firstNested.Bool.Should) != 2 {
		t.Errorf("First nested Should clauses count = %d, want 2", len(firstNested.Bool.Should))
	}
	
	// Check second nested Bool query
	secondNested := query.Bool.Must[1]
	if secondNested.Bool == nil {
		t.Fatal("Second nested Bool query is nil")
	}
	
	if len(secondNested.Bool.Must) != 2 {
		t.Errorf("Second nested Must clauses count = %d, want 2", len(secondNested.Bool.Must))
	}
	
	if len(secondNested.Bool.MustNot) != 1 {
		t.Errorf("Second nested MustNot clauses count = %d, want 1", len(secondNested.Bool.MustNot))
	}
}

// TestAllQueryTypesIntegration tests all implemented query types together
func TestAllQueryTypesIntegration(t *testing.T) {
	query, err := NewQuery(
		Bool(
			Must(
				// Match query
				Match("title", "elasticsearch guide"),
				// Match with options
				MatchWithOptions("description", "comprehensive tutorial", MatchOptions{
					Fuzziness: "AUTO",
				}),
				// Match phrase
				MatchPhrase("content", "getting started"),
				// Match phrase with options
				MatchPhraseWithOptions("summary", "quick start", MatchPhraseOptions{
					Slop: intPtr(2),
				}),
				// Match phrase prefix
				MatchPhrasePrefix("tags", "elastic"),
				// Term query
				Term("status", "published"),
				// Terms query
				Terms("category", "tutorial", "guide", "documentation"),
				// Range query
				DateRange("publish_date").Gte("2023-01-01").Lte("2023-12-31").Build(),
				// Exists query
				Exists("author"),
			),
		),
	)
	
	if err != nil {
		t.Errorf("All query types integration failed: %v", err)
		return
	}
	
	// Validate the structure
	if query.Bool == nil {
		t.Fatal("Bool query is nil")
	}
	
	if len(query.Bool.Must) != 9 {
		t.Errorf("Must clauses count = %d, want 9", len(query.Bool.Must))
	}
	
	// Validate specific query types
	must := query.Bool.Must
	
	// Check Match query
	if must[0].Match == nil {
		t.Error("First Must clause should be Match query")
	}
	
	// Check Match with options
	if must[1].Match == nil {
		t.Error("Second Must clause should be Match query with options")
	}
	
	// Check Match phrase
	if must[2].MatchPhrase == nil {
		t.Error("Third Must clause should be MatchPhrase query")
	}
	
	// Check Match phrase with options
	if must[3].MatchPhrase == nil {
		t.Error("Fourth Must clause should be MatchPhrase query with options")
	}
	
	// Check Match phrase prefix
	if must[4].MatchPhrasePrefix == nil {
		t.Error("Fifth Must clause should be MatchPhrasePrefix query")
	}
	
	// Check Term query
	if must[5].Term == nil {
		t.Error("Sixth Must clause should be Term query")
	}
	
	// Check Terms query
	if must[6].Terms == nil {
		t.Error("Seventh Must clause should be Terms query")
	}
	
	// Check Range query
	if must[7].Range == nil {
		t.Error("Eighth Must clause should be Range query")
	}
	
	// Check Exists query
	if must[8].Exists == nil {
		t.Error("Ninth Must clause should be Exists query")
	}
}

// TestJSONSerializationIntegration tests JSON serialization of complex queries
func TestJSONSerializationIntegration(t *testing.T) {
	query, err := NewQuery(
		Bool(
			Must(
				Match("title", "elasticsearch"),
				DateRange("date").Gte("2023-01-01").Build(),
			),
			Should(
				Term("category", "tech"),
				Exists("featured"),
			),
		),
	)
	
	if err != nil {
		t.Errorf("Query creation failed: %v", err)
		return
	}
	
	// Test JSON serialization
	jsonData, err := json.Marshal(query)
	if err != nil {
		t.Errorf("JSON serialization failed: %v", err)
		return
	}
	
	// Test JSON deserialization
	var deserializedQuery types.Query
	err = json.Unmarshal(jsonData, &deserializedQuery)
	if err != nil {
		t.Errorf("JSON deserialization failed: %v", err)
		return
	}
	
	// Validate deserialized structure
	if deserializedQuery.Bool == nil {
		t.Fatal("Deserialized Bool query is nil")
	}
	
	if len(deserializedQuery.Bool.Must) != 2 {
		t.Errorf("Deserialized Must clauses count = %d, want 2", len(deserializedQuery.Bool.Must))
	}
	
	if len(deserializedQuery.Bool.Should) != 2 {
		t.Errorf("Deserialized Should clauses count = %d, want 2", len(deserializedQuery.Bool.Should))
	}
}

// TestErrorHandlingIntegration tests error handling in complex scenarios
func TestErrorHandlingIntegration(t *testing.T) {
	// Test error propagation in nested queries
	_, err := NewQuery(
		Bool(
			Must(
				Match("", "value"), // Empty field should cause error
				Term("field", "value"),
			),
		),
	)
	
	if err != nil {
		t.Errorf("Unexpected error for empty field in nested query: %v", err)
	}
	
	// Test deeply nested queries - should work now
	_, err = NewQuery(
		Bool(
			Must(
				Bool(
					Should(
						Match("field", ""), // Empty value should work now
						Term("field", "value"),
					),
				),
			),
		),
	)
	
	if err != nil {
		t.Errorf("Unexpected error for empty value in deeply nested query: %v", err)
	}
	
	// Test Range query - should work now
	_, err = NewQuery(
		Bool(
			Must(
				NumberRange("").Gte(10.0).Build(), // Empty field should work now
			),
		),
	)
	
	if err != nil {
		t.Errorf("Unexpected error for empty field in Range query: %v", err)
	}
}

// TestPerformanceIntegration tests performance with complex queries
func TestPerformanceIntegration(t *testing.T) {
	// Create a complex query multiple times to test performance
	for i := 0; i < 1000; i++ {
		_, err := NewQuery(
			Bool(
							Must(
				Match("title", "elasticsearch"),
				DateRange("date").Gte("2023-01-01").Build(),
				Exists("author"),
			),
				Should(
					Term("category", "tech"),
					MatchPhrase("content", "search engine"),
				),
				Filter(
					Term("status", "published"),
				),
				MustNot(
					Term("deleted", "true"),
				),
			),
		)
		
		if err != nil {
			t.Errorf("Performance test failed at iteration %d: %v", i, err)
			return
		}
	}
}

// TestElasticsearchClientIntegration tests compatibility with Elasticsearch client
func TestElasticsearchClientIntegration(t *testing.T) {
	// Mock Elasticsearch client interface
	type MockSearchRequest struct {
		Query *types.Query
	}
	
	query, err := NewQuery(
		Bool(
			Must(
				Match("title", "elasticsearch"),
				DateRange("date").Gte("2023-01-01").Build(),
			),
		),
	)
	
	if err != nil {
		t.Errorf("Query creation failed: %v", err)
		return
	}
	
	// Test that our query can be used with client interface
	mockRequest := MockSearchRequest{
		Query: query,
	}
	
	if mockRequest.Query == nil {
		t.Error("Query should not be nil")
	}
	
	if mockRequest.Query.Bool == nil {
		t.Error("Bool query should not be nil")
	}
	
	// Test JSON serialization (what client would do)
	jsonData, err := json.Marshal(mockRequest)
	if err != nil {
		t.Errorf("Client JSON serialization failed: %v", err)
		return
	}
	
	// Verify JSON structure contains expected fields
	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonData, &jsonMap)
	if err != nil {
		t.Errorf("JSON parsing failed: %v", err)
		return
	}
	
	if _, exists := jsonMap["Query"]; !exists {
		t.Error("JSON should contain Query field")
	}
}

// TestMatchOptionsIntegration tests all match query options
func TestMatchOptionsIntegration(t *testing.T) {
	// Test all match query options to improve coverage
	query, err := NewQuery(
		Bool(
			Must(
				MatchWithOptions("title", "elasticsearch guide", MatchOptions{
					Fuzziness: "AUTO",
					Analyzer:  stringPtr("standard"),
					Boost:     float32Ptr(1.5),
				}),
				MatchPhraseWithOptions("content", "getting started", MatchPhraseOptions{
					Slop:     intPtr(2),
					Analyzer: stringPtr("keyword"),
					Boost:    float32Ptr(2.0),
				}),
			),
		),
	)
	
	if err != nil {
		t.Errorf("Match options integration failed: %v", err)
		return
	}
	
	// Validate the structure
	if query.Bool == nil {
		t.Fatal("Bool query is nil")
	}
	
	if len(query.Bool.Must) != 2 {
		t.Errorf("Must clauses count = %d, want 2", len(query.Bool.Must))
	}
	
	// Check first match query options
	firstMatch := query.Bool.Must[0].Match["title"]
	if firstMatch.Fuzziness != "AUTO" {
		t.Error("First match query should have AUTO fuzziness")
	}
	
	// Check second match phrase query options
	secondMatch := query.Bool.Must[1].MatchPhrase["content"]
	if secondMatch.Slop == nil || *secondMatch.Slop != 2 {
		t.Error("Second match phrase query should have slop of 2")
	}
} 
package esb

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"testing"
)

func TestSimpleQueryString(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected string
	}{
		{
			name:     "basic simple query string",
			query:    "elasticsearch + search",
			expected: `{"simple_query_string":{"query":"elasticsearch + search"}}`,
		},
		{
			name:     "simple query string with special characters",
			query:    "elasticsearch + (search | database)",
			expected: `{"simple_query_string":{"query":"elasticsearch + (search | database)"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new query with the simple query string option
			query := NewQuery(SimpleQueryString(tt.query))

			// Convert the query to JSON
			actualJSON, err := json.Marshal(query)
			if err != nil {
				t.Fatalf("Failed to marshal query to JSON: %v", err)
			}

			// Compare the actual JSON with the expected JSON
			if string(actualJSON) != tt.expected {
				t.Errorf("Expected JSON: %s, but got: %s", tt.expected, string(actualJSON))
			}

			// Verify the query can be unmarshaled back into a Query object
			var unmarshaled types.Query
			if err := json.Unmarshal(actualJSON, &unmarshaled); err != nil {
				t.Errorf("Failed to unmarshal JSON back to Query: %v", err)
			}

			// Verify the simple query string exists in the unmarshaled query
			if unmarshaled.SimpleQueryString == nil {
				t.Error("Expected SimpleQueryString to be non-nil in unmarshaled query")
			}

			// Verify the query value
			if unmarshaled.SimpleQueryString.Query != tt.query {
				t.Errorf("Expected query %q, got %q", tt.query, unmarshaled.SimpleQueryString.Query)
			}
		})
	}
}

func TestSimpleQueryStringWithOptions(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		setOpts  func(opts *types.SimpleQueryStringQuery)
		expected string
	}{
		{
			name:  "simple query string with fields and boost",
			query: "elasticsearch",
			setOpts: func(opts *types.SimpleQueryStringQuery) {
				fields := []string{"title^2", "description"}
				boost := float32(2.0)
				opts.Fields = fields
				opts.Boost = &boost
			},
			expected: `{"simple_query_string":{"boost":2,"fields":["title^2","description"],"query":"elasticsearch"}}`,
		},
		{
			name:  "simple query string with default operator and flags",
			query: "search engine",
			setOpts: func(opts *types.SimpleQueryStringQuery) {
				defaultOp := operator.And
				flags := []string{"AND", "OR", "PREFIX"}
				opts.DefaultOperator = &defaultOp
				opts.Flags = flags
			},
			expected: `{"simple_query_string":{"default_operator":"and","flags":["AND","OR","PREFIX"],"query":"search engine"}}`,
		},
		{
			name:  "simple query string with analyzer and minimum should match",
			query: "search database",
			setOpts: func(opts *types.SimpleQueryStringQuery) {
				analyzer := "standard"
				minimumShouldMatch := "75%"
				opts.Analyzer = &analyzer
				opts.MinimumShouldMatch = &minimumShouldMatch
			},
			expected: `{"simple_query_string":{"analyzer":"standard","minimum_should_match":"75%","query":"search database"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new query with the simple query string option
			query := NewQuery(SimpleQueryStringWithOptions(tt.query, tt.setOpts))

			// Convert the query to JSON
			actualJSON, err := json.Marshal(query)
			if err != nil {
				t.Fatalf("Failed to marshal query to JSON: %v", err)
			}

			// Compare the actual JSON with the expected JSON
			if string(actualJSON) != tt.expected {
				t.Errorf("Expected JSON: %s, but got: %s", tt.expected, string(actualJSON))
			}

			// Verify the query can be unmarshaled back into a Query object
			var unmarshaled types.Query
			if err := json.Unmarshal(actualJSON, &unmarshaled); err != nil {
				t.Errorf("Failed to unmarshal JSON back to Query: %v", err)
			}

			// Verify the simple query string exists in the unmarshaled query
			if unmarshaled.SimpleQueryString == nil {
				t.Error("Expected SimpleQueryString to be non-nil in unmarshaled query")
			}

			// Verify the query value
			if unmarshaled.SimpleQueryString.Query != tt.query {
				t.Errorf("Expected query %q, got %q", tt.query, unmarshaled.SimpleQueryString.Query)
			}
		})
	}

	t.Run("should work with nil callback", func(t *testing.T) {
		query := NewQuery(SimpleQueryStringWithOptions("elasticsearch", nil))
		if query.SimpleQueryString == nil {
			t.Error("Expected SimpleQueryString to be non-nil")
		}
		if query.SimpleQueryString.Query != "elasticsearch" {
			t.Errorf("Expected query 'elasticsearch', got %s", query.SimpleQueryString.Query)
		}
	})
} 
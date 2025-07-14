package esb

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"testing"
)

func TestQueryString(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected string
	}{
		{
			name:     "basic query string",
			query:    "title:elasticsearch",
			expected: `{"query_string":{"query":"title:elasticsearch"}}`,
		},
		{
			name:     "complex query string",
			query:    "title:elasticsearch AND (tags:search OR tags:database)",
			expected: `{"query_string":{"query":"title:elasticsearch AND (tags:search OR tags:database)"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new query with the query string option
			query := NewQuery(QueryString(tt.query))

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

			// Verify the query string exists in the unmarshaled query
			if unmarshaled.QueryString == nil {
				t.Error("Expected QueryString to be non-nil in unmarshaled query")
			}

			// Verify the query value
			if unmarshaled.QueryString.Query != tt.query {
				t.Errorf("Expected query %q, got %q", tt.query, unmarshaled.QueryString.Query)
			}
		})
	}
}

func TestQueryStringWithOptions(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		setOpts  func(opts *types.QueryStringQuery)
		expected string
	}{
		{
			name:  "query string with default field and operator",
			query: "elasticsearch",
			setOpts: func(opts *types.QueryStringQuery) {
				defaultField := "title"
				defaultOp := operator.And
				opts.DefaultField = &defaultField
				opts.DefaultOperator = &defaultOp
			},
			expected: `{"query_string":{"default_field":"title","default_operator":"and","query":"elasticsearch"}}`,
		},
		{
			name:  "query string with multiple fields and boost",
			query: "search engine",
			setOpts: func(opts *types.QueryStringQuery) {
				fields := []string{"title^2", "description"}
				boost := float32(2.0)
				opts.Fields = fields
				opts.Boost = &boost
			},
			expected: `{"query_string":{"boost":2,"fields":["title^2","description"],"query":"search engine"}}`,
		},
		{
			name:  "query string with analyzer and allow leading wildcard",
			query: "*search",
			setOpts: func(opts *types.QueryStringQuery) {
				analyzer := "standard"
				allowLeadingWildcard := true
				opts.Analyzer = &analyzer
				opts.AllowLeadingWildcard = &allowLeadingWildcard
			},
			expected: `{"query_string":{"allow_leading_wildcard":true,"analyzer":"standard","query":"*search"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new query with the query string option
			query := NewQuery(QueryStringWithOptions(tt.query, tt.setOpts))

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

			// Verify the query string exists in the unmarshaled query
			if unmarshaled.QueryString == nil {
				t.Error("Expected QueryString to be non-nil in unmarshaled query")
			}

			// Verify the query value
			if unmarshaled.QueryString.Query != tt.query {
				t.Errorf("Expected query %q, got %q", tt.query, unmarshaled.QueryString.Query)
			}
		})
	}

	t.Run("should work with nil callback", func(t *testing.T) {
		query := NewQuery(QueryStringWithOptions("elasticsearch", nil))
		if query.QueryString == nil {
			t.Error("Expected QueryString to be non-nil")
		}
		if query.QueryString.Query != "elasticsearch" {
			t.Errorf("Expected query 'elasticsearch', got %s", query.QueryString.Query)
		}
	})
} 
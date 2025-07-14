package esb

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/childscoremode"
	"testing"
)

func TestNested(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		query    QueryOption
		expected string
	}{
		{
			name: "basic nested query",
			path: "comments",
			query: Match("comments.author", "john"),
			expected: `{"nested":{"path":"comments","query":{"match":{"comments.author":{"query":"john"}}}}}`,
		},
		{
			name: "nested query with bool query",
			path: "comments",
			query: Bool(
				Must(
					Match("comments.author", "john"),
					Match("comments.content", "great"),
				),
			),
			expected: `{"nested":{"path":"comments","query":{"bool":{"must":[{"match":{"comments.author":{"query":"john"}}},{"match":{"comments.content":{"query":"great"}}}]}}}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new query with the nested option
			query := NewQuery(Nested(tt.path, tt.query))

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

			// Verify the nested query exists
			if unmarshaled.Nested == nil {
				t.Error("Expected Nested to be non-nil in unmarshaled query")
			}

			// Verify the path
			if unmarshaled.Nested.Path != tt.path {
				t.Errorf("Expected path %q, got %q", tt.path, unmarshaled.Nested.Path)
			}
		})
	}
}

func TestNestedWithOptions(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		query    QueryOption
		setOpts  func(opts *types.NestedQuery)
		expected string
	}{
		{
			name: "nested query with score mode",
			path: "comments",
			query: Match("comments.author", "john"),
			setOpts: func(opts *types.NestedQuery) {
				scoreMode := childscoremode.Avg
				opts.ScoreMode = &scoreMode
			},
			expected: `{"nested":{"path":"comments","query":{"match":{"comments.author":{"query":"john"}}},"score_mode":"avg"}}`,
		},
		{
			name: "nested query with ignore unmapped",
			path: "comments",
			query: Match("comments.author", "john"),
			setOpts: func(opts *types.NestedQuery) {
				ignoreUnmapped := true
				opts.IgnoreUnmapped = &ignoreUnmapped
			},
			expected: `{"nested":{"ignore_unmapped":true,"path":"comments","query":{"match":{"comments.author":{"query":"john"}}}}}`,
		},
		{
			name: "nested query with all options",
			path: "comments",
			query: Match("comments.author", "john"),
			setOpts: func(opts *types.NestedQuery) {
				scoreMode := childscoremode.Max
				ignoreUnmapped := true
				boost := float32(2.0)
				opts.ScoreMode = &scoreMode
				opts.IgnoreUnmapped = &ignoreUnmapped
				opts.Boost = &boost
			},
			expected: `{"nested":{"boost":2,"ignore_unmapped":true,"path":"comments","query":{"match":{"comments.author":{"query":"john"}}},"score_mode":"max"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new query with the nested option
			query := NewQuery(NestedWithOptions(tt.path, tt.query, tt.setOpts))

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

			// Verify the nested query exists
			if unmarshaled.Nested == nil {
				t.Error("Expected Nested to be non-nil in unmarshaled query")
			}

			// Verify the path
			if unmarshaled.Nested.Path != tt.path {
				t.Errorf("Expected path %q, got %q", tt.path, unmarshaled.Nested.Path)
			}
		})
	}

	t.Run("should work with nil callback", func(t *testing.T) {
		query := NewQuery(NestedWithOptions("comments", Match("comments.author", "john"), nil))
		if query.Nested == nil {
			t.Error("Expected Nested to be non-nil")
		}
		if query.Nested.Path != "comments" {
			t.Errorf("Expected path 'comments', got %s", query.Nested.Path)
		}
	})
} 
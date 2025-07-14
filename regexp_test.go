package esb

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"testing"
)

func TestRegexp(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		value    string
		expected string
	}{
		{
			name:     "basic regexp query",
			field:    "username",
			value:    "j.*n",
			expected: `{"regexp":{"username":{"value":"j.*n"}}}`,
		},
		{
			name:     "regexp query with complex pattern",
			field:    "email",
			value:    ".*@example\\.com",
			expected: `{"regexp":{"email":{"value":".*@example\\.com"}}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new query with the regexp option
			query := NewQuery(Regexp(tt.field, tt.value))

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

			// Verify the regexp exists in the unmarshaled query
			if unmarshaled.Regexp == nil {
				t.Error("Expected Regexp to be non-nil in unmarshaled query")
			}

			// Verify the specific field exists in the regexp query
			if _, ok := unmarshaled.Regexp[tt.field]; !ok {
				t.Errorf("Expected field %q in regexp query", tt.field)
			}

			// Verify the value
			if unmarshaled.Regexp[tt.field].Value != tt.value {
				t.Errorf("Expected value %q, got %q", tt.value, unmarshaled.Regexp[tt.field].Value)
			}
		})
	}
}

func TestRegexpWithOptions(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		value    string
		setOpts  func(opts *types.RegexpQuery)
		expected string
	}{
		{
			name:  "regexp query with flags and max states",
			field: "username",
			value: "j.*n",
			setOpts: func(opts *types.RegexpQuery) {
				flags := "ALL"
				maxDeterminizedStates := 10000
				opts.Flags = &flags
				opts.MaxDeterminizedStates = &maxDeterminizedStates
			},
			expected: `{"regexp":{"username":{"flags":"ALL","max_determinized_states":10000,"value":"j.*n"}}}`,
		},
		{
			name:  "regexp query with case insensitive flag",
			field: "email",
			value: ".*@example\\.com",
			setOpts: func(opts *types.RegexpQuery) {
				flags := "CASE_INSENSITIVE"
				opts.Flags = &flags
			},
			expected: `{"regexp":{"email":{"flags":"CASE_INSENSITIVE","value":".*@example\\.com"}}}`,
		},
		{
			name:  "regexp query with boost",
			field: "username",
			value: "j.*n",
			setOpts: func(opts *types.RegexpQuery) {
				boost := float32(2.0)
				opts.Boost = &boost
			},
			expected: `{"regexp":{"username":{"boost":2,"value":"j.*n"}}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new query with the regexp option
			query := NewQuery(RegexpWithOptions(tt.field, tt.value, tt.setOpts))

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

			// Verify the regexp exists in the unmarshaled query
			if unmarshaled.Regexp == nil {
				t.Error("Expected Regexp to be non-nil in unmarshaled query")
			}

			// Verify the specific field exists in the regexp query
			if _, ok := unmarshaled.Regexp[tt.field]; !ok {
				t.Errorf("Expected field %q in regexp query", tt.field)
			}

			// Verify the value
			if unmarshaled.Regexp[tt.field].Value != tt.value {
				t.Errorf("Expected value %q, got %q", tt.value, unmarshaled.Regexp[tt.field].Value)
			}
		})
	}

	t.Run("should work with nil callback", func(t *testing.T) {
		query := NewQuery(RegexpWithOptions("username", "j.*n", nil))
		if query.Regexp == nil {
			t.Error("Expected Regexp to be non-nil")
		}
		if query.Regexp["username"].Value != "j.*n" {
			t.Errorf("Expected value 'j.*n', got %s", query.Regexp["username"].Value)
		}
	})
} 
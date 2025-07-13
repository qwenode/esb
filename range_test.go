package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestRangeBuilder_BasicUsage(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		builder  func() *RangeBuilder
		validate func(t *testing.T, query *types.Query)
	}{
		{
			name:  "gte condition",
			field: "age",
			builder: func() *RangeBuilder {
				return Range("age").Gte(18)
			},
			validate: func(t *testing.T, query *types.Query) {
				if query.Range == nil {
					t.Fatal("Range query is nil")
				}
				rangeQuery := query.Range["age"]
				untypedQuery := rangeQuery.(types.UntypedRangeQuery)
				if string(untypedQuery.Gte) != "18" {
					t.Errorf("Gte = %s, want 18", string(untypedQuery.Gte))
				}
			},
		},
		{
			name:  "gt condition",
			field: "price",
			builder: func() *RangeBuilder {
				return Range("price").Gt(10.5)
			},
			validate: func(t *testing.T, query *types.Query) {
				if query.Range == nil {
					t.Fatal("Range query is nil")
				}
				rangeQuery := query.Range["price"]
				untypedQuery := rangeQuery.(types.UntypedRangeQuery)
				if string(untypedQuery.Gt) != "10.5" {
					t.Errorf("Gt = %s, want 10.5", string(untypedQuery.Gt))
				}
			},
		},
		{
			name:  "lte condition",
			field: "score",
			builder: func() *RangeBuilder {
				return Range("score").Lte(100)
			},
			validate: func(t *testing.T, query *types.Query) {
				if query.Range == nil {
					t.Fatal("Range query is nil")
				}
				rangeQuery := query.Range["score"]
				untypedQuery := rangeQuery.(types.UntypedRangeQuery)
				if string(untypedQuery.Lte) != "100" {
					t.Errorf("Lte = %s, want 100", string(untypedQuery.Lte))
				}
			},
		},
		{
			name:  "lt condition",
			field: "temperature",
			builder: func() *RangeBuilder {
				return Range("temperature").Lt(25.5)
			},
			validate: func(t *testing.T, query *types.Query) {
				if query.Range == nil {
					t.Fatal("Range query is nil")
				}
				rangeQuery := query.Range["temperature"]
				untypedQuery := rangeQuery.(types.UntypedRangeQuery)
				if string(untypedQuery.Lt) != "25.5" {
					t.Errorf("Lt = %s, want 25.5", string(untypedQuery.Lt))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query, err := NewQuery(tt.builder().Build())
			if err != nil {
				t.Errorf("NewQuery() error = %v", err)
				return
			}
			tt.validate(t, query)
		})
	}
}

func TestRangeBuilder_ChainedConditions(t *testing.T) {
	query, err := NewQuery(
		Range("age").Gte(18).Lt(65).Build(),
	)
	if err != nil {
		t.Errorf("NewQuery() error = %v", err)
		return
	}

	if query.Range == nil {
		t.Fatal("Range query is nil")
	}

	rangeQuery := query.Range["age"]
	untypedQuery := rangeQuery.(types.UntypedRangeQuery)
	
	if string(untypedQuery.Gte) != "18" {
		t.Errorf("Gte = %s, want 18", string(untypedQuery.Gte))
	}
	
	if string(untypedQuery.Lt) != "65" {
		t.Errorf("Lt = %s, want 65", string(untypedQuery.Lt))
	}
}

func TestRangeBuilder_FromToConditions(t *testing.T) {
	query, err := NewQuery(
		Range("price").From(10.0).To(100.0).Build(),
	)
	if err != nil {
		t.Errorf("NewQuery() error = %v", err)
		return
	}

	if query.Range == nil {
		t.Fatal("Range query is nil")
	}

	rangeQuery := query.Range["price"]
	untypedQuery := rangeQuery.(types.UntypedRangeQuery)
	
	if untypedQuery.From == nil {
		t.Fatal("From is nil")
	}
	if string(*untypedQuery.From) != "10" {
		t.Errorf("From = %s, want 10", string(*untypedQuery.From))
	}
	
	if untypedQuery.To == nil {
		t.Fatal("To is nil")
	}
	if string(*untypedQuery.To) != "100" {
		t.Errorf("To = %s, want 100", string(*untypedQuery.To))
	}
}

func TestRangeBuilder_StringValues(t *testing.T) {
	query, err := NewQuery(
		Range("timestamp").Gte("2023-01-01").Lte("2023-12-31").Build(),
	)
	if err != nil {
		t.Errorf("NewQuery() error = %v", err)
		return
	}

	if query.Range == nil {
		t.Fatal("Range query is nil")
	}

	rangeQuery := query.Range["timestamp"]
	untypedQuery := rangeQuery.(types.UntypedRangeQuery)
	
	if string(untypedQuery.Gte) != `"2023-01-01"` {
		t.Errorf("Gte = %s, want \"2023-01-01\"", string(untypedQuery.Gte))
	}
	
	if string(untypedQuery.Lte) != `"2023-12-31"` {
		t.Errorf("Lte = %s, want \"2023-12-31\"", string(untypedQuery.Lte))
	}
}

func TestRangeBuilder_WithOptions(t *testing.T) {
	query, err := NewQuery(
		Range("timestamp").
			Gte("2023-01-01").
			Lte("2023-12-31").
			Format("yyyy-MM-dd").
			TimeZone("UTC").
			Boost(1.5).
			Build(),
	)
	if err != nil {
		t.Errorf("NewQuery() error = %v", err)
		return
	}

	if query.Range == nil {
		t.Fatal("Range query is nil")
	}

	rangeQuery := query.Range["timestamp"]
	untypedQuery := rangeQuery.(types.UntypedRangeQuery)
	
	if untypedQuery.Format == nil || *untypedQuery.Format != "yyyy-MM-dd" {
		t.Errorf("Format = %v, want yyyy-MM-dd", untypedQuery.Format)
	}
	
	if untypedQuery.TimeZone == nil || *untypedQuery.TimeZone != "UTC" {
		t.Errorf("TimeZone = %v, want UTC", untypedQuery.TimeZone)
	}
	
	if untypedQuery.Boost == nil || *untypedQuery.Boost != 1.5 {
		t.Errorf("Boost = %v, want 1.5", untypedQuery.Boost)
	}
}

func TestRangeBuilder_WithBoolQuery(t *testing.T) {
	query, err := NewQuery(
		Bool(
			Must(
				Range("age").Gte(18).Lt(65).Build(),
				Range("score").Gte(80).Build(),
			),
		),
	)
	if err != nil {
		t.Errorf("NewQuery() error = %v", err)
		return
	}

	if query.Bool == nil {
		t.Fatal("Bool query is nil")
	}

	if len(query.Bool.Must) != 2 {
		t.Errorf("Bool.Must length = %d, want 2", len(query.Bool.Must))
	}

	// Check first range query
	if query.Bool.Must[0].Range == nil {
		t.Fatal("First Must clause should be Range query")
	}
	
	ageRange := query.Bool.Must[0].Range["age"]
	ageQuery := ageRange.(types.UntypedRangeQuery)
	if string(ageQuery.Gte) != "18" {
		t.Errorf("Age Gte = %s, want 18", string(ageQuery.Gte))
	}

	// Check second range query
	if query.Bool.Must[1].Range == nil {
		t.Fatal("Second Must clause should be Range query")
	}
	
	scoreRange := query.Bool.Must[1].Range["score"]
	scoreQuery := scoreRange.(types.UntypedRangeQuery)
	if string(scoreQuery.Gte) != "80" {
		t.Errorf("Score Gte = %s, want 80", string(scoreQuery.Gte))
	}
}

func TestRangeBuilder_EmptyField(t *testing.T) {
	_, err := NewQuery(Range("").Gte(18).Build())
	if err == nil {
		t.Errorf("Expected error for empty field")
	}
	if err != ErrEmptyField {
		t.Errorf("Expected ErrEmptyField, got %v", err)
	}
}

func TestRangeBuilder_WhitespaceField(t *testing.T) {
	_, err := NewQuery(Range("   ").Gte(18).Build())
	if err == nil {
		t.Errorf("Expected error for whitespace field")
	}
	if err != ErrEmptyField {
		t.Errorf("Expected ErrEmptyField, got %v", err)
	}
}

// Benchmark tests
func BenchmarkRangeBuilder_Simple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewQuery(Range("age").Gte(18).Build())
		if err != nil {
			b.Errorf("NewQuery() error = %v", err)
		}
	}
}

func BenchmarkRangeBuilder_Complex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewQuery(
			Range("age").Gte(18).Lt(65).Boost(1.2).Build(),
		)
		if err != nil {
			b.Errorf("NewQuery() error = %v", err)
		}
	}
}

func BenchmarkRangeBuilder_WithBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewQuery(
			Bool(
				Must(
					Range("age").Gte(18).Lt(65).Build(),
					Range("score").Gte(80).Build(),
				),
			),
		)
		if err != nil {
			b.Errorf("NewQuery() error = %v", err)
		}
	}
} 
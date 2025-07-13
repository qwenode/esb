package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/rangerelation"
)

// =============================================================================
// NumberRangeBuilder Tests
// =============================================================================

func TestNumberRangeBuilder_BasicUsage(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		builder  func() *NumberRangeBuilder
		validate func(t *testing.T, query *types.Query)
	}{
		{
			name:  "gte condition",
			field: "age",
			builder: func() *NumberRangeBuilder {
				return NumberRange("age").Gte(18.0)
			},
			validate: func(t *testing.T, query *types.Query) {
				if query.Range == nil {
					t.Fatal("Range query is nil")
				}
				rangeQuery := query.Range["age"]
				numberQuery := rangeQuery.(types.NumberRangeQuery)
				if numberQuery.Gte == nil || float64(*numberQuery.Gte) != 18.0 {
					t.Errorf("Gte = %v, want 18.0", numberQuery.Gte)
				}
			},
		},
		{
			name:  "gt condition",
			field: "price",
			builder: func() *NumberRangeBuilder {
				return NumberRange("price").Gt(10.5)
			},
			validate: func(t *testing.T, query *types.Query) {
				if query.Range == nil {
					t.Fatal("Range query is nil")
				}
				rangeQuery := query.Range["price"]
				numberQuery := rangeQuery.(types.NumberRangeQuery)
				if numberQuery.Gt == nil || float64(*numberQuery.Gt) != 10.5 {
					t.Errorf("Gt = %v, want 10.5", numberQuery.Gt)
				}
			},
		},
		{
			name:  "lte condition",
			field: "score",
			builder: func() *NumberRangeBuilder {
				return NumberRange("score").Lte(100.0)
			},
			validate: func(t *testing.T, query *types.Query) {
				if query.Range == nil {
					t.Fatal("Range query is nil")
				}
				rangeQuery := query.Range["score"]
				numberQuery := rangeQuery.(types.NumberRangeQuery)
				if numberQuery.Lte == nil || float64(*numberQuery.Lte) != 100.0 {
					t.Errorf("Lte = %v, want 100.0", numberQuery.Lte)
				}
			},
		},
		{
			name:  "lt condition",
			field: "temperature",
			builder: func() *NumberRangeBuilder {
				return NumberRange("temperature").Lt(25.5)
			},
			validate: func(t *testing.T, query *types.Query) {
				if query.Range == nil {
					t.Fatal("Range query is nil")
				}
				rangeQuery := query.Range["temperature"]
				numberQuery := rangeQuery.(types.NumberRangeQuery)
				if numberQuery.Lt == nil || float64(*numberQuery.Lt) != 25.5 {
					t.Errorf("Lt = %v, want 25.5", numberQuery.Lt)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &types.Query{}
			builder := tt.builder()
			queryOption := builder.Build()
			
			err := queryOption(query)
			if err != nil {
				t.Fatalf("Build() error = %v", err)
			}
			
			tt.validate(t, query)
		})
	}
}

func TestNumberRangeBuilder_ChainedConditions(t *testing.T) {
	builder := NumberRange("age").Gte(18.0).Lt(65.0)
	query := &types.Query{}
	queryOption := builder.Build()
	
	err := queryOption(query)
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	
	if query.Range == nil {
		t.Fatal("Range query is nil")
	}
	
	rangeQuery := query.Range["age"]
	numberQuery := rangeQuery.(types.NumberRangeQuery)
	
	if numberQuery.Gte == nil || float64(*numberQuery.Gte) != 18.0 {
		t.Errorf("Gte = %v, want 18.0", numberQuery.Gte)
	}
	
	if numberQuery.Lt == nil || float64(*numberQuery.Lt) != 65.0 {
		t.Errorf("Lt = %v, want 65.0", numberQuery.Lt)
	}
}

func TestNumberRangeBuilder_FromToConditions(t *testing.T) {
	builder := NumberRange("price").From(10.0).To(100.0)
	query := &types.Query{}
	queryOption := builder.Build()
	
	err := queryOption(query)
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	
	rangeQuery := query.Range["price"]
	numberQuery := rangeQuery.(types.NumberRangeQuery)
	
	if numberQuery.From == nil || float64(*numberQuery.From) != 10.0 {
		t.Errorf("From = %v, want 10.0", numberQuery.From)
	}
	
	if numberQuery.To == nil || float64(*numberQuery.To) != 100.0 {
		t.Errorf("To = %v, want 100.0", numberQuery.To)
	}
}

func TestNumberRangeBuilder_WithOptions(t *testing.T) {
	boost := float32(2.0)
	queryName := "test_query"
	relation := rangerelation.Within
	
	builder := NumberRange("score").Gte(50.0).Boost(boost).QueryName(queryName).Relation(&relation)
	query := &types.Query{}
	queryOption := builder.Build()
	
	err := queryOption(query)
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	
	rangeQuery := query.Range["score"]
	numberQuery := rangeQuery.(types.NumberRangeQuery)
	
	if numberQuery.Boost == nil || *numberQuery.Boost != boost {
		t.Errorf("Boost = %v, want %v", numberQuery.Boost, boost)
	}
	
	if numberQuery.QueryName_ == nil || *numberQuery.QueryName_ != queryName {
		t.Errorf("QueryName = %v, want %v", numberQuery.QueryName_, queryName)
	}
	
	if numberQuery.Relation == nil || *numberQuery.Relation != relation {
		t.Errorf("Relation = %v, want %v", numberQuery.Relation, relation)
	}
}

// =============================================================================
// DateRangeBuilder Tests
// =============================================================================

func TestDateRangeBuilder_BasicUsage(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		builder  func() *DateRangeBuilder
		validate func(t *testing.T, query *types.Query)
	}{
		{
			name:  "gte condition",
			field: "created_at",
			builder: func() *DateRangeBuilder {
				return DateRange("created_at").Gte("2023-01-01")
			},
			validate: func(t *testing.T, query *types.Query) {
				rangeQuery := query.Range["created_at"]
				dateQuery := rangeQuery.(types.DateRangeQuery)
				if dateQuery.Gte == nil || *dateQuery.Gte != "2023-01-01" {
					t.Errorf("Gte = %v, want 2023-01-01", dateQuery.Gte)
				}
			},
		},
		{
			name:  "gt condition",
			field: "timestamp",
			builder: func() *DateRangeBuilder {
				return DateRange("timestamp").Gt("now-1d")
			},
			validate: func(t *testing.T, query *types.Query) {
				rangeQuery := query.Range["timestamp"]
				dateQuery := rangeQuery.(types.DateRangeQuery)
				if dateQuery.Gt == nil || *dateQuery.Gt != "now-1d" {
					t.Errorf("Gt = %v, want now-1d", dateQuery.Gt)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &types.Query{}
			builder := tt.builder()
			queryOption := builder.Build()
			
			err := queryOption(query)
			if err != nil {
				t.Fatalf("Build() error = %v", err)
			}
			
			tt.validate(t, query)
		})
	}
}

func TestDateRangeBuilder_WithFormatAndTimeZone(t *testing.T) {
	format := "yyyy-MM-dd"
	timeZone := "UTC"
	
	builder := DateRange("created_at").Gte("2023-01-01").Format(format).TimeZone(timeZone)
	query := &types.Query{}
	queryOption := builder.Build()
	
	err := queryOption(query)
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	
	rangeQuery := query.Range["created_at"]
	dateQuery := rangeQuery.(types.DateRangeQuery)
	
	if dateQuery.Format == nil || *dateQuery.Format != format {
		t.Errorf("Format = %v, want %v", dateQuery.Format, format)
	}
	
	if dateQuery.TimeZone == nil || *dateQuery.TimeZone != timeZone {
		t.Errorf("TimeZone = %v, want %v", dateQuery.TimeZone, timeZone)
	}
}

func TestDateRangeBuilder_ChainedConditions(t *testing.T) {
	builder := DateRange("created_at").Gte("2023-01-01").Lt("2023-12-31")
	query := &types.Query{}
	queryOption := builder.Build()
	
	err := queryOption(query)
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	
	rangeQuery := query.Range["created_at"]
	dateQuery := rangeQuery.(types.DateRangeQuery)
	
	if dateQuery.Gte == nil || *dateQuery.Gte != "2023-01-01" {
		t.Errorf("Gte = %v, want 2023-01-01", dateQuery.Gte)
	}
	
	if dateQuery.Lt == nil || *dateQuery.Lt != "2023-12-31" {
		t.Errorf("Lt = %v, want 2023-12-31", dateQuery.Lt)
	}
}

// =============================================================================
// TermRangeBuilder Tests
// =============================================================================

func TestTermRangeBuilder_BasicUsage(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		builder  func() *TermRangeBuilder
		validate func(t *testing.T, query *types.Query)
	}{
		{
			name:  "gte condition",
			field: "username",
			builder: func() *TermRangeBuilder {
				return TermRange("username").Gte("alice")
			},
			validate: func(t *testing.T, query *types.Query) {
				rangeQuery := query.Range["username"]
				termQuery := rangeQuery.(types.TermRangeQuery)
				if termQuery.Gte == nil || *termQuery.Gte != "alice" {
					t.Errorf("Gte = %v, want alice", termQuery.Gte)
				}
			},
		},
		{
			name:  "lt condition",
			field: "category",
			builder: func() *TermRangeBuilder {
				return TermRange("category").Lt("electronics")
			},
			validate: func(t *testing.T, query *types.Query) {
				rangeQuery := query.Range["category"]
				termQuery := rangeQuery.(types.TermRangeQuery)
				if termQuery.Lt == nil || *termQuery.Lt != "electronics" {
					t.Errorf("Lt = %v, want electronics", termQuery.Lt)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &types.Query{}
			builder := tt.builder()
			queryOption := builder.Build()
			
			err := queryOption(query)
			if err != nil {
				t.Fatalf("Build() error = %v", err)
			}
			
			tt.validate(t, query)
		})
	}
}

func TestTermRangeBuilder_ChainedConditions(t *testing.T) {
	builder := TermRange("username").Gte("alice").Lt("bob")
	query := &types.Query{}
	queryOption := builder.Build()
	
	err := queryOption(query)
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	
	rangeQuery := query.Range["username"]
	termQuery := rangeQuery.(types.TermRangeQuery)
	
	if termQuery.Gte == nil || *termQuery.Gte != "alice" {
		t.Errorf("Gte = %v, want alice", termQuery.Gte)
	}
	
	if termQuery.Lt == nil || *termQuery.Lt != "bob" {
		t.Errorf("Lt = %v, want bob", termQuery.Lt)
	}
}

func TestTermRangeBuilder_FromToConditions(t *testing.T) {
	builder := TermRange("category").From("electronics").To("home")
	query := &types.Query{}
	queryOption := builder.Build()
	
	err := queryOption(query)
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	
	rangeQuery := query.Range["category"]
	termQuery := rangeQuery.(types.TermRangeQuery)
	
	if termQuery.From == nil || *termQuery.From != "electronics" {
		t.Errorf("From = %v, want electronics", termQuery.From)
	}
	
	if termQuery.To == nil || *termQuery.To != "home" {
		t.Errorf("To = %v, want home", termQuery.To)
	}
}

// =============================================================================
// Integration Tests
// =============================================================================

func TestRangeBuilder_WithBoolQuery(t *testing.T) {
	boolQuery := Bool(
		Must(
			NumberRange("age").Gte(18.0).Lt(65.0).Build(),
			DateRange("created_at").Gte("2023-01-01").Build(),
			TermRange("status").Gte("active").Build(),
		),
	)
	
	query := &types.Query{}
	queryOption := boolQuery
	
	err := queryOption(query)
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	
	if query.Bool == nil {
		t.Fatal("Bool query is nil")
	}
	
	if len(query.Bool.Must) != 3 {
		t.Errorf("Must conditions count = %d, want 3", len(query.Bool.Must))
	}
	
	// Check first condition (NumberRange)
	firstCondition := query.Bool.Must[0]
	if firstCondition.Range == nil {
		t.Fatal("First condition Range is nil")
	}
	
	// Check second condition (DateRange)
	secondCondition := query.Bool.Must[1]
	if secondCondition.Range == nil {
		t.Fatal("Second condition Range is nil")
	}
	
	// Check third condition (TermRange)
	thirdCondition := query.Bool.Must[2]
	if thirdCondition.Range == nil {
		t.Fatal("Third condition Range is nil")
	}
}

// =============================================================================
// Benchmarks
// =============================================================================

func BenchmarkNumberRangeBuilder_Simple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		query := &types.Query{}
		queryOption := NumberRange("age").Gte(18.0).Build()
		queryOption(query)
	}
}

func BenchmarkDateRangeBuilder_Simple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		query := &types.Query{}
		queryOption := DateRange("created_at").Gte("2023-01-01").Build()
		queryOption(query)
	}
}

func BenchmarkTermRangeBuilder_Simple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		query := &types.Query{}
		queryOption := TermRange("username").Gte("alice").Build()
		queryOption(query)
	}
}

func BenchmarkNumberRangeBuilder_Complex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		query := &types.Query{}
		queryOption := NumberRange("score").Gte(50.0).Lt(100.0).Boost(2.0).QueryName("test").Build()
		queryOption(query)
	}
}

func BenchmarkRangeBuilder_WithBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		query := &types.Query{}
		queryOption := Bool(
			Must(
				NumberRange("age").Gte(18.0).Build(),
				DateRange("created_at").Gte("2023-01-01").Build(),
			),
		)
		queryOption(query)
	}
} 
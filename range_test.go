package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/rangerelation"
)

// =============================================================================
// 数值范围构建器测试
// =============================================================================

func TestNumberRangeBuilder_BasicUsage(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		builder  func() *NumberRangeBuilder
		validate func(t *testing.T, query *types.Query)
	}{
		{
			name:  "大于等于条件",
			field: "age",
			builder: func() *NumberRangeBuilder {
				return NumberRange("age").Gte(18.0)
			},
			validate: func(t *testing.T, query *types.Query) {
				if query.Range == nil {
					t.Fatal("范围查询为 nil")
				}
				rangeQuery := query.Range["age"]
				numberQuery := rangeQuery.(types.NumberRangeQuery)
				if numberQuery.Gte == nil || float64(*numberQuery.Gte) != 18.0 {
					t.Errorf("Gte = %v，期望 18.0", numberQuery.Gte)
				}
			},
		},
		{
			name:  "大于条件",
			field: "price",
			builder: func() *NumberRangeBuilder {
				return NumberRange("price").Gt(10.5)
			},
			validate: func(t *testing.T, query *types.Query) {
				if query.Range == nil {
					t.Fatal("范围查询为 nil")
				}
				rangeQuery := query.Range["price"]
				numberQuery := rangeQuery.(types.NumberRangeQuery)
				if numberQuery.Gt == nil || float64(*numberQuery.Gt) != 10.5 {
					t.Errorf("Gt = %v，期望 10.5", numberQuery.Gt)
				}
			},
		},
		{
			name:  "小于等于条件",
			field: "score",
			builder: func() *NumberRangeBuilder {
				return NumberRange("score").Lte(100.0)
			},
			validate: func(t *testing.T, query *types.Query) {
				if query.Range == nil {
					t.Fatal("范围查询为 nil")
				}
				rangeQuery := query.Range["score"]
				numberQuery := rangeQuery.(types.NumberRangeQuery)
				if numberQuery.Lte == nil || float64(*numberQuery.Lte) != 100.0 {
					t.Errorf("Lte = %v，期望 100.0", numberQuery.Lte)
				}
			},
		},
		{
			name:  "小于条件",
			field: "temperature",
			builder: func() *NumberRangeBuilder {
				return NumberRange("temperature").Lt(25.5)
			},
			validate: func(t *testing.T, query *types.Query) {
				if query.Range == nil {
					t.Fatal("范围查询为 nil")
				}
				rangeQuery := query.Range["temperature"]
				numberQuery := rangeQuery.(types.NumberRangeQuery)
				if numberQuery.Lt == nil || float64(*numberQuery.Lt) != 25.5 {
					t.Errorf("Lt = %v，期望 25.5", numberQuery.Lt)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &types.Query{}
			builder := tt.builder()
			queryOption := builder.Build()
			
			queryOption(query)
			
			tt.validate(t, query)
		})
	}
}

func TestNumberRangeBuilder_ChainedConditions(t *testing.T) {
	builder := NumberRange("age").Gte(18.0).Lt(65.0)
	query := &types.Query{}
	queryOption := builder.Build()
	
	queryOption(query)
	
	if query.Range == nil {
		t.Fatal("范围查询为 nil")
	}
	
	rangeQuery := query.Range["age"]
	numberQuery := rangeQuery.(types.NumberRangeQuery)
	
	if numberQuery.Gte == nil || float64(*numberQuery.Gte) != 18.0 {
		t.Errorf("Gte = %v，期望 18.0", numberQuery.Gte)
	}
	
	if numberQuery.Lt == nil || float64(*numberQuery.Lt) != 65.0 {
		t.Errorf("Lt = %v，期望 65.0", numberQuery.Lt)
	}
}

func TestNumberRangeBuilder_FromToConditions(t *testing.T) {
	builder := NumberRange("price").From(10.0).To(100.0)
	query := &types.Query{}
	queryOption := builder.Build()
	
	queryOption(query)
	
	rangeQuery := query.Range["price"]
	numberQuery := rangeQuery.(types.NumberRangeQuery)
	
	if numberQuery.From == nil || float64(*numberQuery.From) != 10.0 {
		t.Errorf("From = %v，期望 10.0", numberQuery.From)
	}
	
	if numberQuery.To == nil || float64(*numberQuery.To) != 100.0 {
		t.Errorf("To = %v，期望 100.0", numberQuery.To)
	}
}

func TestNumberRangeBuilder_WithOptions(t *testing.T) {
	boost := float32(2.0)
	queryName := "test_query"
	relation := rangerelation.Within
	
	builder := NumberRange("score").Gte(50.0).Boost(boost).QueryName(queryName).Relation(&relation)
	query := &types.Query{}
	queryOption := builder.Build()
	
	queryOption(query)
	
	rangeQuery := query.Range["score"]
	numberQuery := rangeQuery.(types.NumberRangeQuery)
	
	if numberQuery.Boost == nil || *numberQuery.Boost != boost {
		t.Errorf("Boost = %v，期望 %v", numberQuery.Boost, boost)
	}
	
	if numberQuery.QueryName_ == nil || *numberQuery.QueryName_ != queryName {
		t.Errorf("QueryName = %v，期望 %v", numberQuery.QueryName_, queryName)
	}
	
	if numberQuery.Relation == nil || *numberQuery.Relation != relation {
		t.Errorf("Relation = %v，期望 %v", numberQuery.Relation, relation)
	}
}

// =============================================================================
// 日期范围构建器测试
// =============================================================================

func TestDateRangeBuilder_BasicUsage(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		builder  func() *DateRangeBuilder
		validate func(t *testing.T, query *types.Query)
	}{
		{
			name:  "大于等于条件",
			field: "created_at",
			builder: func() *DateRangeBuilder {
				return DateRange("created_at").Gte("2023-01-01")
			},
			validate: func(t *testing.T, query *types.Query) {
				rangeQuery := query.Range["created_at"]
				dateQuery := rangeQuery.(types.DateRangeQuery)
				if dateQuery.Gte == nil || *dateQuery.Gte != "2023-01-01" {
					t.Errorf("Gte = %v，期望 2023-01-01", dateQuery.Gte)
				}
			},
		},
		{
			name:  "大于条件",
			field: "timestamp",
			builder: func() *DateRangeBuilder {
				return DateRange("timestamp").Gt("now-1d")
			},
			validate: func(t *testing.T, query *types.Query) {
				rangeQuery := query.Range["timestamp"]
				dateQuery := rangeQuery.(types.DateRangeQuery)
				if dateQuery.Gt == nil || *dateQuery.Gt != "now-1d" {
					t.Errorf("Gt = %v，期望 now-1d", dateQuery.Gt)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &types.Query{}
			builder := tt.builder()
			queryOption := builder.Build()
			
			queryOption(query)
			
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
	
	queryOption(query)
	
	rangeQuery := query.Range["created_at"]
	dateQuery := rangeQuery.(types.DateRangeQuery)
	
	if dateQuery.Format == nil || *dateQuery.Format != format {
		t.Errorf("Format = %v，期望 %v", dateQuery.Format, format)
	}
	
	if dateQuery.TimeZone == nil || *dateQuery.TimeZone != timeZone {
		t.Errorf("TimeZone = %v，期望 %v", dateQuery.TimeZone, timeZone)
	}
}

func TestDateRangeBuilder_ChainedConditions(t *testing.T) {
	builder := DateRange("created_at").Gte("2023-01-01").Lt("2023-12-31")
	query := &types.Query{}
	queryOption := builder.Build()
	
	queryOption(query)
	
	rangeQuery := query.Range["created_at"]
	dateQuery := rangeQuery.(types.DateRangeQuery)
	
	if dateQuery.Gte == nil || *dateQuery.Gte != "2023-01-01" {
		t.Errorf("Gte = %v，期望 2023-01-01", dateQuery.Gte)
	}
	
	if dateQuery.Lt == nil || *dateQuery.Lt != "2023-12-31" {
		t.Errorf("Lt = %v，期望 2023-12-31", dateQuery.Lt)
	}
}

// =============================================================================
// 词项范围构建器测试
// =============================================================================

func TestTermRangeBuilder_BasicUsage(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		builder  func() *TermRangeBuilder
		validate func(t *testing.T, query *types.Query)
	}{
		{
			name:  "大于等于条件",
			field: "username",
			builder: func() *TermRangeBuilder {
				return TermRange("username").Gte("alice")
			},
			validate: func(t *testing.T, query *types.Query) {
				rangeQuery := query.Range["username"]
				termQuery := rangeQuery.(types.TermRangeQuery)
				if termQuery.Gte == nil || *termQuery.Gte != "alice" {
					t.Errorf("Gte = %v，期望 alice", termQuery.Gte)
				}
			},
		},
		{
			name:  "小于条件",
			field: "category",
			builder: func() *TermRangeBuilder {
				return TermRange("category").Lt("electronics")
			},
			validate: func(t *testing.T, query *types.Query) {
				rangeQuery := query.Range["category"]
				termQuery := rangeQuery.(types.TermRangeQuery)
				if termQuery.Lt == nil || *termQuery.Lt != "electronics" {
					t.Errorf("Lt = %v，期望 electronics", termQuery.Lt)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &types.Query{}
			builder := tt.builder()
			queryOption := builder.Build()
			
			queryOption(query)
			
			tt.validate(t, query)
		})
	}
}

func TestTermRangeBuilder_ChainedConditions(t *testing.T) {
	builder := TermRange("username").Gte("alice").Lt("bob")
	query := &types.Query{}
	queryOption := builder.Build()
	
	queryOption(query)
	
	rangeQuery := query.Range["username"]
	termQuery := rangeQuery.(types.TermRangeQuery)
	
	if termQuery.Gte == nil || *termQuery.Gte != "alice" {
		t.Errorf("Gte = %v，期望 alice", termQuery.Gte)
	}
	
	if termQuery.Lt == nil || *termQuery.Lt != "bob" {
		t.Errorf("Lt = %v，期望 bob", termQuery.Lt)
	}
}

func TestTermRangeBuilder_FromToConditions(t *testing.T) {
	builder := TermRange("category").From("electronics").To("home")
	query := &types.Query{}
	queryOption := builder.Build()
	
	queryOption(query)
	
	rangeQuery := query.Range["category"]
	termQuery := rangeQuery.(types.TermRangeQuery)
	
	if termQuery.From == nil || *termQuery.From != "electronics" {
		t.Errorf("From = %v，期望 electronics", termQuery.From)
	}
	
	if termQuery.To == nil || *termQuery.To != "home" {
		t.Errorf("To = %v，期望 home", termQuery.To)
	}
}

// =============================================================================
// 集成测试
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
	
	queryOption(query)
	
	if query.Bool == nil {
		t.Fatal("布尔查询为 nil")
	}
	
	if len(query.Bool.Must) != 3 {
		t.Errorf("Must 条件数量 = %d，期望 3", len(query.Bool.Must))
	}
	
	// 检查第一个条件（数值范围）
	firstCondition := query.Bool.Must[0]
	if firstCondition.Range == nil {
		t.Fatal("第一个条件的范围为 nil")
	}
	
	// 检查第二个条件（日期范围）
	secondCondition := query.Bool.Must[1]
	if secondCondition.Range == nil {
		t.Fatal("第二个条件的范围为 nil")
	}
	
	// 检查第三个条件（词项范围）
	thirdCondition := query.Bool.Must[2]
	if thirdCondition.Range == nil {
		t.Fatal("第三个条件的范围为 nil")
	}
}

// =============================================================================
// 基准测试
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

// TestRangeBuilderOptions 测试范围查询构建器的通用选项
func TestRangeBuilderOptions(t *testing.T) {
	t.Run("测试数值范围查询的通用选项", func(t *testing.T) {
		relation := rangerelation.Contains
		query := NewQuery(
			NumberRange("age").
				Gte(18).
				Lt(65).
				Boost(2.0).
				QueryName("age_range").
				Relation(&relation).
				Build(),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Range == nil {
			t.Fatal("Range查询不应该为nil")
		}
		
		rangeQuery, exists := query.Range["age"]
		if !exists {
			t.Fatal("期望存在age字段")
		}
		
		numberRange, ok := rangeQuery.(types.NumberRangeQuery)
		if !ok {
			t.Fatal("期望为NumberRangeQuery类型")
		}
		
		if *numberRange.Boost != 2.0 {
			t.Errorf("期望Boost为2.0, 实际得到: %f", *numberRange.Boost)
		}
		
		if *numberRange.QueryName_ != "age_range" {
			t.Errorf("期望QueryName为'age_range', 实际得到: %s", *numberRange.QueryName_)
		}
		
		if *numberRange.Relation != rangerelation.Contains {
			t.Errorf("期望Relation为Contains, 实际得到: %s", *numberRange.Relation)
		}
	})
	
	t.Run("测试日期范围查询的通用选项", func(t *testing.T) {
		relation := rangerelation.Within
		query := NewQuery(
			DateRange("created_at").
				Gte("2023-01-01").
				Lt("2024-01-01").
				Format("yyyy-MM-dd").
				TimeZone("UTC").
				Boost(1.5).
				QueryName("date_range").
				Relation(&relation).
				Build(),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Range == nil {
			t.Fatal("Range查询不应该为nil")
		}
		
		rangeQuery, exists := query.Range["created_at"]
		if !exists {
			t.Fatal("期望存在created_at字段")
		}
		
		dateRange, ok := rangeQuery.(types.DateRangeQuery)
		if !ok {
			t.Fatal("期望为DateRangeQuery类型")
		}
		
		if *dateRange.Format != "yyyy-MM-dd" {
			t.Errorf("期望Format为'yyyy-MM-dd', 实际得到: %s", *dateRange.Format)
		}
		
		if *dateRange.TimeZone != "UTC" {
			t.Errorf("期望TimeZone为'UTC', 实际得到: %s", *dateRange.TimeZone)
		}
		
		if *dateRange.Boost != 1.5 {
			t.Errorf("期望Boost为1.5, 实际得到: %f", *dateRange.Boost)
		}
		
		if *dateRange.QueryName_ != "date_range" {
			t.Errorf("期望QueryName为'date_range', 实际得到: %s", *dateRange.QueryName_)
		}
		
		if *dateRange.Relation != rangerelation.Within {
			t.Errorf("期望Relation为Within, 实际得到: %s", *dateRange.Relation)
		}
	})
	
	t.Run("测试词项范围查询的通用选项", func(t *testing.T) {
		relation := rangerelation.Intersects
		query := NewQuery(
			TermRange("username").
				From("alice").
				To("bob").
				Boost(1.2).
				QueryName("term_range").
				Relation(&relation).
				Build(),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Range == nil {
			t.Fatal("Range查询不应该为nil")
		}
		
		rangeQuery, exists := query.Range["username"]
		if !exists {
			t.Fatal("期望存在username字段")
		}
		
		termRange, ok := rangeQuery.(types.TermRangeQuery)
		if !ok {
			t.Fatal("期望为TermRangeQuery类型")
		}
		
		if *termRange.From != "alice" {
			t.Errorf("期望From为'alice', 实际得到: %s", *termRange.From)
		}
		
		if *termRange.To != "bob" {
			t.Errorf("期望To为'bob', 实际得到: %s", *termRange.To)
		}
		
		if *termRange.Boost != 1.2 {
			t.Errorf("期望Boost为1.2, 实际得到: %f", *termRange.Boost)
		}
		
		if *termRange.QueryName_ != "term_range" {
			t.Errorf("期望QueryName为'term_range', 实际得到: %s", *termRange.QueryName_)
		}
		
		if *termRange.Relation != rangerelation.Intersects {
			t.Errorf("期望Relation为Intersects, 实际得到: %s", *termRange.Relation)
		}
	})
}

// TestRangeBuilderEdgeCases 测试范围查询构建器的边界条件
func TestRangeBuilderEdgeCases(t *testing.T) {
	t.Run("测试数值范围查询的边界条件", func(t *testing.T) {
		query := NewQuery(
			NumberRange("value").
				Gt(0).
				Lt(0).
				Build(),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Range == nil {
			t.Fatal("Range查询不应该为nil")
		}
		
		rangeQuery, exists := query.Range["value"]
		if !exists {
			t.Fatal("期望存在value字段")
		}
		
		numberRange, ok := rangeQuery.(types.NumberRangeQuery)
		if !ok {
			t.Fatal("期望为NumberRangeQuery类型")
		}
		
		if *numberRange.Gt != 0 {
			t.Errorf("期望Gt为0, 实际得到: %f", *numberRange.Gt)
		}
		
		if *numberRange.Lt != 0 {
			t.Errorf("期望Lt为0, 实际得到: %f", *numberRange.Lt)
		}
	})
	
	t.Run("测试日期范围查询的边界条件", func(t *testing.T) {
		query := NewQuery(
			DateRange("timestamp").
				Gte("now").
				Lte("now").
				Build(),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Range == nil {
			t.Fatal("Range查询不应该为nil")
		}
		
		rangeQuery, exists := query.Range["timestamp"]
		if !exists {
			t.Fatal("期望存在timestamp字段")
		}
		
		dateRange, ok := rangeQuery.(types.DateRangeQuery)
		if !ok {
			t.Fatal("期望为DateRangeQuery类型")
		}
		
		if *dateRange.Gte != "now" {
			t.Errorf("期望Gte为'now', 实际得到: %s", *dateRange.Gte)
		}
		
		if *dateRange.Lte != "now" {
			t.Errorf("期望Lte为'now', 实际得到: %s", *dateRange.Lte)
		}
	})
	
	t.Run("测试词项范围查询的边界条件", func(t *testing.T) {
		query := NewQuery(
			TermRange("key").
				Gt("").
				Lt("").
				Build(),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Range == nil {
			t.Fatal("Range查询不应该为nil")
		}
		
		rangeQuery, exists := query.Range["key"]
		if !exists {
			t.Fatal("期望存在key字段")
		}
		
		termRange, ok := rangeQuery.(types.TermRangeQuery)
		if !ok {
			t.Fatal("期望为TermRangeQuery类型")
		}
		
		if *termRange.Gt != "" {
			t.Errorf("期望Gt为空字符串, 实际得到: %s", *termRange.Gt)
		}
		
		if *termRange.Lt != "" {
			t.Errorf("期望Lt为空字符串, 实际得到: %s", *termRange.Lt)
		}
	})
} 

// TestRangeBuilderMethods 测试范围查询构建器的所有方法
func TestRangeBuilderMethods(t *testing.T) {
	t.Run("测试数值范围查询的所有方法", func(t *testing.T) {
		query := NewQuery(
			NumberRange("score").
				Gte(60).
				Gt(50).
				Lte(100).
				Lt(90).
				From(0).
				To(150).
				Build(),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Range == nil {
			t.Fatal("Range查询不应该为nil")
		}
		
		rangeQuery, exists := query.Range["score"]
		if !exists {
			t.Fatal("期望存在score字段")
		}
		
		numberRange, ok := rangeQuery.(types.NumberRangeQuery)
		if !ok {
			t.Fatal("期望为NumberRangeQuery类型")
		}
		
		if *numberRange.Gte != 60 {
			t.Errorf("期望Gte为60, 实际得到: %f", *numberRange.Gte)
		}
		
		if *numberRange.Gt != 50 {
			t.Errorf("期望Gt为50, 实际得到: %f", *numberRange.Gt)
		}
		
		if *numberRange.Lte != 100 {
			t.Errorf("期望Lte为100, 实际得到: %f", *numberRange.Lte)
		}
		
		if *numberRange.Lt != 90 {
			t.Errorf("期望Lt为90, 实际得到: %f", *numberRange.Lt)
		}
		
		if *numberRange.From != 0 {
			t.Errorf("期望From为0, 实际得到: %f", *numberRange.From)
		}
		
		if *numberRange.To != 150 {
			t.Errorf("期望To为150, 实际得到: %f", *numberRange.To)
		}
	})
	
	t.Run("测试日期范围查询的所有方法", func(t *testing.T) {
		query := NewQuery(
			DateRange("timestamp").
				Gte("2023-01-01").
				Gt("2022-12-31").
				Lte("2023-12-31").
				Lt("2024-01-01").
				From("2023-01-01T00:00:00Z").
				To("2023-12-31T23:59:59Z").
				Format("yyyy-MM-dd'T'HH:mm:ssX").
				TimeZone("UTC").
				Build(),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Range == nil {
			t.Fatal("Range查询不应该为nil")
		}
		
		rangeQuery, exists := query.Range["timestamp"]
		if !exists {
			t.Fatal("期望存在timestamp字段")
		}
		
		dateRange, ok := rangeQuery.(types.DateRangeQuery)
		if !ok {
			t.Fatal("期望为DateRangeQuery类型")
		}
		
		if *dateRange.Gte != "2023-01-01" {
			t.Errorf("期望Gte为'2023-01-01', 实际得到: %s", *dateRange.Gte)
		}
		
		if *dateRange.Gt != "2022-12-31" {
			t.Errorf("期望Gt为'2022-12-31', 实际得到: %s", *dateRange.Gt)
		}
		
		if *dateRange.Lte != "2023-12-31" {
			t.Errorf("期望Lte为'2023-12-31', 实际得到: %s", *dateRange.Lte)
		}
		
		if *dateRange.Lt != "2024-01-01" {
			t.Errorf("期望Lt为'2024-01-01', 实际得到: %s", *dateRange.Lt)
		}
		
		if *dateRange.From != "2023-01-01T00:00:00Z" {
			t.Errorf("期望From为'2023-01-01T00:00:00Z', 实际得到: %s", *dateRange.From)
		}
		
		if *dateRange.To != "2023-12-31T23:59:59Z" {
			t.Errorf("期望To为'2023-12-31T23:59:59Z', 实际得到: %s", *dateRange.To)
		}
		
		if *dateRange.Format != "yyyy-MM-dd'T'HH:mm:ssX" {
			t.Errorf("期望Format为'yyyy-MM-dd'T'HH:mm:ssX', 实际得到: %s", *dateRange.Format)
		}
		
		if *dateRange.TimeZone != "UTC" {
			t.Errorf("期望TimeZone为'UTC', 实际得到: %s", *dateRange.TimeZone)
		}
	})
	
	t.Run("测试词项范围查询的所有方法", func(t *testing.T) {
		query := NewQuery(
			TermRange("category").
				Gte("electronics").
				Gt("computers").
				Lte("phones").
				Lt("tablets").
				From("a").
				To("z").
				Build(),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Range == nil {
			t.Fatal("Range查询不应该为nil")
		}
		
		rangeQuery, exists := query.Range["category"]
		if !exists {
			t.Fatal("期望存在category字段")
		}
		
		termRange, ok := rangeQuery.(types.TermRangeQuery)
		if !ok {
			t.Fatal("期望为TermRangeQuery类型")
		}
		
		if *termRange.Gte != "electronics" {
			t.Errorf("期望Gte为'electronics', 实际得到: %s", *termRange.Gte)
		}
		
		if *termRange.Gt != "computers" {
			t.Errorf("期望Gt为'computers', 实际得到: %s", *termRange.Gt)
		}
		
		if *termRange.Lte != "phones" {
			t.Errorf("期望Lte为'phones', 实际得到: %s", *termRange.Lte)
		}
		
		if *termRange.Lt != "tablets" {
			t.Errorf("期望Lt为'tablets', 实际得到: %s", *termRange.Lt)
		}
		
		if *termRange.From != "a" {
			t.Errorf("期望From为'a', 实际得到: %s", *termRange.From)
		}
		
		if *termRange.To != "z" {
			t.Errorf("期望To为'z', 实际得到: %s", *termRange.To)
		}
	})
}

// TestRangeBuilderCombinations 测试范围查询构建器的组合用法
func TestRangeBuilderCombinations(t *testing.T) {
	t.Run("测试数值范围查询的组合用法", func(t *testing.T) {
		relation := rangerelation.Contains
		query := NewQuery(
			NumberRange("price").
				Gte(10).
				Lt(100).
				Boost(2.0).
				QueryName("price_range").
				Relation(&relation).
				Build(),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Range == nil {
			t.Fatal("Range查询不应该为nil")
		}
		
		rangeQuery, exists := query.Range["price"]
		if !exists {
			t.Fatal("期望存在price字段")
		}
		
		numberRange, ok := rangeQuery.(types.NumberRangeQuery)
		if !ok {
			t.Fatal("期望为NumberRangeQuery类型")
		}
		
		if *numberRange.Gte != 10 {
			t.Errorf("期望Gte为10, 实际得到: %f", *numberRange.Gte)
		}
		
		if *numberRange.Lt != 100 {
			t.Errorf("期望Lt为100, 实际得到: %f", *numberRange.Lt)
		}
		
		if *numberRange.Boost != 2.0 {
			t.Errorf("期望Boost为2.0, 实际得到: %f", *numberRange.Boost)
		}
		
		if *numberRange.QueryName_ != "price_range" {
			t.Errorf("期望QueryName为'price_range', 实际得到: %s", *numberRange.QueryName_)
		}
		
		if *numberRange.Relation != rangerelation.Contains {
			t.Errorf("期望Relation为Contains, 实际得到: %s", *numberRange.Relation)
		}
	})
	
	t.Run("测试日期范围查询的组合用法", func(t *testing.T) {
		relation := rangerelation.Within
		query := NewQuery(
			DateRange("created_at").
				Gt("now-1d").
				Lt("now").
				Format("epoch_millis").
				TimeZone("America/New_York").
				Boost(1.5).
				QueryName("recent_docs").
				Relation(&relation).
				Build(),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Range == nil {
			t.Fatal("Range查询不应该为nil")
		}
		
		rangeQuery, exists := query.Range["created_at"]
		if !exists {
			t.Fatal("期望存在created_at字段")
		}
		
		dateRange, ok := rangeQuery.(types.DateRangeQuery)
		if !ok {
			t.Fatal("期望为DateRangeQuery类型")
		}
		
		if *dateRange.Gt != "now-1d" {
			t.Errorf("期望Gt为'now-1d', 实际得到: %s", *dateRange.Gt)
		}
		
		if *dateRange.Lt != "now" {
			t.Errorf("期望Lt为'now', 实际得到: %s", *dateRange.Lt)
		}
		
		if *dateRange.Format != "epoch_millis" {
			t.Errorf("期望Format为'epoch_millis', 实际得到: %s", *dateRange.Format)
		}
		
		if *dateRange.TimeZone != "America/New_York" {
			t.Errorf("期望TimeZone为'America/New_York', 实际得到: %s", *dateRange.TimeZone)
		}
		
		if *dateRange.Boost != 1.5 {
			t.Errorf("期望Boost为1.5, 实际得到: %f", *dateRange.Boost)
		}
		
		if *dateRange.QueryName_ != "recent_docs" {
			t.Errorf("期望QueryName为'recent_docs', 实际得到: %s", *dateRange.QueryName_)
		}
		
		if *dateRange.Relation != rangerelation.Within {
			t.Errorf("期望Relation为Within, 实际得到: %s", *dateRange.Relation)
		}
	})
	
	t.Run("测试词项范围查询的组合用法", func(t *testing.T) {
		relation := rangerelation.Intersects
		query := NewQuery(
			TermRange("status").
				From("active").
				To("completed").
				Boost(1.2).
				QueryName("status_range").
				Relation(&relation).
				Build(),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Range == nil {
			t.Fatal("Range查询不应该为nil")
		}
		
		rangeQuery, exists := query.Range["status"]
		if !exists {
			t.Fatal("期望存在status字段")
		}
		
		termRange, ok := rangeQuery.(types.TermRangeQuery)
		if !ok {
			t.Fatal("期望为TermRangeQuery类型")
		}
		
		if *termRange.From != "active" {
			t.Errorf("期望From为'active', 实际得到: %s", *termRange.From)
		}
		
		if *termRange.To != "completed" {
			t.Errorf("期望To为'completed', 实际得到: %s", *termRange.To)
		}
		
		if *termRange.Boost != 1.2 {
			t.Errorf("期望Boost为1.2, 实际得到: %f", *termRange.Boost)
		}
		
		if *termRange.QueryName_ != "status_range" {
			t.Errorf("期望QueryName为'status_range', 实际得到: %s", *termRange.QueryName_)
		}
		
		if *termRange.Relation != rangerelation.Intersects {
			t.Errorf("期望Relation为Intersects, 实际得到: %s", *termRange.Relation)
		}
	})
} 
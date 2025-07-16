package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestNewAggregations(t *testing.T) {
	t.Run("empty aggregations", func(t *testing.T) {
		aggs := NewAggregations()
		if aggs == nil {
			t.Fatal("Expected aggregations to not be nil")
		}
		if aggs.Aggregations == nil {
			t.Fatal("Expected aggregations map to not be nil")
		}
		if len(aggs.Aggregations) != 0 {
			t.Errorf("Expected empty aggregations, got %d", len(aggs.Aggregations))
		}
	})

	t.Run("with single aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			TermsAgg("categories", "category"),
		)
		if aggs == nil {
			t.Fatal("Expected aggregations to not be nil")
		}
		if len(aggs.Aggregations) != 1 {
			t.Errorf("Expected 1 aggregation, got %d", len(aggs.Aggregations))
		}
		if _, exists := aggs.Aggregations["categories"]; !exists {
			t.Error("Expected 'categories' aggregation to exist")
		}
	})

	t.Run("with multiple aggregations", func(t *testing.T) {
		aggs := NewAggregations(
			TermsAgg("categories", "category"),
			AvgAgg("avg_price", "price"),
			SumAgg("total_sales", "sales"),
		)
		if aggs == nil {
			t.Fatal("Expected aggregations to not be nil")
		}
		if len(aggs.Aggregations) != 3 {
			t.Errorf("Expected 3 aggregations, got %d", len(aggs.Aggregations))
		}
		expectedKeys := []string{"categories", "avg_price", "total_sales"}
		for _, key := range expectedKeys {
			if _, exists := aggs.Aggregations[key]; !exists {
				t.Errorf("Expected '%s' aggregation to exist", key)
			}
		}
	})
}

func TestTermsAgg(t *testing.T) {
	t.Run("basic terms aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			TermsAgg("categories", "category"),
		)
		
		categoryAgg := aggs.Aggregations["categories"]
		if categoryAgg.Terms == nil {
			t.Fatal("Expected Terms aggregation to not be nil")
		}
		if categoryAgg.Terms.Field == nil {
			t.Fatal("Expected Terms field to not be nil")
		}
		if *categoryAgg.Terms.Field != "category" {
			t.Errorf("Expected field to be 'category', got '%s'", *categoryAgg.Terms.Field)
		}
	})

	t.Run("terms aggregation with sub-aggregations", func(t *testing.T) {
		aggs := NewAggregations(
			TermsAgg("categories", "category",
				AvgAgg("avg_price", "price"),
				SumAgg("total_sales", "sales"),
			),
		)
		
		categoryAgg := aggs.Aggregations["categories"]
		if categoryAgg.Terms == nil {
			t.Fatal("Expected Terms aggregation to not be nil")
		}
		if categoryAgg.Terms.Field == nil {
			t.Fatal("Expected Terms field to not be nil")
		}
		if *categoryAgg.Terms.Field != "category" {
			t.Errorf("Expected field to be 'category', got '%s'", *categoryAgg.Terms.Field)
		}
		
		if categoryAgg.Aggregations == nil {
			t.Fatal("Expected sub-aggregations to not be nil")
		}
		if len(categoryAgg.Aggregations) != 2 {
			t.Errorf("Expected 2 sub-aggregations, got %d", len(categoryAgg.Aggregations))
		}
		expectedSubKeys := []string{"avg_price", "total_sales"}
		for _, key := range expectedSubKeys {
			if _, exists := categoryAgg.Aggregations[key]; !exists {
				t.Errorf("Expected '%s' sub-aggregation to exist", key)
			}
		}
	})
}

func TestMetricAggregations(t *testing.T) {
	t.Run("avg aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			AvgAgg("avg_price", "price"),
		)
		
		avgAgg := aggs.Aggregations["avg_price"]
		if avgAgg.Avg == nil {
			t.Fatal("Expected Avg aggregation to not be nil")
		}
		if avgAgg.Avg.Field == nil {
			t.Fatal("Expected Avg field to not be nil")
		}
		if *avgAgg.Avg.Field != "price" {
			t.Errorf("Expected field to be 'price', got '%s'", *avgAgg.Avg.Field)
		}
	})

	t.Run("sum aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			SumAgg("total_sales", "sales"),
		)
		
		sumAgg := aggs.Aggregations["total_sales"]
		if sumAgg.Sum == nil {
			t.Fatal("Expected Sum aggregation to not be nil")
		}
		if sumAgg.Sum.Field == nil {
			t.Fatal("Expected Sum field to not be nil")
		}
		if *sumAgg.Sum.Field != "sales" {
			t.Errorf("Expected field to be 'sales', got '%s'", *sumAgg.Sum.Field)
		}
	})

	t.Run("max aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			MaxAgg("max_price", "price"),
		)
		
		maxAgg := aggs.Aggregations["max_price"]
		if maxAgg.Max == nil {
			t.Fatal("Expected Max aggregation to not be nil")
		}
		if maxAgg.Max.Field == nil {
			t.Fatal("Expected Max field to not be nil")
		}
		if *maxAgg.Max.Field != "price" {
			t.Errorf("Expected field to be 'price', got '%s'", *maxAgg.Max.Field)
		}
	})

	t.Run("min aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			MinAgg("min_price", "price"),
		)
		
		minAgg := aggs.Aggregations["min_price"]
		if minAgg.Min == nil {
			t.Fatal("Expected Min aggregation to not be nil")
		}
		if minAgg.Min.Field == nil {
			t.Fatal("Expected Min field to not be nil")
		}
		if *minAgg.Min.Field != "price" {
			t.Errorf("Expected field to be 'price', got '%s'", *minAgg.Min.Field)
		}
	})

	t.Run("stats aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			StatsAgg("price_stats", "price"),
		)
		
		statsAgg := aggs.Aggregations["price_stats"]
		if statsAgg.Stats == nil {
			t.Fatal("Expected Stats aggregation to not be nil")
		}
		if statsAgg.Stats.Field == nil {
			t.Fatal("Expected Stats field to not be nil")
		}
		if *statsAgg.Stats.Field != "price" {
			t.Errorf("Expected field to be 'price', got '%s'", *statsAgg.Stats.Field)
		}
	})

	t.Run("value count aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			ValueCountAgg("field_count", "category"),
		)
		
		countAgg := aggs.Aggregations["field_count"]
		if countAgg.ValueCount == nil {
			t.Fatal("Expected ValueCount aggregation to not be nil")
		}
		if countAgg.ValueCount.Field == nil {
			t.Fatal("Expected ValueCount field to not be nil")
		}
		if *countAgg.ValueCount.Field != "category" {
			t.Errorf("Expected field to be 'category', got '%s'", *countAgg.ValueCount.Field)
		}
	})

	t.Run("cardinality aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			CardinalityAgg("unique_users", "user_id"),
		)
		
		cardinalityAgg := aggs.Aggregations["unique_users"]
		if cardinalityAgg.Cardinality == nil {
			t.Fatal("Expected Cardinality aggregation to not be nil")
		}
		if cardinalityAgg.Cardinality.Field == nil {
			t.Fatal("Expected Cardinality field to not be nil")
		}
		if *cardinalityAgg.Cardinality.Field != "user_id" {
			t.Errorf("Expected field to be 'user_id', got '%s'", *cardinalityAgg.Cardinality.Field)
		}
	})
}

func TestTermsAggWithOptions(t *testing.T) {
	t.Run("terms aggregation with size option", func(t *testing.T) {
		aggs := NewAggregations(
			TermsAggWithOptions("top_categories", "category", func(opts *types.TermsAggregation) {
				size := 10
				opts.Size = &size
			}),
		)
		
		categoryAgg := aggs.Aggregations["top_categories"]
		if categoryAgg.Terms == nil {
			t.Fatal("Expected Terms aggregation to not be nil")
		}
		if categoryAgg.Terms.Field == nil {
			t.Fatal("Expected Terms field to not be nil")
		}
		if *categoryAgg.Terms.Field != "category" {
			t.Errorf("Expected field to be 'category', got '%s'", *categoryAgg.Terms.Field)
		}
		if categoryAgg.Terms.Size == nil {
			t.Fatal("Expected Terms size to not be nil")
		}
		if *categoryAgg.Terms.Size != 10 {
			t.Errorf("Expected size to be 10, got %d", *categoryAgg.Terms.Size)
		}
	})

	t.Run("terms aggregation with options and sub-aggregations", func(t *testing.T) {
		aggs := NewAggregations(
			TermsAggWithOptions("top_categories", "category", func(opts *types.TermsAggregation) {
				size := 5
				opts.Size = &size
			}, AvgAgg("avg_price", "price")),
		)
		
		categoryAgg := aggs.Aggregations["top_categories"]
		if categoryAgg.Terms == nil {
			t.Fatal("Expected Terms aggregation to not be nil")
		}
		if categoryAgg.Terms.Size == nil || *categoryAgg.Terms.Size != 5 {
			t.Error("Expected size to be 5")
		}
		if categoryAgg.Aggregations == nil {
			t.Fatal("Expected sub-aggregations to not be nil")
		}
		if len(categoryAgg.Aggregations) != 1 {
			t.Errorf("Expected 1 sub-aggregation, got %d", len(categoryAgg.Aggregations))
		}
		if _, exists := categoryAgg.Aggregations["avg_price"]; !exists {
			t.Error("Expected 'avg_price' sub-aggregation to exist")
		}
	})

	t.Run("terms aggregation with nil options", func(t *testing.T) {
		aggs := NewAggregations(
			TermsAggWithOptions("categories", "category", nil),
		)
		
		categoryAgg := aggs.Aggregations["categories"]
		if categoryAgg.Terms == nil {
			t.Fatal("Expected Terms aggregation to not be nil")
		}
		if categoryAgg.Terms.Field == nil {
			t.Fatal("Expected Terms field to not be nil")
		}
		if *categoryAgg.Terms.Field != "category" {
			t.Errorf("Expected field to be 'category', got '%s'", *categoryAgg.Terms.Field)
		}
	})
}

func TestDateHistogramAgg(t *testing.T) {
	t.Run("basic date histogram aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			DateHistogramAgg("sales_over_time", "timestamp", "1d"),
		)
		
		dateHistAgg := aggs.Aggregations["sales_over_time"]
		if dateHistAgg.DateHistogram == nil {
			t.Fatal("Expected DateHistogram aggregation to not be nil")
		}
		if dateHistAgg.DateHistogram.Field == nil {
			t.Fatal("Expected DateHistogram field to not be nil")
		}
		if *dateHistAgg.DateHistogram.Field != "timestamp" {
			t.Errorf("Expected field to be 'timestamp', got '%s'", *dateHistAgg.DateHistogram.Field)
		}
		if dateHistAgg.DateHistogram.CalendarInterval == nil {
			t.Fatal("Expected CalendarInterval to not be nil")
		}
		if dateHistAgg.DateHistogram.CalendarInterval.Name != "1d" {
			t.Errorf("Expected interval to be '1d', got '%s'", dateHistAgg.DateHistogram.CalendarInterval.Name)
		}
	})

	t.Run("date histogram with sub-aggregations", func(t *testing.T) {
		aggs := NewAggregations(
			DateHistogramAgg("sales_over_time", "timestamp", "1d",
				SumAgg("total_sales", "amount"),
				AvgAgg("avg_price", "price"),
			),
		)
		
		dateHistAgg := aggs.Aggregations["sales_over_time"]
		if dateHistAgg.DateHistogram == nil {
			t.Fatal("Expected DateHistogram aggregation to not be nil")
		}
		if dateHistAgg.Aggregations == nil {
			t.Fatal("Expected sub-aggregations to not be nil")
		}
		if len(dateHistAgg.Aggregations) != 2 {
			t.Errorf("Expected 2 sub-aggregations, got %d", len(dateHistAgg.Aggregations))
		}
		expectedSubKeys := []string{"total_sales", "avg_price"}
		for _, key := range expectedSubKeys {
			if _, exists := dateHistAgg.Aggregations[key]; !exists {
				t.Errorf("Expected '%s' sub-aggregation to exist", key)
			}
		}
	})
}

func TestDateHistogramAggWithOptions(t *testing.T) {
	t.Run("date histogram with format option", func(t *testing.T) {
		aggs := NewAggregations(
			DateHistogramAggWithOptions("sales_over_time", "timestamp", "1d", func(opts *types.DateHistogramAggregation) {
				format := "yyyy-MM-dd"
				opts.Format = &format
			}),
		)
		
		dateHistAgg := aggs.Aggregations["sales_over_time"]
		if dateHistAgg.DateHistogram == nil {
			t.Fatal("Expected DateHistogram aggregation to not be nil")
		}
		if dateHistAgg.DateHistogram.Format == nil {
			t.Fatal("Expected Format to not be nil")
		}
		if *dateHistAgg.DateHistogram.Format != "yyyy-MM-dd" {
			t.Errorf("Expected format to be 'yyyy-MM-dd', got '%s'", *dateHistAgg.DateHistogram.Format)
		}
	})

	t.Run("date histogram with nil options", func(t *testing.T) {
		aggs := NewAggregations(
			DateHistogramAggWithOptions("sales_over_time", "timestamp", "1d", nil),
		)
		
		dateHistAgg := aggs.Aggregations["sales_over_time"]
		if dateHistAgg.DateHistogram == nil {
			t.Fatal("Expected DateHistogram aggregation to not be nil")
		}
		if dateHistAgg.DateHistogram.Field == nil {
			t.Fatal("Expected DateHistogram field to not be nil")
		}
		if *dateHistAgg.DateHistogram.Field != "timestamp" {
			t.Errorf("Expected field to be 'timestamp', got '%s'", *dateHistAgg.DateHistogram.Field)
		}
	})
}

func TestHistogramAgg(t *testing.T) {
	t.Run("basic histogram aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			HistogramAgg("price_ranges", "price", 100),
		)
		
		histAgg := aggs.Aggregations["price_ranges"]
		if histAgg.Histogram == nil {
			t.Fatal("Expected Histogram aggregation to not be nil")
		}
		if histAgg.Histogram.Field == nil {
			t.Fatal("Expected Histogram field to not be nil")
		}
		if *histAgg.Histogram.Field != "price" {
			t.Errorf("Expected field to be 'price', got '%s'", *histAgg.Histogram.Field)
		}
		if histAgg.Histogram.Interval == nil {
			t.Fatal("Expected Interval to not be nil")
		}
		if *histAgg.Histogram.Interval != 100 {
			t.Errorf("Expected interval to be 100, got %f", *histAgg.Histogram.Interval)
		}
	})
}

func TestRangeAgg(t *testing.T) {
	t.Run("basic range aggregation", func(t *testing.T) {
		to1 := types.Float64(100)
		from2 := types.Float64(100)
		to2 := types.Float64(500)
		from3 := types.Float64(500)
		ranges := []types.AggregationRange{
			{To: &to1},
			{From: &from2, To: &to2},
			{From: &from3},
		}
		
		aggs := NewAggregations(
			RangeAgg("price_ranges", "price", ranges),
		)
		
		rangeAgg := aggs.Aggregations["price_ranges"]
		if rangeAgg.Range == nil {
			t.Fatal("Expected Range aggregation to not be nil")
		}
		if rangeAgg.Range.Field == nil {
			t.Fatal("Expected Range field to not be nil")
		}
		if *rangeAgg.Range.Field != "price" {
			t.Errorf("Expected field to be 'price', got '%s'", *rangeAgg.Range.Field)
		}
		if len(rangeAgg.Range.Ranges) != 3 {
			t.Errorf("Expected 3 ranges, got %d", len(rangeAgg.Range.Ranges))
		}
	})
}

func TestFilterAgg(t *testing.T) {
	t.Run("basic filter aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			FilterAgg("expensive_products", NumberRange("price").Gte(1000).Build()),
		)
		
		filterAgg := aggs.Aggregations["expensive_products"]
		if filterAgg.Filter == nil {
			t.Fatal("Expected Filter aggregation to not be nil")
		}
	})

	t.Run("filter aggregation with nil query", func(t *testing.T) {
		aggs := NewAggregations(
			FilterAgg("all_products", nil),
		)
		
		filterAgg := aggs.Aggregations["all_products"]
		if filterAgg.Filter == nil {
			t.Fatal("Expected Filter aggregation to not be nil")
		}
	})
}

func TestFiltersAgg(t *testing.T) {
	t.Run("basic filters aggregation", func(t *testing.T) {
		filters := map[string]QueryOption{
			"electronics": Term("category", "electronics"),
			"books":       Term("category", "books"),
		}
		
		aggs := NewAggregations(
			FiltersAgg("product_categories", filters),
		)
		
		filtersAgg := aggs.Aggregations["product_categories"]
		if filtersAgg.Filters == nil {
			t.Fatal("Expected Filters aggregation to not be nil")
		}
		if filtersAgg.Filters.Filters == nil {
			t.Fatal("Expected Filters.Filters to not be nil")
		}
		keyedFilters, ok := filtersAgg.Filters.Filters.(map[string]*types.Query)
		if !ok {
			t.Fatal("Expected Keyed filters to be map[string]*types.Query")
		}
		if len(keyedFilters) != 2 {
			t.Errorf("Expected 2 filters, got %d", len(keyedFilters))
		}
		expectedKeys := []string{"electronics", "books"}
		for _, key := range expectedKeys {
			if _, exists := keyedFilters[key]; !exists {
				t.Errorf("Expected '%s' filter to exist", key)
			}
		}
	})

	t.Run("filters aggregation with nil query", func(t *testing.T) {
		filters := map[string]QueryOption{
			"all": nil,
		}
		
		aggs := NewAggregations(
			FiltersAgg("categories", filters),
		)
		
		filtersAgg := aggs.Aggregations["categories"]
		if filtersAgg.Filters == nil {
			t.Fatal("Expected Filters aggregation to not be nil")
		}
		keyedFilters2, ok2 := filtersAgg.Filters.Filters.(map[string]*types.Query)
		if !ok2 {
			t.Fatal("Expected Keyed filters to be map[string]*types.Query")
		}
		if len(keyedFilters2) != 1 {
			t.Errorf("Expected 1 filter, got %d", len(keyedFilters2))
		}
	})
}

func TestNestedAgg(t *testing.T) {
	t.Run("basic nested aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			NestedAgg("nested_products", "products"),
		)
		
		nestedAgg := aggs.Aggregations["nested_products"]
		if nestedAgg.Nested == nil {
			t.Fatal("Expected Nested aggregation to not be nil")
		}
		if nestedAgg.Nested.Path == nil || *nestedAgg.Nested.Path != "products" {
			if nestedAgg.Nested.Path == nil {
				t.Error("Expected path to be 'products', got nil")
			} else {
				t.Errorf("Expected path to be 'products', got '%s'", *nestedAgg.Nested.Path)
			}
		}
	})
}

func TestGlobalAgg(t *testing.T) {
	t.Run("basic global aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			GlobalAgg("all_docs"),
		)
		
		globalAgg := aggs.Aggregations["all_docs"]
		if globalAgg.Global == nil {
			t.Fatal("Expected Global aggregation to not be nil")
		}
	})
}

func TestSubAgg(t *testing.T) {
	t.Run("add sub-aggregation to existing parent", func(t *testing.T) {
		aggs := NewAggregations(
			TermsAgg("categories", "category"),
		)
		
		// 添加子聚合
		SubAgg("categories", AvgAgg("avg_price", "price"))(aggs)
		
		categoryAgg := aggs.Aggregations["categories"]
		if categoryAgg.Aggregations == nil {
			t.Fatal("Expected sub-aggregations to not be nil")
		}
		if len(categoryAgg.Aggregations) != 1 {
			t.Errorf("Expected 1 sub-aggregation, got %d", len(categoryAgg.Aggregations))
		}
		if _, exists := categoryAgg.Aggregations["avg_price"]; !exists {
			t.Error("Expected 'avg_price' sub-aggregation to exist")
		}
	})

	t.Run("add sub-aggregation to non-existing parent", func(t *testing.T) {
		aggs := NewAggregations()
		
		// 添加子聚合到不存在的父聚合
		SubAgg("non_existing", AvgAgg("avg_price", "price"))(aggs)
		
		parentAgg := aggs.Aggregations["non_existing"]
		if parentAgg.Aggregations == nil {
			t.Fatal("Expected sub-aggregations to not be nil")
		}
		if len(parentAgg.Aggregations) != 1 {
			t.Errorf("Expected 1 sub-aggregation, got %d", len(parentAgg.Aggregations))
		}
		if _, exists := parentAgg.Aggregations["avg_price"]; !exists {
			t.Error("Expected 'avg_price' sub-aggregation to exist")
		}
	})

	t.Run("add multiple sub-aggregations", func(t *testing.T) {
		aggs := NewAggregations(
			TermsAgg("categories", "category"),
		)
		
		// 添加多个子聚合
		SubAgg("categories", 
			AvgAgg("avg_price", "price"),
			SumAgg("total_sales", "sales"),
			MaxAgg("max_price", "price"),
		)(aggs)
		
		categoryAgg := aggs.Aggregations["categories"]
		if categoryAgg.Aggregations == nil {
			t.Fatal("Expected sub-aggregations to not be nil")
		}
		if len(categoryAgg.Aggregations) != 3 {
			t.Errorf("Expected 3 sub-aggregations, got %d", len(categoryAgg.Aggregations))
		}
		expectedSubKeys := []string{"avg_price", "total_sales", "max_price"}
		for _, key := range expectedSubKeys {
			if _, exists := categoryAgg.Aggregations[key]; !exists {
				t.Errorf("Expected '%s' sub-aggregation to exist", key)
			}
		}
	})
}

func TestTopTermsAgg(t *testing.T) {
	t.Run("basic top terms aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			TopTermsAgg("top_categories", "category", 10),
		)
		
		categoryAgg := aggs.Aggregations["top_categories"]
		if categoryAgg.Terms == nil {
			t.Fatal("Expected Terms aggregation to not be nil")
		}
		if categoryAgg.Terms.Field == nil {
			t.Fatal("Expected Terms field to not be nil")
		}
		if *categoryAgg.Terms.Field != "category" {
			t.Errorf("Expected field to be 'category', got '%s'", *categoryAgg.Terms.Field)
		}
		if categoryAgg.Terms.Size == nil {
			t.Fatal("Expected Terms size to not be nil")
		}
		if *categoryAgg.Terms.Size != 10 {
			t.Errorf("Expected size to be 10, got %d", *categoryAgg.Terms.Size)
		}
	})

	t.Run("top terms aggregation with sub-aggregations", func(t *testing.T) {
		aggs := NewAggregations(
			TopTermsAgg("top_categories", "category", 5,
				AvgAgg("avg_price", "price"),
				SumAgg("total_sales", "sales"),
			),
		)
		
		categoryAgg := aggs.Aggregations["top_categories"]
		if categoryAgg.Terms == nil {
			t.Fatal("Expected Terms aggregation to not be nil")
		}
		if categoryAgg.Terms.Size == nil || *categoryAgg.Terms.Size != 5 {
			t.Error("Expected size to be 5")
		}
		if categoryAgg.Aggregations == nil {
			t.Fatal("Expected sub-aggregations to not be nil")
		}
		if len(categoryAgg.Aggregations) != 2 {
			t.Errorf("Expected 2 sub-aggregations, got %d", len(categoryAgg.Aggregations))
		}
	})
}

func TestDailyHistogramAgg(t *testing.T) {
	t.Run("basic daily histogram aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			DailyHistogramAgg("daily_sales", "timestamp"),
		)
		
		dateHistAgg := aggs.Aggregations["daily_sales"]
		if dateHistAgg.DateHistogram == nil {
			t.Fatal("Expected DateHistogram aggregation to not be nil")
		}
		if dateHistAgg.DateHistogram.Field == nil {
			t.Fatal("Expected DateHistogram field to not be nil")
		}
		if *dateHistAgg.DateHistogram.Field != "timestamp" {
			t.Errorf("Expected field to be 'timestamp', got '%s'", *dateHistAgg.DateHistogram.Field)
		}
		if dateHistAgg.DateHistogram.CalendarInterval == nil {
			t.Fatal("Expected CalendarInterval to not be nil")
		}
		if dateHistAgg.DateHistogram.CalendarInterval.Name != "1d" {
			t.Errorf("Expected interval to be '1d', got '%s'", dateHistAgg.DateHistogram.CalendarInterval.Name)
		}
	})

	t.Run("daily histogram with sub-aggregations", func(t *testing.T) {
		aggs := NewAggregations(
			DailyHistogramAgg("daily_sales", "timestamp",
				SumAgg("total", "amount"),
			),
		)
		
		dateHistAgg := aggs.Aggregations["daily_sales"]
		if dateHistAgg.Aggregations == nil {
			t.Fatal("Expected sub-aggregations to not be nil")
		}
		if len(dateHistAgg.Aggregations) != 1 {
			t.Errorf("Expected 1 sub-aggregation, got %d", len(dateHistAgg.Aggregations))
		}
		if _, exists := dateHistAgg.Aggregations["total"]; !exists {
			t.Error("Expected 'total' sub-aggregation to exist")
		}
	})
}

func TestMonthlyHistogramAgg(t *testing.T) {
	t.Run("basic monthly histogram aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			MonthlyHistogramAgg("monthly_sales", "timestamp"),
		)
		
		dateHistAgg := aggs.Aggregations["monthly_sales"]
		if dateHistAgg.DateHistogram == nil {
			t.Fatal("Expected DateHistogram aggregation to not be nil")
		}
		if dateHistAgg.DateHistogram.Field == nil {
			t.Fatal("Expected DateHistogram field to not be nil")
		}
		if *dateHistAgg.DateHistogram.Field != "timestamp" {
			t.Errorf("Expected field to be 'timestamp', got '%s'", *dateHistAgg.DateHistogram.Field)
		}
		if dateHistAgg.DateHistogram.CalendarInterval == nil {
			t.Fatal("Expected CalendarInterval to not be nil")
		}
		if dateHistAgg.DateHistogram.CalendarInterval.Name != "1M" {
			t.Errorf("Expected interval to be '1M', got '%s'", dateHistAgg.DateHistogram.CalendarInterval.Name)
		}
	})
}

func TestPriceRangeAgg(t *testing.T) {
	t.Run("basic price range aggregation", func(t *testing.T) {
		boundaries := []float64{0, 100, 500, 1000}
		aggs := NewAggregations(
			PriceRangeAgg("price_segments", "price", boundaries),
		)
		
		rangeAgg := aggs.Aggregations["price_segments"]
		if rangeAgg.Range == nil {
			t.Fatal("Expected Range aggregation to not be nil")
		}
		if rangeAgg.Range.Field == nil {
			t.Fatal("Expected Range field to not be nil")
		}
		if *rangeAgg.Range.Field != "price" {
			t.Errorf("Expected field to be 'price', got '%s'", *rangeAgg.Range.Field)
		}
		if len(rangeAgg.Range.Ranges) != 4 {
			t.Errorf("Expected 4 ranges, got %d", len(rangeAgg.Range.Ranges))
		}
		
		// 验证第一个范围 (< 0)
		firstRange := rangeAgg.Range.Ranges[0]
		if firstRange.To == nil || *firstRange.To != 0 {
			t.Error("Expected first range to have To=0")
		}
		if firstRange.From != nil {
			t.Error("Expected first range to have no From")
		}
		
		// 验证最后一个范围 (>= 1000)
		lastRange := rangeAgg.Range.Ranges[3]
		if lastRange.From == nil || *lastRange.From != 1000 {
			t.Error("Expected last range to have From=1000")
		}
		if lastRange.To != nil {
			t.Error("Expected last range to have no To")
		}
	})

	t.Run("single boundary price range", func(t *testing.T) {
		boundaries := []float64{100}
		aggs := NewAggregations(
			PriceRangeAgg("price_segments", "price", boundaries),
		)
		
		rangeAgg := aggs.Aggregations["price_segments"]
		if len(rangeAgg.Range.Ranges) != 1 {
			t.Errorf("Expected 1 range, got %d", len(rangeAgg.Range.Ranges))
		}
	})

	t.Run("empty boundaries", func(t *testing.T) {
		boundaries := []float64{}
		aggs := NewAggregations(
			PriceRangeAgg("price_segments", "price", boundaries),
		)
		
		rangeAgg := aggs.Aggregations["price_segments"]
		if len(rangeAgg.Range.Ranges) != 0 {
			t.Errorf("Expected 0 ranges, got %d", len(rangeAgg.Range.Ranges))
		}
	})
}

func TestPercentilesAgg(t *testing.T) {
	t.Run("basic percentiles aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			PercentilesAgg("response_time_percentiles", "response_time"),
		)
		
		percentilesAgg := aggs.Aggregations["response_time_percentiles"]
		if percentilesAgg.Percentiles == nil {
			t.Fatal("Expected Percentiles aggregation to not be nil")
		}
		if percentilesAgg.Percentiles.Field == nil {
			t.Fatal("Expected Percentiles field to not be nil")
		}
		if *percentilesAgg.Percentiles.Field != "response_time" {
			t.Errorf("Expected field to be 'response_time', got '%s'", *percentilesAgg.Percentiles.Field)
		}
	})

	t.Run("percentiles aggregation with custom percentiles", func(t *testing.T) {
		customPercentiles := []float64{25, 50, 75, 95, 99}
		aggs := NewAggregations(
			PercentilesAgg("price_percentiles", "price", customPercentiles),
		)
		
		percentilesAgg := aggs.Aggregations["price_percentiles"]
		if percentilesAgg.Percentiles == nil {
			t.Fatal("Expected Percentiles aggregation to not be nil")
		}
		if percentilesAgg.Percentiles.Field == nil {
			t.Fatal("Expected Percentiles field to not be nil")
		}
		if *percentilesAgg.Percentiles.Field != "price" {
			t.Errorf("Expected field to be 'price', got '%s'", *percentilesAgg.Percentiles.Field)
		}
		if len(percentilesAgg.Percentiles.Percents) != 5 {
			t.Errorf("Expected 5 percentiles, got %d", len(percentilesAgg.Percentiles.Percents))
		}
		expectedPercentiles := []float64{25, 50, 75, 95, 99}
		for i, expected := range expectedPercentiles {
			if float64(percentilesAgg.Percentiles.Percents[i]) != expected {
				t.Errorf("Expected percentile %d to be %f, got %f", i, expected, float64(percentilesAgg.Percentiles.Percents[i]))
			}
		}
	})
}

func TestExtendedStatsAgg(t *testing.T) {
	t.Run("basic extended stats aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			ExtendedStatsAgg("price_extended_stats", "price"),
		)
		
		extendedStatsAgg := aggs.Aggregations["price_extended_stats"]
		if extendedStatsAgg.ExtendedStats == nil {
			t.Fatal("Expected ExtendedStats aggregation to not be nil")
		}
		if extendedStatsAgg.ExtendedStats.Field == nil {
			t.Fatal("Expected ExtendedStats field to not be nil")
		}
		if *extendedStatsAgg.ExtendedStats.Field != "price" {
			t.Errorf("Expected field to be 'price', got '%s'", *extendedStatsAgg.ExtendedStats.Field)
		}
	})
}

func TestTopHitsAgg(t *testing.T) {
	t.Run("basic top hits aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			TopHitsAgg("top_products", 3),
		)
		
		topHitsAgg := aggs.Aggregations["top_products"]
		if topHitsAgg.TopHits == nil {
			t.Fatal("Expected TopHits aggregation to not be nil")
		}
		if topHitsAgg.TopHits.Size == nil {
			t.Fatal("Expected TopHits size to not be nil")
		}
		if *topHitsAgg.TopHits.Size != 3 {
			t.Errorf("Expected size to be 3, got %d", *topHitsAgg.TopHits.Size)
		}
	})
}

func TestSignificantTermsAgg(t *testing.T) {
	t.Run("basic significant terms aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			SignificantTermsAgg("significant_tags", "tags"),
		)
		
		significantTermsAgg := aggs.Aggregations["significant_tags"]
		if significantTermsAgg.SignificantTerms == nil {
			t.Fatal("Expected SignificantTerms aggregation to not be nil")
		}
		if significantTermsAgg.SignificantTerms.Field == nil {
			t.Fatal("Expected SignificantTerms field to not be nil")
		}
		if *significantTermsAgg.SignificantTerms.Field != "tags" {
			t.Errorf("Expected field to be 'tags', got '%s'", *significantTermsAgg.SignificantTerms.Field)
		}
	})
}

func TestGeoBoundsAgg(t *testing.T) {
	t.Run("basic geo bounds aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			GeoBoundsAgg("viewport", "location"),
		)
		
		geoBoundsAgg := aggs.Aggregations["viewport"]
		if geoBoundsAgg.GeoBounds == nil {
			t.Fatal("Expected GeoBounds aggregation to not be nil")
		}
		if geoBoundsAgg.GeoBounds.Field == nil {
			t.Fatal("Expected GeoBounds field to not be nil")
		}
		if *geoBoundsAgg.GeoBounds.Field != "location" {
			t.Errorf("Expected field to be 'location', got '%s'", *geoBoundsAgg.GeoBounds.Field)
		}
	})
}

func TestGeoCentroidAgg(t *testing.T) {
	t.Run("basic geo centroid aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			GeoCentroidAgg("centroid", "location"),
		)
		
		geoCentroidAgg := aggs.Aggregations["centroid"]
		if geoCentroidAgg.GeoCentroid == nil {
			t.Fatal("Expected GeoCentroid aggregation to not be nil")
		}
		if geoCentroidAgg.GeoCentroid.Field == nil {
			t.Fatal("Expected GeoCentroid field to not be nil")
		}
		if *geoCentroidAgg.GeoCentroid.Field != "location" {
			t.Errorf("Expected field to be 'location', got '%s'", *geoCentroidAgg.GeoCentroid.Field)
		}
	})
}

func TestGeoDistanceAgg(t *testing.T) {
	t.Run("basic geo distance aggregation", func(t *testing.T) {
		rangeKeys := []string{"0-1km", "1-5km", "5-10km", "10km+"}
		distances := []float64{1000, 5000, 10000}
		
		aggs := NewAggregations(
			GeoDistanceAgg("distance_ranges", "location", "40.7128,-74.0060", rangeKeys, distances),
		)
		
		geoDistanceAgg := aggs.Aggregations["distance_ranges"]
		if geoDistanceAgg.GeoDistance == nil {
			t.Fatal("Expected GeoDistance aggregation to not be nil")
		}
		if geoDistanceAgg.GeoDistance.Field == nil {
			t.Fatal("Expected GeoDistance field to not be nil")
		}
		if *geoDistanceAgg.GeoDistance.Field != "location" {
			t.Errorf("Expected field to be 'location', got '%s'", *geoDistanceAgg.GeoDistance.Field)
		}
		if geoDistanceAgg.GeoDistance.Origin == nil {
			t.Fatal("Expected GeoDistance origin to not be nil")
		}
		if originStr, ok := geoDistanceAgg.GeoDistance.Origin.(string); !ok || originStr != "40.7128,-74.0060" {
			t.Errorf("Expected origin to be '40.7128,-74.0060', got '%v'", geoDistanceAgg.GeoDistance.Origin)
		}
		if len(geoDistanceAgg.GeoDistance.Ranges) != 4 {
			t.Errorf("Expected 4 ranges, got %d", len(geoDistanceAgg.GeoDistance.Ranges))
		}
		
		// 验证第一个范围 (0-1000)
		firstRange := geoDistanceAgg.GeoDistance.Ranges[0]
		if firstRange.To == nil || *firstRange.To != 1000 {
			t.Error("Expected first range to have To=1000")
		}
		if firstRange.Key == nil || *firstRange.Key != "0-1km" {
			t.Error("Expected first range to have Key='0-1km'")
		}
		
		// 验证最后一个范围 (>= 10000)
		lastRange := geoDistanceAgg.GeoDistance.Ranges[3]
		if lastRange.From == nil || *lastRange.From != 10000 {
			t.Error("Expected last range to have From=10000")
		}
		if lastRange.Key == nil || *lastRange.Key != "10km+" {
			t.Error("Expected last range to have Key='10km+'")
		}
	})

	t.Run("geo distance with empty distances", func(t *testing.T) {
		rangeKeys := []string{}
		distances := []float64{}
		
		aggs := NewAggregations(
			GeoDistanceAgg("distance_ranges", "location", "40.7128,-74.0060", rangeKeys, distances),
		)
		
		geoDistanceAgg := aggs.Aggregations["distance_ranges"]
		if len(geoDistanceAgg.GeoDistance.Ranges) != 1 {
			t.Errorf("Expected 1 range, got %d", len(geoDistanceAgg.GeoDistance.Ranges))
		}
	})
}

func TestComplexAggregationCombinations(t *testing.T) {
	t.Run("complex nested aggregations", func(t *testing.T) {
		aggs := NewAggregations(
			TermsAgg("categories", "category",
				AvgAgg("avg_price", "price"),
				DateHistogramAgg("sales_over_time", "timestamp", "1d",
					SumAgg("daily_sales", "amount"),
				),
				TopTermsAgg("top_brands", "brand", 5,
					MaxAgg("max_price", "price"),
					MinAgg("min_price", "price"),
				),
			),
			GlobalAgg("all_stats"),
			FilterAgg("expensive_items", NumberRange("price").Gte(1000).Build()),
		)
		
		// 验证主聚合数量
		if len(aggs.Aggregations) != 3 {
			t.Errorf("Expected 3 main aggregations, got %d", len(aggs.Aggregations))
		}
		
		// 验证 categories 聚合及其子聚合
		categoryAgg := aggs.Aggregations["categories"]
		if categoryAgg.Terms == nil {
			t.Fatal("Expected Terms aggregation to not be nil")
		}
		if len(categoryAgg.Aggregations) != 3 {
			t.Errorf("Expected 3 sub-aggregations in categories, got %d", len(categoryAgg.Aggregations))
		}
		
		// 验证嵌套的日期直方图聚合
		dateHistAgg := categoryAgg.Aggregations["sales_over_time"]
		if dateHistAgg.DateHistogram == nil {
			t.Fatal("Expected DateHistogram aggregation to not be nil")
		}
		if len(dateHistAgg.Aggregations) != 1 {
			t.Errorf("Expected 1 sub-aggregation in date histogram, got %d", len(dateHistAgg.Aggregations))
		}
		
		// 验证嵌套的 top terms 聚合
		topBrandsAgg := categoryAgg.Aggregations["top_brands"]
		if topBrandsAgg.Terms == nil {
			t.Fatal("Expected Terms aggregation to not be nil")
		}
		if len(topBrandsAgg.Aggregations) != 2 {
			t.Errorf("Expected 2 sub-aggregations in top brands, got %d", len(topBrandsAgg.Aggregations))
		}
	})

	t.Run("multiple metric aggregations", func(t *testing.T) {
		aggs := NewAggregations(
			AvgAgg("avg_price", "price"),
			SumAgg("total_sales", "sales"),
			MaxAgg("max_price", "price"),
			MinAgg("min_price", "price"),
			StatsAgg("price_stats", "price"),
			ValueCountAgg("product_count", "product_id"),
			CardinalityAgg("unique_categories", "category"),
			ExtendedStatsAgg("extended_price_stats", "price"),
			PercentilesAgg("price_percentiles", "price"),
		)
		
		if len(aggs.Aggregations) != 9 {
			t.Errorf("Expected 9 aggregations, got %d", len(aggs.Aggregations))
		}
		
		expectedKeys := []string{
			"avg_price", "total_sales", "max_price", "min_price", 
			"price_stats", "product_count", "unique_categories", 
			"extended_price_stats", "price_percentiles",
		}
		for _, key := range expectedKeys {
			if _, exists := aggs.Aggregations[key]; !exists {
				t.Errorf("Expected '%s' aggregation to exist", key)
			}
		}
	})
}

// 测试新增的聚合类型

func TestDateRangeAgg(t *testing.T) {
	t.Run("basic date range aggregation", func(t *testing.T) {
		ranges := []types.DateRangeExpression{
			{To: "2023-01-01"},
			{From: "2023-01-01", To: "2023-12-31"},
			{From: "2023-12-31"},
		}
		
		aggs := NewAggregations(
			DateRangeAgg("date_ranges", "timestamp", ranges),
		)
		
		dateRangeAgg := aggs.Aggregations["date_ranges"]
		if dateRangeAgg.DateRange == nil {
			t.Fatal("Expected DateRange aggregation to not be nil")
		}
		if dateRangeAgg.DateRange.Field == nil {
			t.Fatal("Expected DateRange field to not be nil")
		}
		if *dateRangeAgg.DateRange.Field != "timestamp" {
			t.Errorf("Expected field to be 'timestamp', got '%s'", *dateRangeAgg.DateRange.Field)
		}
		if len(dateRangeAgg.DateRange.Ranges) != 3 {
			t.Errorf("Expected 3 ranges, got %d", len(dateRangeAgg.DateRange.Ranges))
		}
	})
}

func TestIpRangeAgg(t *testing.T) {
	t.Run("basic ip range aggregation", func(t *testing.T) {
		to1 := "192.168.1.0/24"
		from2 := "10.0.0.0/8"
		ranges := []types.IpRangeAggregationRange{
			{To: &to1},
			{From: &from2},
		}
		
		aggs := NewAggregations(
			IpRangeAgg("ip_ranges", "client_ip", ranges),
		)
		
		ipRangeAgg := aggs.Aggregations["ip_ranges"]
		if ipRangeAgg.IpRange == nil {
			t.Fatal("Expected IpRange aggregation to not be nil")
		}
		if ipRangeAgg.IpRange.Field == nil {
			t.Fatal("Expected IpRange field to not be nil")
		}
		if *ipRangeAgg.IpRange.Field != "client_ip" {
			t.Errorf("Expected field to be 'client_ip', got '%s'", *ipRangeAgg.IpRange.Field)
		}
		if len(ipRangeAgg.IpRange.Ranges) != 2 {
			t.Errorf("Expected 2 ranges, got %d", len(ipRangeAgg.IpRange.Ranges))
		}
	})
}

func TestMissingAgg(t *testing.T) {
	t.Run("basic missing aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			MissingAgg("missing_emails", "email"),
		)
		
		missingAgg := aggs.Aggregations["missing_emails"]
		if missingAgg.Missing == nil {
			t.Fatal("Expected Missing aggregation to not be nil")
		}
		if missingAgg.Missing.Field == nil {
			t.Fatal("Expected Missing field to not be nil")
		}
		if *missingAgg.Missing.Field != "email" {
			t.Errorf("Expected field to be 'email', got '%s'", *missingAgg.Missing.Field)
		}
	})
}

func TestRareTermsAgg(t *testing.T) {
	t.Run("basic rare terms aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			RareTermsAgg("rare_categories", "category"),
		)
		
		rareTermsAgg := aggs.Aggregations["rare_categories"]
		if rareTermsAgg.RareTerms == nil {
			t.Fatal("Expected RareTerms aggregation to not be nil")
		}
		if rareTermsAgg.RareTerms.Field == nil {
			t.Fatal("Expected RareTerms field to not be nil")
		}
		if *rareTermsAgg.RareTerms.Field != "category" {
			t.Errorf("Expected field to be 'category', got '%s'", *rareTermsAgg.RareTerms.Field)
		}
	})
}

func TestSamplerAgg(t *testing.T) {
	t.Run("basic sampler aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			SamplerAgg("sample", 1000),
		)
		
		samplerAgg := aggs.Aggregations["sample"]
		if samplerAgg.Sampler == nil {
			t.Fatal("Expected Sampler aggregation to not be nil")
		}
		if samplerAgg.Sampler.ShardSize == nil {
			t.Fatal("Expected Sampler ShardSize to not be nil")
		}
		if *samplerAgg.Sampler.ShardSize != 1000 {
			t.Errorf("Expected shard size to be 1000, got %d", *samplerAgg.Sampler.ShardSize)
		}
	})
}

func TestDiversifiedSamplerAgg(t *testing.T) {
	t.Run("basic diversified sampler aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			DiversifiedSamplerAgg("diversified_sample", "category", 1000),
		)
		
		diversifiedSamplerAgg := aggs.Aggregations["diversified_sample"]
		if diversifiedSamplerAgg.DiversifiedSampler == nil {
			t.Fatal("Expected DiversifiedSampler aggregation to not be nil")
		}
		if diversifiedSamplerAgg.DiversifiedSampler.Field == nil {
			t.Fatal("Expected DiversifiedSampler field to not be nil")
		}
		if *diversifiedSamplerAgg.DiversifiedSampler.Field != "category" {
			t.Errorf("Expected field to be 'category', got '%s'", *diversifiedSamplerAgg.DiversifiedSampler.Field)
		}
		if diversifiedSamplerAgg.DiversifiedSampler.ShardSize == nil {
			t.Fatal("Expected DiversifiedSampler ShardSize to not be nil")
		}
		if *diversifiedSamplerAgg.DiversifiedSampler.ShardSize != 1000 {
			t.Errorf("Expected shard size to be 1000, got %d", *diversifiedSamplerAgg.DiversifiedSampler.ShardSize)
		}
	})
}

func TestReverseNestedAgg(t *testing.T) {
	t.Run("basic reverse nested aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			ReverseNestedAgg("back_to_parent"),
		)
		
		reverseNestedAgg := aggs.Aggregations["back_to_parent"]
		if reverseNestedAgg.ReverseNested == nil {
			t.Fatal("Expected ReverseNested aggregation to not be nil")
		}
	})

	t.Run("reverse nested aggregation with path", func(t *testing.T) {
		aggs := NewAggregations(
			ReverseNestedAgg("back_to_parent", "parent_path"),
		)
		
		reverseNestedAgg := aggs.Aggregations["back_to_parent"]
		if reverseNestedAgg.ReverseNested == nil {
			t.Fatal("Expected ReverseNested aggregation to not be nil")
		}
		if reverseNestedAgg.ReverseNested.Path == nil {
			t.Fatal("Expected ReverseNested path to not be nil")
		}
		if *reverseNestedAgg.ReverseNested.Path != "parent_path" {
			t.Errorf("Expected path to be 'parent_path', got '%s'", *reverseNestedAgg.ReverseNested.Path)
		}
	})
}

func TestChildrenAgg(t *testing.T) {
	t.Run("basic children aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			ChildrenAgg("child_products", "product"),
		)
		
		childrenAgg := aggs.Aggregations["child_products"]
		if childrenAgg.Children == nil {
			t.Fatal("Expected Children aggregation to not be nil")
		}
		if childrenAgg.Children.Type == nil {
			t.Fatal("Expected Children type to not be nil")
		}
		if *childrenAgg.Children.Type != "product" {
			t.Errorf("Expected type to be 'product', got '%s'", *childrenAgg.Children.Type)
		}
	})
}

func TestParentAgg(t *testing.T) {
	t.Run("basic parent aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			ParentAgg("parent_categories", "category"),
		)
		
		parentAgg := aggs.Aggregations["parent_categories"]
		if parentAgg.Parent == nil {
			t.Fatal("Expected Parent aggregation to not be nil")
		}
		if parentAgg.Parent.Type == nil {
			t.Fatal("Expected Parent type to not be nil")
		}
		if *parentAgg.Parent.Type != "category" {
			t.Errorf("Expected type to be 'category', got '%s'", *parentAgg.Parent.Type)
		}
	})
}

func TestAutoDateHistogramAgg(t *testing.T) {
	t.Run("basic auto date histogram aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			AutoDateHistogramAgg("auto_dates", "timestamp", 10),
		)
		
		autoDateHistAgg := aggs.Aggregations["auto_dates"]
		if autoDateHistAgg.AutoDateHistogram == nil {
			t.Fatal("Expected AutoDateHistogram aggregation to not be nil")
		}
		if autoDateHistAgg.AutoDateHistogram.Field == nil {
			t.Fatal("Expected AutoDateHistogram field to not be nil")
		}
		if *autoDateHistAgg.AutoDateHistogram.Field != "timestamp" {
			t.Errorf("Expected field to be 'timestamp', got '%s'", *autoDateHistAgg.AutoDateHistogram.Field)
		}
		if autoDateHistAgg.AutoDateHistogram.Buckets == nil {
			t.Fatal("Expected AutoDateHistogram buckets to not be nil")
		}
		if *autoDateHistAgg.AutoDateHistogram.Buckets != 10 {
			t.Errorf("Expected buckets to be 10, got %d", *autoDateHistAgg.AutoDateHistogram.Buckets)
		}
	})
}

func TestVariableWidthHistogramAgg(t *testing.T) {
	t.Run("basic variable width histogram aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			VariableWidthHistogramAgg("variable_histogram", "price", 10),
		)
		
		variableHistAgg := aggs.Aggregations["variable_histogram"]
		if variableHistAgg.VariableWidthHistogram == nil {
			t.Fatal("Expected VariableWidthHistogram aggregation to not be nil")
		}
		if variableHistAgg.VariableWidthHistogram.Field == nil {
			t.Fatal("Expected VariableWidthHistogram field to not be nil")
		}
		if *variableHistAgg.VariableWidthHistogram.Field != "price" {
			t.Errorf("Expected field to be 'price', got '%s'", *variableHistAgg.VariableWidthHistogram.Field)
		}
		if variableHistAgg.VariableWidthHistogram.Buckets == nil {
			t.Fatal("Expected VariableWidthHistogram buckets to not be nil")
		}
		if *variableHistAgg.VariableWidthHistogram.Buckets != 10 {
			t.Errorf("Expected buckets to be 10, got %d", *variableHistAgg.VariableWidthHistogram.Buckets)
		}
	})
}

func TestSignificantTextAgg(t *testing.T) {
	t.Run("basic significant text aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			SignificantTextAgg("significant_text", "description"),
		)
		
		significantTextAgg := aggs.Aggregations["significant_text"]
		if significantTextAgg.SignificantText == nil {
			t.Fatal("Expected SignificantText aggregation to not be nil")
		}
		if significantTextAgg.SignificantText.Field == nil {
			t.Fatal("Expected SignificantText field to not be nil")
		}
		if *significantTextAgg.SignificantText.Field != "description" {
			t.Errorf("Expected field to be 'description', got '%s'", *significantTextAgg.SignificantText.Field)
		}
	})
}

func TestPipelineAggregations(t *testing.T) {
	t.Run("avg bucket aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			AvgBucketAgg("avg_monthly_sales", "monthly_sales>total_sales"),
		)
		
		avgBucketAgg := aggs.Aggregations["avg_monthly_sales"]
		if avgBucketAgg.AvgBucket == nil {
			t.Fatal("Expected AvgBucket aggregation to not be nil")
		}
		if bucketsPath, ok := avgBucketAgg.AvgBucket.BucketsPath.(string); !ok || bucketsPath != "monthly_sales>total_sales" {
			t.Errorf("Expected buckets path to be 'monthly_sales>total_sales', got '%v'", avgBucketAgg.AvgBucket.BucketsPath)
		}
	})

	t.Run("max bucket aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			MaxBucketAgg("max_monthly_sales", "monthly_sales>total_sales"),
		)
		
		maxBucketAgg := aggs.Aggregations["max_monthly_sales"]
		if maxBucketAgg.MaxBucket == nil {
			t.Fatal("Expected MaxBucket aggregation to not be nil")
		}
		if bucketsPath, ok := maxBucketAgg.MaxBucket.BucketsPath.(string); !ok || bucketsPath != "monthly_sales>total_sales" {
			t.Errorf("Expected buckets path to be 'monthly_sales>total_sales', got '%v'", maxBucketAgg.MaxBucket.BucketsPath)
		}
	})

	t.Run("min bucket aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			MinBucketAgg("min_monthly_sales", "monthly_sales>total_sales"),
		)
		
		minBucketAgg := aggs.Aggregations["min_monthly_sales"]
		if minBucketAgg.MinBucket == nil {
			t.Fatal("Expected MinBucket aggregation to not be nil")
		}
		if bucketsPath, ok := minBucketAgg.MinBucket.BucketsPath.(string); !ok || bucketsPath != "monthly_sales>total_sales" {
			t.Errorf("Expected buckets path to be 'monthly_sales>total_sales', got '%v'", minBucketAgg.MinBucket.BucketsPath)
		}
	})

	t.Run("sum bucket aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			SumBucketAgg("total_monthly_sales", "monthly_sales>total_sales"),
		)
		
		sumBucketAgg := aggs.Aggregations["total_monthly_sales"]
		if sumBucketAgg.SumBucket == nil {
			t.Fatal("Expected SumBucket aggregation to not be nil")
		}
		if bucketsPath, ok := sumBucketAgg.SumBucket.BucketsPath.(string); !ok || bucketsPath != "monthly_sales>total_sales" {
			t.Errorf("Expected buckets path to be 'monthly_sales>total_sales', got '%v'", sumBucketAgg.SumBucket.BucketsPath)
		}
	})

	t.Run("stats bucket aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			StatsBucketAgg("sales_stats", "monthly_sales>total_sales"),
		)
		
		statsBucketAgg := aggs.Aggregations["sales_stats"]
		if statsBucketAgg.StatsBucket == nil {
			t.Fatal("Expected StatsBucket aggregation to not be nil")
		}
		if bucketsPath, ok := statsBucketAgg.StatsBucket.BucketsPath.(string); !ok || bucketsPath != "monthly_sales>total_sales" {
			t.Errorf("Expected buckets path to be 'monthly_sales>total_sales', got '%v'", statsBucketAgg.StatsBucket.BucketsPath)
		}
	})

	t.Run("extended stats bucket aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			ExtendedStatsBucketAgg("sales_extended_stats", "monthly_sales>total_sales"),
		)
		
		extendedStatsBucketAgg := aggs.Aggregations["sales_extended_stats"]
		if extendedStatsBucketAgg.ExtendedStatsBucket == nil {
			t.Fatal("Expected ExtendedStatsBucket aggregation to not be nil")
		}
		if bucketsPath, ok := extendedStatsBucketAgg.ExtendedStatsBucket.BucketsPath.(string); !ok || bucketsPath != "monthly_sales>total_sales" {
			t.Errorf("Expected buckets path to be 'monthly_sales>total_sales', got '%v'", extendedStatsBucketAgg.ExtendedStatsBucket.BucketsPath)
		}
	})

	t.Run("percentiles bucket aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			PercentilesBucketAgg("sales_percentiles", "monthly_sales>total_sales"),
		)
		
		percentilesBucketAgg := aggs.Aggregations["sales_percentiles"]
		if percentilesBucketAgg.PercentilesBucket == nil {
			t.Fatal("Expected PercentilesBucket aggregation to not be nil")
		}
		if bucketsPath, ok := percentilesBucketAgg.PercentilesBucket.BucketsPath.(string); !ok || bucketsPath != "monthly_sales>total_sales" {
			t.Errorf("Expected buckets path to be 'monthly_sales>total_sales', got '%v'", percentilesBucketAgg.PercentilesBucket.BucketsPath)
		}
	})

	t.Run("derivative aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			DerivativeAgg("sales_derivative", "sales"),
		)
		
		derivativeAgg := aggs.Aggregations["sales_derivative"]
		if derivativeAgg.Derivative == nil {
			t.Fatal("Expected Derivative aggregation to not be nil")
		}
		if bucketsPath, ok := derivativeAgg.Derivative.BucketsPath.(string); !ok || bucketsPath != "sales" {
			t.Errorf("Expected buckets path to be 'sales', got '%v'", derivativeAgg.Derivative.BucketsPath)
		}
	})

	t.Run("cumulative sum aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			CumulativeSumAgg("cumulative_sales", "sales"),
		)
		
		cumulativeSumAgg := aggs.Aggregations["cumulative_sales"]
		if cumulativeSumAgg.CumulativeSum == nil {
			t.Fatal("Expected CumulativeSum aggregation to not be nil")
		}
		if bucketsPath, ok := cumulativeSumAgg.CumulativeSum.BucketsPath.(string); !ok || bucketsPath != "sales" {
			t.Errorf("Expected buckets path to be 'sales', got '%v'", cumulativeSumAgg.CumulativeSum.BucketsPath)
		}
	})
}

func TestGeoGridAggregations(t *testing.T) {
	t.Run("geohash grid aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			GeohashGridAgg("location_grid", "location", 5),
		)
		
		geohashGridAgg := aggs.Aggregations["location_grid"]
		if geohashGridAgg.GeohashGrid == nil {
			t.Fatal("Expected GeohashGrid aggregation to not be nil")
		}
		if geohashGridAgg.GeohashGrid.Field == nil {
			t.Fatal("Expected GeohashGrid field to not be nil")
		}
		if *geohashGridAgg.GeohashGrid.Field != "location" {
			t.Errorf("Expected field to be 'location', got '%s'", *geohashGridAgg.GeohashGrid.Field)
		}
		if geohashGridAgg.GeohashGrid.Precision == nil {
			t.Fatal("Expected GeohashGrid precision to not be nil")
		}
		// GeohashGrid precision 可能是 int 或 *int 类型，我们需要检查实际值
		switch p := geohashGridAgg.GeohashGrid.Precision.(type) {
		case int:
			if p != 5 {
				t.Errorf("Expected precision to be 5, got %d", p)
			}
		case *int:
			if *p != 5 {
				t.Errorf("Expected precision to be 5, got %d", *p)
			}
		default:
			t.Errorf("Expected precision to be int or *int, got %T", p)
		}
	})

	t.Run("geotile grid aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			GeotileGridAgg("tile_grid", "location", 8),
		)
		
		geotileGridAgg := aggs.Aggregations["tile_grid"]
		if geotileGridAgg.GeotileGrid == nil {
			t.Fatal("Expected GeotileGrid aggregation to not be nil")
		}
		if geotileGridAgg.GeotileGrid.Field == nil {
			t.Fatal("Expected GeotileGrid field to not be nil")
		}
		if *geotileGridAgg.GeotileGrid.Field != "location" {
			t.Errorf("Expected field to be 'location', got '%s'", *geotileGridAgg.GeotileGrid.Field)
		}
		if geotileGridAgg.GeotileGrid.Precision == nil {
			t.Fatal("Expected GeotileGrid precision to not be nil")
		}
		if *geotileGridAgg.GeotileGrid.Precision != 8 {
			t.Errorf("Expected precision to be 8, got %d", *geotileGridAgg.GeotileGrid.Precision)
		}
	})
}

func TestAdvancedMetricAggregations(t *testing.T) {
	t.Run("weighted avg aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			WeightedAvgAgg("weighted_avg_price", "price", "quantity"),
		)
		
		weightedAvgAgg := aggs.Aggregations["weighted_avg_price"]
		if weightedAvgAgg.WeightedAvg == nil {
			t.Fatal("Expected WeightedAvg aggregation to not be nil")
		}
		if weightedAvgAgg.WeightedAvg.Value == nil {
			t.Fatal("Expected WeightedAvg value to not be nil")
		}
		if weightedAvgAgg.WeightedAvg.Value.Field == nil {
			t.Fatal("Expected WeightedAvg value field to not be nil")
		}
		if *weightedAvgAgg.WeightedAvg.Value.Field != "price" {
			t.Errorf("Expected value field to be 'price', got '%s'", *weightedAvgAgg.WeightedAvg.Value.Field)
		}
		if weightedAvgAgg.WeightedAvg.Weight == nil {
			t.Fatal("Expected WeightedAvg weight to not be nil")
		}
		if weightedAvgAgg.WeightedAvg.Weight.Field == nil {
			t.Fatal("Expected WeightedAvg weight field to not be nil")
		}
		if *weightedAvgAgg.WeightedAvg.Weight.Field != "quantity" {
			t.Errorf("Expected weight field to be 'quantity', got '%s'", *weightedAvgAgg.WeightedAvg.Weight.Field)
		}
	})

	t.Run("median absolute deviation aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			MedianAbsoluteDeviationAgg("price_mad", "price"),
		)
		
		madAgg := aggs.Aggregations["price_mad"]
		if madAgg.MedianAbsoluteDeviation == nil {
			t.Fatal("Expected MedianAbsoluteDeviation aggregation to not be nil")
		}
		if madAgg.MedianAbsoluteDeviation.Field == nil {
			t.Fatal("Expected MedianAbsoluteDeviation field to not be nil")
		}
		if *madAgg.MedianAbsoluteDeviation.Field != "price" {
			t.Errorf("Expected field to be 'price', got '%s'", *madAgg.MedianAbsoluteDeviation.Field)
		}
	})

	t.Run("string stats aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			StringStatsAgg("description_stats", "description"),
		)
		
		stringStatsAgg := aggs.Aggregations["description_stats"]
		if stringStatsAgg.StringStats == nil {
			t.Fatal("Expected StringStats aggregation to not be nil")
		}
		if stringStatsAgg.StringStats.Field == nil {
			t.Fatal("Expected StringStats field to not be nil")
		}
		if *stringStatsAgg.StringStats.Field != "description" {
			t.Errorf("Expected field to be 'description', got '%s'", *stringStatsAgg.StringStats.Field)
		}
	})

	t.Run("t test aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			TTestAgg("t_test", "score", Term("group", "A"), Term("group", "B")),
		)
		
		tTestAgg := aggs.Aggregations["t_test"]
		if tTestAgg.TTest == nil {
			t.Fatal("Expected TTest aggregation to not be nil")
		}
		if tTestAgg.TTest.A == nil {
			t.Fatal("Expected TTest A population to not be nil")
		}
		if tTestAgg.TTest.A.Field != "score" {
			t.Errorf("Expected A field to be 'score', got '%s'", tTestAgg.TTest.A.Field)
		}
		if tTestAgg.TTest.B == nil {
			t.Fatal("Expected TTest B population to not be nil")
		}
		if tTestAgg.TTest.B.Field != "score" {
			t.Errorf("Expected B field to be 'score', got '%s'", tTestAgg.TTest.B.Field)
		}
	})
}


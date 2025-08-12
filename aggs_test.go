package esb

import (
	"testing"
)

func TestNewAggregations(t *testing.T) {
	t.Run("empty aggregations", func(t *testing.T) {
		aggs := NewAggregations()
		if aggs == nil {
			t.Fatal("Expected aggregations to not be nil")
		}
		if len(aggs) != 0 {
			t.Errorf("Expected empty aggregations, got %d", len(aggs))
		}
	})

	t.Run("with single aggregation", func(t *testing.T) {
		aggs := NewAggregations(
			TermsAgg("categories", "category"),
		)
		if aggs == nil {
			t.Fatal("Expected aggregations to not be nil")
		}
		if len(aggs) != 1 {
			t.Errorf("Expected 1 aggregation, got %d", len(aggs))
		}
		if _, exists := aggs["categories"]; !exists {
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
		if len(aggs) != 3 {
			t.Errorf("Expected 3 aggregations, got %d", len(aggs))
		}
		expectedKeys := []string{"categories", "avg_price", "total_sales"}
		for _, key := range expectedKeys {
			if _, exists := aggs[key]; !exists {
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
		
		categoryAgg := aggs["categories"]
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
		
		categoryAgg := aggs["categories"]
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
		
		avgAgg := aggs["avg_price"]
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
		
		sumAgg := aggs["total_sales"]
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
}
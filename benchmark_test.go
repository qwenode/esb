package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
)

// BenchmarkSimpleQuery compares simple query construction
func BenchmarkSimpleQuery(b *testing.B) {
	b.Run("ESB_Term", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(Term("status", "published"))
			if err != nil {
				b.Errorf("ESB Term query failed: %v", err)
			}
		}
	})
	
	b.Run("Native_Term", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			query := &types.Query{
				Term: map[string]types.TermQuery{
					"status": {
						Value: "published",
					},
				},
			}
			_ = query
		}
	})
}

// BenchmarkMatchQueryConstruction compares match query construction
func BenchmarkMatchQueryConstruction(b *testing.B) {
	b.Run("ESB_Match", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(Match("title", "elasticsearch"))
			if err != nil {
				b.Errorf("ESB Match query failed: %v", err)
			}
		}
	})
	
	b.Run("Native_Match", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			query := &types.Query{
				Match: map[string]types.MatchQuery{
					"title": {
						Query: "elasticsearch",
					},
				},
			}
			_ = query
		}
	})
}

// BenchmarkRangeQuery compares range query construction
func BenchmarkRangeQuery(b *testing.B) {
	b.Run("ESB_Range", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(NumberRange("age").Gte(18.0).Lt(65.0).Build())
			if err != nil {
				b.Errorf("ESB Range query failed: %v", err)
			}
		}
	})
	
	b.Run("Native_Range", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			query := &types.Query{
				Range: map[string]types.RangeQuery{
					"age": types.UntypedRangeQuery{
						Gte: []byte("18"),
						Lt:  []byte("65"),
					},
				},
			}
			_ = query
		}
	})
}

// BenchmarkBoolQueryConstruction compares bool query construction
func BenchmarkBoolQueryConstruction(b *testing.B) {
	b.Run("ESB_Bool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(
				Bool(
					Must(
						Match("title", "elasticsearch"),
						Term("status", "published"),
					),
					Should(
						Term("category", "tech"),
						NumberRange("score").Gte(4.0).Build(),
					),
				),
			)
			if err != nil {
				b.Errorf("ESB Bool query failed: %v", err)
			}
		}
	})
	
	b.Run("Native_Bool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			query := &types.Query{
				Bool: &types.BoolQuery{
					Must: []types.Query{
						{
							Match: map[string]types.MatchQuery{
								"title": {
									Query: "elasticsearch",
								},
							},
						},
						{
							Term: map[string]types.TermQuery{
								"status": {
									Value: "published",
								},
							},
						},
					},
					Should: []types.Query{
						{
							Term: map[string]types.TermQuery{
								"category": {
									Value: "tech",
								},
							},
						},
						{
							Range: map[string]types.RangeQuery{
								"score": types.UntypedRangeQuery{
									Gte: []byte("4"),
								},
							},
						},
					},
				},
			}
			_ = query
		}
	})
}

// BenchmarkComplexQuery compares complex query construction
func BenchmarkComplexQuery(b *testing.B) {
	b.Run("ESB_Complex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(
				Bool(
					Must(
						Match("title", "elasticsearch guide"),
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
				b.Errorf("ESB Complex query failed: %v", err)
			}
		}
	})
	
	b.Run("Native_Complex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			query := &types.Query{
				Bool: &types.BoolQuery{
					Must: []types.Query{
						{
							Match: map[string]types.MatchQuery{
								"title": {
									Query: "elasticsearch guide",
								},
							},
						},
						{
							Range: map[string]types.RangeQuery{
								"publish_date": types.UntypedRangeQuery{
									Gte: []byte(`"2023-01-01"`),
									Lte: []byte(`"2023-12-31"`),
								},
							},
						},
						{
							Exists: &types.ExistsQuery{
								Field: "author",
							},
						},
					},
					Should: []types.Query{
						{
							MatchPhrase: map[string]types.MatchPhraseQuery{
								"content": {
									Query: "search engine",
								},
							},
						},
						{
							Term: map[string]types.TermQuery{
								"category": {
									Value: "technology",
								},
							},
						},
						{
							Range: map[string]types.RangeQuery{
								"views": types.UntypedRangeQuery{
									Gte: []byte("1000"),
								},
							},
						},
					},
					Filter: []types.Query{
						{
							Term: map[string]types.TermQuery{
								"status": {
									Value: "published",
								},
							},
						},
						{
							Range: map[string]types.RangeQuery{
								"score": types.UntypedRangeQuery{
									Gte: []byte("4"),
								},
							},
						},
					},
					MustNot: []types.Query{
						{
							Term: map[string]types.TermQuery{
								"deleted": {
									Value: "true",
								},
							},
						},
						{
							Exists: &types.ExistsQuery{
								Field: "spam_flag",
							},
						},
					},
				},
			}
			_ = query
		}
	})
}

// BenchmarkNestedBoolQuery compares nested bool query construction
func BenchmarkNestedBoolQuery(b *testing.B) {
	b.Run("ESB_NestedBool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(
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
				b.Errorf("ESB Nested Bool query failed: %v", err)
			}
		}
	})
	
	b.Run("Native_NestedBool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			query := &types.Query{
				Bool: &types.BoolQuery{
					Must: []types.Query{
						{
							Bool: &types.BoolQuery{
								Should: []types.Query{
									{
										Match: map[string]types.MatchQuery{
											"title": {
												Query: "elasticsearch",
											},
										},
									},
									{
										Match: map[string]types.MatchQuery{
											"content": {
												Query: "search",
											},
										},
									},
								},
							},
						},
						{
							Bool: &types.BoolQuery{
								Must: []types.Query{
									{
										Range: map[string]types.RangeQuery{
											"date": types.UntypedRangeQuery{
												Gte: []byte(`"2023-01-01"`),
											},
										},
									},
									{
										Exists: &types.ExistsQuery{
											Field: "author",
										},
									},
								},
								MustNot: []types.Query{
									{
										Term: map[string]types.TermQuery{
											"deleted": {
												Value: "true",
											},
										},
									},
								},
							},
						},
					},
				},
			}
			_ = query
		}
	})
}

// BenchmarkMatchWithOptions compares match query with options
func BenchmarkMatchWithOptions(b *testing.B) {
	b.Run("ESB_MatchWithOptions", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(
				MatchWithOptions("title", "elasticsearch guide", MatchOptions{
					Boost:    float32Ptr(2.0),
					Analyzer: stringPtr("standard"),
					Operator: &operator.And,
				}),
			)
			if err != nil {
				b.Errorf("ESB MatchWithOptions query failed: %v", err)
			}
		}
	})
	
	b.Run("Native_MatchWithOptions", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			boost := float32(2.0)
			analyzer := "standard"
			op := operator.And
			query := &types.Query{
				Match: map[string]types.MatchQuery{
					"title": {
						Query:    "elasticsearch guide",
						Boost:    &boost,
						Analyzer: &analyzer,
						Operator: &op,
					},
				},
			}
			_ = query
		}
	})
}

// BenchmarkTermsQuery compares terms query construction
func BenchmarkTermsQuery(b *testing.B) {
	b.Run("ESB_Terms", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(Terms("category", "tech", "science", "programming"))
			if err != nil {
				b.Errorf("ESB Terms query failed: %v", err)
			}
		}
	})
	
	b.Run("Native_Terms", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			query := &types.Query{
				Terms: &types.TermsQuery{
					TermsQuery: map[string]types.TermsQueryField{
						"category": []string{"tech", "science", "programming"},
					},
				},
			}
			_ = query
		}
	})
}

// BenchmarkExistsQuery compares exists query construction
func BenchmarkExistsQuery(b *testing.B) {
	b.Run("ESB_Exists", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(Exists("author"))
			if err != nil {
				b.Errorf("ESB Exists query failed: %v", err)
			}
		}
	})
	
	b.Run("Native_Exists", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			query := &types.Query{
				Exists: &types.ExistsQuery{
					Field: "author",
				},
			}
			_ = query
		}
	})
}

// BenchmarkRangeWithOptions compares range query with options
func BenchmarkRangeWithOptions(b *testing.B) {
	b.Run("ESB_RangeWithOptions", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(
				DateRange("timestamp").
					Gte("2023-01-01").
					Lte("2023-12-31").
					Format("yyyy-MM-dd").
					TimeZone("UTC").
					Boost(1.5).
					Build(),
			)
			if err != nil {
				b.Errorf("ESB Range with options query failed: %v", err)
			}
		}
	})
	
	b.Run("Native_RangeWithOptions", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			format := "yyyy-MM-dd"
			timezone := "UTC"
			boost := float32(1.5)
			query := &types.Query{
				Range: map[string]types.RangeQuery{
					"timestamp": types.UntypedRangeQuery{
						Gte:      []byte(`"2023-01-01"`),
						Lte:      []byte(`"2023-12-31"`),
						Format:   &format,
						TimeZone: &timezone,
						Boost:    &boost,
					},
				},
			}
			_ = query
		}
	})
}

// BenchmarkMemoryAllocation compares memory allocation
func BenchmarkMemoryAllocation(b *testing.B) {
	b.Run("ESB_Memory", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, err := NewQuery(
				Bool(
					Must(
						Match("title", "elasticsearch"),
						DateRange("date").Gte("2023-01-01").Build(),
					),
				),
			)
			if err != nil {
				b.Errorf("ESB Memory test failed: %v", err)
			}
		}
	})
	
	b.Run("Native_Memory", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			query := &types.Query{
				Bool: &types.BoolQuery{
					Must: []types.Query{
						{
							Match: map[string]types.MatchQuery{
								"title": {
									Query: "elasticsearch",
								},
							},
						},
						{
							Range: map[string]types.RangeQuery{
								"date": types.UntypedRangeQuery{
									Gte: []byte(`"2023-01-01"`),
								},
							},
						},
					},
				},
			}
			_ = query
		}
	})
}

// BenchmarkConcurrentUsage tests concurrent usage performance
func BenchmarkConcurrentUsage(b *testing.B) {
	b.Run("ESB_Concurrent", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_, err := NewQuery(
					Bool(
						Must(
							Match("title", "elasticsearch"),
							Term("status", "published"),
						),
					),
				)
				if err != nil {
					b.Errorf("ESB Concurrent test failed: %v", err)
				}
			}
		})
	})
	
	b.Run("Native_Concurrent", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				query := &types.Query{
					Bool: &types.BoolQuery{
						Must: []types.Query{
							{
								Match: map[string]types.MatchQuery{
									"title": {
										Query: "elasticsearch",
									},
								},
							},
							{
								Term: map[string]types.TermQuery{
									"status": {
										Value: "published",
									},
								},
							},
						},
					},
				}
				_ = query
			}
		})
	})
} 
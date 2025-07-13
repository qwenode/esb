package main

import (
	"fmt"
	"log"
	"encoding/json"
	
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/qwenode/esb"
)

func main() {
	// Initialize Elasticsearch client
	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Bool Query Builder Examples ===\n")

	// Example 1: Simple Bool Query with Must clause
	fmt.Println("1. Simple Bool Query with Must clause:")
	
	// Using our builder
	builderQuery1, err := esb.NewQuery(
		esb.Bool(
			esb.Must(
				esb.Term("status", "published"),
				esb.Term("active", "true"),
			),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// Manual construction (original way)
	manualQuery1 := &types.Query{
		Bool: &types.BoolQuery{
			Must: []types.Query{
				{
					Term: map[string]types.TermQuery{
						"status": {Value: "published"},
					},
				},
				{
					Term: map[string]types.TermQuery{
						"active": {Value: "true"},
					},
				},
			},
		},
	}
	
	fmt.Printf("Builder Query: %+v\n", builderQuery1)
	fmt.Printf("Manual Query:  %+v\n", manualQuery1)
	fmt.Println()

	// Example 2: Complex Bool Query with all clauses
	fmt.Println("2. Complex Bool Query with Must, Should, Filter, MustNot:")
	
	// Using our builder
	builderQuery2, err := esb.NewQuery(
		esb.Bool(
			esb.Must(
				esb.Term("status", "published"),
			),
			esb.Should(
				esb.Term("category", "tech"),
				esb.Term("category", "science"),
			),
			esb.Filter(
				esb.Term("active", "true"),
				esb.Terms("type", "article", "blog"),
			),
			esb.MustNot(
				esb.Term("deleted", "true"),
			),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// Manual construction would be much more verbose...
	manualQuery2 := &types.Query{
		Bool: &types.BoolQuery{
			Must: []types.Query{
				{
					Term: map[string]types.TermQuery{
						"status": {Value: "published"},
					},
				},
			},
			Should: []types.Query{
				{
					Term: map[string]types.TermQuery{
						"category": {Value: "tech"},
					},
				},
				{
					Term: map[string]types.TermQuery{
						"category": {Value: "science"},
					},
				},
			},
			Filter: []types.Query{
				{
					Term: map[string]types.TermQuery{
						"active": {Value: "true"},
					},
				},
				{
					Terms: &types.TermsQuery{
						TermsQuery: map[string]types.TermsQueryField{
							"type": []types.FieldValue{"article", "blog"},
						},
					},
				},
			},
			MustNot: []types.Query{
				{
					Term: map[string]types.TermQuery{
						"deleted": {Value: "true"},
					},
				},
			},
		},
	}
	
	fmt.Printf("Builder Query: %+v\n", builderQuery2)
	fmt.Printf("Manual Query:  %+v\n", manualQuery2)
	fmt.Println()

	// Example 3: Nested Bool Queries
	fmt.Println("3. Nested Bool Queries:")
	
	// Using our builder
	builderQuery3, err := esb.NewQuery(
		esb.Bool(
			esb.Must(
				esb.Term("status", "published"),
				esb.Bool(
					esb.Should(
						esb.Term("category", "tech"),
						esb.Term("category", "science"),
					),
				),
			),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Nested Bool Query: %+v\n", builderQuery3)
	fmt.Println()

	// Example 4: Using with Elasticsearch client
	fmt.Println("4. Using with Elasticsearch client:")
	
	// Our builder query can be used directly with the client
	searchQuery, err := esb.NewQuery(
		esb.Bool(
			esb.Must(
				esb.Term("status", "published"),
			),
			esb.Filter(
				esb.Term("active", "true"),
			),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// This would work with a real Elasticsearch instance
	response := client.Search().
		Index("my-index").
		Query(searchQuery)
	
	fmt.Printf("Search Request: %+v\n", response)
	fmt.Println()

	// Example 5: JSON serialization comparison
	fmt.Println("5. JSON Serialization Comparison:")
	
	// Serialize both queries to JSON for comparison
	builderJSON, _ := json.MarshalIndent(builderQuery2, "", "  ")
	manualJSON, _ := json.MarshalIndent(manualQuery2, "", "  ")
	
	fmt.Println("Builder Query JSON:")
	fmt.Println(string(builderJSON))
	fmt.Println()
	
	fmt.Println("Manual Query JSON:")
	fmt.Println(string(manualJSON))
	fmt.Println()

	// Code length comparison
	fmt.Println("=== Code Length Comparison ===")
	fmt.Println("Builder approach: ~8 lines for complex query")
	fmt.Println("Manual approach:  ~40+ lines for same query")
	fmt.Println("Code reduction:   ~80% less code!")
} 
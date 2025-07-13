package main

import (
	"fmt"
	"log"
	
	"github.com/qwenode/esb"
)

func main() {
	// Example 1: Basic numeric range query
	fmt.Println("=== Basic Numeric Range Query ===")
	query1 := esb.NewQuery(
		esb.Range("age").Gte(18).Lt(65).Build(),
	)
	fmt.Printf("Age range: 18 <= age < 65\n")
	
	// Example 2: Price range with decimals
	fmt.Println("\n=== Price Range Query ===")
	query2 := esb.NewQuery(
		esb.Range("price").Gt(10.0).Lte(100.0).Build(),
	)
	fmt.Printf("Price range: 10.0 < price <= 100.0\n")
	
	// Example 3: Date range query
	fmt.Println("\n=== Date Range Query ===")
	query3 := esb.NewQuery(
		esb.Range("created_at").
			Gte("2023-01-01").
			Lte("2023-12-31").
			Format("yyyy-MM-dd").
			TimeZone("UTC").
			Build(),
	)
	fmt.Printf("Date range: 2023-01-01 to 2023-12-31 (UTC)\n")
	
	// Example 4: Using From/To syntax
	fmt.Println("\n=== From/To Range Query ===")
	query4 := esb.NewQuery(
		esb.Range("score").From(80).To(100).Build(),
	)
	fmt.Printf("Score range: [80, 100)\n")
	
	// Example 5: Range with boost
	fmt.Println("\n=== Range Query with Boost ===")
	query5 := esb.NewQuery(
		esb.Range("priority").Gte(5).Boost(2.0).Build(),
	)
	fmt.Printf("Priority >= 5 with boost 2.0\n")
	
	// Example 6: Multiple range queries in Bool query
	fmt.Println("\n=== Multiple Range Queries in Bool ===")
	query6 := esb.NewQuery(
		esb.Bool(
			esb.Must(
				esb.Range("age").Gte(18).Lt(65).Build(),
				esb.Range("salary").Gte(30000).Build(),
				esb.Range("experience").Gte(2).Build(),
			),
			esb.Should(
				esb.Range("bonus").Gt(5000).Build(),
				esb.Range("rating").Gte(4.5).Build(),
			),
		),
	)
	fmt.Printf("Complex query with multiple range conditions\n")
	
	// Example 7: Open-ended ranges
	fmt.Println("\n=== Open-ended Range Queries ===")
	query7 := esb.NewQuery(
		esb.Bool(
			esb.Must(
				esb.Range("created_at").Gte("2023-01-01").Build(), // Only lower bound
				esb.Range("views").Lt(1000).Build(),               // Only upper bound
			),
		),
	)
	fmt.Printf("Open-ended ranges: created_at >= 2023-01-01 AND views < 1000\n")
	
	// Example 8: String range (lexicographic)
	fmt.Println("\n=== String Range Query ===")
	query8 := esb.NewQuery(
		esb.Range("category").Gte("A").Lt("M").Build(),
	)
	fmt.Printf("Category range: A <= category < M (lexicographic)\n")
	
	fmt.Println("\n=== All Range query examples completed successfully! ===")
	
	// Suppress unused variable warnings
	_ = query1
	_ = query2
	_ = query3
	_ = query4
	_ = query5
	_ = query6
	_ = query7
	_ = query8
} 
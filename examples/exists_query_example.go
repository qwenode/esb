package main

import (
	"fmt"
	"log"
	
	"github.com/qwenode/esb"
)

func main() {
	// Example 1: Basic exists query
	fmt.Println("=== Basic Exists Query ===")
	query1 := esb.NewQuery(
		esb.Exists("user.name"),
	)
	fmt.Printf("Query: %+v\n", query1.Exists)
	
	// Example 2: Exists query with Bool query
	fmt.Println("\n=== Exists Query with Bool ===")
	query2 := esb.NewQuery(
		esb.Bool(
			esb.Must(
				esb.Exists("user.name"),
				esb.Term("status", "active"),
			),
			esb.MustNot(
				esb.Exists("deleted_at"),
			),
		),
	)
	fmt.Printf("Bool Query Must: %d clauses\n", len(query2.Bool.Must))
	fmt.Printf("Bool Query MustNot: %d clauses\n", len(query2.Bool.MustNot))
	
	// Example 3: Multiple exists queries
	fmt.Println("\n=== Multiple Exists Queries ===")
	query3 := esb.NewQuery(
		esb.Bool(
			esb.Must(
				esb.Exists("user.name"),
				esb.Exists("user.email"),
				esb.Exists("metadata.timestamp"),
			),
		),
	)
	fmt.Printf("Checking for %d required fields\n", len(query3.Bool.Must))
	
	// Example 4: Nested field exists
	fmt.Println("\n=== Nested Field Exists ===")
	query4 := esb.NewQuery(
		esb.Exists("user.profile.avatar.url"),
	)
	fmt.Printf("Nested field: %s\n", query4.Exists.Field)
	
	fmt.Println("\n=== All examples completed successfully! ===")
} 
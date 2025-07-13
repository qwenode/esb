package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// BoolOption represents a function that modifies a types.BoolQuery.
// It is used to build complex boolean queries with Must, Should, Filter, and MustNot clauses.
type BoolOption func(*types.BoolQuery)

// Bool creates a boolean query with the specified options.
// Boolean queries are used to combine multiple queries using boolean logic.
//
// Example:
//   query, err := esb.NewQuery(
//       esb.Bool(
//           esb.Must(
//               esb.Term("status", "published"),
//               esb.Range("date").Gte("2023-01-01").Build(),
//           ),
//           esb.Filter(
//               esb.Term("category", "tech"),
//           ),
//           esb.Should(
//               esb.Match("title", "elasticsearch"),
//               esb.Match("content", "search"),
//           ),
//       ),
//   )
func Bool(opts ...BoolOption) QueryOption {
	return func(q *types.Query) {
		boolQuery := &types.BoolQuery{}
		for _, opt := range opts {
			opt(boolQuery)
		}
		q.Bool = boolQuery
	}
}

// Must specifies queries that must match for a document to be returned.
// All queries in the Must clause are required to match (AND logic).
//
// Example:
//   esb.Must(
//       esb.Term("status", "published"),
//       esb.Range("date").Gte("2023-01-01").Build(),
//   )
func Must(opts ...QueryOption) BoolOption {
	return func(bq *types.BoolQuery) {
		for _, opt := range opts {
			subQuery := &types.Query{}
			opt(subQuery)
			bq.Must = append(bq.Must, *subQuery)
		}
	}
}

// Should specifies queries that should match for a document to be returned.
// At least one query in the Should clause should match (OR logic).
// The more Should clauses that match, the higher the document's score.
//
// Example:
//   esb.Should(
//       esb.Match("title", "elasticsearch"),
//       esb.Match("content", "search engine"),
//   )
func Should(opts ...QueryOption) BoolOption {
	return func(bq *types.BoolQuery) {
		for _, opt := range opts {
			subQuery := &types.Query{}
			opt(subQuery)
			bq.Should = append(bq.Should, *subQuery)
		}
	}
}

// Filter specifies queries that must match for a document to be returned,
// but unlike Must, Filter queries do not contribute to the score.
// Filter queries are cached and faster than Must queries.
//
// Example:
//   esb.Filter(
//       esb.Term("status", "published"),
//       esb.Range("publish_date").Gte("2023-01-01").Build(),
//   )
func Filter(opts ...QueryOption) BoolOption {
	return func(bq *types.BoolQuery) {
		for _, opt := range opts {
			subQuery := &types.Query{}
			opt(subQuery)
			bq.Filter = append(bq.Filter, *subQuery)
		}
	}
}

// MustNot specifies queries that must not match for a document to be returned.
// Documents matching any MustNot query will be excluded from the results.
//
// Example:
//   esb.MustNot(
//       esb.Term("status", "deleted"),
//       esb.Term("hidden", true),
//   )
func MustNot(opts ...QueryOption) BoolOption {
	return func(bq *types.BoolQuery) {
		for _, opt := range opts {
			subQuery := &types.Query{}
			opt(subQuery)
			bq.MustNot = append(bq.MustNot, *subQuery)
		}
	}
} 
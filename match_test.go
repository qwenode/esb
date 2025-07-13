package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
)

func TestMatch(t *testing.T) {
	t.Run("should create basic match query", func(t *testing.T) {
		query := NewQuery(Match("title", "elasticsearch search"))
		if query == nil {
			t.Error("expected non-nil query")
		}
		if query.Match == nil {
			t.Error("expected Match query")
		}
		if query.Match["title"].Query != "elasticsearch search" {
			t.Errorf("expected query 'elasticsearch search', got %s", query.Match["title"].Query)
		}
	})



	t.Run("should support multiple match queries in Bool", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Must(
					Match("title", "elasticsearch"),
					Match("content", "search engine"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("expected Bool query")
		}
		if len(query.Bool.Must) != 2 {
			t.Errorf("expected 2 must queries, got %d", len(query.Bool.Must))
		}
		// Check first match query
		if query.Bool.Must[0].Match == nil {
			t.Error("expected first query to be Match")
		}
		// Check second match query
		if query.Bool.Must[1].Match == nil {
			t.Error("expected second query to be Match")
		}
	})
}

func TestMatchWithOptions(t *testing.T) {
	t.Run("should create match query with operator", func(t *testing.T) {
		op := operator.And
		query := NewQuery(
			MatchWithOptions("title", "elasticsearch search", MatchOptions{
				Operator: &op,
			}),
		)
		if query.Match == nil {
			t.Error("expected Match query")
		}
		if query.Match["title"].Operator == nil {
			t.Error("expected operator to be set")
		}
		if *query.Match["title"].Operator != operator.And {
			t.Errorf("expected operator AND, got %v", *query.Match["title"].Operator)
		}
	})

	t.Run("should create match query with fuzziness", func(t *testing.T) {
		fuzziness := types.Fuzziness("AUTO")
		query := NewQuery(
			MatchWithOptions("title", "elasticsearch", MatchOptions{
				Fuzziness: fuzziness,
			}),
		)
		if query.Match["title"].Fuzziness == nil {
			t.Error("expected fuzziness to be set")
		}
		if query.Match["title"].Fuzziness != types.Fuzziness("AUTO") {
			t.Errorf("expected fuzziness AUTO, got %v", query.Match["title"].Fuzziness)
		}
	})

	t.Run("should create match query with multiple options", func(t *testing.T) {
		op := operator.And
		analyzer := "standard"
		boost := float32(1.5)
		lenient := true
		
		query := NewQuery(
			MatchWithOptions("title", "elasticsearch search", MatchOptions{
				Operator: &op,
				Analyzer: &analyzer,
				Boost:    &boost,
				Lenient:  &lenient,
			}),
		)
		
		matchQuery := query.Match["title"]
		if matchQuery.Operator == nil || *matchQuery.Operator != operator.And {
			t.Error("expected operator to be AND")
		}
		if matchQuery.Analyzer == nil || *matchQuery.Analyzer != "standard" {
			t.Error("expected analyzer to be 'standard'")
		}
		if matchQuery.Boost == nil || *matchQuery.Boost != 1.5 {
			t.Error("expected boost to be 1.5")
		}
		if matchQuery.Lenient == nil || *matchQuery.Lenient != true {
			t.Error("expected lenient to be true")
		}
	})

	t.Run("should work with empty options", func(t *testing.T) {
		query := NewQuery(
			MatchWithOptions("title", "elasticsearch", MatchOptions{}),
		)
		if query.Match == nil {
			t.Error("expected Match query")
		}
		if query.Match["title"].Query != "elasticsearch" {
			t.Errorf("expected query 'elasticsearch', got %s", query.Match["title"].Query)
		}
	})
}

func TestMatchPhrase(t *testing.T) {
	t.Run("should create match phrase query", func(t *testing.T) {
		query := NewQuery(MatchPhrase("content", "elasticsearch is awesome"))
		if query == nil {
			t.Error("expected non-nil query")
		}
		if query.MatchPhrase == nil {
			t.Error("expected MatchPhrase query")
		}
		if query.MatchPhrase["content"].Query != "elasticsearch is awesome" {
			t.Errorf("expected query 'elasticsearch is awesome', got %s", query.MatchPhrase["content"].Query)
		}
	})


}

func TestMatchPhraseWithOptions(t *testing.T) {
	t.Run("should create match phrase query with slop", func(t *testing.T) {
		slop := 2
		query := NewQuery(
			MatchPhraseWithOptions("content", "elasticsearch search", MatchPhraseOptions{
				Slop: &slop,
			}),
		)
		if query.MatchPhrase == nil {
			t.Error("expected MatchPhrase query")
		}
		if query.MatchPhrase["content"].Slop == nil {
			t.Error("expected slop to be set")
		}
		if *query.MatchPhrase["content"].Slop != 2 {
			t.Errorf("expected slop 2, got %d", *query.MatchPhrase["content"].Slop)
		}
	})

	t.Run("should create match phrase query with analyzer and boost", func(t *testing.T) {
		analyzer := "keyword"
		boost := float32(2.0)
		
		query := NewQuery(
			MatchPhraseWithOptions("content", "exact phrase", MatchPhraseOptions{
				Analyzer: &analyzer,
				Boost:    &boost,
			}),
		)
		
		matchPhraseQuery := query.MatchPhrase["content"]
		if matchPhraseQuery.Analyzer == nil || *matchPhraseQuery.Analyzer != "keyword" {
			t.Error("expected analyzer to be 'keyword'")
		}
		if matchPhraseQuery.Boost == nil || *matchPhraseQuery.Boost != 2.0 {
			t.Error("expected boost to be 2.0")
		}
	})
}

func TestMatchPhrasePrefix(t *testing.T) {
	t.Run("should create match phrase prefix query", func(t *testing.T) {
		query := NewQuery(MatchPhrasePrefix("title", "elasticsearch sea"))
		if query == nil {
			t.Error("expected non-nil query")
		}
		if query.MatchPhrasePrefix == nil {
			t.Error("expected MatchPhrasePrefix query")
		}
		if query.MatchPhrasePrefix["title"].Query != "elasticsearch sea" {
			t.Errorf("expected query 'elasticsearch sea', got %s", query.MatchPhrasePrefix["title"].Query)
		}
	})


}

func TestMatchInBoolQuery(t *testing.T) {
	t.Run("should work with Bool query", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Must(
					Match("title", "elasticsearch"),
					MatchPhrase("content", "search engine"),
				),
				Should(
					Match("tags", "database"),
					MatchPhrasePrefix("description", "fast sea"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("expected Bool query")
		}
		if len(query.Bool.Must) != 2 {
			t.Errorf("expected 2 Must clauses, got %d", len(query.Bool.Must))
		}
		if len(query.Bool.Should) != 2 {
			t.Errorf("expected 2 Should clauses, got %d", len(query.Bool.Should))
		}
		
		// Check that Match queries are properly nested
		mustQuery1 := query.Bool.Must[0]
		if mustQuery1.Match == nil {
			t.Error("expected Match query in Must clause")
		}
		
		mustQuery2 := query.Bool.Must[1]
		if mustQuery2.MatchPhrase == nil {
			t.Error("expected MatchPhrase query in Must clause")
		}
		
		shouldQuery1 := query.Bool.Should[0]
		if shouldQuery1.Match == nil {
			t.Error("expected Match query in Should clause")
		}
		
		shouldQuery2 := query.Bool.Should[1]
		if shouldQuery2.MatchPhrasePrefix == nil {
			t.Error("expected MatchPhrasePrefix query in Should clause")
		}
	})
}

func TestMatchCompatibility(t *testing.T) {
	t.Run("should generate compatible Match query structure", func(t *testing.T) {
		query := NewQuery(Match("title", "elasticsearch search"))
		
		// Verify the structure matches what elasticsearch expects
		if query.Match == nil {
			t.Error("expected Match query")
		}
		
		matchQuery := query.Match["title"]
		if matchQuery.Query != "elasticsearch search" {
			t.Errorf("expected query 'elasticsearch search', got %s", matchQuery.Query)
		}
	})

	t.Run("should match manual Match query construction", func(t *testing.T) {
		// Our builder approach
		builderQuery := NewQuery(Match("title", "elasticsearch"))

		// Manual construction
		manualQuery := &types.Query{
			Match: map[string]types.MatchQuery{
				"title": {
					Query: "elasticsearch",
				},
			},
		}

		// Compare structures
		if builderQuery.Match == nil || manualQuery.Match == nil {
			t.Error("both queries should have Match queries")
		}
		
		if builderQuery.Match["title"].Query != manualQuery.Match["title"].Query {
			t.Errorf("Query mismatch: builder=%s, manual=%s", 
				builderQuery.Match["title"].Query, manualQuery.Match["title"].Query)
		}
	})
}

// Benchmark tests for Match queries
func BenchmarkMatchQuery(b *testing.B) {
	b.Run("Simple Match", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(Match("title", "elasticsearch search"))
		}
	})

	b.Run("Match with Options", func(b *testing.B) {
		op := operator.And
		fuzziness := types.Fuzziness("AUTO")
		for i := 0; i < b.N; i++ {
			_ = NewQuery(
				MatchWithOptions("title", "elasticsearch search", MatchOptions{
					Operator:  &op,
					Fuzziness: fuzziness,
				}),
			)
		}
	})

	b.Run("Match Phrase", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(MatchPhrase("content", "elasticsearch is awesome"))
		}
	})

	b.Run("Match Phrase Prefix", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(MatchPhrasePrefix("title", "elasticsearch sea"))
		}
	})

	b.Run("Complex Match in Bool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(
				Bool(
					Must(
						Match("title", "elasticsearch"),
						MatchPhrase("content", "search engine"),
					),
					Should(
						Match("tags", "database"),
						MatchPhrasePrefix("description", "fast sea"),
					),
				),
			)
		}
	})
} 
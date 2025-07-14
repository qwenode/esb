package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/textquerytype"
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

func TestMatchWith(t *testing.T) {
	t.Run("should create match query with callback options", func(t *testing.T) {
		analyzer := "standard"
		boost := float32(1.5)
		fuzziness := types.Fuzziness("AUTO")
		query := NewQuery(
			MatchWith("title", "elasticsearch guide", func(q *types.MatchQuery) {
				q.Fuzziness = fuzziness
				q.Analyzer = &analyzer
				q.Boost = &boost
			}),
		)
		if query.Match == nil {
			t.Error("expected Match query")
		}
		matchQuery := query.Match["title"]
		if matchQuery.Fuzziness != fuzziness {
			t.Errorf("expected fuzziness %v, got %v", fuzziness, matchQuery.Fuzziness)
		}
		if matchQuery.Analyzer == nil || *matchQuery.Analyzer != "standard" {
			t.Error("expected analyzer to be 'standard'")
		}
		if matchQuery.Boost == nil || *matchQuery.Boost != 1.5 {
			t.Error("expected boost to be 1.5")
		}
	})

	t.Run("should work with nil callback", func(t *testing.T) {
		query := NewQuery(MatchWith("title", "elasticsearch", nil))
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

func TestMatchPhraseWith(t *testing.T) {
	t.Run("should create match phrase query with callback options", func(t *testing.T) {
		analyzer := "keyword"
		boost := float32(2.0)
		slop := 2
		query := NewQuery(
			MatchPhraseWith("content", "elasticsearch is awesome", func(q *types.MatchPhraseQuery) {
				q.Slop = &slop
				q.Analyzer = &analyzer
				q.Boost = &boost
			}),
		)
		if query.MatchPhrase == nil {
			t.Error("expected MatchPhrase query")
		}
		matchPhraseQuery := query.MatchPhrase["content"]
		if matchPhraseQuery.Slop == nil || *matchPhraseQuery.Slop != 2 {
			t.Error("expected slop to be 2")
		}
		if matchPhraseQuery.Analyzer == nil || *matchPhraseQuery.Analyzer != "keyword" {
			t.Error("expected analyzer to be 'keyword'")
		}
		if matchPhraseQuery.Boost == nil || *matchPhraseQuery.Boost != 2.0 {
			t.Error("expected boost to be 2.0")
		}
	})

	t.Run("should work with nil callback", func(t *testing.T) {
		query := NewQuery(MatchPhraseWith("content", "elasticsearch is awesome", nil))
		if query.MatchPhrase == nil {
			t.Error("expected MatchPhrase query")
		}
		if query.MatchPhrase["content"].Query != "elasticsearch is awesome" {
			t.Errorf("expected query 'elasticsearch is awesome', got %s", query.MatchPhrase["content"].Query)
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

func TestMultiMatchWith(t *testing.T) {
	t.Run("should create multi match query with callback options", func(t *testing.T) {
		analyzer := "standard"
		boost := float32(1.5)
		typeVal := textquerytype.Bestfields
		query := NewQuery(
			MultiMatchWith("elasticsearch", []string{"title", "content"}, func(q *types.MultiMatchQuery) {
				q.Analyzer = &analyzer
				q.Boost = &boost
				q.Type = &typeVal
			}),
		)
		if query.MultiMatch == nil {
			t.Error("expected MultiMatch query")
		}
		if query.MultiMatch.Analyzer == nil || *query.MultiMatch.Analyzer != "standard" {
			t.Error("expected analyzer to be 'standard'")
		}
		if query.MultiMatch.Boost == nil || *query.MultiMatch.Boost != 1.5 {
			t.Error("expected boost to be 1.5")
		}
		if query.MultiMatch.Type == nil || *query.MultiMatch.Type != textquerytype.Bestfields {
			t.Error("expected type to be Bestfields")
		}
	})

	t.Run("should work with nil callback", func(t *testing.T) {
		query := NewQuery(MultiMatchWith("elasticsearch", []string{"title"}, nil))
		if query.MultiMatch == nil {
			t.Error("expected MultiMatch query")
		}
		if query.MultiMatch.Query != "elasticsearch" {
			t.Errorf("expected query 'elasticsearch', got %s", query.MultiMatch.Query)
		}
	})
}

func TestWildcardWith(t *testing.T) {
	t.Run("should create wildcard query with callback options", func(t *testing.T) {
		boost := float32(2.0)
		caseInsensitive := true
		query := NewQuery(
			WildcardWith("username", "john*", func(q *types.WildcardQuery) {
				q.Boost = &boost
				q.CaseInsensitive = &caseInsensitive
			}),
		)
		if query.Wildcard == nil {
			t.Error("expected Wildcard query")
		}
		wildcardQuery := query.Wildcard["username"]
		if wildcardQuery.Boost == nil || *wildcardQuery.Boost != 2.0 {
			t.Error("expected boost to be 2.0")
		}
		if wildcardQuery.CaseInsensitive == nil || *wildcardQuery.CaseInsensitive != true {
			t.Error("expected caseInsensitive to be true")
		}
		if wildcardQuery.Value == nil || *wildcardQuery.Value != "john*" {
			t.Errorf("expected value 'john*', got %v", wildcardQuery.Value)
		}
	})

	t.Run("should work with nil callback", func(t *testing.T) {
		query := NewQuery(WildcardWith("username", "john*", nil))
		if query.Wildcard == nil {
			t.Error("expected Wildcard query")
		}
		if query.Wildcard["username"].Value == nil || *query.Wildcard["username"].Value != "john*" {
			t.Errorf("expected value 'john*', got %v", query.Wildcard["username"].Value)
		}
	})
}

func TestFuzzyWith(t *testing.T) {
	t.Run("should create fuzzy query with callback options", func(t *testing.T) {
		boost := float32(1.5)
		fuzziness := types.Fuzziness("AUTO")
		query := NewQuery(
			FuzzyWith("username", "john", func(q *types.FuzzyQuery) {
				q.Fuzziness = fuzziness
				q.Boost = &boost
			}),
		)
		if query.Fuzzy == nil {
			t.Error("expected Fuzzy query")
		}
		fuzzyQuery := query.Fuzzy["username"]
		if fuzzyQuery.Fuzziness != fuzziness {
			t.Errorf("expected fuzziness %v, got %v", fuzziness, fuzzyQuery.Fuzziness)
		}
		if fuzzyQuery.Boost == nil || *fuzzyQuery.Boost != 1.5 {
			t.Error("expected boost to be 1.5")
		}
	})

	t.Run("should work with nil callback", func(t *testing.T) {
		query := NewQuery(FuzzyWith("username", "john", nil))
		if query.Fuzzy == nil {
			t.Error("expected Fuzzy query")
		}
		if query.Fuzzy["username"].Value != "john" {
			t.Errorf("expected value 'john', got %s", query.Fuzzy["username"].Value)
		}
	})
}

func TestPrefixWith(t *testing.T) {
	t.Run("should create prefix query with callback options", func(t *testing.T) {
		boost := float32(1.5)
		caseInsensitive := true
		query := NewQuery(
			PrefixWith("username", "john", func(q *types.PrefixQuery) {
				q.Boost = &boost
				q.CaseInsensitive = &caseInsensitive
			}),
		)
		if query.Prefix == nil {
			t.Error("expected Prefix query")
		}
		prefixQuery := query.Prefix["username"]
		if prefixQuery.Boost == nil || *prefixQuery.Boost != 1.5 {
			t.Error("expected boost to be 1.5")
		}
		if prefixQuery.CaseInsensitive == nil || *prefixQuery.CaseInsensitive != true {
			t.Error("expected caseInsensitive to be true")
		}
	})

	t.Run("should work with nil callback", func(t *testing.T) {
		query := NewQuery(PrefixWith("username", "john", nil))
		if query.Prefix == nil {
			t.Error("expected Prefix query")
		}
		if query.Prefix["username"].Value != "john" {
			t.Errorf("expected value 'john', got %s", query.Prefix["username"].Value)
		}
	})
}

func TestMatchWithOptionsFunc(t *testing.T) {
	t.Run("should create match query with callback options", func(t *testing.T) {
		analyzer := "standard"
		boost := float32(1.5)
		fuzziness := types.Fuzziness("AUTO")
		query := NewQuery(
			MatchWithOptionsFunc("title", "elasticsearch guide", func(opts *types.MatchQuery) {
				opts.Fuzziness = fuzziness
				opts.Analyzer = &analyzer
				opts.Boost = &boost
			}),
		)
		if query.Match == nil {
			t.Error("expected Match query")
		}
		matchQuery := query.Match["title"]
		if matchQuery.Fuzziness != fuzziness {
			t.Errorf("expected fuzziness %v, got %v", fuzziness, matchQuery.Fuzziness)
		}
		if matchQuery.Analyzer == nil || *matchQuery.Analyzer != "standard" {
			t.Error("expected analyzer to be 'standard'")
		}
		if matchQuery.Boost == nil || *matchQuery.Boost != 1.5 {
			t.Error("expected boost to be 1.5")
		}
	})

	t.Run("should work with nil callback", func(t *testing.T) {
		query := NewQuery(MatchWithOptionsFunc("title", "elasticsearch", nil))
		if query.Match == nil {
			t.Error("expected Match query")
		}
		if query.Match["title"].Query != "elasticsearch" {
			t.Errorf("expected query 'elasticsearch', got %s", query.Match["title"].Query)
		}
	})
}

func TestMatchPhraseWithOptionsFunc(t *testing.T) {
	t.Run("should create match phrase query with callback options", func(t *testing.T) {
		analyzer := "keyword"
		boost := float32(2.0)
		slop := 2
		query := NewQuery(
			MatchPhraseWithOptionsFunc("content", "elasticsearch is awesome", func(opts *types.MatchPhraseQuery) {
				opts.Slop = &slop
				opts.Analyzer = &analyzer
				opts.Boost = &boost
			}),
		)
		if query.MatchPhrase == nil {
			t.Error("expected MatchPhrase query")
		}
		matchPhraseQuery := query.MatchPhrase["content"]
		if matchPhraseQuery.Slop == nil || *matchPhraseQuery.Slop != 2 {
			t.Error("expected slop to be 2")
		}
		if matchPhraseQuery.Analyzer == nil || *matchPhraseQuery.Analyzer != "keyword" {
			t.Error("expected analyzer to be 'keyword'")
		}
		if matchPhraseQuery.Boost == nil || *matchPhraseQuery.Boost != 2.0 {
			t.Error("expected boost to be 2.0")
		}
	})

	t.Run("should work with nil callback", func(t *testing.T) {
		query := NewQuery(MatchPhraseWithOptionsFunc("content", "elasticsearch is awesome", nil))
		if query.MatchPhrase == nil {
			t.Error("expected MatchPhrase query")
		}
		if query.MatchPhrase["content"].Query != "elasticsearch is awesome" {
			t.Errorf("expected query 'elasticsearch is awesome', got %s", query.MatchPhrase["content"].Query)
		}
	})
}

func TestMultiMatchWithOptionsFunc(t *testing.T) {
	t.Run("should create multi match query with callback options", func(t *testing.T) {
		analyzer := "standard"
		boost := float32(1.5)
		typeVal := textquerytype.Bestfields
		query := NewQuery(
			MultiMatchWithOptionsFunc("elasticsearch", []string{"title", "content"}, func(opts *types.MultiMatchQuery) {
				opts.Analyzer = &analyzer
				opts.Boost = &boost
				opts.Type = &typeVal
			}),
		)
		if query.MultiMatch == nil {
			t.Error("expected MultiMatch query")
		}
		if query.MultiMatch.Analyzer == nil || *query.MultiMatch.Analyzer != "standard" {
			t.Error("expected analyzer to be 'standard'")
		}
		if query.MultiMatch.Boost == nil || *query.MultiMatch.Boost != 1.5 {
			t.Error("expected boost to be 1.5")
		}
		if query.MultiMatch.Type == nil || *query.MultiMatch.Type != textquerytype.Bestfields {
			t.Error("expected type to be Bestfields")
		}
	})

	t.Run("should work with nil callback", func(t *testing.T) {
		query := NewQuery(MultiMatchWithOptionsFunc("elasticsearch", []string{"title"}, nil))
		if query.MultiMatch == nil {
			t.Error("expected MultiMatch query")
		}
		if query.MultiMatch.Query != "elasticsearch" {
			t.Errorf("expected query 'elasticsearch', got %s", query.MultiMatch.Query)
		}
	})
}

func TestWildcardWithOptionsFunc(t *testing.T) {
	t.Run("should create wildcard query with callback options", func(t *testing.T) {
		boost := float32(2.0)
		caseInsensitive := true
		query := NewQuery(
			WildcardWithOptionsFunc("username", "john*", func(opts *types.WildcardQuery) {
				opts.Boost = &boost
				opts.CaseInsensitive = &caseInsensitive
			}),
		)
		if query.Wildcard == nil {
			t.Error("expected Wildcard query")
		}
		wildcardQuery := query.Wildcard["username"]
		if wildcardQuery.Boost == nil || *wildcardQuery.Boost != 2.0 {
			t.Error("expected boost to be 2.0")
		}
		if wildcardQuery.CaseInsensitive == nil || *wildcardQuery.CaseInsensitive != true {
			t.Error("expected caseInsensitive to be true")
		}
		if wildcardQuery.Value == nil || *wildcardQuery.Value != "john*" {
			t.Errorf("expected value 'john*', got %v", wildcardQuery.Value)
		}
	})

	t.Run("should work with nil callback", func(t *testing.T) {
		query := NewQuery(WildcardWithOptionsFunc("username", "john*", nil))
		if query.Wildcard == nil {
			t.Error("expected Wildcard query")
		}
		if query.Wildcard["username"].Value == nil || *query.Wildcard["username"].Value != "john*" {
			t.Errorf("expected value 'john*', got %v", query.Wildcard["username"].Value)
		}
	})
}

func TestFuzzyWithOptionsFunc(t *testing.T) {
	t.Run("should create fuzzy query with callback options", func(t *testing.T) {
		boost := float32(1.5)
		fuzziness := types.Fuzziness("AUTO")
		query := NewQuery(
			FuzzyWithOptionsFunc("username", "john", func(opts *types.FuzzyQuery) {
				opts.Fuzziness = fuzziness
				opts.Boost = &boost
			}),
		)
		if query.Fuzzy == nil {
			t.Error("expected Fuzzy query")
		}
		fuzzyQuery := query.Fuzzy["username"]
		if fuzzyQuery.Fuzziness != fuzziness {
			t.Errorf("expected fuzziness %v, got %v", fuzziness, fuzzyQuery.Fuzziness)
		}
		if fuzzyQuery.Boost == nil || *fuzzyQuery.Boost != 1.5 {
			t.Error("expected boost to be 1.5")
		}
	})

	t.Run("should work with nil callback", func(t *testing.T) {
		query := NewQuery(FuzzyWithOptionsFunc("username", "john", nil))
		if query.Fuzzy == nil {
			t.Error("expected Fuzzy query")
		}
		if query.Fuzzy["username"].Value != "john" {
			t.Errorf("expected value 'john', got %s", query.Fuzzy["username"].Value)
		}
	})
}

func TestPrefixWithOptionsFunc(t *testing.T) {
	t.Run("should create prefix query with callback options", func(t *testing.T) {
		boost := float32(1.5)
		caseInsensitive := true
		query := NewQuery(
			PrefixWithOptionsFunc("username", "john", func(opts *types.PrefixQuery) {
				opts.Boost = &boost
				opts.CaseInsensitive = &caseInsensitive
			}),
		)
		if query.Prefix == nil {
			t.Error("expected Prefix query")
		}
		prefixQuery := query.Prefix["username"]
		if prefixQuery.Boost == nil || *prefixQuery.Boost != 1.5 {
			t.Error("expected boost to be 1.5")
		}
		if prefixQuery.CaseInsensitive == nil || *prefixQuery.CaseInsensitive != true {
			t.Error("expected caseInsensitive to be true")
		}
	})

	t.Run("should work with nil callback", func(t *testing.T) {
		query := NewQuery(PrefixWithOptionsFunc("username", "john", nil))
		if query.Prefix == nil {
			t.Error("expected Prefix query")
		}
		if query.Prefix["username"].Value != "john" {
			t.Errorf("expected value 'john', got %s", query.Prefix["username"].Value)
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
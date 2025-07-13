package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"
)

// Match creates a match query for full-text search.
// Match queries analyze the query text and find documents that match.
//
// Example:
//   esb.Match("title", "elasticsearch search")
func Match(field, query string) QueryOption {
	return func(q *types.Query) error {
		if err := validateField(field); err != nil {
			return err
		}
		if err := validateValue(query); err != nil {
			return err
		}
		
		if q.Match == nil {
			q.Match = make(map[string]types.MatchQuery)
		}
		
		q.Match[field] = types.MatchQuery{
			Query: query,
		}
		
		return nil
	}
}

// MatchOptions represents options for configuring a Match query.
type MatchOptions struct {
	Operator                        *operator.Operator
	MinimumShouldMatch              *types.MinimumShouldMatch
	Fuzziness                       types.Fuzziness
	FuzzyTranspositions             *bool
	FuzzyRewrite                    *string
	Lenient                         *bool
	Analyzer                        *string
	AutoGenerateSynonymsPhraseQuery *bool
	Boost                           *float32
	CutoffFrequency                 *types.Float64
	MaxExpansions                   *int
	PrefixLength                    *int
	ZeroTermsQuery                  *zerotermsquery.ZeroTermsQuery
}

// MatchWithOptions creates a match query with advanced options.
// This allows for more complex full-text search configurations.
//
// Example:
//   esb.MatchWithOptions("title", "elasticsearch search", esb.MatchOptions{
//       Operator: &types.OperatorAnd,
//       Fuzziness: &types.FuzzinessAuto,
//       MinimumShouldMatch: &types.MinimumShouldMatch("75%"),
//   })
func MatchWithOptions(field, query string, options MatchOptions) QueryOption {
	return func(q *types.Query) error {
		if err := validateField(field); err != nil {
			return err
		}
		if err := validateValue(query); err != nil {
			return err
		}
		
		if q.Match == nil {
			q.Match = make(map[string]types.MatchQuery)
		}
		
		matchQuery := types.MatchQuery{
			Query: query,
		}
		
		// Apply options if provided
		if options.Operator != nil {
			matchQuery.Operator = options.Operator
		}
		if options.MinimumShouldMatch != nil {
			matchQuery.MinimumShouldMatch = *options.MinimumShouldMatch
		}
		if options.Fuzziness != nil {
			matchQuery.Fuzziness = options.Fuzziness
		}
		if options.FuzzyTranspositions != nil {
			matchQuery.FuzzyTranspositions = options.FuzzyTranspositions
		}
		if options.FuzzyRewrite != nil {
			matchQuery.FuzzyRewrite = options.FuzzyRewrite
		}
		if options.Lenient != nil {
			matchQuery.Lenient = options.Lenient
		}
		if options.Analyzer != nil {
			matchQuery.Analyzer = options.Analyzer
		}
		if options.AutoGenerateSynonymsPhraseQuery != nil {
			matchQuery.AutoGenerateSynonymsPhraseQuery = options.AutoGenerateSynonymsPhraseQuery
		}
		if options.Boost != nil {
			matchQuery.Boost = options.Boost
		}
		if options.CutoffFrequency != nil {
			matchQuery.CutoffFrequency = options.CutoffFrequency
		}
		if options.MaxExpansions != nil {
			matchQuery.MaxExpansions = options.MaxExpansions
		}
		if options.PrefixLength != nil {
			matchQuery.PrefixLength = options.PrefixLength
		}
		if options.ZeroTermsQuery != nil {
			matchQuery.ZeroTermsQuery = options.ZeroTermsQuery
		}
		
		q.Match[field] = matchQuery
		
		return nil
	}
}

// MatchPhrase creates a match phrase query for exact phrase matching.
// Match phrase queries find documents containing the exact phrase.
//
// Example:
//   esb.MatchPhrase("content", "elasticsearch is awesome")
func MatchPhrase(field, phrase string) QueryOption {
	return func(q *types.Query) error {
		if err := validateField(field); err != nil {
			return err
		}
		if err := validateValue(phrase); err != nil {
			return err
		}
		
		if q.MatchPhrase == nil {
			q.MatchPhrase = make(map[string]types.MatchPhraseQuery)
		}
		
		q.MatchPhrase[field] = types.MatchPhraseQuery{
			Query: phrase,
		}
		
		return nil
	}
}

// MatchPhraseOptions represents options for configuring a MatchPhrase query.
type MatchPhraseOptions struct {
	Slop     *int
	Analyzer *string
	Boost    *float32
}

// MatchPhraseWithOptions creates a match phrase query with advanced options.
//
// Example:
//   esb.MatchPhraseWithOptions("content", "elasticsearch search", esb.MatchPhraseOptions{
//       Slop: intPtr(2),
//       Analyzer: stringPtr("standard"),
//   })
func MatchPhraseWithOptions(field, phrase string, options MatchPhraseOptions) QueryOption {
	return func(q *types.Query) error {
		if err := validateField(field); err != nil {
			return err
		}
		if err := validateValue(phrase); err != nil {
			return err
		}
		
		if q.MatchPhrase == nil {
			q.MatchPhrase = make(map[string]types.MatchPhraseQuery)
		}
		
		matchPhraseQuery := types.MatchPhraseQuery{
			Query: phrase,
		}
		
		// Apply options if provided
		if options.Slop != nil {
			matchPhraseQuery.Slop = options.Slop
		}
		if options.Analyzer != nil {
			matchPhraseQuery.Analyzer = options.Analyzer
		}
		if options.Boost != nil {
			matchPhraseQuery.Boost = options.Boost
		}
		
		q.MatchPhrase[field] = matchPhraseQuery
		
		return nil
	}
}

// MatchPhrasePrefix creates a match phrase prefix query.
// This is useful for autocomplete and "search as you type" functionality.
//
// Example:
//   esb.MatchPhrasePrefix("title", "elasticsearch sea")
func MatchPhrasePrefix(field, prefix string) QueryOption {
	return func(q *types.Query) error {
		if err := validateField(field); err != nil {
			return err
		}
		if err := validateValue(prefix); err != nil {
			return err
		}
		
		if q.MatchPhrasePrefix == nil {
			q.MatchPhrasePrefix = make(map[string]types.MatchPhrasePrefixQuery)
		}
		
		q.MatchPhrasePrefix[field] = types.MatchPhrasePrefixQuery{
			Query: prefix,
		}
		
		return nil
	}
}

// Helper functions for creating pointers to basic types
func intPtr(i int) *int {
	return &i
}

func stringPtr(s string) *string {
	return &s
}

func float32Ptr(f float32) *float32 {
	return &f
}

func boolPtr(b bool) *bool {
	return &b
} 
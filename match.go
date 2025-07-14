package esb

import (
    "github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Match creates a match query for full-text search.
// Match queries analyze the query text and find documents that match.
//
// Example:
//   esb.Match("title", "elasticsearch search")
func Match(field, query string) QueryOption {
    return func(q *types.Query) {
        q.Match = map[string]types.MatchQuery{
            field: {
                Query: query,
            },
        }
    }
}

// MatchWithOptions 提供回调函数式的 Match 查询配置。
// 示例：
//   esb.MatchWith("title", "elasticsearch guide", func(q *types.MatchQuery) {
//       q.Fuzziness = "AUTO"
//       q.Analyzer = &analyzer
//       q.Boost = &boost
//   })
func MatchWithOptions(field, query string, setOpts func(opts *types.MatchQuery)) QueryOption {
    return func(q *types.Query) {
        matchQuery := types.MatchQuery{
            Query: query,
        }
        if setOpts != nil {
            setOpts(&matchQuery)
        }
        q.Match = map[string]types.MatchQuery{
            field: matchQuery,
        }
    }
}

// MatchPhrase creates a match phrase query for exact phrase matching.
// Match phrase queries find documents containing the exact phrase.
//
// Example:
//   esb.MatchPhrase("content", "elasticsearch is awesome")
func MatchPhrase(field, phrase string) QueryOption {
    return func(q *types.Query) {
        q.MatchPhrase = map[string]types.MatchPhraseQuery{
            field: {
                Query: phrase,
            },
        }
    }
}

// MatchPhraseWithOptions 提供回调函数式的 MatchPhrase 查询配置。
// 示例：
//   esb.MatchPhraseWith("content", "elasticsearch is awesome", func(q *types.MatchPhraseQuery) {
//       q.Slop = &slop
//       q.Analyzer = &analyzer
//       q.Boost = &boost
//   })
func MatchPhraseWithOptions(field, phrase string, setOpts func(opts *types.MatchPhraseQuery)) QueryOption {
    return func(q *types.Query) {
        matchPhraseQuery := types.MatchPhraseQuery{
            Query: phrase,
        }
        if setOpts != nil {
            setOpts(&matchPhraseQuery)
        }
        q.MatchPhrase = map[string]types.MatchPhraseQuery{
            field: matchPhraseQuery,
        }
    }
}

// MatchPhrasePrefix creates a match phrase prefix query.
// This is useful for autocomplete and "search as you type" functionality.
//
// Example:
//   esb.MatchPhrasePrefix("title", "elasticsearch sea")
func MatchPhrasePrefix(field, prefix string) QueryOption {
    return func(q *types.Query) {
        q.MatchPhrasePrefix = map[string]types.MatchPhrasePrefixQuery{
            field: {
                Query: prefix,
            },
        }
    }
}

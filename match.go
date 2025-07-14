package esb

import (
    "github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Match 创建一个用于全文搜索的匹配查询。
// Match 查询会分析查询文本并查找匹配的文档。
//
// 示例：
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

// MatchPhrase 创建一个用于精确短语匹配的查询。
// Match phrase 查询会查找包含完整短语的文档。
//
// 示例：
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

// MatchPhrasePrefix 创建一个短语前缀匹配查询。
// 这对于自动补全和"输入即搜索"功能非常有用。
//
// 示例：
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

// MatchPhrasePrefixWithOptions 提供回调函数式的 MatchPhrasePrefix 查询配置。
// 示例：
//   esb.MatchPhrasePrefixWithOptions("title", "elasticsearch sea", func(q *types.MatchPhrasePrefixQuery) {
//       q.MaxExpansions = &maxExpansions
//       q.Analyzer = &analyzer
//       q.Slop = &slop
//   })
func MatchPhrasePrefixWithOptions(field, prefix string, setOpts func(opts *types.MatchPhrasePrefixQuery)) QueryOption {
    return func(q *types.Query) {
        matchPhrasePrefixQuery := types.MatchPhrasePrefixQuery{
            Query: prefix,
        }
        if setOpts != nil {
            setOpts(&matchPhrasePrefixQuery)
        }
        q.MatchPhrasePrefix = map[string]types.MatchPhrasePrefixQuery{
            field: matchPhrasePrefixQuery,
        }
    }
}

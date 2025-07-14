package esb

import (
    "github.com/elastic/go-elasticsearch/v8/typedapi/types"
    "github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/textquerytype"
)

// MultiMatch 创建多字段匹配查询，允许在多个字段中搜索文本。
// 该查询会分析提供的文本，然后在指定的字段中进行匹配。
//
// 示例：
//   esb.MultiMatch("elasticsearch", "title", "content")          // 在标题和内容字段中搜索
//   esb.MultiMatch("john doe", "first_name", "last_name")        // 在姓名字段中搜索
//   esb.MultiMatch("java programming", "title^2", "content")     // 标题字段权重为2
func MultiMatch(query string, fields ...string) QueryOption {
    return func(q *types.Query) {
        q.MultiMatch = &types.MultiMatchQuery{
            Query:  query,
            Fields: fields,
        }
    }
}

// MultiMatchSlice 创建一个多字段匹配查询，功能与MultiMatch相同，但接收切片参数而不是可变参数。
// MultiMatchSlice查询用于在多个字段中搜索相同的文本。
//
// 示例：
//   esb.MultiMatchSlice("quick brown fox", []string{"title", "description"})
func MultiMatchSlice(query string, fields []string) QueryOption {
	return func(q *types.Query) {
		q.MultiMatch = &types.MultiMatchQuery{
			Query:  query,
			Fields: fields,
		}
	}
}

// MultiMatchWithOptions 提供回调函数式的 MultiMatch 查询配置。
func MultiMatchWithOptions(query string, fields []string, setOpts func(opts *types.MultiMatchQuery)) QueryOption {
    return func(qy *types.Query) {
        multiMatchQuery := &types.MultiMatchQuery{
            Query:  query,
            Fields: fields,
        }
        if setOpts != nil {
            setOpts(multiMatchQuery)
        }
        qy.MultiMatch = multiMatchQuery
    }
}

// MultiMatchBestFields 创建 best_fields 类型的多字段匹配查询。
// 这是默认类型，查找匹配任何字段的文档，但使用最佳字段的分数。
//
// 示例：
//   esb.MultiMatchBestFields("java programming", "title", "content")
func MultiMatchBestFields(query string, fields ...string) QueryOption {
    return MultiMatchWithOptions(
        query, fields, func(opts *types.MultiMatchQuery) {
            opts.Type = &textquerytype.Bestfields
        },
    )
}

// MultiMatchBestFieldsSlice 创建一个best_fields类型的多字段匹配查询，功能与MultiMatchBestFields相同，但接收切片参数而不是可变参数。
// MultiMatchBestFieldsSlice查询会找到至少一个字段完全匹配的文档。
//
// 示例：
//   esb.MultiMatchBestFieldsSlice("quick brown fox", []string{"title", "description"})
func MultiMatchBestFieldsSlice(query string, fields []string) QueryOption {
	return func(q *types.Query) {
		q.MultiMatch = &types.MultiMatchQuery{
			Query:  query,
			Fields: fields,
			Type:   &textquerytype.Bestfields,
		}
	}
}

// MultiMatchMostFields 创建 most_fields 类型的多字段匹配查询。
// 查找匹配任何字段的文档，并结合每个字段的分数。
//
// 示例：
//   esb.MultiMatchMostFields("java programming", "title", "content", "tags")
func MultiMatchMostFields(query string, fields ...string) QueryOption {
    return MultiMatchWithOptions(
        query, fields, func(opts *types.MultiMatchQuery) {
            opts.Type = &textquerytype.Mostfields
        },
    )
}

// MultiMatchMostFieldsSlice 创建一个most_fields类型的多字段匹配查询，功能与MultiMatchMostFields相同，但接收切片参数而不是可变参数。
// MultiMatchMostFieldsSlice查询会匹配多个字段并合并分数。
//
// 示例：
//   esb.MultiMatchMostFieldsSlice("quick brown fox", []string{"title", "description"})
func MultiMatchMostFieldsSlice(query string, fields []string) QueryOption {
	return func(q *types.Query) {
		q.MultiMatch = &types.MultiMatchQuery{
			Query:  query,
			Fields: fields,
			Type:   &textquerytype.Mostfields,
		}
	}
}

// MultiMatchCrossFields 创建 cross_fields 类型的多字段匹配查询。
// 将字段视为一个大字段，查找每个词条在任何字段中的匹配。
//
// 示例：
//   esb.MultiMatchCrossFields("john doe", "first_name", "last_name")
func MultiMatchCrossFields(query string, fields ...string) QueryOption {
    return MultiMatchWithOptions(
        query, fields, func(opts *types.MultiMatchQuery) {
            opts.Type = &textquerytype.Crossfields
        },
    )
}

// MultiMatchCrossFieldsSlice 创建一个cross_fields类型的多字段匹配查询，功能与MultiMatchCrossFields相同，但接收切片参数而不是可变参数。
// MultiMatchCrossFieldsSlice查询会在多个字段中寻找词项。
//
// 示例：
//   esb.MultiMatchCrossFieldsSlice("quick brown fox", []string{"title", "description"})
func MultiMatchCrossFieldsSlice(query string, fields []string) QueryOption {
	return func(q *types.Query) {
		q.MultiMatch = &types.MultiMatchQuery{
			Query:  query,
			Fields: fields,
			Type:   &textquerytype.Crossfields,
		}
	}
}

// MultiMatchPhrase 创建 phrase 类型的多字段匹配查询。
// 对每个字段运行 match_phrase 查询，使用最佳字段的分数。
//
// 示例：
//   esb.MultiMatchPhrase("elasticsearch guide", "title", "content")
func MultiMatchPhrase(query string, fields ...string) QueryOption {
    return MultiMatchWithOptions(
        query, fields, func(opts *types.MultiMatchQuery) {
            opts.Type = &textquerytype.Phrase
        },
    )
}

// MultiMatchPhraseSlice 创建一个phrase类型的多字段匹配查询，功能与MultiMatchPhrase相同，但接收切片参数而不是可变参数。
// MultiMatchPhraseSlice查询会在多个字段中寻找短语。
//
// 示例：
//   esb.MultiMatchPhraseSlice("quick brown fox", []string{"title", "description"})
func MultiMatchPhraseSlice(query string, fields []string) QueryOption {
	return func(q *types.Query) {
		q.MultiMatch = &types.MultiMatchQuery{
			Query:  query,
			Fields: fields,
			Type:   &textquerytype.Phrase,
		}
	}
}

// MultiMatchPhrasePrefix 创建 phrase_prefix 类型的多字段匹配查询。
// 对每个字段运行 match_phrase_prefix 查询，使用最佳字段的分数。
//
// 示例：
//   esb.MultiMatchPhrasePrefix("elasticsearch sea", "title", "content")
func MultiMatchPhrasePrefix(query string, fields ...string) QueryOption {
    return MultiMatchWithOptions(
        query, fields, func(opts *types.MultiMatchQuery) {
            opts.Type = &textquerytype.Phraseprefix
        },
    )
}

// MultiMatchPhrasePrefixSlice 创建一个phrase_prefix类型的多字段匹配查询，功能与MultiMatchPhrasePrefix相同，但接收切片参数而不是可变参数。
// MultiMatchPhrasePrefixSlice查询会在多个字段中寻找以指定词项开头的短语。
//
// 示例：
//   esb.MultiMatchPhrasePrefixSlice("quick brown", []string{"title", "description"})
func MultiMatchPhrasePrefixSlice(query string, fields []string) QueryOption {
	return func(q *types.Query) {
		q.MultiMatch = &types.MultiMatchQuery{
			Query:  query,
			Fields: fields,
			Type:   &textquerytype.Phraseprefix,
		}
	}
}

package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// MoreLikeThis 创建一个相似度查询，用于查找与给定文档相似的文档。
//
// 示例：
//   esb.MoreLikeThis([]string{"title", "content"}, "This is a sample text")
func MoreLikeThis(fields []string, likeText string) QueryOption {
	return func(q *types.Query) {
		q.MoreLikeThis = &types.MoreLikeThisQuery{
			Fields: fields,
			Like:   []types.Like{likeText},
		}
	}
}

// MoreLikeThisWithOptions 提供回调函数式的相似度查询配置。
//
// 示例：
//   esb.MoreLikeThisWithOptions([]string{"title", "content"}, "This is a sample text", func(opts *types.MoreLikeThisQuery) {
//       minTermFreq := 2
//       opts.MinTermFreq = &minTermFreq
//       maxQueryTerms := 25
//       opts.MaxQueryTerms = &maxQueryTerms
//   })
func MoreLikeThisWithOptions(fields []string, likeText string, setOpts func(opts *types.MoreLikeThisQuery)) QueryOption {
	return func(q *types.Query) {
		moreLikeThisQuery := &types.MoreLikeThisQuery{
			Fields: fields,
			Like:   []types.Like{likeText},
		}
		
		if setOpts != nil {
			setOpts(moreLikeThisQuery)
		}
		
		q.MoreLikeThis = moreLikeThisQuery
	}
}

// MoreLikeThisWithDocument 创建一个基于文档的相似度查询。
//
// 示例：
//   esb.MoreLikeThisWithDocument([]string{"title", "content"}, "my-index", "1")
func MoreLikeThisWithDocument(fields []string, index, id string) QueryOption {
	return func(q *types.Query) {
		like := types.LikeDocument{
			Index_: &index,
			Id_:    &id,
		}
		
		q.MoreLikeThis = &types.MoreLikeThisQuery{
			Fields: fields,
			Like:   []types.Like{like},
		}
	}
}

// MoreLikeThisWithMultipleLikes 创建一个支持多个相似对象的查询。
//
// 示例：
//   likes := []types.Like{"text1", "text2"}
//   esb.MoreLikeThisWithMultipleLikes([]string{"title", "content"}, likes)
func MoreLikeThisWithMultipleLikes(fields []string, likes []types.Like) QueryOption {
	return func(q *types.Query) {
		q.MoreLikeThis = &types.MoreLikeThisQuery{
			Fields: fields,
			Like:   likes,
		}
	}
}

// MoreLikeThisWithUnlike 创建一个带有排除条件的相似度查询。
//
// 示例：
//   esb.MoreLikeThisWithUnlike([]string{"title", "content"}, "This is a sample text", "unwanted text")
func MoreLikeThisWithUnlike(fields []string, likeText string, unlikeText string) QueryOption {
	return func(q *types.Query) {
		q.MoreLikeThis = &types.MoreLikeThisQuery{
			Fields: fields,
			Like:   []types.Like{likeText},
			Unlike: []types.Like{unlikeText},
		}
	}
} 
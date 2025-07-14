package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Fuzzy 创建模糊查询，返回包含与搜索词相似的词条的文档。
// 模糊查询基于 Levenshtein 编辑距离来衡量词条的相似性。
//
// 示例：
//   esb.Fuzzy("username", "john")    // 匹配 "john", "jhon", "joh" 等相似词条
//   esb.Fuzzy("title", "elasticsearch") // 匹配 "elasticsearch", "elasticsearh" 等
//   esb.Fuzzy("product", "iphone")   // 匹配 "iphone", "iphon", "ipone" 等
func Fuzzy(field, value string) QueryOption {
	return func(q *types.Query) {
		q.Fuzzy = map[string]types.FuzzyQuery{
			field: {
				Value: value,
			},
		}
	}
}



 

// FuzzyWithOptions 提供回调函数式的 Fuzzy 查询配置。
func FuzzyWithOptions(field, value string, setOpts func(opts *types.FuzzyQuery)) QueryOption {
    return func(q *types.Query) {
        fuzzyQuery := types.FuzzyQuery{
            Value: value,
        }
        if setOpts != nil {
            setOpts(&fuzzyQuery)
        }
        q.Fuzzy = map[string]types.FuzzyQuery{
            field: fuzzyQuery,
        }
    }
} 
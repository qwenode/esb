package esb

import (
    "github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Nested 创建一个嵌套查询。
// 此查询用于搜索由对象数组组成的嵌套字段。
//
// 示例：
//
//	query := esb.NewQuery(
//		esb.Nested("comments", // 嵌套字段的路径
//			esb.Match("comments.author", "john"), // 对嵌套字段的查询
//		),
//	)
func Nested(path string, query QueryOption) QueryOption {
    return func(q *types.Query) {
        nestedQuery := &types.Query{}
        query(nestedQuery)
        
        q.Nested = &types.NestedQuery{
            Path:  path,
            Query: nestedQuery,
        }
    }
}

// NestedWithOptions 创建一个带有附加选项的嵌套查询。
//
// 示例：
//
//	query := esb.NewQuery(
//		esb.NestedWithOptions("comments",
//			esb.Match("comments.author", "john"),
//			func(opts *types.NestedQuery) {
//				scoreMode := types.NestedScoremodeAvg
//				ignoreUnmapped := true
//				opts.ScoreMode = &scoreMode
//				opts.IgnoreUnmapped = &ignoreUnmapped
//			},
//		),
//	)
func NestedWithOptions(path string, query QueryOption, setOpts func(opts *types.NestedQuery)) QueryOption {
    return func(q *types.Query) {
        nestedQuery := &types.Query{}
        query(nestedQuery)
        
        nested := &types.NestedQuery{
            Path:  path,
            Query: nestedQuery,
        }
        
        if setOpts != nil {
            setOpts(nested)
        }
        
        q.Nested = nested
    }
}

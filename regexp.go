package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Regexp 创建一个正则表达式查询。
// 此查询允许使用正则表达式来匹配字段值。
//
// 示例：
//
//	query := esb.NewQuery(
//		esb.Regexp("username", "j.*n"),
//	)
func Regexp(field string, value string) QueryOption {
	return func(q *types.Query) {
		if q.Regexp == nil {
			q.Regexp = make(map[string]types.RegexpQuery)
		}
		q.Regexp[field] = types.RegexpQuery{
			Value: value,
		}
	}
}

// RegexpWithOptions 创建一个带有附加选项的正则表达式查询。
//
// 示例：
//
//	query := esb.NewQuery(
//		esb.RegexpWithOptions("username", "j.*n", func(opts *types.RegexpQuery) {
//			flags := "ALL"
//			maxDeterminizedStates := 10000
//			opts.Flags = &flags
//			opts.MaxDeterminizedStates = &maxDeterminizedStates
//		}),
//	)
func RegexpWithOptions(field string, value string, setOpts func(opts *types.RegexpQuery)) QueryOption {
	return func(q *types.Query) {
		if q.Regexp == nil {
			q.Regexp = make(map[string]types.RegexpQuery)
		}
		
		regexpQuery := types.RegexpQuery{
			Value: value,
		}
		if setOpts != nil {
			setOpts(&regexpQuery)
		}
		q.Regexp[field] = regexpQuery
	}
} 
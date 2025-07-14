package esb

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"
)

// Script 创建一个脚本查询，用于执行自定义脚本来匹配文档。
//
// 示例：
//   esb.Script("doc['age'].value > 30")
func Script(source string) QueryOption {
	return func(q *types.Query) {
		q.Script = &types.ScriptQuery{
			Script: types.Script{
				Source: &source,
			},
		}
	}
}

// ScriptWithParams 创建一个带参数的脚本查询。
//
// 示例：
//   params := map[string]any{"min_age": 30}
//   esb.ScriptWithParams("doc['age'].value > params.min_age", params)
func ScriptWithParams(source string, params map[string]any) QueryOption {
	return func(q *types.Query) {
		// 将参数转换为json.RawMessage
		jsonParams := make(map[string]json.RawMessage)
		for k, v := range params {
			jsonBytes, _ := json.Marshal(v)
			jsonParams[k] = json.RawMessage(jsonBytes)
		}
		
		q.Script = &types.ScriptQuery{
			Script: types.Script{
				Source: &source,
				Params: jsonParams,
			},
		}
	}
}

// ScriptWithLang 创建一个指定语言的脚本查询。
//
// 示例：
//   esb.ScriptWithLang("doc['age'].value > 30", scriptlanguage.Painless)
func ScriptWithLang(source string, lang scriptlanguage.ScriptLanguage) QueryOption {
	return func(q *types.Query) {
		q.Script = &types.ScriptQuery{
			Script: types.Script{
				Source: &source,
				Lang:   &lang,
			},
		}
	}
}

// ScriptWithOptions 提供回调函数式的脚本查询配置。
//
// 示例：
//   esb.ScriptWithOptions("doc['age'].value > 30", func(opts *types.ScriptQuery) {
//       lang := scriptlanguage.Painless
//       opts.Script.Lang = &lang
//       boost := float32(1.5)
//       opts.Boost = &boost
//   })
func ScriptWithOptions(source string, setOpts func(opts *types.ScriptQuery)) QueryOption {
	return func(q *types.Query) {
		scriptQuery := &types.ScriptQuery{
			Script: types.Script{
				Source: &source,
			},
		}
		
		if setOpts != nil {
			setOpts(scriptQuery)
		}
		
		q.Script = scriptQuery
	}
}

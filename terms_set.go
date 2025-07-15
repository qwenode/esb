package esb

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

// TermsSet 创建一个Terms Set查询，使用默认的最小匹配数。
func TermsSet(field string, values []string) QueryOption {
    return func(q *types.Query) {
        
        q.TermsSet = map[string]types.TermsSetQuery{
            field: {
                Terms: values,
            },
        }
    }
}

// TermsSetWithOptions 提供一个函数式选项的Terms Set查询配置器
func TermsSetWithOptions(field string, values []string, setOpts func(*types.TermsSetQuery)) QueryOption {
    return func(q *types.Query) {
        query := types.TermsSetQuery{
            Terms: values,
        }
        // 调用配置函数自定义设置（如果不为nil）
        if setOpts != nil {
            setOpts(&query)
        }
        
        q.TermsSet = map[string]types.TermsSetQuery{
            field: query,
        }
    }
}

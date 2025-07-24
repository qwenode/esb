package esb

import (
    "encoding/json"
    "errors"

    "github.com/elastic/go-elasticsearch/v8/typedapi/core/get"
    "github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

var (
    ErrNoData = errors.New("no data")
)

func IsNoData(err error) bool {
    return errors.Is(err, ErrNoData)
}

// 解析单条数据 20250722
func FormatOne[T any](response *get.Response, err error) (T, error) {
    var v T
    if err != nil {
        return v, err
    }
    if !response.Found {
        return v, ErrNoData
    }
    err = json.Unmarshal(response.Source_, &v)
    return v, err
}

// 解析后的回调处理,如果_append=false,则数据不会返回 20250722
type FormatSearchPostProcessor[T any] func(src T) (_append bool)

// 解析多条数据返回值 20250722
func FormatSearch[T any](response types.HitsMetadata, postprocessor FormatSearchPostProcessor[T]) []T {
    if response.Hits == nil {
        return nil
    }
    hits := response.Hits
    hLen := len(hits)
    if hLen <= 0 {
        return nil
    }
    if postprocessor == nil {
        postprocessor = func(src T) (_append bool) {
            return true
        }
    }
    ts := make([]T, 0, hLen)
    for _, hit := range hits {
        var v T
        err := json.Unmarshal(hit.Source_, &v)
        if err != nil {
            continue
        }
        if postprocessor(v) {
            ts = append(ts, v)
        }
    }
    return ts
}

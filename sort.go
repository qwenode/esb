package esb

import (
    "github.com/elastic/go-elasticsearch/v8/typedapi/types"
    "github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

// 按字段正序
func SortFieldAsc(field string) *types.SortOptions {
    s := &types.SortOptions{
        SortOptions: map[string]types.FieldSort{
            field: {
                Order: &sortorder.Asc,
            },
        },
    }
    return s
}

// 按字段倒序
func SortFieldDesc(field string) *types.SortOptions {
    s := &types.SortOptions{
        SortOptions: map[string]types.FieldSort{
            field: {
                Order: &sortorder.Desc,
            },
        },
    }
    return s
}

package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/calendarinterval"
)

// AggregationOption 表示一个修改 types.Aggregations 的函数。
// 它遵循函数式选项模式来构建 Elasticsearch 聚合。
type AggregationOption func(*types.Aggregations)

// NewAggregations 通过应用提供的选项创建一个新的 Elasticsearch 聚合。
// 它返回一个可以直接用于 go-elasticsearch 客户端的 *types.Aggregations。
//
// 示例：
//   aggs := esb.NewAggregations(
//       esb.TermsAgg("categories", "category"),
//       esb.AvgAgg("avg_price", "price"),
//   )
func NewAggregations(opts ...AggregationOption) *types.Aggregations {
	aggs := &types.Aggregations{
		Aggregations: make(map[string]types.Aggregations),
	}
	for _, opt := range opts {
		opt(aggs)
	}
	return aggs
}

// TermsAgg 创建一个词项聚合，用于统计字段的不同值及其文档数量。
// Terms 聚合是最常用的桶聚合之一。
//
// 示例：
//   esb.TermsAgg("categories", "category")
//   esb.TermsAgg("categories", "category", esb.AvgAgg("avg_price", "price"))
func TermsAgg(name, field string, subAggs ...AggregationOption) AggregationOption {
	return func(aggs *types.Aggregations) {
		termsAgg := types.Aggregations{
			Terms: &types.TermsAggregation{
				Field: &field,
			},
		}
		
		// 添加子聚合
		if len(subAggs) > 0 {
			termsAgg.Aggregations = make(map[string]types.Aggregations)
			for _, subAgg := range subAggs {
				subAgg(&termsAgg)
			}
		}
		
		aggs.Aggregations[name] = termsAgg
	}
}

// TermsAggWithOptions 提供回调函数式的 Terms 聚合配置。
// 允许设置 size、order 等高级选项。
//
// 示例：
//   esb.TermsAggWithOptions("top_categories", "category", func(opts *types.TermsAggregation) {
//       size := 10
//       opts.Size = &size
//   })
//   esb.TermsAggWithOptions("top_categories", "category", func(opts *types.TermsAggregation) {
//       size := 10
//       opts.Size = &size
//   }, esb.AvgAgg("avg_price", "price"))
func TermsAggWithOptions(name, field string, setOpts func(opts *types.TermsAggregation), subAggs ...AggregationOption) AggregationOption {
	return func(aggs *types.Aggregations) {
		termsAgg := &types.TermsAggregation{
			Field: &field,
		}
		
		if setOpts != nil {
			setOpts(termsAgg)
		}
		
		aggContainer := types.Aggregations{
			Terms: termsAgg,
		}
		
		// 添加子聚合
		if len(subAggs) > 0 {
			aggContainer.Aggregations = make(map[string]types.Aggregations)
			for _, subAgg := range subAggs {
				subAgg(&aggContainer)
			}
		}
		
		aggs.Aggregations[name] = aggContainer
	}
}

// AvgAgg 创建一个平均值聚合，计算数值字段的平均值。
//
// 示例：
//   esb.AvgAgg("avg_price", "price")
func AvgAgg(name, field string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			Avg: &types.AverageAggregation{
				Field: &field,
			},
		}
	}
}

// SumAgg 创建一个求和聚合，计算数值字段的总和。
//
// 示例：
//   esb.SumAgg("total_sales", "sales")
func SumAgg(name, field string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			Sum: &types.SumAggregation{
				Field: &field,
			},
		}
	}
}

// MaxAgg 创建一个最大值聚合，找出数值字段的最大值。
//
// 示例：
//   esb.MaxAgg("max_price", "price")
func MaxAgg(name, field string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			Max: &types.MaxAggregation{
				Field: &field,
			},
		}
	}
}

// MinAgg 创建一个最小值聚合，找出数值字段的最小值。
//
// 示例：
//   esb.MinAgg("min_price", "price")
func MinAgg(name, field string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			Min: &types.MinAggregation{
				Field: &field,
			},
		}
	}
}

// StatsAgg 创建一个统计聚合，计算数值字段的统计信息（count、min、max、avg、sum）。
//
// 示例：
//   esb.StatsAgg("price_stats", "price")
func StatsAgg(name, field string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			Stats: &types.StatsAggregation{
				Field: &field,
			},
		}
	}
}

// ValueCountAgg 创建一个值计数聚合，计算字段的非空值数量。
//
// 示例：
//   esb.ValueCountAgg("field_count", "category")
func ValueCountAgg(name, field string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			ValueCount: &types.ValueCountAggregation{
				Field: &field,
			},
		}
	}
}

// CardinalityAgg 创建一个基数聚合，计算字段的唯一值数量（近似值）。
//
// 示例：
//   esb.CardinalityAgg("unique_users", "user_id")
func CardinalityAgg(name, field string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			Cardinality: &types.CardinalityAggregation{
				Field: &field,
			},
		}
	}
}

// DateHistogramAgg 创建一个日期直方图聚合，按时间间隔对文档进行分组。
//
// 示例：
//   esb.DateHistogramAgg("sales_over_time", "timestamp", "1d")
//   esb.DateHistogramAgg("sales_over_time", "timestamp", "1d", esb.SumAgg("total_sales", "amount"))
func DateHistogramAgg(name, field, interval string, subAggs ...AggregationOption) AggregationOption {
	return func(aggs *types.Aggregations) {
		calInterval := calendarinterval.CalendarInterval{Name: interval}
		dateHistAgg := types.Aggregations{
			DateHistogram: &types.DateHistogramAggregation{
				Field:            &field,
				CalendarInterval: &calInterval,
			},
		}
		
		// 添加子聚合
		if len(subAggs) > 0 {
			dateHistAgg.Aggregations = make(map[string]types.Aggregations)
			for _, subAgg := range subAggs {
				subAgg(&dateHistAgg)
			}
		}
		
		aggs.Aggregations[name] = dateHistAgg
	}
}

// DateHistogramAggWithOptions 提供回调函数式的 DateHistogram 聚合配置。
//
// 示例：
//   esb.DateHistogramAggWithOptions("sales_over_time", "timestamp", "1d", func(opts *types.DateHistogramAggregation) {
//       format := "yyyy-MM-dd"
//       opts.Format = &format
//   })
func DateHistogramAggWithOptions(name, field, interval string, setOpts func(opts *types.DateHistogramAggregation)) AggregationOption {
	return func(aggs *types.Aggregations) {
		calInterval := calendarinterval.CalendarInterval{Name: interval}
		dateHistAgg := &types.DateHistogramAggregation{
			Field:            &field,
			CalendarInterval: &calInterval,
		}
		
		if setOpts != nil {
			setOpts(dateHistAgg)
		}
		
		aggs.Aggregations[name] = types.Aggregations{
			DateHistogram: dateHistAgg,
		}
	}
}

// HistogramAgg 创建一个数值直方图聚合，按数值间隔对文档进行分组。
//
// 示例：
//   esb.HistogramAgg("price_ranges", "price", 100)
func HistogramAgg(name, field string, interval float64) AggregationOption {
	return func(aggs *types.Aggregations) {
		intervalFloat := types.Float64(interval)
		aggs.Aggregations[name] = types.Aggregations{
			Histogram: &types.HistogramAggregation{
				Field:    &field,
				Interval: &intervalFloat,
			},
		}
	}
}

// RangeAgg 创建一个范围聚合，根据指定的范围对文档进行分组。
//
// 示例：
//   esb.RangeAgg("price_ranges", "price", []types.AggregationRange{
//       {To: types.Float64(100)},
//       {From: types.Float64(100), To: types.Float64(500)},
//       {From: types.Float64(500)},
//   })
func RangeAgg(name, field string, ranges []types.AggregationRange) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			Range: &types.RangeAggregation{
				Field:  &field,
				Ranges: ranges,
			},
		}
	}
}

// FilterAgg 创建一个过滤器聚合，只对匹配查询的文档进行聚合。
//
// 示例：
//   esb.FilterAgg("expensive_products", esb.Range("price", 1000, nil))
func FilterAgg(name string, query QueryOption) AggregationOption {
	return func(aggs *types.Aggregations) {
		filterQuery := &types.Query{}
		if query != nil {
			query(filterQuery)
		}
		
		aggs.Aggregations[name] = types.Aggregations{
			Filter: filterQuery,
		}
	}
}

// FiltersAgg 创建一个多过滤器聚合，为每个过滤器创建一个桶。
//
// 示例：
//   esb.FiltersAgg("product_categories", map[string]QueryOption{
//       "electronics": esb.Term("category", "electronics"),
//       "books": esb.Term("category", "books"),
//   })
func FiltersAgg(name string, filters map[string]QueryOption) AggregationOption {
	return func(aggs *types.Aggregations) {
		namedFilters := make(map[string]*types.Query)
		for key, queryOpt := range filters {
			query := &types.Query{}
			if queryOpt != nil {
				queryOpt(query)
			}
			namedFilters[key] = query
		}
		
		aggs.Aggregations[name] = types.Aggregations{
			Filters: &types.FiltersAggregation{
				Filters: types.BucketsQuery(namedFilters),
			},
		}
	}
}

// NestedAgg 创建一个嵌套聚合，用于聚合嵌套对象。
//
// 示例：
//   esb.NestedAgg("nested_products", "products")
func NestedAgg(name, path string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			Nested: &types.NestedAggregation{
				Path: &path,
			},
		}
	}
}

// GlobalAgg 创建一个全局聚合，忽略查询上下文，对所有文档进行聚合。
//
// 示例：
//   esb.GlobalAgg("all_docs")
func GlobalAgg(name string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			Global: &types.GlobalAggregation{},
		}
	}
}

// SubAgg 为聚合添加子聚合。
//
// 示例：
//   esb.TermsAgg("categories", "category").SubAgg(
//       esb.AvgAgg("avg_price", "price"),
//   )
func SubAgg(parentName string, subAggs ...AggregationOption) AggregationOption {
	return func(aggs *types.Aggregations) {
		// 获取或创建父聚合
		parentAgg, exists := aggs.Aggregations[parentName]
		if !exists {
			parentAgg = types.Aggregations{
				Aggregations: make(map[string]types.Aggregations),
			}
		}
		
		// 确保父聚合有 Aggregations map
		if parentAgg.Aggregations == nil {
			parentAgg.Aggregations = make(map[string]types.Aggregations)
		}
		
		// 应用子聚合
		for _, subAgg := range subAggs {
			subAgg(&parentAgg)
		}
		
		aggs.Aggregations[parentName] = parentAgg
	}
}

// 便捷函数：常用聚合组合

// TopTermsAgg 创建一个获取前N个词项的聚合，这是最常用的场景。
//
// 示例：
//   esb.TopTermsAgg("top_categories", "category", 10)
//   esb.TopTermsAgg("top_categories", "category", 10, esb.AvgAgg("avg_price", "price"))
func TopTermsAgg(name, field string, size int, subAggs ...AggregationOption) AggregationOption {
	return TermsAggWithOptions(name, field, func(opts *types.TermsAggregation) {
		opts.Size = &size
	}, subAggs...)
}

// DailyHistogramAgg 创建按天分组的日期直方图聚合，这是最常用的时间聚合。
//
// 示例：
//   esb.DailyHistogramAgg("daily_sales", "timestamp")
//   esb.DailyHistogramAgg("daily_sales", "timestamp", esb.SumAgg("total", "amount"))
func DailyHistogramAgg(name, field string, subAggs ...AggregationOption) AggregationOption {
	return DateHistogramAgg(name, field, "1d", subAggs...)
}

// MonthlyHistogramAgg 创建按月分组的日期直方图聚合。
//
// 示例：
//   esb.MonthlyHistogramAgg("monthly_sales", "timestamp")
func MonthlyHistogramAgg(name, field string, subAggs ...AggregationOption) AggregationOption {
	return DateHistogramAgg(name, field, "1M", subAggs...)
}

// PriceRangeAgg 创建常用的价格范围聚合。
//
// 示例：
//   esb.PriceRangeAgg("price_segments", "price", []float64{0, 100, 500, 1000})
func PriceRangeAgg(name, field string, boundaries []float64) AggregationOption {
	ranges := make([]types.AggregationRange, 0, len(boundaries))
	
	for i := 0; i < len(boundaries); i++ {
		var aggRange types.AggregationRange
		
		if i == 0 {
			// 第一个范围：< boundaries[0]
			to := types.Float64(boundaries[i])
			aggRange.To = &to
		} else if i == len(boundaries)-1 {
			// 最后一个范围：>= boundaries[i]
			from := types.Float64(boundaries[i])
			aggRange.From = &from
		} else {
			// 中间范围：boundaries[i-1] <= x < boundaries[i]
			from := types.Float64(boundaries[i-1])
			to := types.Float64(boundaries[i])
			aggRange.From = &from
			aggRange.To = &to
		}
		
		ranges = append(ranges, aggRange)
	}
	
	return RangeAgg(name, field, ranges)
}

// 补充重要的聚合类型

// PercentilesAgg 创建百分位数聚合，计算字段值的百分位数。
//
// 示例：
//   esb.PercentilesAgg("response_time_percentiles", "response_time")
//   esb.PercentilesAgg("price_percentiles", "price", []float64{25, 50, 75, 95, 99})
func PercentilesAgg(name, field string, percentiles ...[]float64) AggregationOption {
	return func(aggs *types.Aggregations) {
		percentilesAgg := &types.PercentilesAggregation{
			Field: &field,
		}
		
		// 如果指定了百分位数，使用指定的值
		if len(percentiles) > 0 && len(percentiles[0]) > 0 {
			percents := make([]types.Float64, len(percentiles[0]))
			for i, p := range percentiles[0] {
				percents[i] = types.Float64(p)
			}
			percentilesAgg.Percents = percents
		}
		
		aggs.Aggregations[name] = types.Aggregations{
			Percentiles: percentilesAgg,
		}
	}
}

// ExtendedStatsAgg 创建扩展统计聚合，提供更详细的统计信息。
//
// 示例：
//   esb.ExtendedStatsAgg("price_extended_stats", "price")
func ExtendedStatsAgg(name, field string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			ExtendedStats: &types.ExtendedStatsAggregation{
				Field: &field,
			},
		}
	}
}

// TopHitsAgg 创建 top hits 聚合，返回每个桶中的顶部文档。
//
// 示例：
//   esb.TopHitsAgg("top_products", 3)
func TopHitsAgg(name string, size int) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			TopHits: &types.TopHitsAggregation{
				Size: &size,
			},
		}
	}
}

// SignificantTermsAgg 创建重要词项聚合，找出在子集中比在整个数据集中更频繁出现的词项。
//
// 示例：
//   esb.SignificantTermsAgg("significant_tags", "tags")
func SignificantTermsAgg(name, field string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			SignificantTerms: &types.SignificantTermsAggregation{
				Field: &field,
			},
		}
	}
}

// 地理位置相关聚合

// GeoBoundsAgg 创建地理边界聚合，计算所有地理点的边界框。
//
// 示例：
//   esb.GeoBoundsAgg("viewport", "location")
func GeoBoundsAgg(name, field string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			GeoBounds: &types.GeoBoundsAggregation{
				Field: &field,
			},
		}
	}
}

// GeoCentroidAgg 创建地理中心点聚合，计算所有地理点的中心点。
//
// 示例：
//   esb.GeoCentroidAgg("centroid", "location")
func GeoCentroidAgg(name, field string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			GeoCentroid: &types.GeoCentroidAggregation{
				Field: &field,
			},
		}
	}
}

// GeoDistanceAgg 创建地理距离聚合，根据距离某个点的距离对文档进行分组。
//
// 示例：
//   esb.GeoDistanceAgg("distance_ranges", "location", "40.7128,-74.0060", 
//       []string{"0-1km", "1-5km", "5-10km", "10km+"}, 
//       []float64{1000, 5000, 10000})
func GeoDistanceAgg(name, field, origin string, rangeKeys []string, distances []float64) AggregationOption {
	return func(aggs *types.Aggregations) {
		ranges := make([]types.AggregationRange, len(distances)+1)
		
		// 第一个范围：0 到第一个距离
		if len(distances) > 0 {
			to := types.Float64(distances[0])
			ranges[0] = types.AggregationRange{
				To: &to,
			}
			if len(rangeKeys) > 0 {
				ranges[0].Key = &rangeKeys[0]
			}
		}
		
		// 中间范围
		for i := 1; i < len(distances); i++ {
			from := types.Float64(distances[i-1])
			to := types.Float64(distances[i])
			ranges[i] = types.AggregationRange{
				From: &from,
				To:   &to,
			}
			if i < len(rangeKeys) {
				ranges[i].Key = &rangeKeys[i]
			}
		}
		
		// 最后一个范围：最后一个距离到无穷大
		if len(distances) > 0 {
			from := types.Float64(distances[len(distances)-1])
			ranges[len(distances)] = types.AggregationRange{
				From: &from,
			}
			if len(distances) < len(rangeKeys) {
				ranges[len(distances)].Key = &rangeKeys[len(distances)]
			}
		}
		
		aggs.Aggregations[name] = types.Aggregations{
			GeoDistance: &types.GeoDistanceAggregation{
				Field:  &field,
				Origin: origin,
				Ranges: ranges,
			},
		}
	}
}
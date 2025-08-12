package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/calendarinterval"
)

// AggregationOption 表示一个修改 types.Aggregations 的函数。
// 它遵循函数式选项模式来构建 Elasticsearch 聚合。
type AggregationOption func(*types.Aggregations)

// NewAggregations 通过应用提供的选项创建一个新的 Elasticsearch 聚合。
// 它返回一个可以直接用于 go-elasticsearch 客户端 Search.Aggregations() 方法的 map[string]types.Aggregations。
//
// 示例：
//   aggs := esb.NewAggregations(
//       esb.TermsAgg("categories", "category"),
//       esb.AvgAgg("avg_price", "price"),
//   )
//   search.Aggregations(aggs)
func NewAggregations(opts ...AggregationOption) map[string]types.Aggregations {
	aggs := &types.Aggregations{
		Aggregations: make(map[string]types.Aggregations),
	}
	for _, opt := range opts {
		opt(aggs)
	}
	return aggs.Aggregations
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

// 更多重要的聚合类型

// DateRangeAgg 创建日期范围聚合，根据日期范围对文档进行分组。
//
// 示例：
//   esb.DateRangeAgg("date_ranges", "timestamp", []types.DateRangeExpression{
//       {To: "2023-01-01"},
//       {From: "2023-01-01", To: "2023-12-31"},
//       {From: "2023-12-31"},
//   })
func DateRangeAgg(name, field string, ranges []types.DateRangeExpression) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			DateRange: &types.DateRangeAggregation{
				Field:  &field,
				Ranges: ranges,
			},
		}
	}
}

// IpRangeAgg 创建 IP 范围聚合，根据 IP 地址范围对文档进行分组。
//
// 示例：
//   to1 := "192.168.1.0/24"
//   from2 := "10.0.0.0/8"
//   esb.IpRangeAgg("ip_ranges", "client_ip", []types.IpRangeAggregationRange{
//       {To: &to1},
//       {From: &from2},
//   })
func IpRangeAgg(name, field string, ranges []types.IpRangeAggregationRange) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			IpRange: &types.IpRangeAggregation{
				Field:  &field,
				Ranges: ranges,
			},
		}
	}
}

// MissingAgg 创建缺失值聚合，统计指定字段缺失值的文档数量。
//
// 示例：
//   esb.MissingAgg("missing_emails", "email")
func MissingAgg(name, field string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			Missing: &types.MissingAggregation{
				Field: &field,
			},
		}
	}
}

// RareTermsAgg 创建稀有词项聚合，找出出现频率较低的词项。
//
// 示例：
//   esb.RareTermsAgg("rare_categories", "category")
func RareTermsAgg(name, field string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			RareTerms: &types.RareTermsAggregation{
				Field: &field,
			},
		}
	}
}

// SamplerAgg 创建采样聚合，对文档进行采样以提高聚合性能。
//
// 示例：
//   esb.SamplerAgg("sample", 1000)
func SamplerAgg(name string, shardSize int) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			Sampler: &types.SamplerAggregation{
				ShardSize: &shardSize,
			},
		}
	}
}

// DiversifiedSamplerAgg 创建多样化采样聚合，确保采样结果的多样性。
//
// 示例：
//   esb.DiversifiedSamplerAgg("diversified_sample", "category", 1000)
func DiversifiedSamplerAgg(name, field string, shardSize int) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			DiversifiedSampler: &types.DiversifiedSamplerAggregation{
				Field:     &field,
				ShardSize: &shardSize,
			},
		}
	}
}

// ReverseNestedAgg 创建反向嵌套聚合，从嵌套文档聚合回到父文档。
//
// 示例：
//   esb.ReverseNestedAgg("back_to_parent")
func ReverseNestedAgg(name string, path ...string) AggregationOption {
	return func(aggs *types.Aggregations) {
		reverseNested := &types.ReverseNestedAggregation{}
		if len(path) > 0 {
			reverseNested.Path = &path[0]
		}
		
		aggs.Aggregations[name] = types.Aggregations{
			ReverseNested: reverseNested,
		}
	}
}

// ChildrenAgg 创建子文档聚合，聚合指定类型的子文档。
//
// 示例：
//   esb.ChildrenAgg("child_products", "product")
func ChildrenAgg(name, childType string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			Children: &types.ChildrenAggregation{
				Type: &childType,
			},
		}
	}
}

// ParentAgg 创建父文档聚合，聚合指定类型的父文档。
//
// 示例：
//   esb.ParentAgg("parent_categories", "category")
func ParentAgg(name, parentType string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			Parent: &types.ParentAggregation{
				Type: &parentType,
			},
		}
	}
}

// AutoDateHistogramAgg 创建自动日期直方图聚合，自动选择合适的时间间隔。
//
// 示例：
//   esb.AutoDateHistogramAgg("auto_dates", "timestamp", 10)
func AutoDateHistogramAgg(name, field string, buckets int) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			AutoDateHistogram: &types.AutoDateHistogramAggregation{
				Field:   &field,
				Buckets: &buckets,
			},
		}
	}
}

// VariableWidthHistogramAgg 创建可变宽度直方图聚合。
//
// 示例：
//   esb.VariableWidthHistogramAgg("variable_histogram", "price", 10)
func VariableWidthHistogramAgg(name, field string, buckets int) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			VariableWidthHistogram: &types.VariableWidthHistogramAggregation{
				Field:   &field,
				Buckets: &buckets,
			},
		}
	}
}

// CompositeAgg 创建复合聚合，支持多维度分组和分页。
//
// 示例：
//   sources := []map[string]types.CompositeAggregationSource{
//       {"category": {Terms: &types.CompositeTermsAggregation{Field: "category"}}},
//   }
//   esb.CompositeAgg("composite", sources)
func CompositeAgg(name string, sources []map[string]types.CompositeAggregationSource) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			Composite: &types.CompositeAggregation{
				Sources: sources,
			},
		}
	}
}

// MultiTermsAgg 创建多词项聚合，基于多个字段的组合进行分组。
//
// 示例：
//   esb.MultiTermsAgg("multi_terms", []types.MultiTermLookup{
//       {Field: "category"},
//       {Field: "brand"},
//   })
func MultiTermsAgg(name string, terms []types.MultiTermLookup) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			MultiTerms: &types.MultiTermsAggregation{
				Terms: terms,
			},
		}
	}
}

// SignificantTextAgg 创建重要文本聚合，分析文本字段中的重要词汇。
//
// 示例：
//   esb.SignificantTextAgg("significant_text", "description")
func SignificantTextAgg(name, field string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			SignificantText: &types.SignificantTextAggregation{
				Field: &field,
			},
		}
	}
}

// 管道聚合（Pipeline Aggregations）

// AvgBucketAgg 创建平均桶聚合，计算兄弟聚合中所有桶的平均值。
//
// 示例：
//   esb.AvgBucketAgg("avg_monthly_sales", "monthly_sales>total_sales")
func AvgBucketAgg(name, bucketsPath string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			AvgBucket: &types.AverageBucketAggregation{
				BucketsPath: types.BucketsPath(bucketsPath),
			},
		}
	}
}

// MaxBucketAgg 创建最大桶聚合，找出兄弟聚合中值最大的桶。
//
// 示例：
//   esb.MaxBucketAgg("max_monthly_sales", "monthly_sales>total_sales")
func MaxBucketAgg(name, bucketsPath string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			MaxBucket: &types.MaxBucketAggregation{
				BucketsPath: types.BucketsPath(bucketsPath),
			},
		}
	}
}

// MinBucketAgg 创建最小桶聚合，找出兄弟聚合中值最小的桶。
//
// 示例：
//   esb.MinBucketAgg("min_monthly_sales", "monthly_sales>total_sales")
func MinBucketAgg(name, bucketsPath string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			MinBucket: &types.MinBucketAggregation{
				BucketsPath: types.BucketsPath(bucketsPath),
			},
		}
	}
}

// SumBucketAgg 创建求和桶聚合，计算兄弟聚合中所有桶的总和。
//
// 示例：
//   esb.SumBucketAgg("total_monthly_sales", "monthly_sales>total_sales")
func SumBucketAgg(name, bucketsPath string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			SumBucket: &types.SumBucketAggregation{
				BucketsPath: types.BucketsPath(bucketsPath),
			},
		}
	}
}

// StatsBucketAgg 创建统计桶聚合，计算兄弟聚合中所有桶的统计信息。
//
// 示例：
//   esb.StatsBucketAgg("sales_stats", "monthly_sales>total_sales")
func StatsBucketAgg(name, bucketsPath string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			StatsBucket: &types.StatsBucketAggregation{
				BucketsPath: types.BucketsPath(bucketsPath),
			},
		}
	}
}

// ExtendedStatsBucketAgg 创建扩展统计桶聚合，计算兄弟聚合中所有桶的扩展统计信息。
//
// 示例：
//   esb.ExtendedStatsBucketAgg("sales_extended_stats", "monthly_sales>total_sales")
func ExtendedStatsBucketAgg(name, bucketsPath string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			ExtendedStatsBucket: &types.ExtendedStatsBucketAggregation{
				BucketsPath: types.BucketsPath(bucketsPath),
			},
		}
	}
}

// PercentilesBucketAgg 创建百分位桶聚合，计算兄弟聚合中所有桶的百分位数。
//
// 示例：
//   esb.PercentilesBucketAgg("sales_percentiles", "monthly_sales>total_sales")
func PercentilesBucketAgg(name, bucketsPath string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			PercentilesBucket: &types.PercentilesBucketAggregation{
				BucketsPath: types.BucketsPath(bucketsPath),
			},
		}
	}
}

// MovingAvgAgg 创建移动平均聚合，计算时间序列数据的移动平均值。
//
// 示例：
//   esb.MovingAvgAgg("moving_avg", "sales", 7) // 7天移动平均
func MovingAvgAgg(name, bucketsPath string, window int) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			MovingAvg: types.NewSimpleMovingAverageAggregation(),
		}
		// 注意：MovingAvg 的具体实现需要根据模型类型来设置
	}
}

// DerivativeAgg 创建导数聚合，计算时间序列数据的变化率。
//
// 示例：
//   esb.DerivativeAgg("sales_derivative", "sales")
func DerivativeAgg(name, bucketsPath string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			Derivative: &types.DerivativeAggregation{
				BucketsPath: types.BucketsPath(bucketsPath),
			},
		}
	}
}

// CumulativeSumAgg 创建累积求和聚合，计算时间序列数据的累积和。
//
// 示例：
//   esb.CumulativeSumAgg("cumulative_sales", "sales")
func CumulativeSumAgg(name, bucketsPath string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			CumulativeSum: &types.CumulativeSumAggregation{
				BucketsPath: types.BucketsPath(bucketsPath),
			},
		}
	}
}

// 地理网格聚合

// GeohashGridAgg 创建 Geohash 网格聚合，将地理点按 geohash 网格分组。
//
// 示例：
//   esb.GeohashGridAgg("location_grid", "location", 5)
func GeohashGridAgg(name, field string, precision int) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			GeohashGrid: &types.GeoHashGridAggregation{
				Field:     &field,
				Precision: &precision,
			},
		}
	}
}

// GeotileGridAgg 创建地理瓦片网格聚合，将地理点按地图瓦片分组。
//
// 示例：
//   esb.GeotileGridAgg("tile_grid", "location", 8)
func GeotileGridAgg(name, field string, precision int) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			GeotileGrid: &types.GeoTileGridAggregation{
				Field:     &field,
				Precision: &precision,
			},
		}
	}
}

// 高级度量聚合

// WeightedAvgAgg 创建加权平均聚合，计算带权重的平均值。
//
// 示例：
//   esb.WeightedAvgAgg("weighted_avg_price", "price", "quantity")
func WeightedAvgAgg(name, valueField, weightField string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			WeightedAvg: &types.WeightedAverageAggregation{
				Value:  &types.WeightedAverageValue{Field: &valueField},
				Weight: &types.WeightedAverageValue{Field: &weightField},
			},
		}
	}
}

// MedianAbsoluteDeviationAgg 创建中位数绝对偏差聚合。
//
// 示例：
//   esb.MedianAbsoluteDeviationAgg("price_mad", "price")
func MedianAbsoluteDeviationAgg(name, field string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			MedianAbsoluteDeviation: &types.MedianAbsoluteDeviationAggregation{
				Field: &field,
			},
		}
	}
}

// StringStatsAgg 创建字符串统计聚合，分析字符串字段的统计信息。
//
// 示例：
//   esb.StringStatsAgg("description_stats", "description")
func StringStatsAgg(name, field string) AggregationOption {
	return func(aggs *types.Aggregations) {
		aggs.Aggregations[name] = types.Aggregations{
			StringStats: &types.StringStatsAggregation{
				Field: &field,
			},
		}
	}
}

// TTestAgg 创建 T 检验聚合，进行统计假设检验。
//
// 示例：
//   esb.TTestAgg("t_test", "score", esb.Term("group", "A"), esb.Term("group", "B"))
func TTestAgg(name, field string, filterA, filterB QueryOption) AggregationOption {
	return func(aggs *types.Aggregations) {
		queryA := &types.Query{}
		queryB := &types.Query{}
		
		if filterA != nil {
			filterA(queryA)
		}
		if filterB != nil {
			filterB(queryB)
		}
		
		aggs.Aggregations[name] = types.Aggregations{
			TTest: &types.TTestAggregation{
				A: &types.TestPopulation{
					Field:  field,
					Filter: queryA,
				},
				B: &types.TestPopulation{
					Field:  field,
					Filter: queryB,
				},
			},
		}
	}
}
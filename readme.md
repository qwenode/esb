# ESB (Elasticsearch Builder) 使用文档

ESB 是一个基于函数式选项模式的 Elasticsearch 流式查询构建器，简化了复杂 Elasticsearch 查询的构建过程，同时保持与 `github.com/elastic/go-elasticsearch/v8` 的完全兼容性。

## 目录

- [快速开始](#快速开始)
- [基础查询](#基础查询)
- [全文搜索查询](#全文搜索查询)
- [词项级查询](#词项级查询)
- [复合查询](#复合查询)
- [地理位置查询](#地理位置查询)
- [特殊查询](#特殊查询)
- [聚合查询](#聚合查询)
- [最佳实践](#最佳实践)

## 快速开始

```go
import (
    "github.com/qwenode/esb"
    "github.com/elastic/go-elasticsearch/v8"
)

// 创建 Elasticsearch 客户端
client, _ := elasticsearch.NewDefaultClient()

// 构建查询
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Match("title", "elasticsearch"),
            esb.NumberRange("price").Gte(10.0).Lte(100.0).Build(),
        ),
        esb.Filter(
            esb.Term("status", "published"),
        ),
    ),
)

// 执行搜索
res, err := client.Search(
    client.Search.WithIndex("products"),
    client.Search.WithBody(esquery.Wrap(query)),
)
```

## 基础查询

### Match All 查询

匹配所有文档，通常用作默认查询或与过滤器结合使用。

```go
// 示例1: 获取所有文档
query := esb.NewQuery(esb.MatchAll())

// 示例2: 获取所有文档并设置权重
query := esb.NewQuery(
    esb.MatchAllWithOptions(func(opts *types.MatchAllQuery) {
        boost := float32(1.2)
        opts.Boost = &boost
    }),
)

// 示例3: 与过滤器结合使用
query := esb.NewQuery(
    esb.Bool(
        esb.Must(esb.MatchAll()),
        esb.Filter(
            esb.Term("category", "electronics"),
            esb.NumberRange("price").Lte(1000.0).Build(),
        ),
    ),
)
```

### Match None 查询

不匹配任何文档，通常用于测试或特殊场景。

```go
// 示例1: 不匹配任何文档
query := esb.NewQuery(esb.MatchNone())

// 示例2: 条件性排除所有结果
query := esb.NewQuery(
    esb.Bool(
        esb.Should(
            esb.Match("title", "search term"),
            esb.MatchNone(), // 在某些条件下使用
        ),
    ),
)
```

## 全文搜索查询

### Match 查询

最常用的全文搜索查询，会分析查询文本并查找匹配的文档。

```go
// 示例1: 基础文本搜索
query := esb.NewQuery(esb.Match("title", "elasticsearch guide"))

// 示例2: 带模糊匹配的搜索
query := esb.NewQuery(
    esb.MatchWithOptions("title", "elasticsearch guide", func(q *types.MatchQuery) {
        q.Fuzziness = "AUTO"
        analyzer := "standard"
        q.Analyzer = &analyzer
        boost := float32(2.0)
        q.Boost = &boost
    }),
)

// 示例3: 产品搜索场景
query := esb.NewQuery(
    esb.Bool(
        esb.Should(
            esb.MatchWithOptions("name", "iPhone 15", func(q *types.MatchQuery) {
                boost := float32(3.0)
                q.Boost = &boost
            }),
            esb.Match("description", "iPhone 15"),
            esb.Match("brand", "Apple"),
        ),
        esb.Filter(
            esb.Term("status", "available"),
        ),
    ),
)
```

### Match Phrase 查询

用于精确短语匹配，保持词条的顺序和位置。

```go
// 示例1: 精确短语搜索
query := esb.NewQuery(esb.MatchPhrase("content", "machine learning algorithms"))

// 示例2: 带容错的短语搜索
query := esb.NewQuery(
    esb.MatchPhraseWithOptions("content", "artificial intelligence", func(q *types.MatchPhraseQuery) {
        slop := 2 // 允许词条间有2个位置的间隔
        q.Slop = &slop
        analyzer := "english"
        q.Analyzer = &analyzer
    }),
)

// 示例3: 新闻标题精确匹配
query := esb.NewQuery(
    esb.Bool(
        esb.Should(
            esb.MatchPhrase("headline", "COVID-19 vaccine update"),
            esb.MatchPhrase("summary", "COVID-19 vaccine update"),
        ),
        esb.Filter(
            esb.DateRange("publish_date").Gte("2023-01-01").Build(),
        ),
    ),
)
```

### Match Phrase Prefix 查询

用于自动补全和"输入即搜索"功能。

```go
// 示例1: 搜索建议/自动补全
query := esb.NewQuery(esb.MatchPhrasePrefix("title", "elasticsearch sea"))

// 示例2: 高级自动补全配置
query := esb.NewQuery(
    esb.MatchPhrasePrefixWithOptions("product_name", "macbook pro", func(q *types.MatchPhrasePrefixQuery) {
        maxExpansions := 50
        q.MaxExpansions = &maxExpansions
        slop := 1
        q.Slop = &slop
    }),
)

// 示例3: 用户搜索实时建议
userInput := "wireless headph"
query := esb.NewQuery(
    esb.Bool(
        esb.Should(
            esb.MatchPhrasePrefixWithOptions("title", userInput, func(q *types.MatchPhrasePrefixQuery) {
                boost := float32(3.0)
                q.Boost = &boost
            }),
            esb.MatchPhrasePrefix("description", userInput),
            esb.MatchPhrasePrefix("tags", userInput),
        ),
    ),
)
```

### Multi Match 查询

在多个字段中搜索相同的文本，是最实用的搜索查询之一。

```go
// 示例1: 基础多字段搜索
query := esb.NewQuery(esb.MultiMatch("elasticsearch", "title", "content", "tags"))

// 示例2: 带权重的多字段搜索
query := esb.NewQuery(esb.MultiMatch("java programming", "title^3", "content^2", "tags"))

// 示例3: 用户姓名搜索 (cross_fields)
query := esb.NewQuery(esb.MultiMatchCrossFields("john doe", "first_name", "last_name"))

// 示例4: 电商产品搜索 (best_fields)
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.MultiMatchBestFields("wireless headphones", "name^3", "description^2", "brand", "category"),
        ),
        esb.Filter(
            esb.Term("status", "active"),
            esb.NumberRange("price").Gte(0.0).Build(),
        ),
    ),
)
```

### Query String 查询

支持完整的 Lucene 查询语法，适合高级用户。

```go
// 示例1: 基础查询字符串
query := esb.NewQuery(esb.QueryString("title:elasticsearch AND status:published"))

// 示例2: 复杂的布尔查询
query := esb.NewQuery(esb.QueryString("(title:java OR title:python) AND category:programming NOT deprecated:true"))

// 示例3: 带配置的查询字符串
query := esb.NewQuery(
    esb.QueryStringWithOptions("elasticsearch OR \"elastic search\"", func(opts *types.QueryStringQuery) {
        defaultField := "content"
        opts.DefaultField = &defaultField
        opts.Fields = []string{"title^2", "content", "tags"}
        allowLeadingWildcard := true
        opts.AllowLeadingWildcard = &allowLeadingWildcard
    }),
)
```

## 词项级查询

### Term 查询

精确词项匹配，不会对查询词项进行分析。

```go
// 示例1: 状态过滤
query := esb.NewQuery(esb.Term("status", "published"))

// 示例2: 用户ID精确匹配
query := esb.NewQuery(esb.Term("user_id", "12345"))

// 示例3: 多条件精确匹配
query := esb.NewQuery(
    esb.Bool(
        esb.Filter(
            esb.Term("category", "electronics"),
            esb.Term("brand", "apple"),
            esb.Term("in_stock", "true"),
        ),
    ),
)
```

### Terms 查询

匹配多个精确词项中的任意一个。

```go
// 示例1: 多状态过滤
query := esb.NewQuery(esb.Terms("status", "published", "featured", "promoted"))

// 示例2: 多分类过滤
query := esb.NewQuery(esb.Terms("category", "electronics", "computers", "mobile"))

// 示例3: 用户权限检查
query := esb.NewQuery(
    esb.Bool(
        esb.Must(esb.Match("content", "search term")),
        esb.Filter(
            esb.Terms("visibility", "public", "internal"),
            esb.Terms("department", "engineering", "product", "design"),
        ),
    ),
)
```

### Range 查询

范围查询，支持数值、日期和字符串范围。

```go
// 示例1: 价格范围过滤
query := esb.NewQuery(
    esb.Bool(
        esb.Must(esb.Match("name", "laptop")),
        esb.Filter(
            esb.NumberRange("price").Gte(500.0).Lte(2000.0).Build(),
        ),
    ),
)

// 示例2: 日期范围查询
query := esb.NewQuery(
    esb.Bool(
        esb.Filter(
            esb.DateRange("created_at").Gte("2023-01-01").Lte("2023-12-31").Build(),
            esb.DateRange("updated_at").Gte("now-7d").Build(),
        ),
    ),
)

// 示例3: 年龄和评分组合查询
query := esb.NewQuery(
    esb.Bool(
        esb.Must(esb.Match("bio", "software engineer")),
        esb.Filter(
            esb.NumberRange("age").Gte(25.0).Lt(45.0).Build(),
            esb.NumberRange("rating").Gt(4.0).Build(),
        ),
    ),
)
```

### Exists 查询

检查字段是否存在且有值。

```go
// 示例1: 检查必填字段
query := esb.NewQuery(esb.Exists("email"))

// 示例2: 过滤有图片的产品
query := esb.NewQuery(
    esb.Bool(
        esb.Must(esb.Match("category", "electronics")),
        esb.Filter(
            esb.Exists("image_url"),
            esb.Exists("price"),
        ),
    ),
)

// 示例3: 查找完整资料的用户
query := esb.NewQuery(
    esb.Bool(
        esb.Filter(
            esb.Exists("profile.avatar"),
            esb.Exists("profile.bio"),
            esb.Exists("profile.location"),
        ),
    ),
)
```

### Prefix 查询

前缀匹配查询。

```go
// 示例1: 用户名前缀搜索
query := esb.NewQuery(esb.Prefix("username", "john"))

// 示例2: 产品编码搜索
query := esb.NewQuery(esb.Prefix("product_code", "ELEC"))

// 示例3: 多前缀组合搜索
query := esb.NewQuery(
    esb.Bool(
        esb.Should(
            esb.Prefix("title", "elasticsearch"),
            esb.Prefix("tags", "elastic"),
        ),
        esb.Filter(
            esb.Term("status", "active"),
        ),
    ),
)
```

### Wildcard 查询

通配符查询，支持 `*` 和 `?` 通配符。

```go
// 示例1: 邮箱域名搜索
query := esb.NewQuery(esb.Wildcard("email", "*@gmail.com"))

// 示例2: 产品型号模糊搜索
query := esb.NewQuery(esb.Wildcard("model", "iPhone*"))

// 示例3: 文件名搜索
query := esb.NewQuery(
    esb.Bool(
        esb.Should(
            esb.Wildcard("filename", "*.pdf"),
            esb.Wildcard("filename", "*.doc*"),
        ),
        esb.Filter(
            esb.NumberRange("file_size").Lte(10485760.0).Build(), // 10MB
        ),
    ),
)
```

### Fuzzy 查询

模糊查询，基于编辑距离查找相似词条。

```go
// 示例1: 用户名模糊搜索
query := esb.NewQuery(esb.Fuzzy("username", "john"))

// 示例2: 产品名称容错搜索
query := esb.NewQuery(
    esb.FuzzyWithOptions("product_name", "iphone", func(opts *types.FuzzyQuery) {
        fuzziness := "2"
        opts.Fuzziness = &fuzziness
        prefixLength := 1
        opts.PrefixLength = &prefixLength
    }),
)

// 示例3: 多字段模糊搜索
query := esb.NewQuery(
    esb.Bool(
        esb.Should(
            esb.Fuzzy("title", "elasticsearch"),
            esb.Fuzzy("description", "elasticsearch"),
        ),
        esb.Filter(
            esb.Term("language", "en"),
        ),
    ),
)
```

## 复合查询

### Bool 查询

最重要的复合查询，使用布尔逻辑组合多个查询。

```go
// 示例1: 电商产品搜索
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.MultiMatch("wireless headphones", "name^3", "description^2"),
        ),
        esb.Filter(
            esb.Term("category", "electronics"),
            esb.NumberRange("price").Gte(50.0).Lte(300.0).Build(),
            esb.Term("in_stock", "true"),
        ),
        esb.Should(
            esb.Term("brand", "sony"),
            esb.Term("brand", "bose"),
            esb.Term("brand", "apple"),
        ),
        esb.MustNot(
            esb.Term("status", "discontinued"),
            esb.Term("restricted", "true"),
        ),
    ),
)

// 示例2: 内容管理系统
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Match("content", "artificial intelligence"),
        ),
        esb.Filter(
            esb.Terms("status", "published", "featured"),
            esb.DateRange("publish_date").Gte("2023-01-01").Build(),
            esb.Exists("author"),
        ),
        esb.Should(
            esb.Term("category", "technology"),
            esb.Term("category", "science"),
            esb.Match("tags", "AI machine learning"),
        ),
        esb.MustNot(
            esb.Term("draft", "true"),
            esb.Term("private", "true"),
        ),
    ),
)
```

## 聚合查询

### 基础聚合

```go
// 示例1: 词项聚合 - 统计分类
aggs := esb.NewAggregations(
    esb.TermsAgg("categories", "category"),
    esb.AvgAgg("avg_price", "price"),
    esb.MaxAgg("max_price", "price"),
    esb.MinAgg("min_price", "price"),
)

// 示例2: 嵌套聚合 - 每个分类的平均价格
aggs := esb.NewAggregations(
    esb.TermsAgg("categories", "category",
        esb.AvgAgg("avg_price", "price"),
        esb.ValueCountAgg("product_count", "id"),
    ),
)

// 示例3: 日期直方图 - 销售趋势
aggs := esb.NewAggregations(
    esb.DateHistogramAgg("sales_over_time", "order_date", "1d",
        esb.SumAgg("daily_revenue", "amount"),
        esb.ValueCountAgg("daily_orders", "id"),
    ),
)
```

### 高级聚合

```go
// 示例1: 价格范围分析
aggs := esb.NewAggregations(
    esb.PriceRangeAgg("price_segments", "price", []float64{0, 100, 500, 1000}),
    esb.PercentilesAgg("price_percentiles", "price", []float64{25, 50, 75, 95, 99}),
)

// 示例2: 地理位置聚合
aggs := esb.NewAggregations(
    esb.GeoDistanceAgg("distance_ranges", "location", "40.7128,-74.0060",
        []string{"0-1km", "1-5km", "5-10km", "10km+"},
        []float64{1000, 5000, 10000}),
    esb.GeoBoundsAgg("viewport", "location"),
)
```

## 最佳实践

### 1. 性能优化

```go
// 使用 Filter 而不是 Must 来提高性能（不需要评分时）
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Match("title", "search term"), // 需要评分
        ),
        esb.Filter( // 不需要评分，会被缓存
            esb.Term("status", "published"),
            esb.NumberRange("price").Gte(10.0).Build(),
        ),
    ),
)
```

### 2. 搜索相关性优化

```go
// 使用字段权重提升重要字段
query := esb.NewQuery(
    esb.MultiMatch("elasticsearch guide", "title^3", "summary^2", "content"),
)

// 使用 Should 子句增加相关性
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.MultiMatch("laptop", "name", "description"),
        ),
        esb.Should( // 这些条件会增加分数但不是必需的
            esb.Term("brand", "apple"),
            esb.Term("featured", "true"),
            esb.NumberRange("rating").Gte(4.5).Build(),
        ),
    ),
)
```

### 3. 复杂业务场景

```go
// 电商搜索完整示例
type ProductFilters struct {
    Categories      []string
    MinPrice        float64
    MaxPrice        float64
    PreferredBrands []string
}

func BuildProductSearchQuery(searchTerm string, filters ProductFilters) *types.Query {
    boolOptions := []esb.BoolOption{}
    
    // 主搜索条件
    if searchTerm != "" {
        boolOptions = append(boolOptions, esb.Must(
            esb.MultiMatchBestFields(searchTerm, "name^3", "description^2", "brand", "category"),
        ))
    }
    
    // 过滤条件
    filterConditions := []esb.QueryOption{}
    
    if len(filters.Categories) > 0 {
        filterConditions = append(filterConditions, esb.TermsSlice("category", filters.Categories))
    }
    
    if filters.MinPrice > 0 || filters.MaxPrice > 0 {
        rangeQuery := esb.NumberRange("price")
        if filters.MinPrice > 0 {
            rangeQuery = rangeQuery.Gte(filters.MinPrice)
        }
        if filters.MaxPrice > 0 {
            rangeQuery = rangeQuery.Lte(filters.MaxPrice)
        }
        filterConditions = append(filterConditions, rangeQuery.Build())
    }
    
    if len(filterConditions) > 0 {
        boolOptions = append(boolOptions, esb.Filter(filterConditions...))
    }
    
    // 提升条件
    if len(filters.PreferredBrands) > 0 {
        boolOptions = append(boolOptions, esb.Should(
            esb.TermsSlice("brand", filters.PreferredBrands),
        ))
    }
    
    // 排除条件
    boolOptions = append(boolOptions, esb.MustNot(
        esb.Term("status", "discontinued"),
        esb.Term("hidden", "true"),
    ))
    
    return esb.NewQuery(esb.Bool(boolOptions...))
}
```

这个文档涵盖了 ESB 库中所有主要的查询类型和聚合功能，每个部分都包含了实际的业务场景示例，展示了如何在真实项目中使用这些查询。
### S
imple Query String 查询

简化版的查询字符串，语法更简单，容错性更好。

```go
// 示例1: 简单查询字符串
query := esb.NewQuery(esb.SimpleQueryString("elasticsearch +guide -deprecated"))

// 示例2: 多字段简单查询
query := esb.NewQuery(
    esb.SimpleQueryStringWithOptions("java programming", func(opts *types.SimpleQueryStringQuery) {
        opts.Fields = []string{"title^2", "content", "tags"}
        defaultOp := operator.And
        opts.DefaultOperator = &defaultOp
    }),
)

// 示例3: 用户友好的搜索界面
query := esb.NewQuery(
    esb.SimpleQueryStringWithOptions("machine learning python", func(opts *types.SimpleQueryStringQuery) {
        opts.Fields = []string{"title^3", "description^2", "tags"}
        lenient := true
        opts.Lenient = &lenient // 忽略格式错误
        analyzeWildcard := true
        opts.AnalyzeWildcard = &analyzeWildcard
    }),
)
```

### IDs 查询

根据文档ID查询。

```go
// 示例1: 单个文档查询
query := esb.NewQuery(esb.IDs("doc123"))

// 示例2: 多个文档查询
query := esb.NewQuery(esb.IDs("doc1", "doc2", "doc3"))

// 示例3: 批量文档查询
docIDs := []string{"user_123", "user_456", "user_789"}
query := esb.NewQuery(esb.IDsSlice(docIDs))

// 示例4: 结合其他条件的ID查询
query := esb.NewQuery(
    esb.Bool(
        esb.Should(
            esb.IDs("important_doc_1", "important_doc_2"),
            esb.Bool(
                esb.Must(
                    esb.Match("content", "urgent"),
                    esb.Term("priority", "high"),
                ),
            ),
        ),
    ),
)
```

### Regexp 查询

正则表达式查询。

```go
// 示例1: 手机号格式验证
query := esb.NewQuery(esb.Regexp("phone", "[0-9]{3}-[0-9]{3}-[0-9]{4}"))

// 示例2: 邮箱格式搜索
query := esb.NewQuery(esb.Regexp("email", "[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}"))

// 示例3: 产品编码规则匹配
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Regexp("product_id", "PROD-[0-9]{4}-[A-Z]{2}"),
        ),
        esb.Filter(
            esb.Term("status", "active"),
        ),
    ),
)

// 示例4: 复杂的文本模式匹配
query := esb.NewQuery(
    esb.Bool(
        esb.Should(
            esb.Regexp("username", "[a-zA-Z][a-zA-Z0-9_]{2,15}"), // 用户名格式
            esb.Regexp("display_name", "[A-Z][a-z]+ [A-Z][a-z]+"), // 真实姓名格式
        ),
        esb.Filter(
            esb.Term("verified", "true"),
        ),
    ),
)
```

### Terms Set 查询

匹配指定数量的词项。

```go
// 示例1: 技能匹配 - 至少匹配3个技能
query := esb.NewQuery(
    esb.TermsSetWithOptions("skills", []string{"java", "python", "javascript", "react", "spring"}, func(opts *types.TermsSetQuery) {
        minimumShouldMatchField := "required_skills_count"
        opts.MinimumShouldMatchField = &minimumShouldMatchField
    }),
)

// 示例2: 标签匹配 - 动态最小匹配数
query := esb.NewQuery(
    esb.TermsSetWithOptions("tags", []string{"urgent", "important", "customer", "bug"}, func(opts *types.TermsSetQuery) {
        script := types.Script{
            Source: "Math.min(params.num_terms, doc['priority_level'].value)",
        }
        opts.MinimumShouldMatchScript = &script
    }),
)

// 示例3: 产品特性匹配
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Match("category", "smartphone"),
            esb.TermsSetWithOptions("features", 
                []string{"5G", "wireless_charging", "waterproof", "dual_camera"}, 
                func(opts *types.TermsSetQuery) {
                    minimumShouldMatch := "2" // 至少匹配2个特性
                    opts.MinimumShouldMatch = &minimumShouldMatch
                },
            ),
        ),
        esb.Filter(
            esb.NumberRange("price").Lte(1000.0).Build(),
        ),
    ),
)
```

## 地理位置查询

### Geo Distance 查询

基于地理距离的查询。

```go
// 示例1: 查找附近的餐厅
query := esb.NewQuery(
    esb.Bool(
        esb.Must(esb.Match("type", "restaurant")),
        esb.Filter(
            esb.GeoDistance("location", 40.7128, -74.0060, "5km"),
        ),
    ),
)

// 示例2: 多距离范围查询
query := esb.NewQuery(
    esb.Bool(
        esb.Should(
            esb.GeoDistanceWithOptions("location", 40.7128, -74.0060, "1km", func(opts *types.GeoDistanceQuery) {
                boost := float32(2.0)
                opts.Boost = &boost // 1km内高权重
            }),
            esb.GeoDistanceWithOptions("location", 40.7128, -74.0060, "10km", func(opts *types.GeoDistanceQuery) {
                boost := float32(0.5)
                opts.Boost = &boost // 10km内低权重
            }),
        ),
        esb.Filter(
            esb.Term("status", "open"),
        ),
    ),
)

// 示例3: 配送服务范围查询
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Terms("service_type", "delivery", "pickup"),
        ),
        esb.Filter(
            esb.GeoDistance("service_area", 37.7749, -122.4194, "15km"),
            esb.Term("available", "true"),
        ),
        esb.Should(
            esb.GeoDistance("location", 37.7749, -122.4194, "2km"), // 优先显示近距离
        ),
    ),
)
```

### Geo Bounding Box 查询

在地理边界框内查找文档。

```go
// 示例1: 城市范围内搜索
query := esb.NewQuery(
    esb.Bool(
        esb.Must(esb.Match("category", "hotel")),
        esb.Filter(
            esb.GeoBoundingBox("location", 40.8, -74.1, 40.7, -73.9), // NYC 边界框
        ),
    ),
)

// 示例2: 配送范围查询
query := esb.NewQuery(
    esb.Bool(
        esb.Must(esb.Match("service", "delivery")),
        esb.Filter(
            esb.GeoBoundingBoxWithOptions("service_area", 37.8, -122.5, 37.7, -122.3, func(opts *types.GeoBoundingBoxQuery) {
                validationMethod := geovalidationmethod.Strict
                opts.ValidationMethod = &validationMethod
            }),
        ),
    ),
)

// 示例3: 房地产搜索
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Match("property_type", "apartment"),
            esb.NumberRange("bedrooms").Gte(2.0).Build(),
        ),
        esb.Filter(
            esb.GeoBoundingBox("coordinates", 40.75, -74.0, 40.70, -73.95), // 曼哈顿中城
            esb.NumberRange("price").Lte(5000.0).Build(),
        ),
        esb.Should(
            esb.Term("amenities", "gym"),
            esb.Term("amenities", "parking"),
            esb.Term("pet_friendly", "true"),
        ),
    ),
)
```

### Geo Polygon 查询

在地理多边形内查找文档。

```go
// 示例1: 自定义区域搜索
points := [][]float64{
    {40.8, -74.1},
    {40.8, -73.9},
    {40.7, -73.9},
    {40.7, -74.1},
}
query := esb.NewQuery(
    esb.Bool(
        esb.Must(esb.Match("type", "store")),
        esb.Filter(
            esb.GeoPolygon("location", points),
        ),
    ),
)

// 示例2: 学区范围查询
schoolDistrictPoints := [][]float64{
    {40.85, -73.88},
    {40.82, -73.85},
    {40.78, -73.87},
    {40.81, -73.91},
}
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Match("property_type", "house"),
            esb.NumberRange("bedrooms").Gte(3.0).Build(),
        ),
        esb.Filter(
            esb.GeoPolygon("address.coordinates", schoolDistrictPoints),
            esb.Term("school_district", "excellent"),
        ),
    ),
)

// 示例3: 商业区域分析
businessZonePoints := [][]float64{
    {40.76, -73.99},
    {40.76, -73.97},
    {40.74, -73.97},
    {40.74, -73.99},
}
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Terms("business_type", "restaurant", "retail", "office"),
        ),
        esb.Filter(
            esb.GeoPolygon("business_location", businessZonePoints),
            esb.Term("status", "active"),
        ),
        esb.Should(
            esb.NumberRange("rating").Gte(4.0).Build(),
            esb.Term("featured", "true"),
        ),
    ),
)
```

### Geo Shape 查询

查询地理形状。

```go
// 示例1: 点与形状相交查询
query := esb.NewQuery(
    esb.Bool(
        esb.Must(esb.Match("type", "landmark")),
        esb.Filter(
            esb.GeoShape("area", "POINT(13.0 53.0)"),
        ),
    ),
)

// 示例2: 复杂地理形状查询
polygonWKT := "POLYGON((0 0, 10 0, 10 10, 0 10, 0 0))"
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Terms("category", "park", "recreation", "nature"),
        ),
        esb.Filter(
            esb.GeoShape("boundaries", polygonWKT),
            esb.Term("public_access", "true"),
        ),
    ),
)
```

## 特殊查询

### Nested 查询

查询嵌套对象。

```go
// 示例1: 查询产品评论
query := esb.NewQuery(
    esb.Nested("reviews",
        esb.Bool(
            esb.Must(
                esb.NumberRange("reviews.rating").Gte(4.0).Build(),
                esb.Match("reviews.comment", "excellent quality"),
            ),
        ),
    ),
)

// 示例2: 复杂嵌套查询 - 产品规格
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Match("name", "laptop"),
            esb.Nested("specifications",
                esb.Bool(
                    esb.Must(
                        esb.Term("specifications.name", "RAM"),
                        esb.NumberRange("specifications.value").Gte(16.0).Build(),
                    ),
                ),
            ),
        ),
        esb.Filter(
            esb.Term("category", "computers"),
            esb.NumberRange("price").Lte(2000.0).Build(),
        ),
    ),
)

// 示例3: 多层嵌套查询 - 订单详情
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.DateRange("order_date").Gte("2023-01-01").Build(),
            esb.Nested("items",
                esb.Bool(
                    esb.Must(
                        esb.Term("items.category", "electronics"),
                        esb.NumberRange("items.quantity").Gte(2.0).Build(),
                    ),
                    esb.Nested("items.reviews",
                        esb.NumberRange("items.reviews.rating").Gte(4.0).Build(),
                    ),
                ),
            ),
        ),
        esb.Filter(
            esb.Term("status", "completed"),
        ),
    ),
)

// 示例4: 带评分模式的嵌套查询
query := esb.NewQuery(
    esb.NestedWithOptions("comments",
        esb.Bool(
            esb.Must(
                esb.Match("comments.text", "helpful review"),
                esb.NumberRange("comments.upvotes").Gte(5.0).Build(),
            ),
        ),
        func(opts *types.NestedQuery) {
            scoreMode := types.NestedScoremodeMax
            opts.ScoreMode = &scoreMode
            ignoreUnmapped := true
            opts.IgnoreUnmapped = &ignoreUnmapped
        },
    ),
)
```

### Script 查询

使用脚本进行复杂查询。

```go
// 示例1: 自定义评分脚本
query := esb.NewQuery(
    esb.Script("doc['likes'].value > doc['dislikes'].value * 2"),
)

// 示例2: 带参数的脚本查询
params := map[string]any{
    "discount": 0.8,
    "max_price": 100,
}
query := esb.NewQuery(
    esb.ScriptWithParams("doc['price'].value * params.discount < params.max_price", params),
)

// 示例3: 复杂业务逻辑脚本
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Match("category", "electronics"),
            esb.ScriptWithParams(
                "doc['stock'].value > 0 && (doc['price'].value * params.tax_rate + params.shipping_cost) <= params.budget",
                map[string]any{
                    "tax_rate": 0.08,
                    "shipping_cost": 15.0,
                    "budget": 500.0,
                },
            ),
        ),
    ),
)

// 示例4: 时间相关的脚本查询
query := esb.NewQuery(
    esb.ScriptWithOptions("doc['created_at'].value.millis > System.currentTimeMillis() - params.days * 24 * 60 * 60 * 1000", func(opts *types.ScriptQuery) {
        lang := scriptlanguage.Painless
        opts.Script.Lang = &lang
        params := map[string]json.RawMessage{
            "days": json.RawMessage("30"),
        }
        opts.Script.Params = params
        boost := float32(1.5)
        opts.Boost = &boost
    }),
)
```

### More Like This 查询

查找相似文档。

```go
// 示例1: 基于文档ID查找相似内容
query := esb.NewQuery(
    esb.MoreLikeThisWithOptions(func(opts *types.MoreLikeThisQuery) {
        opts.Fields = []string{"title", "content"}
        opts.Like = []types.Like{
            {Document: &types.LikeDocument{
                Index: "articles",
                Id:    "article_123",
            }},
        }
        minTermFreq := 1
        opts.MinTermFreq = &minTermFreq
        maxQueryTerms := 25
        opts.MaxQueryTerms = &maxQueryTerms
    }),
)

// 示例2: 基于文本查找相似内容
query := esb.NewQuery(
    esb.MoreLikeThisWithOptions(func(opts *types.MoreLikeThisQuery) {
        opts.Fields = []string{"title", "description"}
        opts.Like = []types.Like{
            {Text: "artificial intelligence machine learning"},
        }
        minDocFreq := 2
        opts.MinDocFreq = &minDocFreq
        maxDocFreq := 1000
        opts.MaxDocFreq = &maxDocFreq
    }),
)

// 示例3: 推荐系统 - 基于用户行为
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.MoreLikeThisWithMultipleLikes(
                []string{"title", "description", "tags"},
                []types.Like{
                    {Document: &types.LikeDocument{Index: "products", Id: "viewed_product_1"}},
                    {Document: &types.LikeDocument{Index: "products", Id: "purchased_product_1"}},
                    {Text: "user search history keywords"},
                },
            ),
        ),
        esb.Filter(
            esb.Term("status", "available"),
            esb.NumberRange("rating").Gte(3.5).Build(),
        ),
        esb.MustNot(
            esb.IDs("already_viewed_1", "already_purchased_1"), // 排除已查看/购买的商品
        ),
    ),
)

// 示例4: 内容推荐 - 排除不相关内容
query := esb.NewQuery(
    esb.MoreLikeThisWithUnlike(
        []string{"title", "content", "summary"},
        "machine learning artificial intelligence",
        "advertisement promotion spam",
    ),
)
```

### Boosting 查询

提升或降低匹配特定查询的文档分数。

```go
// 示例1: 提升新内容，降低旧内容
query := esb.NewQuery(
    esb.Boosting(
        esb.Match("content", "elasticsearch guide"), // positive query
        esb.DateRange("publish_date").Lt("2022-01-01").Build(), // negative query
        0.5, // negative boost
    ),
)

// 示例2: 产品搜索中降低缺货商品
query := esb.NewQuery(
    esb.BoostingWithOptions(
        esb.MultiMatch("laptop computer", "name", "description"), // positive
        esb.Term("stock_status", "out_of_stock"), // negative
        0.2, // negative boost
        func(opts *types.BoostingQuery) {
            boost := float32(1.5)
            opts.Boost = &boost
        },
    ),
)

// 示例3: 内容质量评分调整
query := esb.NewQuery(
    esb.Boosting(
        esb.Bool( // positive: 高质量内容
            esb.Must(
                esb.Match("content", "comprehensive tutorial"),
                esb.NumberRange("word_count").Gte(1000.0).Build(),
            ),
        ),
        esb.Bool( // negative: 低质量指标
            esb.Should(
                esb.Term("auto_generated", "true"),
                esb.NumberRange("spelling_errors").Gt(5.0).Build(),
                esb.Term("duplicate_content", "true"),
            ),
        ),
        0.3, // 大幅降低低质量内容的分数
    ),
)
```

### Dis Max 查询

返回匹配任一查询的文档，使用最高分数。

```go
// 示例1: 多字段搜索取最高分
query := esb.NewQuery(
    esb.DisMax(
        esb.Match("title", "elasticsearch guide"),
        esb.Match("content", "elasticsearch guide"),
        esb.Match("summary", "elasticsearch guide"),
    ),
)

// 示例2: 带 tie_breaker 的 DisMax
query := esb.NewQuery(
    esb.DisMaxWithOptions(
        func(opts *types.DisMaxQuery) {
            tieBreaker := 0.3
            opts.TieBreaker = &tieBreaker
            boost := float32(1.2)
            opts.Boost = &boost
        },
        esb.Match("title", "machine learning"),
        esb.Match("abstract", "machine learning"),
        esb.Match("keywords", "machine learning"),
    ),
)

// 示例3: 多种搜索策略组合
query := esb.NewQuery(
    esb.DisMaxWithOptions(
        func(opts *types.DisMaxQuery) {
            tieBreaker := 0.4
            opts.TieBreaker = &tieBreaker
        },
        esb.MultiMatchBestFields("python programming", "title^3", "content^2"),
        esb.MultiMatchPhrase("python programming", "title", "content"),
        esb.Bool(
            esb.Must(
                esb.Term("tags", "python"),
                esb.Term("tags", "programming"),
            ),
        ),
    ),
)
```

### Constant Score 查询

为所有匹配的文档返回恒定分数。

```go
// 示例1: 过滤查询不需要评分
query := esb.NewQuery(
    esb.ConstantScore(
        esb.Bool(
            esb.Filter(
                esb.Term("category", "electronics"),
                esb.NumberRange("price").Lte(1000.0).Build(),
            ),
        ),
        1.0,
    ),
)

// 示例2: 组合评分和过滤
query := esb.NewQuery(
    esb.Bool(
        esb.Should(
            esb.Match("title", "smartphone"), // 正常评分
            esb.ConstantScore( // 恒定分数
                esb.Term("featured", "true"),
                2.0,
            ),
            esb.ConstantScore( // 另一个恒定分数
                esb.Term("on_sale", "true"),
                1.5,
            ),
        ),
    ),
)

// 示例3: 分类权重控制
query := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Match("content", "search query"),
        ),
        esb.Should(
            esb.ConstantScore(esb.Term("category", "premium"), 3.0),
            esb.ConstantScore(esb.Term("category", "standard"), 2.0),
            esb.ConstantScore(esb.Term("category", "basic"), 1.0),
        ),
    ),
)
```

## 高级聚合示例

### 电商分析仪表板

```go
// 完整的电商销售分析聚合
func BuildEcommerceAnalytics() *types.Aggregations {
    return esb.NewAggregations(
        // 1. 销售趋势分析
        esb.DateHistogramAgg("sales_trend", "order_date", "1d",
            esb.SumAgg("daily_revenue", "total_amount"),
            esb.ValueCountAgg("daily_orders", "order_id"),
            esb.AvgAgg("avg_order_value", "total_amount"),
            esb.CardinalityAgg("unique_customers", "customer_id"),
        ),
        
        // 2. 产品分类性能
        esb.TopTermsAgg("category_performance", "category", 15,
            esb.SumAgg("category_revenue", "total_amount"),
            esb.ValueCountAgg("orders_count", "order_id"),
            esb.AvgAgg("avg_price", "unit_price"),
            esb.CardinalityAgg("unique_products", "product_id"),
            // 子聚合：每个分类的热门产品
            esb.TopTermsAgg("top_products", "product_name", 5,
                esb.SumAgg("product_revenue", "total_amount"),
            ),
        ),
        
        // 3. 价格分布分析
        esb.HistogramAgg("price_distribution", "unit_price", 50),
        esb.PercentilesAgg("price_percentiles", "unit_price", []float64{25, 50, 75, 90, 95, 99}),
        esb.ExtendedStatsAgg("price_stats", "unit_price"),
        
        // 4. 地理销售分析
        esb.TermsAgg("sales_by_region", "shipping_region",
            esb.SumAgg("region_revenue", "total_amount"),
            esb.CardinalityAgg("unique_customers", "customer_id"),
            esb.AvgAgg("avg_order_value", "total_amount"),
            // 每个地区的城市分布
            esb.TopTermsAgg("top_cities", "shipping_city", 10,
                esb.SumAgg("city_revenue", "total_amount"),
            ),
        ),
        
        // 5. 客户分析
        esb.TermsAgg("customer_segments", "customer_tier",
            esb.AvgAgg("avg_order_value", "total_amount"),
            esb.ValueCountAgg("order_count", "order_id"),
            esb.SumAgg("segment_revenue", "total_amount"),
        ),
        
        // 6. 时间段分析
        esb.TermsAgg("hourly_sales", "hour_of_day",
            esb.SumAgg("hourly_revenue", "total_amount"),
            esb.ValueCountAgg("hourly_orders", "order_id"),
        ),
        
        // 7. 支付方式分析
        esb.TermsAgg("payment_methods", "payment_method",
            esb.SumAgg("payment_revenue", "total_amount"),
            esb.AvgAgg("avg_transaction", "total_amount"),
        ),
        
        // 8. 折扣影响分析
        esb.RangeAgg("discount_impact", "discount_percentage", []types.AggregationRange{
            {To: types.Float64(5)},   // 0-5%
            {From: types.Float64(5), To: types.Float64(15)},  // 5-15%
            {From: types.Float64(15), To: types.Float64(30)}, // 15-30%
            {From: types.Float64(30)}, // 30%+
        }),
    )
}
```

### 内容分析仪表板

```go
// 内容管理系统分析聚合
func BuildContentAnalytics() *types.Aggregations {
    return esb.NewAggregations(
        // 1. 内容发布趋势
        esb.DateHistogramAgg("publishing_trend", "publish_date", "1w",
            esb.ValueCountAgg("weekly_posts", "id"),
            esb.CardinalityAgg("active_authors", "author_id"),
            esb.AvgAgg("avg_word_count", "word_count"),
        ),
        
        // 2. 内容分类分析
        esb.TopTermsAgg("content_categories", "category", 20,
            esb.ValueCountAgg("posts_count", "id"),
            esb.AvgAgg("avg_views", "view_count"),
            esb.AvgAgg("avg_engagement", "engagement_score"),
            // 每个分类的热门标签
            esb.TopTermsAgg("popular_tags", "tags", 10,
                esb.ValueCountAgg("tag_usage", "id"),
            ),
        ),
        
        // 3. 作者表现分析
        esb.TopTermsAgg("top_authors", "author_id", 50,
            esb.ValueCountAgg("posts_count", "id"),
            esb.SumAgg("total_views", "view_count"),
            esb.AvgAgg("avg_engagement", "engagement_score"),
            esb.MaxAgg("best_post_views", "view_count"),
        ),
        
        // 4. 内容长度分析
        esb.HistogramAgg("content_length_dist", "word_count", 500),
        esb.RangeAgg("content_length_ranges", "word_count", []types.AggregationRange{
            {To: types.Float64(500)},     // 短文章
            {From: types.Float64(500), To: types.Float64(1500)},  // 中等文章
            {From: types.Float64(1500), To: types.Float64(3000)}, // 长文章
            {From: types.Float64(3000)},  // 超长文章
        }),
        
        // 5. 参与度分析
        esb.PercentilesAgg("engagement_percentiles", "engagement_score", []float64{50, 75, 90, 95, 99}),
        esb.TermsAgg("engagement_levels", "engagement_level",
            esb.ValueCountAgg("posts_count", "id"),
            esb.AvgAgg("avg_views", "view_count"),
        ),
        
        // 6. 语言分布
        esb.TermsAgg("content_languages", "language",
            esb.ValueCountAgg("posts_count", "id"),
            esb.AvgAgg("avg_word_count", "word_count"),
        ),
        
        // 7. 发布时间分析
        esb.TermsAgg("publish_hour", "publish_hour",
            esb.ValueCountAgg("posts_count", "id"),
            esb.AvgAgg("avg_views", "view_count"),
        ),
        
        // 8. 内容状态分析
        esb.TermsAgg("content_status", "status",
            esb.ValueCountAgg("count", "id"),
            esb.DateHistogramAgg("status_trend", "updated_at", "1d",
                esb.ValueCountAgg("daily_count", "id"),
            ),
        ),
    )
}
```

### 用户行为分析

```go
// 用户行为分析聚合
func BuildUserBehaviorAnalytics() *types.Aggregations {
    return esb.NewAggregations(
        // 1. 用户活跃度趋势
        esb.DateHistogramAgg("user_activity_trend", "timestamp", "1h",
            esb.CardinalityAgg("unique_users", "user_id"),
            esb.ValueCountAgg("total_actions", "action_id"),
            esb.AvgAgg("avg_session_duration", "session_duration"),
        ),
        
        // 2. 用户行为类型分析
        esb.TermsAgg("action_types", "action_type",
            esb.ValueCountAgg("action_count", "action_id"),
            esb.CardinalityAgg("unique_users", "user_id"),
            esb.AvgAgg("avg_duration", "duration"),
            // 每种行为的设备分布
            esb.TermsAgg("device_breakdown", "device_type",
                esb.ValueCountAgg("device_actions", "action_id"),
            ),
        ),
        
        // 3. 页面访问分析
        esb.TopTermsAgg("popular_pages", "page_url", 50,
            esb.ValueCountAgg("page_views", "action_id"),
            esb.CardinalityAgg("unique_visitors", "user_id"),
            esb.AvgAgg("avg_time_on_page", "time_on_page"),
            esb.PercentilesAgg("bounce_rate_percentiles", "bounce_rate", []float64{25, 50, 75, 90}),
        ),
        
        // 4. 用户设备和浏览器分析
        esb.TermsAgg("devices", "device_type",
            esb.ValueCountAgg("sessions", "session_id"),
            esb.CardinalityAgg("unique_users", "user_id"),
            esb.TermsAgg("browsers", "browser",
                esb.ValueCountAgg("browser_sessions", "session_id"),
                esb.TermsAgg("os", "operating_system",
                    esb.ValueCountAgg("os_sessions", "session_id"),
                ),
            ),
        ),
        
        // 5. 地理位置分析
        esb.TermsAgg("countries", "country",
            esb.CardinalityAgg("unique_users", "user_id"),
            esb.ValueCountAgg("sessions", "session_id"),
            esb.TermsAgg("cities", "city",
                esb.ValueCountAgg("city_sessions", "session_id"),
            ),
        ),
        
        // 6. 用户留存分析
        esb.DateHistogramAgg("user_retention", "first_visit_date", "1w",
            esb.CardinalityAgg("new_users", "user_id"),
            esb.FilterAgg("returning_users", 
                esb.NumberRange("visit_count").Gt(1.0).Build(),
            ),
        ),
        
        // 7. 转化漏斗分析
        esb.FiltersAgg("conversion_funnel", map[string]esb.QueryOption{
            "page_view": esb.Term("action_type", "page_view"),
            "add_to_cart": esb.Term("action_type", "add_to_cart"),
            "checkout": esb.Term("action_type", "checkout"),
            "purchase": esb.Term("action_type", "purchase"),
        }),
        
        // 8. 会话时长分析
        esb.HistogramAgg("session_duration_dist", "session_duration", 300), // 5分钟间隔
        esb.PercentilesAgg("session_duration_percentiles", "session_duration", []float64{25, 50, 75, 90, 95}),
    )
}
```

### 管道聚合高级示例

```go
// 高级管道聚合示例
func BuildAdvancedPipelineAggs() *types.Aggregations {
    return esb.NewAggregations(
        // 基础时间序列数据
        esb.DateHistogramAgg("daily_sales", "date", "1d",
            esb.SumAgg("daily_total", "amount"),
            esb.ValueCountAgg("daily_orders", "order_count"),
            esb.AvgAgg("daily_avg", "amount"),
        ),
        
        // 1. 移动平均 - 7天移动平均
        esb.MovingAvgAgg("sales_moving_avg", "daily_sales>daily_total", 7),
        
        // 2. 累积求和
        esb.CumulativeSumAgg("cumulative_sales", "daily_sales>daily_total"),
        
        // 3. 导数 - 销售增长率
        esb.DerivativeAgg("sales_growth", "daily_sales>daily_total"),
        
        // 4. 桶统计分析
        esb.StatsBucketAgg("monthly_sales_stats", "daily_sales>daily_total"),
        esb.ExtendedStatsBucketAgg("detailed_sales_stats", "daily_sales>daily_total"),
        
        // 5. 最值分析
        esb.MaxBucketAgg("best_sales_day", "daily_sales>daily_total"),
        esb.MinBucketAgg("worst_sales_day", "daily_sales>daily_total"),
        
        // 6. 平均值分析
        esb.AvgBucketAgg("avg_daily_sales", "daily_sales>daily_total"),
        
        // 7. 百分位桶分析
        esb.PercentilesBucketAgg("sales_percentiles", "daily_sales>daily_total"),
        
        // 8. 求和桶分析
        esb.SumBucketAgg("total_period_sales", "daily_sales>daily_total"),
    )
}
```

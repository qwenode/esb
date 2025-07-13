# ESB API 文档

## 概述

ESB (Elasticsearch Query Builder) 提供了一套简洁的 API 来构建 Elasticsearch 查询。所有查询都基于函数式选项模式，支持链式调用。

## 核心类型

### QueryOption

```go
type QueryOption func(*types.Query) error
```

`QueryOption` 是所有查询构建器的基础类型，它是一个函数，接收 `*types.Query` 并返回错误。

### 错误类型

```go
var (
    ErrInvalidQuery = errors.New("invalid query")
    
    ErrNoOptions    = errors.New("no query options provided")
)
```

## 核心函数

### NewQuery

```go
func NewQuery(opts ...QueryOption) (*types.Query, error)
```

创建新的 Elasticsearch 查询。

**参数：**
- `opts ...QueryOption`: 可变数量的查询选项

**返回：**
- `*types.Query`: 兼容 go-elasticsearch 的查询对象
- `error`: 错误信息

**示例：**
```go
query, err := esb.NewQuery(
    esb.Term("status", "published"),
)
```

## 查询类型

### Term 查询

#### Term

```go
func Term(field, value string) QueryOption
```

创建精确匹配的 Term 查询。

**参数：**
- `field string`: 字段名
- `value string`: 匹配值

**示例：**
```go
esb.Term("status", "published")
esb.Term("category", "tech")
```

#### Terms

```go
func Terms(field string, values ...string) QueryOption
```

创建多值匹配的 Terms 查询。

**参数：**
- `field string`: 字段名
- `values ...string`: 匹配值列表

**示例：**
```go
esb.Terms("category", "tech", "science", "programming")
```

### Match 查询

#### Match

```go
func Match(field, query string) QueryOption
```

创建基本的 Match 查询。

**参数：**
- `field string`: 字段名
- `query string`: 查询字符串

**示例：**
```go
esb.Match("title", "elasticsearch guide")
```

#### MatchOptions

```go
type MatchOptions struct {
    Operator                        *operator.Operator
    MinimumShouldMatch              *types.MinimumShouldMatch
    Fuzziness                       types.Fuzziness
    FuzzyTranspositions             *bool
    FuzzyRewrite                    *string
    Lenient                         *bool
    Analyzer                        *string
    AutoGenerateSynonymsPhraseQuery *bool
    Boost                           *float32
    CutoffFrequency                 *types.Float64
    MaxExpansions                   *int
    PrefixLength                    *int
    ZeroTermsQuery                  *zerotermsquery.ZeroTermsQuery
}
```

#### MatchWithOptions

```go
func MatchWithOptions(field, query string, options MatchOptions) QueryOption
```

创建带选项的 Match 查询。

**参数：**
- `field string`: 字段名
- `query string`: 查询字符串
- `options MatchOptions`: 查询选项

**示例：**
```go
esb.MatchWithOptions("title", "elasticsearch guide", esb.MatchOptions{
    Fuzziness: "AUTO",
    Analyzer:  esb.StringPtr("standard"),
    Boost:     esb.Float32Ptr(1.5),
})
```

#### MatchPhrase

```go
func MatchPhrase(field, phrase string) QueryOption
```

创建短语匹配查询。

**参数：**
- `field string`: 字段名
- `phrase string`: 短语

**示例：**
```go
esb.MatchPhrase("content", "getting started")
```

#### MatchPhraseOptions

```go
type MatchPhraseOptions struct {
    Slop     *int
    Analyzer *string
    Boost    *float32
}
```

#### MatchPhraseWithOptions

```go
func MatchPhraseWithOptions(field, phrase string, options MatchPhraseOptions) QueryOption
```

创建带选项的短语匹配查询。

**参数：**
- `field string`: 字段名
- `phrase string`: 短语
- `options MatchPhraseOptions`: 查询选项

**示例：**
```go
esb.MatchPhraseWithOptions("content", "getting started", esb.MatchPhraseOptions{
    Slop:     esb.IntPtr(2),
    Analyzer: esb.StringPtr("keyword"),
    Boost:    esb.Float32Ptr(2.0),
})
```

#### MatchPhrasePrefix

```go
func MatchPhrasePrefix(field, prefix string) QueryOption
```

创建短语前缀匹配查询。

**参数：**
- `field string`: 字段名
- `prefix string`: 前缀

**示例：**
```go
esb.MatchPhrasePrefix("tags", "elastic")
```

### Range 查询

#### Range

```go
func Range(field string) *RangeBuilder
```

创建范围查询构建器。

**参数：**
- `field string`: 字段名

**返回：**
- `*RangeBuilder`: 范围查询构建器

#### RangeBuilder

```go
type RangeBuilder struct {
    field string
    query types.UntypedRangeQuery
}
```

范围查询构建器，支持链式调用。

**方法：**

##### Gte

```go
func (rb *RangeBuilder) Gte(value interface{}) *RangeBuilder
```

设置大于等于条件。

##### Gt

```go
func (rb *RangeBuilder) Gt(value interface{}) *RangeBuilder
```

设置大于条件。

##### Lte

```go
func (rb *RangeBuilder) Lte(value interface{}) *RangeBuilder
```

设置小于等于条件。

##### Lt

```go
func (rb *RangeBuilder) Lt(value interface{}) *RangeBuilder
```

设置小于条件。

##### From

```go
func (rb *RangeBuilder) From(value interface{}) *RangeBuilder
```

设置起始值（包含）。

##### To

```go
func (rb *RangeBuilder) To(value interface{}) *RangeBuilder
```

设置结束值（不包含）。

##### Boost

```go
func (rb *RangeBuilder) Boost(boost float32) *RangeBuilder
```

设置权重。

##### Format

```go
func (rb *RangeBuilder) Format(format string) *RangeBuilder
```

设置日期格式。

##### TimeZone

```go
func (rb *RangeBuilder) TimeZone(timeZone string) *RangeBuilder
```

设置时区。

##### Build

```go
func (rb *RangeBuilder) Build() QueryOption
```

构建最终的查询选项。

**示例：**
```go
// 基本范围查询
esb.Range("age").Gte(18).Lt(65).Build()

// 日期范围查询
esb.Range("timestamp").
    Gte("2023-01-01").
    Lte("2023-12-31").
    Format("yyyy-MM-dd").
    TimeZone("UTC").
    Build()

// 价格范围查询
esb.Range("price").From(10.0).To(100.0).Boost(1.5).Build()
```

### Exists 查询

#### Exists

```go
func Exists(field string) QueryOption
```

创建字段存在性查询。

**参数：**
- `field string`: 字段名

**示例：**
```go
esb.Exists("author")
esb.Exists("metadata.timestamp")
```

### Bool 查询

#### Bool

```go
func Bool(clauses ...BoolClause) QueryOption
```

创建布尔查询。

**参数：**
- `clauses ...BoolClause`: 布尔子句

#### BoolClause

```go
type BoolClause func(*types.BoolQuery) error
```

布尔子句类型。

#### Must

```go
func Must(queries ...QueryOption) BoolClause
```

创建 Must 子句（必须匹配）。

**参数：**
- `queries ...QueryOption`: 查询选项

#### Should

```go
func Should(queries ...QueryOption) BoolClause
```

创建 Should 子句（应该匹配）。

**参数：**
- `queries ...QueryOption`: 查询选项

#### Filter

```go
func Filter(queries ...QueryOption) BoolClause
```

创建 Filter 子句（过滤，不影响评分）。

**参数：**
- `queries ...QueryOption`: 查询选项

#### MustNot

```go
func MustNot(queries ...QueryOption) BoolClause
```

创建 MustNot 子句（必须不匹配）。

**参数：**
- `queries ...QueryOption`: 查询选项

**示例：**
```go
esb.Bool(
    esb.Must(
        esb.Match("title", "elasticsearch"),
        esb.Range("date").Gte("2023-01-01").Build(),
    ),
    esb.Should(
        esb.Term("category", "tech"),
        esb.Exists("featured"),
    ),
    esb.Filter(
        esb.Term("status", "published"),
    ),
    esb.MustNot(
        esb.Term("deleted", "true"),
    ),
)
```

## Helper 函数

### IntPtr

```go
func intPtr(i int) *int
```

创建 int 指针。

### StringPtr

```go
func stringPtr(s string) *string
```

创建 string 指针。

### Float32Ptr

```go
func float32Ptr(f float32) *float32
```

创建 float32 指针。

### BoolPtr

```go
func boolPtr(b bool) *bool
```

创建 bool 指针。



## 使用模式

### 基本查询

```go
// 单个查询
query, err := esb.NewQuery(
    esb.Term("status", "published"),
)

// 多个查询组合
query, err := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Match("title", "elasticsearch"),
            esb.Range("date").Gte("2023-01-01").Build(),
        ),
    ),
)
```

### 嵌套查询

```go
query, err := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Bool(
                esb.Should(
                    esb.Match("title", "elasticsearch"),
                    esb.Match("content", "search"),
                ),
            ),
            esb.Bool(
                esb.Must(
                    esb.Range("date").Gte("2023-01-01").Build(),
                    esb.Exists("author"),
                ),
            ),
        ),
    ),
)
```

### 错误处理

```go
query, err := esb.NewQuery(
    esb.Term("status", "published"),
)
if err != nil {
    // 处理错误
    return fmt.Errorf("failed to build query: %w", err)
}
```

## 最佳实践

1. **使用 Filter 而不是 Must 进行精确匹配**
2. **将复杂查询封装成函数**
3. **及时检查错误**
4. **避免过度嵌套**
5. **合理使用 Helper 函数**

更多示例请参考 [examples](../examples/) 目录。 
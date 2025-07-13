# ESB - Elasticsearch Query Builder

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.18-blue.svg)](https://golang.org/)
[![Test Coverage](https://img.shields.io/badge/coverage-92.4%25-brightgreen.svg)](https://github.com/qwenode/esb)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

**ESB** (Elasticsearch Query Builder) 是一个用于构建 Elasticsearch 查询的 Go 库，采用函数式选项模式提供链式调用 API，简化 `github.com/elastic/go-elasticsearch/v8` 的查询构建过程。

## ✨ 特性

- 🚀 **简洁易用**：链式 API 设计，减少 50% 以上样板代码
- 🔒 **类型安全**：完全兼容原生 `types.Query`，编译时类型检查
- 🎯 **功能完整**：支持主要查询类型（Term、Match、Range、Bool、Exists 等）
- 🧪 **高质量**：92.4% 测试覆盖率，全面的集成测试和基准测试
- 📚 **文档完善**：详细的 API 文档和使用示例
- ⚡ **性能可控**：虽然比原生方式慢 4-13 倍，但提供更好的可读性和维护性

## 🚀 快速开始

### 安装

```bash
go get github.com/qwenode/esb
```

### 基本使用

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/qwenode/esb"
)

func main() {
    // 创建一个简单的 Term 查询
    query, err := esb.NewQuery(
        esb.Term("status", "published"),
    )
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Query: %+v\n", query)
}
```

### 复杂查询示例

```go
// 创建复杂的 Bool 查询
query, err := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Match("title", "elasticsearch guide"),
            esb.Range("publish_date").Gte("2023-01-01").Lte("2023-12-31").Build(),
            esb.Exists("author"),
        ),
        esb.Should(
            esb.MatchPhrase("content", "search engine"),
            esb.Term("category", "technology"),
        ),
        esb.Filter(
            esb.Term("status", "published"),
            esb.Range("score").Gte(4.0).Build(),
        ),
        esb.MustNot(
            esb.Term("deleted", "true"),
        ),
    ),
)
```

## 📖 API 文档

### 核心函数

#### `NewQuery(opts ...QueryOption) (*types.Query, error)`

创建新的 Elasticsearch 查询。

**参数：**
- `opts`: 查询选项，可以是任何支持的查询类型

**返回：**
- `*types.Query`: 完全兼容 go-elasticsearch 的查询对象
- `error`: 错误信息

### 查询类型

#### Term 查询

精确匹配查询，用于 keyword 字段。

```go
// 基本 Term 查询
esb.Term("status", "published")

// Terms 查询（多值匹配）
esb.Terms("category", "tech", "science", "programming")
```

#### Match 查询

全文搜索查询，支持分析器处理。

```go
// 基本 Match 查询
esb.Match("title", "elasticsearch guide")

// 带选项的 Match 查询
esb.MatchWithOptions("title", "elasticsearch guide", esb.MatchOptions{
    Fuzziness: "AUTO",
    Analyzer:  esb.StringPtr("standard"),
    Boost:     esb.Float32Ptr(1.5),
})

// Match Phrase 查询
esb.MatchPhrase("content", "getting started")

// 带选项的 Match Phrase 查询
esb.MatchPhraseWithOptions("content", "getting started", esb.MatchPhraseOptions{
    Slop:     esb.IntPtr(2),
    Analyzer: esb.StringPtr("keyword"),
})

// Match Phrase Prefix 查询
esb.MatchPhrasePrefix("tags", "elastic")
```

#### Range 查询

范围查询，支持数值、日期和字符串。

```go
// 基本 Range 查询
esb.Range("age").Gte(18).Lt(65).Build()

// 带选项的 Range 查询
esb.Range("timestamp").
    Gte("2023-01-01").
    Lte("2023-12-31").
    Format("yyyy-MM-dd").
    TimeZone("UTC").
    Boost(1.5).
    Build()
```

**Range 方法：**
- `Gte(value)`: 大于等于
- `Gt(value)`: 大于
- `Lte(value)`: 小于等于
- `Lt(value)`: 小于
- `From(value)`: 起始值（包含）
- `To(value)`: 结束值（不包含）
- `Boost(boost)`: 权重
- `Format(format)`: 日期格式
- `TimeZone(tz)`: 时区

#### Exists 查询

检查字段是否存在。

```go
esb.Exists("author")
```

#### Bool 查询

布尔查询，支持复杂的逻辑组合。

```go
esb.Bool(
    esb.Must(
        // 必须匹配的条件
        esb.Match("title", "elasticsearch"),
        esb.Range("date").Gte("2023-01-01").Build(),
    ),
    esb.Should(
        // 应该匹配的条件（可选）
        esb.Term("category", "tech"),
        esb.Exists("featured"),
    ),
    esb.Filter(
        // 过滤条件（不影响评分）
        esb.Term("status", "published"),
    ),
    esb.MustNot(
        // 必须不匹配的条件
        esb.Term("deleted", "true"),
    ),
)
```

### Helper 函数

```go
// 指针 Helper 函数
esb.IntPtr(42)         // *int
esb.StringPtr("test")  // *string
esb.Float32Ptr(3.14)   // *float32
esb.BoolPtr(true)      // *bool
```

## 🎯 使用示例

### 与 Elasticsearch 客户端集成

```go
package main

import (
    "context"
    "log"
    
    "github.com/elastic/go-elasticsearch/v8"
    "github.com/qwenode/esb"
)

func main() {
    // 创建 Elasticsearch 客户端
    client, err := elasticsearch.NewDefaultClient()
    if err != nil {
        log.Fatal(err)
    }
    
    // 使用 ESB 构建查询
    query, err := esb.NewQuery(
        esb.Bool(
            esb.Must(
                esb.Match("title", "elasticsearch"),
                esb.Range("date").Gte("2023-01-01").Build(),
            ),
        ),
    )
    if err != nil {
        log.Fatal(err)
    }
    
    // 执行搜索
    res, err := client.Search(
        client.Search.WithContext(context.Background()),
        client.Search.WithIndex("my_index"),
        client.Search.WithQuery(query), // 直接使用 ESB 查询
    )
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()
    
    // 处理结果...
}
```

### 高级查询示例

```go
// 电商搜索示例
func buildProductSearchQuery(keyword string, minPrice, maxPrice float64, categories []string) (*types.Query, error) {
    return esb.NewQuery(
        esb.Bool(
            esb.Must(
                esb.Match("name", keyword),
                esb.Range("price").Gte(minPrice).Lte(maxPrice).Build(),
            ),
            esb.Should(
                esb.Terms("category", categories...),
                esb.Exists("discount"),
            ),
            esb.Filter(
                esb.Term("status", "active"),
                esb.Range("stock").Gt(0).Build(),
            ),
            esb.MustNot(
                esb.Term("deleted", "true"),
            ),
        ),
    )
}

// 内容管理系统搜索示例
func buildCMSSearchQuery(searchTerm string, authorID string, dateRange DateRange) (*types.Query, error) {
    return esb.NewQuery(
        esb.Bool(
            esb.Must(
                esb.MatchWithOptions("content", searchTerm, esb.MatchOptions{
                    Fuzziness: "AUTO",
                    Boost:     esb.Float32Ptr(1.2),
                }),
                esb.Term("author_id", authorID),
                esb.Range("created_at").
                    Gte(dateRange.Start).
                    Lte(dateRange.End).
                    Format("yyyy-MM-dd").
                    Build(),
            ),
            esb.Should(
                esb.MatchPhrase("title", searchTerm),
                esb.Exists("featured_image"),
            ),
            esb.Filter(
                esb.Term("status", "published"),
                esb.Range("views").Gte(100).Build(),
            ),
        ),
    )
}
```

## 📊 性能对比

基准测试结果（ESB vs 原生方式）：

| 查询类型 | ESB 性能 | 原生性能 | 性能比率 |
|---------|----------|----------|----------|
| Term | 321.1 ns/op | 24.21 ns/op | ~13x |
| Match | 297.4 ns/op | 75.34 ns/op | ~4x |
| Range | 488.9 ns/op | 98.71 ns/op | ~5x |
| Bool | 2607 ns/op | 256.5 ns/op | ~10x |
| 复杂查询 | 7051 ns/op | 581.8 ns/op | ~12x |

**性能说明：**
- ESB 比原生方式慢 4-13 倍
- 考虑到提供的便利性和可读性，这个性能代价是可接受的
- 对于大多数应用场景，这个性能差异不会成为瓶颈
- 复杂查询的构建时间通常远小于网络请求时间

## 🛠️ 最佳实践

### 1. 查询构建

```go
// ✅ 推荐：使用链式调用
query, err := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Match("title", "elasticsearch"),
            esb.Range("date").Gte("2023-01-01").Build(),
        ),
    ),
)

// ❌ 避免：过度嵌套
query, err := esb.NewQuery(
    esb.Bool(
        esb.Must(
            esb.Bool(
                esb.Should(
                    esb.Bool(
                        esb.Must(
                            esb.Term("field", "value"),
                        ),
                    ),
                ),
            ),
        ),
    ),
)
```

### 2. 错误处理

```go
// ✅ 推荐：及时检查错误
query, err := esb.NewQuery(
    esb.Term("status", "published"),
)
if err != nil {
    return fmt.Errorf("failed to build query: %w", err)
}

// ❌ 避免：忽略错误
query, _ := esb.NewQuery(
    esb.Term("status", "published"),
)
```

### 3. 性能优化

```go
// ✅ 推荐：使用 Filter 而不是 Must 进行精确匹配
esb.Bool(
    esb.Must(
        esb.Match("content", "search term"),
    ),
    esb.Filter(
        esb.Term("status", "published"),
        esb.Range("date").Gte("2023-01-01").Build(),
    ),
)

// ❌ 避免：在 Must 中使用不需要评分的查询
esb.Bool(
    esb.Must(
        esb.Match("content", "search term"),
        esb.Term("status", "published"),
        esb.Range("date").Gte("2023-01-01").Build(),
    ),
)
```

### 4. 代码组织

```go
// ✅ 推荐：将复杂查询封装成函数
func buildUserSearchQuery(name string, ageRange AgeRange, active bool) (*types.Query, error) {
    return esb.NewQuery(
        esb.Bool(
            esb.Must(
                esb.Match("name", name),
                esb.Range("age").Gte(ageRange.Min).Lte(ageRange.Max).Build(),
            ),
            esb.Filter(
                esb.Term("active", fmt.Sprintf("%t", active)),
            ),
        ),
    )
}

// ❌ 避免：在业务逻辑中直接构建复杂查询
func searchUsers(name string, minAge, maxAge int) {
    query, err := esb.NewQuery(
        esb.Bool(
            esb.Must(
                esb.Match("name", name),
                esb.Range("age").Gte(minAge).Lte(maxAge).Build(),
            ),
        ),
    )
    // ... 业务逻辑
}
```

## 🧪 测试

运行测试：

```bash
# 运行所有测试
go test -v

# 运行测试并生成覆盖率报告
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# 运行基准测试
go test -bench=. -benchmem
```

## 📁 项目结构

```
esb/
├── README.md              # 项目文档
├── LICENSE                # 许可证
├── go.mod                 # Go 模块文件
├── go.sum                 # Go 依赖校验
├── query.go               # 核心查询构建器
├── bool.go                # Bool 查询实现
├── match.go               # Match 查询实现
├── range.go               # Range 查询实现
├── exists.go              # Exists 查询实现
├── term.go                # Term 查询实现
├── *_test.go              # 单元测试
├── integration_test.go    # 集成测试
├── benchmark_test.go      # 基准测试
├── coverage_test.go       # 覆盖率测试
└── examples/              # 使用示例
    ├── bool_query_example.go
    ├── match_query_example.go
    ├── range_query_example.go
    └── exists_query_example.go
```

## 🤝 贡献

我们欢迎所有形式的贡献！

### 如何贡献

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

### 开发指南

1. 确保所有测试通过
2. 保持测试覆盖率在 90% 以上
3. 遵循 Go 代码规范
4. 添加必要的文档和示例

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🔗 相关链接

- [Elasticsearch 官方文档](https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html)
- [go-elasticsearch 客户端](https://github.com/elastic/go-elasticsearch)
- [问题反馈](https://github.com/qwenode/esb/issues)

## 📞 支持

如果您有任何问题或建议，请：

1. 查看 [文档](README.md)
2. 搜索 [已有问题](https://github.com/qwenode/esb/issues)
3. 创建 [新问题](https://github.com/qwenode/esb/issues/new)

---

**Made with ❤️ by the ESB team** 
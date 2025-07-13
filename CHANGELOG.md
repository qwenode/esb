# 变更日志

本文档记录了 ESB (Elasticsearch Query Builder) 项目的所有重要变更。

格式基于 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.0.0/)，
并且本项目遵循 [语义化版本](https://semver.org/lang/zh-CN/)。

## [未发布]

### 新增
- 初始项目设置和核心架构
- 核心 QueryOption 类型和 NewQuery 函数
- Bool 查询支持（Must、Should、Filter、MustNot）
- Match 查询支持（Match、MatchPhrase、MatchPhrasePrefix）
- Range 查询支持（链式构建器）
- Exists 查询支持
- Term 和 Terms 查询支持
- 完整的测试套件（92.4% 覆盖率）
- 基准测试和性能对比
- 完整的文档和 API 参考
- 使用示例和最佳实践指南

### 特性
- 🚀 **简洁易用**：链式 API 设计，减少 50% 以上样板代码
- 🔒 **类型安全**：完全兼容原生 `types.Query`，编译时类型检查
- 🎯 **功能完整**：支持主要查询类型（Term、Match、Range、Bool、Exists 等）
- 🧪 **高质量**：92.4% 测试覆盖率，全面的集成测试和基准测试
- 📚 **文档完善**：详细的 API 文档和使用示例
- ⚡ **性能可控**：虽然比原生方式慢 4-13 倍，但提供更好的可读性和维护性

### 技术细节
- 使用函数式选项模式设计
- 完全兼容 `github.com/elastic/go-elasticsearch/v8/typedapi/types`
- 支持嵌套查询和复杂布尔逻辑
- 提供 Helper 函数简化指针操作
- 完整的错误处理和输入验证

### 文档
- 📖 README.md - 项目主文档
- 📚 docs/API.md - 详细 API 文档
- 🤝 CONTRIBUTING.md - 贡献指南
- 📄 LICENSE - MIT 许可证
- 📋 CHANGELOG.md - 版本变更日志

### 示例
- examples/bool_query_example.go - Bool 查询示例
- examples/match_query_example.go - Match 查询示例
- examples/range_query_example.go - Range 查询示例
- examples/exists_query_example.go - Exists 查询示例

---

## 版本计划

### v1.0.0 (计划中)
- 稳定的 API 接口
- 完整的查询类型支持
- 生产就绪的质量保证

### v1.1.0 (未来)
- 聚合查询支持
- 更多查询类型（Wildcard、Regex 等）
- 性能优化

### v1.2.0 (未来)
- 查询模板支持
- 查询缓存机制
- 更多高级功能

---

*注意：此项目目前处于开发阶段，API 可能会有变化。* 
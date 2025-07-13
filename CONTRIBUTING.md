# 贡献指南

感谢您对 ESB (Elasticsearch Query Builder) 项目的关注！我们欢迎所有形式的贡献，包括但不限于：

- 🐛 Bug 报告
- 💡 功能建议
- 📖 文档改进
- 🔧 代码贡献
- 🧪 测试用例
- 📝 使用示例

## 开发环境设置

### 前置要求

- Go 1.18 或更高版本
- Git

### 克隆项目

```bash
git clone https://github.com/qwenode/esb.git
cd esb
```

### 安装依赖

```bash
go mod tidy
```

### 运行测试

```bash
# 运行所有测试
go test -v

# 运行测试并生成覆盖率报告
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# 运行基准测试
go test -bench=. -benchmem
```

## 贡献流程

### 1. 提交 Issue

在开始编码之前，请先提交一个 Issue 来描述：

- 🐛 **Bug 报告**：详细描述问题、复现步骤和预期行为
- 💡 **功能请求**：描述功能需求、使用场景和预期效果
- 📖 **文档改进**：指出文档中的问题或改进建议

### 2. Fork 和分支

```bash
# Fork 项目到您的 GitHub 账户
# 克隆您的 fork
git clone https://github.com/YOUR_USERNAME/esb.git
cd esb

# 创建功能分支
git checkout -b feature/your-feature-name
# 或者创建修复分支
git checkout -b fix/your-fix-name
```

### 3. 开发规范

#### 代码风格

- 遵循 Go 官方代码规范
- 使用 `gofmt` 格式化代码
- 使用 `golint` 检查代码质量
- 变量和函数命名要有意义

#### 提交规范

使用语义化提交信息：

```
type(scope): description

[optional body]

[optional footer]
```

**类型 (type)：**
- `feat`: 新功能
- `fix`: Bug 修复
- `docs`: 文档更新
- `style`: 代码格式化
- `refactor`: 代码重构
- `test`: 测试相关
- `chore`: 构建或辅助工具的变动

**示例：**
```
feat(range): add timezone support for date ranges

Add TimeZone() method to RangeBuilder for better date handling
in different timezones.

Closes #123
```

#### 测试要求

- 所有新功能必须包含测试
- 保持测试覆盖率在 90% 以上
- 测试应该覆盖正常情况和边界情况
- 包含基准测试（如果适用）

#### 文档要求

- 所有公开 API 必须有 Go 文档注释
- 复杂功能需要提供使用示例
- 更新 README.md（如果需要）
- 更新 API 文档（如果需要）

### 4. 提交 Pull Request

```bash
# 提交更改
git add .
git commit -m "feat(query): add new query type support"

# 推送到您的 fork
git push origin feature/your-feature-name
```

然后在 GitHub 上创建 Pull Request，包含：

- 📝 **清晰的标题和描述**
- 🔗 **关联的 Issue 链接**
- 📋 **更改内容说明**
- 🧪 **测试结果证明**
- 📖 **文档更新说明**

### 5. 代码审查

- 维护者会审查您的代码
- 根据反馈进行必要的修改
- 保持代码质量和项目一致性

## 开发指南

### 项目结构

```
esb/
├── README.md              # 项目文档
├── CONTRIBUTING.md        # 贡献指南
├── LICENSE               # 许可证
├── go.mod               # Go 模块文件
├── query.go             # 核心查询构建器
├── bool.go              # Bool 查询实现
├── match.go             # Match 查询实现
├── range.go             # Range 查询实现
├── exists.go            # Exists 查询实现
├── term.go              # Term 查询实现
├── *_test.go            # 测试文件
├── docs/                # 详细文档
│   └── API.md          # API 文档
└── examples/            # 使用示例
    ├── bool_query_example.go
    ├── match_query_example.go
    ├── range_query_example.go
    └── exists_query_example.go
```

### 添加新查询类型

如果您想添加新的查询类型，请遵循以下模式：

1. **创建新文件**：`new_query.go`
2. **实现 QueryOption**：
   ```go
   func NewQuery(field string, options ...Option) QueryOption {
       return func(q *types.Query) error {
           // 实现逻辑
       }
   }
   ```
3. **添加测试**：`new_query_test.go`
4. **添加示例**：`examples/new_query_example.go`
5. **更新文档**：README.md 和 docs/API.md

### 性能考虑

- 避免不必要的内存分配
- 使用基准测试验证性能
- 考虑大规模使用场景
- 保持 API 的简洁性

### 兼容性

- 保持与 `github.com/elastic/go-elasticsearch/v8` 的兼容性
- 不要破坏现有 API
- 新功能应该是向后兼容的

## 发布流程

### 版本号规则

遵循 [语义化版本](https://semver.org/lang/zh-CN/)：

- `MAJOR.MINOR.PATCH`
- **MAJOR**：不兼容的 API 修改
- **MINOR**：向后兼容的功能性新增
- **PATCH**：向后兼容的问题修正

### 发布检查清单

- [ ] 所有测试通过
- [ ] 测试覆盖率 ≥ 90%
- [ ] 文档已更新
- [ ] CHANGELOG.md 已更新
- [ ] 版本号已更新

## 社区准则

### 行为准则

- 🤝 **尊重他人**：友善对待所有参与者
- 💬 **建设性沟通**：提供有用的反馈和建议
- 🎯 **专注目标**：保持讨论的相关性
- 📚 **乐于学习**：开放接受新想法和反馈

### 获取帮助

- 📖 查看 [文档](README.md)
- 🔍 搜索 [已有 Issues](https://github.com/qwenode/esb/issues)
- 💬 创建 [新 Issue](https://github.com/qwenode/esb/issues/new)
- 📧 联系维护者

## 致谢

感谢所有为 ESB 项目做出贡献的开发者！您的参与让这个项目变得更好。

---

**再次感谢您的贡献！** 🎉 
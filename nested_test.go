package esb

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/childscoremode"
	"testing"
)

func TestNested(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		query    QueryOption
		expected string
	}{
		{
			name: "基本嵌套查询",
			path: "comments",
			query: Match("comments.author", "john"),
			expected: `{"nested":{"path":"comments","query":{"match":{"comments.author":{"query":"john"}}}}}`,
		},
		{
			name: "带布尔查询的嵌套查询",
			path: "comments",
			query: Bool(
				Must(
					Match("comments.author", "john"),
					Match("comments.content", "great"),
				),
			),
			expected: `{"nested":{"path":"comments","query":{"bool":{"must":[{"match":{"comments.author":{"query":"john"}}},{"match":{"comments.content":{"query":"great"}}}]}}}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 使用嵌套选项创建新查询
			query := NewQuery(Nested(tt.path, tt.query))

			// 将查询转换为 JSON
			actualJSON, err := json.Marshal(query)
			if err != nil {
				t.Fatalf("将查询转换为 JSON 失败：%v", err)
			}

			// 比较实际 JSON 与预期 JSON
			if string(actualJSON) != tt.expected {
				t.Errorf("预期 JSON：%s，但得到：%s", tt.expected, string(actualJSON))
			}

			// 验证查询可以被反序列化回 Query 对象
			var unmarshaled types.Query
			if err := json.Unmarshal(actualJSON, &unmarshaled); err != nil {
				t.Errorf("将 JSON 反序列化回 Query 失败：%v", err)
			}

			// 验证嵌套查询存在
			if unmarshaled.Nested == nil {
				t.Error("在反序列化的查询中，预期 Nested 不为 nil")
			}

			// 验证路径
			if unmarshaled.Nested.Path != tt.path {
				t.Errorf("预期路径 %q，得到 %q", tt.path, unmarshaled.Nested.Path)
			}
		})
	}
}

func TestNestedWithOptions(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		query    QueryOption
		setOpts  func(opts *types.NestedQuery)
		expected string
	}{
		{
			name: "带评分模式的嵌套查询",
			path: "comments",
			query: Match("comments.author", "john"),
			setOpts: func(opts *types.NestedQuery) {
				scoreMode := childscoremode.Avg
				opts.ScoreMode = &scoreMode
			},
			expected: `{"nested":{"path":"comments","query":{"match":{"comments.author":{"query":"john"}}},"score_mode":"avg"}}`,
		},
		{
			name: "带忽略未映射的嵌套查询",
			path: "comments",
			query: Match("comments.author", "john"),
			setOpts: func(opts *types.NestedQuery) {
				ignoreUnmapped := true
				opts.IgnoreUnmapped = &ignoreUnmapped
			},
			expected: `{"nested":{"ignore_unmapped":true,"path":"comments","query":{"match":{"comments.author":{"query":"john"}}}}}`,
		},
		{
			name: "带所有选项的嵌套查询",
			path: "comments",
			query: Match("comments.author", "john"),
			setOpts: func(opts *types.NestedQuery) {
				scoreMode := childscoremode.Max
				ignoreUnmapped := true
				boost := float32(2.0)
				opts.ScoreMode = &scoreMode
				opts.IgnoreUnmapped = &ignoreUnmapped
				opts.Boost = &boost
			},
			expected: `{"nested":{"boost":2,"ignore_unmapped":true,"path":"comments","query":{"match":{"comments.author":{"query":"john"}}},"score_mode":"max"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 使用嵌套选项创建新查询
			query := NewQuery(NestedWithOptions(tt.path, tt.query, tt.setOpts))

			// 将查询转换为 JSON
			actualJSON, err := json.Marshal(query)
			if err != nil {
				t.Fatalf("将查询转换为 JSON 失败：%v", err)
			}

			// 比较实际 JSON 与预期 JSON
			if string(actualJSON) != tt.expected {
				t.Errorf("预期 JSON：%s，但得到：%s", tt.expected, string(actualJSON))
			}

			// 验证查询可以被反序列化回 Query 对象
			var unmarshaled types.Query
			if err := json.Unmarshal(actualJSON, &unmarshaled); err != nil {
				t.Errorf("将 JSON 反序列化回 Query 失败：%v", err)
			}

			// 验证嵌套查询存在
			if unmarshaled.Nested == nil {
				t.Error("在反序列化的查询中，预期 Nested 不为 nil")
			}

			// 验证路径
			if unmarshaled.Nested.Path != tt.path {
				t.Errorf("预期路径 %q，得到 %q", tt.path, unmarshaled.Nested.Path)
			}
		})
	}

	t.Run("应该可以处理空回调", func(t *testing.T) {
		query := NewQuery(NestedWithOptions("comments", Match("comments.author", "john"), nil))
		if query.Nested == nil {
			t.Error("预期 Nested 不为 nil")
		}
		if query.Nested.Path != "comments" {
			t.Errorf("预期路径为 'comments'，得到 %s", query.Nested.Path)
		}
	})
} 
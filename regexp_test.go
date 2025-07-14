package esb

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"testing"
)

func TestRegexp(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		value    string
		expected string
	}{
		{
			name:     "基本正则表达式查询",
			field:    "username",
			value:    "j.*n",
			expected: `{"regexp":{"username":{"value":"j.*n"}}}`,
		},
		{
			name:     "带复杂模式的正则表达式查询",
			field:    "email",
			value:    ".*@example\\.com",
			expected: `{"regexp":{"email":{"value":".*@example\\.com"}}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 使用正则表达式选项创建新查询
			query := NewQuery(Regexp(tt.field, tt.value))

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

			// 验证正则表达式在反序列化的查询中存在
			if unmarshaled.Regexp == nil {
				t.Error("在反序列化的查询中，预期 Regexp 不为 nil")
			}

			// 验证特定字段在正则表达式查询中存在
			if _, ok := unmarshaled.Regexp[tt.field]; !ok {
				t.Errorf("预期在正则表达式查询中存在字段 %q", tt.field)
			}

			// 验证值
			if unmarshaled.Regexp[tt.field].Value != tt.value {
				t.Errorf("预期值 %q，得到 %q", tt.value, unmarshaled.Regexp[tt.field].Value)
			}
		})
	}
}

func TestRegexpWithOptions(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		value    string
		setOpts  func(opts *types.RegexpQuery)
		expected string
	}{
		{
			name:  "带标志和最大状态数的正则表达式查询",
			field: "username",
			value: "j.*n",
			setOpts: func(opts *types.RegexpQuery) {
				flags := "ALL"
				maxDeterminizedStates := 10000
				opts.Flags = &flags
				opts.MaxDeterminizedStates = &maxDeterminizedStates
			},
			expected: `{"regexp":{"username":{"flags":"ALL","max_determinized_states":10000,"value":"j.*n"}}}`,
		},
		{
			name:  "带大小写不敏感标志的正则表达式查询",
			field: "email",
			value: ".*@example\\.com",
			setOpts: func(opts *types.RegexpQuery) {
				flags := "CASE_INSENSITIVE"
				opts.Flags = &flags
			},
			expected: `{"regexp":{"email":{"flags":"CASE_INSENSITIVE","value":".*@example\\.com"}}}`,
		},
		{
			name:  "带权重的正则表达式查询",
			field: "username",
			value: "j.*n",
			setOpts: func(opts *types.RegexpQuery) {
				boost := float32(2.0)
				opts.Boost = &boost
			},
			expected: `{"regexp":{"username":{"boost":2,"value":"j.*n"}}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 使用正则表达式选项创建新查询
			query := NewQuery(RegexpWithOptions(tt.field, tt.value, tt.setOpts))

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

			// 验证正则表达式在反序列化的查询中存在
			if unmarshaled.Regexp == nil {
				t.Error("在反序列化的查询中，预期 Regexp 不为 nil")
			}

			// 验证特定字段在正则表达式查询中存在
			if _, ok := unmarshaled.Regexp[tt.field]; !ok {
				t.Errorf("预期在正则表达式查询中存在字段 %q", tt.field)
			}

			// 验证值
			if unmarshaled.Regexp[tt.field].Value != tt.value {
				t.Errorf("预期值 %q，得到 %q", tt.value, unmarshaled.Regexp[tt.field].Value)
			}
		})
	}

	t.Run("应该可以处理空回调", func(t *testing.T) {
		query := NewQuery(RegexpWithOptions("username", "j.*n", nil))
		if query.Regexp == nil {
			t.Error("预期 Regexp 不为 nil")
		}
		if query.Regexp["username"].Value != "j.*n" {
			t.Errorf("预期值为 'j.*n'，得到 %s", query.Regexp["username"].Value)
		}
	})
} 
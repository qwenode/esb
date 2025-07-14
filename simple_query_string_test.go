package esb

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"testing"
)

func TestSimpleQueryString(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected string
	}{
		{
			name:     "基本简单查询字符串",
			query:    "elasticsearch + search",
			expected: `{"simple_query_string":{"query":"elasticsearch + search"}}`,
		},
		{
			name:     "带特殊字符的简单查询字符串",
			query:    "elasticsearch + (search | database)",
			expected: `{"simple_query_string":{"query":"elasticsearch + (search | database)"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 使用简单查询字符串选项创建新查询
			query := NewQuery(SimpleQueryString(tt.query))

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

			// 验证简单查询字符串在反序列化的查询中存在
			if unmarshaled.SimpleQueryString == nil {
				t.Error("在反序列化的查询中，预期 SimpleQueryString 不为 nil")
			}

			// 验证查询值
			if unmarshaled.SimpleQueryString.Query != tt.query {
				t.Errorf("预期查询 %q，得到 %q", tt.query, unmarshaled.SimpleQueryString.Query)
			}
		})
	}
}

func TestSimpleQueryStringWithOptions(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		setOpts  func(opts *types.SimpleQueryStringQuery)
		expected string
	}{
		{
			name:  "带字段和权重的简单查询字符串",
			query: "elasticsearch",
			setOpts: func(opts *types.SimpleQueryStringQuery) {
				fields := []string{"title^2", "description"}
				boost := float32(2.0)
				opts.Fields = fields
				opts.Boost = &boost
			},
			expected: `{"simple_query_string":{"boost":2,"fields":["title^2","description"],"query":"elasticsearch"}}`,
		},
		{
			name:  "带默认操作符和标志的简单查询字符串",
			query: "search engine",
			setOpts: func(opts *types.SimpleQueryStringQuery) {
				defaultOp := operator.And
				flags := []string{"AND", "OR", "PREFIX"}
				opts.DefaultOperator = &defaultOp
				opts.Flags = flags
			},
			expected: `{"simple_query_string":{"default_operator":"and","flags":["AND","OR","PREFIX"],"query":"search engine"}}`,
		},
		{
			name:  "带分析器和最小匹配数的简单查询字符串",
			query: "search database",
			setOpts: func(opts *types.SimpleQueryStringQuery) {
				analyzer := "standard"
				minimumShouldMatch := "75%"
				opts.Analyzer = &analyzer
				opts.MinimumShouldMatch = &minimumShouldMatch
			},
			expected: `{"simple_query_string":{"analyzer":"standard","minimum_should_match":"75%","query":"search database"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 使用简单查询字符串选项创建新查询
			query := NewQuery(SimpleQueryStringWithOptions(tt.query, tt.setOpts))

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

			// 验证简单查询字符串在反序列化的查询中存在
			if unmarshaled.SimpleQueryString == nil {
				t.Error("在反序列化的查询中，预期 SimpleQueryString 不为 nil")
			}

			// 验证查询值
			if unmarshaled.SimpleQueryString.Query != tt.query {
				t.Errorf("预期查询 %q，得到 %q", tt.query, unmarshaled.SimpleQueryString.Query)
			}
		})
	}

	t.Run("应该可以处理空回调", func(t *testing.T) {
		query := NewQuery(SimpleQueryStringWithOptions("elasticsearch", nil))
		if query.SimpleQueryString == nil {
			t.Error("预期 SimpleQueryString 不为 nil")
		}
		if query.SimpleQueryString.Query != "elasticsearch" {
			t.Errorf("预期查询为 'elasticsearch'，得到 %s", query.SimpleQueryString.Query)
		}
	})
} 
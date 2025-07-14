package esb

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"testing"
)

func TestQueryString(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected string
	}{
		{
			name:     "基本查询字符串",
			query:    "title:elasticsearch",
			expected: `{"query_string":{"query":"title:elasticsearch"}}`,
		},
		{
			name:     "复杂查询字符串",
			query:    "title:elasticsearch AND (tags:search OR tags:database)",
			expected: `{"query_string":{"query":"title:elasticsearch AND (tags:search OR tags:database)"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 使用查询字符串选项创建新查询
			query := NewQuery(QueryString(tt.query))

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

			// 验证查询字符串在反序列化的查询中存在
			if unmarshaled.QueryString == nil {
				t.Error("在反序列化的查询中，预期 QueryString 不为 nil")
			}

			// 验证查询值
			if unmarshaled.QueryString.Query != tt.query {
				t.Errorf("预期查询 %q，得到 %q", tt.query, unmarshaled.QueryString.Query)
			}
		})
	}
}

func TestQueryStringWithOptions(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		setOpts  func(opts *types.QueryStringQuery)
		expected string
	}{
		{
			name:  "带默认字段和操作符的查询字符串",
			query: "elasticsearch",
			setOpts: func(opts *types.QueryStringQuery) {
				defaultField := "title"
				defaultOp := operator.And
				opts.DefaultField = &defaultField
				opts.DefaultOperator = &defaultOp
			},
			expected: `{"query_string":{"default_field":"title","default_operator":"and","query":"elasticsearch"}}`,
		},
		{
			name:  "带多个字段和权重的查询字符串",
			query: "search engine",
			setOpts: func(opts *types.QueryStringQuery) {
				fields := []string{"title^2", "description"}
				boost := float32(2.0)
				opts.Fields = fields
				opts.Boost = &boost
			},
			expected: `{"query_string":{"boost":2,"fields":["title^2","description"],"query":"search engine"}}`,
		},
		{
			name:  "带分析器和允许前导通配符的查询字符串",
			query: "*search",
			setOpts: func(opts *types.QueryStringQuery) {
				analyzer := "standard"
				allowLeadingWildcard := true
				opts.Analyzer = &analyzer
				opts.AllowLeadingWildcard = &allowLeadingWildcard
			},
			expected: `{"query_string":{"allow_leading_wildcard":true,"analyzer":"standard","query":"*search"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 使用查询字符串选项创建新查询
			query := NewQuery(QueryStringWithOptions(tt.query, tt.setOpts))

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

			// 验证查询字符串在反序列化的查询中存在
			if unmarshaled.QueryString == nil {
				t.Error("在反序列化的查询中，预期 QueryString 不为 nil")
			}

			// 验证查询值
			if unmarshaled.QueryString.Query != tt.query {
				t.Errorf("预期查询 %q，得到 %q", tt.query, unmarshaled.QueryString.Query)
			}
		})
	}

	t.Run("应该可以处理空回调", func(t *testing.T) {
		query := NewQuery(QueryStringWithOptions("elasticsearch", nil))
		if query.QueryString == nil {
			t.Error("预期 QueryString 不为 nil")
		}
		if query.QueryString.Query != "elasticsearch" {
			t.Errorf("预期查询为 'elasticsearch'，得到 %s", query.QueryString.Query)
		}
	})
} 
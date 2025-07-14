package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// TestWildcardWithOptions 测试函数式选项模式的通配符查询
func TestWildcardWithOptions(t *testing.T) {
	t.Run("测试基本函数式选项配置", func(t *testing.T) {
		// 测试使用函数式选项模式的 WildcardWithOptions
		query := NewQuery(
			WildcardWithOptions("title", "java*", func(opts *types.WildcardQuery) {
				boost := float32(2.0)
				opts.Boost = &boost
				caseInsensitive := true
				opts.CaseInsensitive = &caseInsensitive
				rewrite := "top_terms_10"
				opts.Rewrite = &rewrite
				queryName := "test_wildcard"
				opts.QueryName_ = &queryName
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		// 验证字段存在
		wildcardQuery, exists := query.Wildcard["title"]
		if !exists {
			t.Fatal("应该存在title字段的通配符查询")
		}
		
		// 验证基本值
		if wildcardQuery.Value == nil || *wildcardQuery.Value != "java*" {
			t.Errorf("期望Value为'java*', 实际得到: %v", wildcardQuery.Value)
		}
		
		// 验证选项配置
		if wildcardQuery.Boost == nil || *wildcardQuery.Boost != 2.0 {
			t.Errorf("期望Boost为2.0, 实际得到: %v", wildcardQuery.Boost)
		}
		
		if wildcardQuery.CaseInsensitive == nil || *wildcardQuery.CaseInsensitive != true {
			t.Errorf("期望CaseInsensitive为true, 实际得到: %v", wildcardQuery.CaseInsensitive)
		}
		
		if wildcardQuery.Rewrite == nil || *wildcardQuery.Rewrite != "top_terms_10" {
			t.Errorf("期望Rewrite为'top_terms_10', 实际得到: %v", wildcardQuery.Rewrite)
		}
		
		if wildcardQuery.QueryName_ == nil || *wildcardQuery.QueryName_ != "test_wildcard" {
			t.Errorf("期望QueryName为'test_wildcard', 实际得到: %v", wildcardQuery.QueryName_)
		}
	})
	
	t.Run("测试nil选项配置", func(t *testing.T) {
		// 测试传递nil选项配置函数
		query := NewQuery(
			WildcardWithOptions("content", "test*", nil),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		// 验证字段存在
		wildcardQuery, exists := query.Wildcard["content"]
		if !exists {
			t.Fatal("应该存在content字段的通配符查询")
		}
		
		// 验证基本值
		if wildcardQuery.Value == nil || *wildcardQuery.Value != "test*" {
			t.Errorf("期望Value为'test*', 实际得到: %v", wildcardQuery.Value)
		}
		
		// 验证选项为默认值
		if wildcardQuery.Boost != nil {
			t.Errorf("期望Boost为nil, 实际得到: %v", wildcardQuery.Boost)
		}
	})
	
	t.Run("测试空选项配置", func(t *testing.T) {
		// 测试传递空的选项配置函数
		query := NewQuery(
			WildcardWithOptions("name", "john*", func(opts *types.WildcardQuery) {
				// 不设置任何选项
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		// 验证字段存在
		wildcardQuery, exists := query.Wildcard["name"]
		if !exists {
			t.Fatal("应该存在name字段的通配符查询")
		}
		
		// 验证基本值
		if wildcardQuery.Value == nil || *wildcardQuery.Value != "john*" {
			t.Errorf("期望Value为'john*', 实际得到: %v", wildcardQuery.Value)
		}
		
		// 验证选项为默认值
		if wildcardQuery.Boost != nil {
			t.Errorf("期望Boost为nil, 实际得到: %v", wildcardQuery.Boost)
		}
		if wildcardQuery.CaseInsensitive != nil {
			t.Errorf("期望CaseInsensitive为nil, 实际得到: %v", wildcardQuery.CaseInsensitive)
		}
	})
} 
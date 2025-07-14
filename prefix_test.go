package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// TestPrefixWithOptions 测试函数式选项模式的前缀查询
func TestPrefixWithOptions(t *testing.T) {
	t.Run("测试基本函数式选项配置", func(t *testing.T) {
		// 测试使用函数式选项模式的 PrefixWithOptions
		query := NewQuery(
			PrefixWithOptions("username", "john", func(opts *types.PrefixQuery) {
				boost := float32(1.5)
				opts.Boost = &boost
				caseInsensitive := true
				opts.CaseInsensitive = &caseInsensitive
				rewrite := "top_terms_10"
				opts.Rewrite = &rewrite
				queryName := "test_prefix"
				opts.QueryName_ = &queryName
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Prefix == nil {
			t.Fatal("Prefix查询不应该为nil")
		}
		
		// 验证字段存在
		prefixQuery, exists := query.Prefix["username"]
		if !exists {
			t.Fatal("应该存在username字段的前缀查询")
		}
		
		// 验证基本值
		if prefixQuery.Value != "john" {
			t.Errorf("期望Value为'john', 实际得到: %v", prefixQuery.Value)
		}
		
		// 验证选项配置
		if prefixQuery.Boost == nil || *prefixQuery.Boost != 1.5 {
			t.Errorf("期望Boost为1.5, 实际得到: %v", prefixQuery.Boost)
		}
		
		if prefixQuery.CaseInsensitive == nil || *prefixQuery.CaseInsensitive != true {
			t.Errorf("期望CaseInsensitive为true, 实际得到: %v", prefixQuery.CaseInsensitive)
		}
		
		if prefixQuery.Rewrite == nil || *prefixQuery.Rewrite != "top_terms_10" {
			t.Errorf("期望Rewrite为'top_terms_10', 实际得到: %v", prefixQuery.Rewrite)
		}
		
		if prefixQuery.QueryName_ == nil || *prefixQuery.QueryName_ != "test_prefix" {
			t.Errorf("期望QueryName为'test_prefix', 实际得到: %v", prefixQuery.QueryName_)
		}
	})
	
	t.Run("测试nil选项配置", func(t *testing.T) {
		// 测试传递nil选项配置函数
		query := NewQuery(
			PrefixWithOptions("title", "java", nil),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Prefix == nil {
			t.Fatal("Prefix查询不应该为nil")
		}
		
		// 验证字段存在
		prefixQuery, exists := query.Prefix["title"]
		if !exists {
			t.Fatal("应该存在title字段的前缀查询")
		}
		
		// 验证基本值
		if prefixQuery.Value != "java" {
			t.Errorf("期望Value为'java', 实际得到: %v", prefixQuery.Value)
		}
		
		// 验证选项为默认值
		if prefixQuery.Boost != nil {
			t.Errorf("期望Boost为nil, 实际得到: %v", prefixQuery.Boost)
		}
	})
	
	t.Run("测试空选项配置", func(t *testing.T) {
		// 测试传递空的选项配置函数
		query := NewQuery(
			PrefixWithOptions("category", "tech", func(opts *types.PrefixQuery) {
				// 不设置任何选项
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Prefix == nil {
			t.Fatal("Prefix查询不应该为nil")
		}
		
		// 验证字段存在
		prefixQuery, exists := query.Prefix["category"]
		if !exists {
			t.Fatal("应该存在category字段的前缀查询")
		}
		
		// 验证基本值
		if prefixQuery.Value != "tech" {
			t.Errorf("期望Value为'tech', 实际得到: %v", prefixQuery.Value)
		}
		
		// 验证选项为默认值
		if prefixQuery.Boost != nil {
			t.Errorf("期望Boost为nil, 实际得到: %v", prefixQuery.Boost)
		}
		if prefixQuery.CaseInsensitive != nil {
			t.Errorf("期望CaseInsensitive为nil, 实际得到: %v", prefixQuery.CaseInsensitive)
		}
	})
} 
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

// TestPrefixBasic 测试基本的前缀查询功能
func TestPrefixBasic(t *testing.T) {
	t.Run("测试基本前缀查询", func(t *testing.T) {
		query := NewQuery(Prefix("username", "john"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Prefix == nil {
			t.Fatal("Prefix查询不应该为nil")
		}
		
		prefixQuery, exists := query.Prefix["username"]
		if !exists {
			t.Fatal("期望存在username字段")
		}
		
		if prefixQuery.Value != "john" {
			t.Errorf("期望Value为'john', 实际得到: %s", prefixQuery.Value)
		}
	})
	
	t.Run("测试空前缀", func(t *testing.T) {
		query := NewQuery(Prefix("username", ""))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Prefix == nil {
			t.Fatal("Prefix查询不应该为nil")
		}
		
		prefixQuery, exists := query.Prefix["username"]
		if !exists {
			t.Fatal("期望存在username字段")
		}
		
		if prefixQuery.Value != "" {
			t.Errorf("期望Value为空字符串, 实际得到: %s", prefixQuery.Value)
		}
	})
	
	t.Run("测试特殊字符前缀", func(t *testing.T) {
		query := NewQuery(Prefix("path", "/usr/"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Prefix == nil {
			t.Fatal("Prefix查询不应该为nil")
		}
		
		prefixQuery, exists := query.Prefix["path"]
		if !exists {
			t.Fatal("期望存在path字段")
		}
		
		if prefixQuery.Value != "/usr/" {
			t.Errorf("期望Value为'/usr/', 实际得到: %s", prefixQuery.Value)
		}
	})
	
	t.Run("测试Unicode字符前缀", func(t *testing.T) {
		query := NewQuery(Prefix("name", "张"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Prefix == nil {
			t.Fatal("Prefix查询不应该为nil")
		}
		
		prefixQuery, exists := query.Prefix["name"]
		if !exists {
			t.Fatal("期望存在name字段")
		}
		
		if prefixQuery.Value != "张" {
			t.Errorf("期望Value为'张', 实际得到: %s", prefixQuery.Value)
		}
	})
}

// TestPrefixWithOptionsExtended 测试前缀查询的扩展选项
func TestPrefixWithOptionsExtended(t *testing.T) {
	t.Run("测试所有选项", func(t *testing.T) {
		boost := float32(2.0)
		queryName := "test_prefix"
		rewrite := "constant_score"
		
		query := NewQuery(
			PrefixWithOptions("username", "john", func(q *types.PrefixQuery) {
				q.Boost = &boost
				q.QueryName_ = &queryName
				q.Rewrite = &rewrite
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Prefix == nil {
			t.Fatal("Prefix查询不应该为nil")
		}
		
		prefixQuery, exists := query.Prefix["username"]
		if !exists {
			t.Fatal("期望存在username字段")
		}
		
		if prefixQuery.Value != "john" {
			t.Errorf("期望Value为'john', 实际得到: %s", prefixQuery.Value)
		}
		
		if *prefixQuery.Boost != 2.0 {
			t.Errorf("期望Boost为2.0, 实际得到: %f", *prefixQuery.Boost)
		}
		
		if *prefixQuery.QueryName_ != "test_prefix" {
			t.Errorf("期望QueryName为'test_prefix', 实际得到: %s", *prefixQuery.QueryName_)
		}
		
		if *prefixQuery.Rewrite != "constant_score" {
			t.Errorf("期望Rewrite为'constant_score', 实际得到: %s", *prefixQuery.Rewrite)
		}
	})
	
	t.Run("测试部分选项", func(t *testing.T) {
		boost := float32(1.5)
		
		query := NewQuery(
			PrefixWithOptions("username", "john", func(q *types.PrefixQuery) {
				q.Boost = &boost
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Prefix == nil {
			t.Fatal("Prefix查询不应该为nil")
		}
		
		prefixQuery, exists := query.Prefix["username"]
		if !exists {
			t.Fatal("期望存在username字段")
		}
		
		if prefixQuery.Value != "john" {
			t.Errorf("期望Value为'john', 实际得到: %s", prefixQuery.Value)
		}
		
		if *prefixQuery.Boost != 1.5 {
			t.Errorf("期望Boost为1.5, 实际得到: %f", *prefixQuery.Boost)
		}
		
		if prefixQuery.QueryName_ != nil {
			t.Error("期望QueryName为nil")
		}
		
		if prefixQuery.Rewrite != nil {
			t.Error("期望Rewrite为nil")
		}
	})
	
	t.Run("测试nil选项函数", func(t *testing.T) {
		query := NewQuery(
			PrefixWithOptions("username", "john", nil),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Prefix == nil {
			t.Fatal("Prefix查询不应该为nil")
		}
		
		prefixQuery, exists := query.Prefix["username"]
		if !exists {
			t.Fatal("期望存在username字段")
		}
		
		if prefixQuery.Value != "john" {
			t.Errorf("期望Value为'john', 实际得到: %s", prefixQuery.Value)
		}
		
		if prefixQuery.Boost != nil {
			t.Error("期望Boost为nil")
		}
		
		if prefixQuery.QueryName_ != nil {
			t.Error("期望QueryName为nil")
		}
		
		if prefixQuery.Rewrite != nil {
			t.Error("期望Rewrite为nil")
		}
	})
}

// TestPrefixEdgeCases 测试前缀查询的边界条件
func TestPrefixEdgeCases(t *testing.T) {
	t.Run("测试长前缀", func(t *testing.T) {
		longPrefix := "this_is_a_very_long_prefix_that_should_still_work_as_expected"
		query := NewQuery(Prefix("field", longPrefix))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Prefix == nil {
			t.Fatal("Prefix查询不应该为nil")
		}
		
		prefixQuery, exists := query.Prefix["field"]
		if !exists {
			t.Fatal("期望存在field字段")
		}
		
		if prefixQuery.Value != longPrefix {
			t.Errorf("期望Value为'%s', 实际得到: %s", longPrefix, prefixQuery.Value)
		}
	})
	
	t.Run("测试特殊字段名", func(t *testing.T) {
		query := NewQuery(Prefix("field.with.dots", "test"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Prefix == nil {
			t.Fatal("Prefix查询不应该为nil")
		}
		
		prefixQuery, exists := query.Prefix["field.with.dots"]
		if !exists {
			t.Fatal("期望存在field.with.dots字段")
		}
		
		if prefixQuery.Value != "test" {
			t.Errorf("期望Value为'test', 实际得到: %s", prefixQuery.Value)
		}
	})
	
	t.Run("测试空字段名", func(t *testing.T) {
		query := NewQuery(Prefix("", "test"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Prefix == nil {
			t.Fatal("Prefix查询不应该为nil")
		}
		
		prefixQuery, exists := query.Prefix[""]
		if !exists {
			t.Fatal("期望存在空字段")
		}
		
		if prefixQuery.Value != "test" {
			t.Errorf("期望Value为'test', 实际得到: %s", prefixQuery.Value)
		}
	})
} 
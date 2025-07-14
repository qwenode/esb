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

// TestWildcard 测试Wildcard查询功能
func TestWildcard(t *testing.T) {
	t.Run("测试基本通配符查询", func(t *testing.T) {
		query := NewQuery(Wildcard("username", "john*"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		if len(query.Wildcard) != 1 {
			t.Errorf("期望Wildcard查询长度为1, 实际得到: %d", len(query.Wildcard))
		}
		
		if wildcardQuery, exists := query.Wildcard["username"]; !exists {
			t.Error("期望存在username字段")
		} else if *wildcardQuery.Value != "john*" {
			t.Errorf("期望Value为'john*', 实际得到: %s", *wildcardQuery.Value)
		}
	})
	
	t.Run("测试问号通配符查询", func(t *testing.T) {
		query := NewQuery(Wildcard("product", "iphone-?"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		if wildcardQuery, exists := query.Wildcard["product"]; !exists {
			t.Error("期望存在product字段")
		} else if *wildcardQuery.Value != "iphone-?" {
			t.Errorf("期望Value为'iphone-?', 实际得到: %s", *wildcardQuery.Value)
		}
	})
	
	t.Run("测试多个通配符查询", func(t *testing.T) {
		query := NewQuery(Wildcard("email", "*@*.com"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		if wildcardQuery, exists := query.Wildcard["email"]; !exists {
			t.Error("期望存在email字段")
		} else if *wildcardQuery.Value != "*@*.com" {
			t.Errorf("期望Value为'*@*.com', 实际得到: %s", *wildcardQuery.Value)
		}
	})
	
	t.Run("测试空值的通配符查询", func(t *testing.T) {
		query := NewQuery(Wildcard("username", ""))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		if wildcardQuery, exists := query.Wildcard["username"]; !exists {
			t.Error("期望存在username字段")
		} else if *wildcardQuery.Value != "" {
			t.Errorf("期望Value为空字符串, 实际得到: %s", *wildcardQuery.Value)
		}
	})
}

// TestWildcardWithOptionsComplete 测试WildcardWithOptions的完整功能
func TestWildcardWithOptionsComplete(t *testing.T) {
	t.Run("测试所有选项的通配符查询", func(t *testing.T) {
		boost := float32(2.0)
		queryName := "test_wildcard"
		rewrite := "constant_score"
		
		query := NewQuery(
			WildcardWithOptions("username", "john*", func(q *types.WildcardQuery) {
				q.Boost = &boost
				q.QueryName_ = &queryName
				q.Rewrite = &rewrite
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		wildcardQuery, exists := query.Wildcard["username"]
		if !exists {
			t.Fatal("期望存在username字段")
		}
		
		if *wildcardQuery.Value != "john*" {
			t.Errorf("期望Value为'john*', 实际得到: %s", *wildcardQuery.Value)
		}
		
		if *wildcardQuery.Boost != 2.0 {
			t.Errorf("期望Boost为2.0, 实际得到: %f", *wildcardQuery.Boost)
		}
		
		if *wildcardQuery.QueryName_ != "test_wildcard" {
			t.Errorf("期望QueryName为'test_wildcard', 实际得到: %s", *wildcardQuery.QueryName_)
		}
		
		if *wildcardQuery.Rewrite != "constant_score" {
			t.Errorf("期望Rewrite为'constant_score', 实际得到: %s", *wildcardQuery.Rewrite)
		}
	})
	
	t.Run("测试部分选项的通配符查询", func(t *testing.T) {
		boost := float32(1.5)
		
		query := NewQuery(
			WildcardWithOptions("username", "john*", func(q *types.WildcardQuery) {
				q.Boost = &boost
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		wildcardQuery, exists := query.Wildcard["username"]
		if !exists {
			t.Fatal("期望存在username字段")
		}
		
		if *wildcardQuery.Value != "john*" {
			t.Errorf("期望Value为'john*', 实际得到: %s", *wildcardQuery.Value)
		}
		
		if *wildcardQuery.Boost != 1.5 {
			t.Errorf("期望Boost为1.5, 实际得到: %f", *wildcardQuery.Boost)
		}
		
		if wildcardQuery.QueryName_ != nil {
			t.Error("期望QueryName为nil")
		}
		
		if wildcardQuery.Rewrite != nil {
			t.Error("期望Rewrite为nil")
		}
	})
	
	t.Run("测试特殊字符的通配符查询", func(t *testing.T) {
		query := NewQuery(
			WildcardWithOptions("path", "/usr/*/bin/?", func(q *types.WildcardQuery) {
				// 不设置任何选项
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		wildcardQuery, exists := query.Wildcard["path"]
		if !exists {
			t.Fatal("期望存在path字段")
		}
		
		if *wildcardQuery.Value != "/usr/*/bin/?" {
			t.Errorf("期望Value为'/usr/*/bin/?', 实际得到: %s", *wildcardQuery.Value)
		}
	})
} 

// TestWildcardBasic 测试基本的通配符查询功能
func TestWildcardBasic(t *testing.T) {
	t.Run("测试星号通配符", func(t *testing.T) {
		query := NewQuery(Wildcard("username", "john*"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		wildcardQuery, exists := query.Wildcard["username"]
		if !exists {
			t.Fatal("期望存在username字段")
		}
		
		if wildcardQuery.Value == nil || *wildcardQuery.Value != "john*" {
			t.Errorf("期望Value为'john*', 实际得到: %v", wildcardQuery.Value)
		}
	})
	
	t.Run("测试问号通配符", func(t *testing.T) {
		query := NewQuery(Wildcard("product", "iphone-?"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		wildcardQuery, exists := query.Wildcard["product"]
		if !exists {
			t.Fatal("期望存在product字段")
		}
		
		if wildcardQuery.Value == nil || *wildcardQuery.Value != "iphone-?" {
			t.Errorf("期望Value为'iphone-?', 实际得到: %v", wildcardQuery.Value)
		}
	})
	
	t.Run("测试多个通配符", func(t *testing.T) {
		query := NewQuery(Wildcard("email", "*@*.com"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		wildcardQuery, exists := query.Wildcard["email"]
		if !exists {
			t.Fatal("期望存在email字段")
		}
		
		if wildcardQuery.Value == nil || *wildcardQuery.Value != "*@*.com" {
			t.Errorf("期望Value为'*@*.com', 实际得到: %v", wildcardQuery.Value)
		}
	})
	
	t.Run("测试空值", func(t *testing.T) {
		query := NewQuery(Wildcard("username", ""))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		wildcardQuery, exists := query.Wildcard["username"]
		if !exists {
			t.Fatal("期望存在username字段")
		}
		
		if wildcardQuery.Value == nil || *wildcardQuery.Value != "" {
			t.Errorf("期望Value为空字符串, 实际得到: %v", wildcardQuery.Value)
		}
	})
}

// TestWildcardWithOptionsExtended 测试通配符查询的扩展选项
func TestWildcardWithOptionsExtended(t *testing.T) {
	t.Run("测试所有选项", func(t *testing.T) {
		boost := float32(2.0)
		queryName := "test_wildcard"
		rewrite := "constant_score"
		
		query := NewQuery(
			WildcardWithOptions("username", "john*", func(q *types.WildcardQuery) {
				q.Boost = &boost
				q.QueryName_ = &queryName
				q.Rewrite = &rewrite
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		wildcardQuery, exists := query.Wildcard["username"]
		if !exists {
			t.Fatal("期望存在username字段")
		}
		
		if wildcardQuery.Value == nil || *wildcardQuery.Value != "john*" {
			t.Errorf("期望Value为'john*', 实际得到: %v", wildcardQuery.Value)
		}
		
		if wildcardQuery.Boost == nil || *wildcardQuery.Boost != 2.0 {
			t.Errorf("期望Boost为2.0, 实际得到: %v", wildcardQuery.Boost)
		}
		
		if wildcardQuery.QueryName_ == nil || *wildcardQuery.QueryName_ != "test_wildcard" {
			t.Errorf("期望QueryName为'test_wildcard', 实际得到: %v", wildcardQuery.QueryName_)
		}
		
		if wildcardQuery.Rewrite == nil || *wildcardQuery.Rewrite != "constant_score" {
			t.Errorf("期望Rewrite为'constant_score', 实际得到: %v", wildcardQuery.Rewrite)
		}
	})
	
	t.Run("测试部分选项", func(t *testing.T) {
		boost := float32(1.5)
		
		query := NewQuery(
			WildcardWithOptions("username", "john*", func(q *types.WildcardQuery) {
				q.Boost = &boost
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		wildcardQuery, exists := query.Wildcard["username"]
		if !exists {
			t.Fatal("期望存在username字段")
		}
		
		if wildcardQuery.Value == nil || *wildcardQuery.Value != "john*" {
			t.Errorf("期望Value为'john*', 实际得到: %v", wildcardQuery.Value)
		}
		
		if wildcardQuery.Boost == nil || *wildcardQuery.Boost != 1.5 {
			t.Errorf("期望Boost为1.5, 实际得到: %v", wildcardQuery.Boost)
		}
		
		if wildcardQuery.QueryName_ != nil {
			t.Error("期望QueryName为nil")
		}
		
		if wildcardQuery.Rewrite != nil {
			t.Error("期望Rewrite为nil")
		}
	})
	
	t.Run("测试nil选项函数", func(t *testing.T) {
		query := NewQuery(
			WildcardWithOptions("username", "john*", nil),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		wildcardQuery, exists := query.Wildcard["username"]
		if !exists {
			t.Fatal("期望存在username字段")
		}
		
		if wildcardQuery.Value == nil || *wildcardQuery.Value != "john*" {
			t.Errorf("期望Value为'john*', 实际得到: %v", wildcardQuery.Value)
		}
		
		if wildcardQuery.Boost != nil {
			t.Error("期望Boost为nil")
		}
		
		if wildcardQuery.QueryName_ != nil {
			t.Error("期望QueryName为nil")
		}
		
		if wildcardQuery.Rewrite != nil {
			t.Error("期望Rewrite为nil")
		}
	})
}

// TestWildcardEdgeCases 测试通配符查询的边界条件
func TestWildcardEdgeCases(t *testing.T) {
	t.Run("测试特殊字符", func(t *testing.T) {
		query := NewQuery(Wildcard("path", "/usr/*/bin/?"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		wildcardQuery, exists := query.Wildcard["path"]
		if !exists {
			t.Fatal("期望存在path字段")
		}
		
		if wildcardQuery.Value == nil || *wildcardQuery.Value != "/usr/*/bin/?" {
			t.Errorf("期望Value为'/usr/*/bin/?', 实际得到: %v", wildcardQuery.Value)
		}
	})
	
	t.Run("测试Unicode字符", func(t *testing.T) {
		query := NewQuery(Wildcard("name", "张*"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		wildcardQuery, exists := query.Wildcard["name"]
		if !exists {
			t.Fatal("期望存在name字段")
		}
		
		if wildcardQuery.Value == nil || *wildcardQuery.Value != "张*" {
			t.Errorf("期望Value为'张*', 实际得到: %v", wildcardQuery.Value)
		}
	})
	
	t.Run("测试只有通配符", func(t *testing.T) {
		query := NewQuery(Wildcard("field", "*"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Wildcard == nil {
			t.Fatal("Wildcard查询不应该为nil")
		}
		
		wildcardQuery, exists := query.Wildcard["field"]
		if !exists {
			t.Fatal("期望存在field字段")
		}
		
		if wildcardQuery.Value == nil || *wildcardQuery.Value != "*" {
			t.Errorf("期望Value为'*', 实际得到: %v", wildcardQuery.Value)
		}
	})
} 
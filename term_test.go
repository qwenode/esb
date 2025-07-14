package esb

import (
	"testing"
)

// TestTerm 测试基本的Term查询功能
func TestTerm(t *testing.T) {
	t.Run("测试字符串值的Term查询", func(t *testing.T) {
		query := NewQuery(Term("status", "published"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Term == nil {
			t.Fatal("Term查询不应该为nil")
		}
		
		// 验证字段存在
		termQuery, exists := query.Term["status"]
		if !exists {
			t.Fatal("应该存在status字段的term查询")
		}
		
		// 验证值
		if termQuery.Value == nil {
			t.Fatal("Term查询的Value不应该为nil")
		}
		
		if strVal, ok := termQuery.Value.(string); !ok || strVal != "published" {
			t.Errorf("期望Value为'published', 实际得到: %v", termQuery.Value)
		}
	})
	
	t.Run("测试其他字符串值的Term查询", func(t *testing.T) {
		query := NewQuery(Term("category", "electronics"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Term == nil {
			t.Fatal("Term查询不应该为nil")
		}
		
		// 验证字段存在
		termQuery, exists := query.Term["category"]
		if !exists {
			t.Fatal("应该存在category字段的term查询")
		}
		
		// 验证值
		if termQuery.Value == nil {
			t.Fatal("Term查询的Value不应该为nil")
		}
		
		if strVal, ok := termQuery.Value.(string); !ok || strVal != "electronics" {
			t.Errorf("期望Value为'electronics', 实际得到: %v", termQuery.Value)
		}
	})
	
	t.Run("测试空字符串值的Term查询", func(t *testing.T) {
		query := NewQuery(Term("empty_field", ""))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Term == nil {
			t.Fatal("Term查询不应该为nil")
		}
		
		// 验证字段存在
		termQuery, exists := query.Term["empty_field"]
		if !exists {
			t.Fatal("应该存在empty_field字段的term查询")
		}
		
		// 验证值
		if termQuery.Value == nil {
			t.Fatal("Term查询的Value不应该为nil")
		}
		
		if strVal, ok := termQuery.Value.(string); !ok || strVal != "" {
			t.Errorf("期望Value为空字符串, 实际得到: %v", termQuery.Value)
		}
	})
}

// TestTerms 测试Terms查询功能
func TestTerms(t *testing.T) {
	t.Run("测试基本Terms查询", func(t *testing.T) {
		query := NewQuery(Terms("category", "tech", "science", "programming"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Terms == nil {
			t.Fatal("Terms查询不应该为nil")
		}
		
		// 验证字段存在
		_, exists := query.Terms.TermsQuery["category"]
		if !exists {
			t.Fatal("应该存在category字段的terms查询")
		}
		
		// 由于TermsQueryField是interface类型，这里我们只验证查询结构是否正确创建
		// 实际的值验证在集成测试中进行
	})
	
	t.Run("测试单个值的Terms查询", func(t *testing.T) {
		query := NewQuery(Terms("status", "active"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Terms == nil {
			t.Fatal("Terms查询不应该为nil")
		}
		
		// 验证字段存在
		_, exists := query.Terms.TermsQuery["status"]
		if !exists {
			t.Fatal("应该存在status字段的terms查询")
		}
	})
	
	t.Run("测试空值的Terms查询", func(t *testing.T) {
		query := NewQuery(Terms("empty_field"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Terms == nil {
			t.Fatal("Terms查询不应该为nil")
		}
		
		// 验证字段存在
		_, exists := query.Terms.TermsQuery["empty_field"]
		if !exists {
			t.Fatal("应该存在empty_field字段的terms查询")
		}
	})
} 

// TestTermsSlice 测试TermsSlice查询功能
func TestTermsSlice(t *testing.T) {
	t.Run("测试基本TermsSlice查询", func(t *testing.T) {
		values := []string{"tech", "science", "programming"}
		query := NewQuery(TermsSlice("category", values))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Terms == nil {
			t.Fatal("Terms查询不应该为nil")
		}
		
		// 验证字段存在
		_, exists := query.Terms.TermsQuery["category"]
		if !exists {
			t.Fatal("应该存在category字段的terms查询")
		}
	})
	
	t.Run("测试单个值的TermsSlice查询", func(t *testing.T) {
		values := []string{"active"}
		query := NewQuery(TermsSlice("status", values))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Terms == nil {
			t.Fatal("Terms查询不应该为nil")
		}
		
		// 验证字段存在
		_, exists := query.Terms.TermsQuery["status"]
		if !exists {
			t.Fatal("应该存在status字段的terms查询")
		}
	})
	
	t.Run("测试空切片的TermsSlice查询", func(t *testing.T) {
		var values []string
		query := NewQuery(TermsSlice("empty_field", values))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Terms == nil {
			t.Fatal("Terms查询不应该为nil")
		}
		
		// 验证字段存在
		_, exists := query.Terms.TermsQuery["empty_field"]
		if !exists {
			t.Fatal("应该存在empty_field字段的terms查询")
		}
	})
	
	t.Run("测试nil切片的TermsSlice查询", func(t *testing.T) {
		var values []string = nil
		query := NewQuery(TermsSlice("nil_field", values))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Terms == nil {
			t.Fatal("Terms查询不应该为nil")
		}
		
		// 验证字段存在
		_, exists := query.Terms.TermsQuery["nil_field"]
		if !exists {
			t.Fatal("应该存在nil_field字段的terms查询")
		}
	})
} 
package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// TestFuzzyWithOptions 测试函数式选项模式的模糊查询
func TestFuzzyWithOptions(t *testing.T) {
	t.Run("测试基本函数式选项配置", func(t *testing.T) {
		// 测试使用函数式选项模式的 FuzzyWithOptions
		query := NewQuery(
			FuzzyWithOptions("content", "elasticsearch", func(opts *types.FuzzyQuery) {
				boost := float32(1.8)
				opts.Boost = &boost
				maxExpansions := 100
				opts.MaxExpansions = &maxExpansions
				prefixLength := 2
				opts.PrefixLength = &prefixLength
				opts.Fuzziness = types.Fuzziness("AUTO")
				transpositions := true
				opts.Transpositions = &transpositions
				rewrite := "top_terms_10"
				opts.Rewrite = &rewrite
				queryName := "test_fuzzy"
				opts.QueryName_ = &queryName
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Fuzzy == nil {
			t.Fatal("Fuzzy查询不应该为nil")
		}
		
		// 验证字段存在
		fuzzyQuery, exists := query.Fuzzy["content"]
		if !exists {
			t.Fatal("应该存在content字段的模糊查询")
		}
		
		// 验证基本值
		if fuzzyQuery.Value != "elasticsearch" {
			t.Errorf("期望Value为'elasticsearch', 实际得到: %v", fuzzyQuery.Value)
		}
		
		// 验证选项配置
		if fuzzyQuery.Boost == nil || *fuzzyQuery.Boost != 1.8 {
			t.Errorf("期望Boost为1.8, 实际得到: %v", fuzzyQuery.Boost)
		}
		
		if fuzzyQuery.MaxExpansions == nil || *fuzzyQuery.MaxExpansions != 100 {
			t.Errorf("期望MaxExpansions为100, 实际得到: %v", fuzzyQuery.MaxExpansions)
		}
		
		if fuzzyQuery.PrefixLength == nil || *fuzzyQuery.PrefixLength != 2 {
			t.Errorf("期望PrefixLength为2, 实际得到: %v", fuzzyQuery.PrefixLength)
		}
		
		if fuzzyQuery.Fuzziness == nil {
			t.Error("期望Fuzziness不为nil")
		} else if fuzzStr, ok := fuzzyQuery.Fuzziness.(string); !ok || fuzzStr != "AUTO" {
			t.Errorf("期望Fuzziness为'AUTO', 实际得到: %v", fuzzyQuery.Fuzziness)
		}
		
		if fuzzyQuery.Transpositions == nil || *fuzzyQuery.Transpositions != true {
			t.Errorf("期望Transpositions为true, 实际得到: %v", fuzzyQuery.Transpositions)
		}
		
		if fuzzyQuery.Rewrite == nil || *fuzzyQuery.Rewrite != "top_terms_10" {
			t.Errorf("期望Rewrite为'top_terms_10', 实际得到: %v", fuzzyQuery.Rewrite)
		}
		
		if fuzzyQuery.QueryName_ == nil || *fuzzyQuery.QueryName_ != "test_fuzzy" {
			t.Errorf("期望QueryName为'test_fuzzy', 实际得到: %v", fuzzyQuery.QueryName_)
		}
	})
	
	t.Run("测试nil选项配置", func(t *testing.T) {
		// 测试传递nil选项配置函数
		query := NewQuery(
			FuzzyWithOptions("title", "search", nil),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Fuzzy == nil {
			t.Fatal("Fuzzy查询不应该为nil")
		}
		
		// 验证字段存在
		fuzzyQuery, exists := query.Fuzzy["title"]
		if !exists {
			t.Fatal("应该存在title字段的模糊查询")
		}
		
		// 验证基本值
		if fuzzyQuery.Value != "search" {
			t.Errorf("期望Value为'search', 实际得到: %v", fuzzyQuery.Value)
		}
		
		// 验证选项为默认值
		if fuzzyQuery.Boost != nil {
			t.Errorf("期望Boost为nil, 实际得到: %v", fuzzyQuery.Boost)
		}
	})
	
	t.Run("测试空选项配置", func(t *testing.T) {
		// 测试传递空的选项配置函数
		query := NewQuery(
			FuzzyWithOptions("description", "text", func(opts *types.FuzzyQuery) {
				// 不设置任何选项
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Fuzzy == nil {
			t.Fatal("Fuzzy查询不应该为nil")
		}
		
		// 验证字段存在
		fuzzyQuery, exists := query.Fuzzy["description"]
		if !exists {
			t.Fatal("应该存在description字段的模糊查询")
		}
		
		// 验证基本值
		if fuzzyQuery.Value != "text" {
			t.Errorf("期望Value为'text', 实际得到: %v", fuzzyQuery.Value)
		}
		
		// 验证选项为默认值
		if fuzzyQuery.Boost != nil {
			t.Errorf("期望Boost为nil, 实际得到: %v", fuzzyQuery.Boost)
		}
		if fuzzyQuery.MaxExpansions != nil {
			t.Errorf("期望MaxExpansions为nil, 实际得到: %v", fuzzyQuery.MaxExpansions)
		}
		if fuzzyQuery.PrefixLength != nil {
			t.Errorf("期望PrefixLength为nil, 实际得到: %v", fuzzyQuery.PrefixLength)
		}
		if fuzzyQuery.Fuzziness != nil {
			t.Errorf("期望Fuzziness为nil, 实际得到: %v", fuzzyQuery.Fuzziness)
		}
		if fuzzyQuery.Transpositions != nil {
			t.Errorf("期望Transpositions为nil, 实际得到: %v", fuzzyQuery.Transpositions)
		}
	})
} 

// TestFuzzy 测试Fuzzy查询功能
func TestFuzzy(t *testing.T) {
	t.Run("测试基本Fuzzy查询", func(t *testing.T) {
		query := NewQuery(Fuzzy("username", "john"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Fuzzy == nil {
			t.Fatal("Fuzzy查询不应该为nil")
		}
		
		if len(query.Fuzzy) != 1 {
			t.Errorf("期望Fuzzy查询长度为1, 实际得到: %d", len(query.Fuzzy))
		}
		
		if fuzzyQuery, exists := query.Fuzzy["username"]; !exists {
			t.Error("期望存在username字段")
		} else if fuzzyQuery.Value != "john" {
			t.Errorf("期望Value为'john', 实际得到: %s", fuzzyQuery.Value)
		}
	})
	
	t.Run("测试空值的Fuzzy查询", func(t *testing.T) {
		query := NewQuery(Fuzzy("username", ""))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Fuzzy == nil {
			t.Fatal("Fuzzy查询不应该为nil")
		}
		
		if fuzzyQuery, exists := query.Fuzzy["username"]; !exists {
			t.Error("期望存在username字段")
		} else if fuzzyQuery.Value != "" {
			t.Errorf("期望Value为空字符串, 实际得到: %s", fuzzyQuery.Value)
		}
	})
	
	t.Run("测试特殊字符的Fuzzy查询", func(t *testing.T) {
		query := NewQuery(Fuzzy("username", "john.doe@example.com"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Fuzzy == nil {
			t.Fatal("Fuzzy查询不应该为nil")
		}
		
		if fuzzyQuery, exists := query.Fuzzy["username"]; !exists {
			t.Error("期望存在username字段")
		} else if fuzzyQuery.Value != "john.doe@example.com" {
			t.Errorf("期望Value为'john.doe@example.com', 实际得到: %s", fuzzyQuery.Value)
		}
	})
	
	t.Run("测试Unicode字符的Fuzzy查询", func(t *testing.T) {
		query := NewQuery(Fuzzy("username", "张三"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Fuzzy == nil {
			t.Fatal("Fuzzy查询不应该为nil")
		}
		
		if fuzzyQuery, exists := query.Fuzzy["username"]; !exists {
			t.Error("期望存在username字段")
		} else if fuzzyQuery.Value != "张三" {
			t.Errorf("期望Value为'张三', 实际得到: %s", fuzzyQuery.Value)
		}
	})
}

// TestFuzzyWithOptionsComplete 测试FuzzyWithOptions的完整功能
func TestFuzzyWithOptionsComplete(t *testing.T) {
	t.Run("测试所有选项的Fuzzy查询", func(t *testing.T) {
		fuzziness := types.Fuzziness("AUTO")
		prefixLength := 2
		maxExpansions := 50
		transpositions := true
		rewrite := "constant_score"
		
		query := NewQuery(
			FuzzyWithOptions("username", "john", func(q *types.FuzzyQuery) {
				q.Fuzziness = fuzziness
				q.PrefixLength = &prefixLength
				q.MaxExpansions = &maxExpansions
				q.Transpositions = &transpositions
				q.Rewrite = &rewrite
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Fuzzy == nil {
			t.Fatal("Fuzzy查询不应该为nil")
		}
		
		fuzzyQuery, exists := query.Fuzzy["username"]
		if !exists {
			t.Fatal("期望存在username字段")
		}
		
		if fuzzyQuery.Fuzziness != "AUTO" {
			t.Errorf("期望Fuzziness为'AUTO', 实际得到: %s", fuzzyQuery.Fuzziness)
		}
		
		if *fuzzyQuery.PrefixLength != 2 {
			t.Errorf("期望PrefixLength为2, 实际得到: %d", *fuzzyQuery.PrefixLength)
		}
		
		if *fuzzyQuery.MaxExpansions != 50 {
			t.Errorf("期望MaxExpansions为50, 实际得到: %d", *fuzzyQuery.MaxExpansions)
		}
		
		if !*fuzzyQuery.Transpositions {
			t.Error("期望Transpositions为true")
		}
		
		if *fuzzyQuery.Rewrite != "constant_score" {
			t.Errorf("期望Rewrite为'constant_score', 实际得到: %s", *fuzzyQuery.Rewrite)
		}
	})
	
	t.Run("测试部分选项的Fuzzy查询", func(t *testing.T) {
		fuzziness := types.Fuzziness("2")
		
		query := NewQuery(
			FuzzyWithOptions("username", "john", func(q *types.FuzzyQuery) {
				q.Fuzziness = fuzziness
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Fuzzy == nil {
			t.Fatal("Fuzzy查询不应该为nil")
		}
		
		fuzzyQuery, exists := query.Fuzzy["username"]
		if !exists {
			t.Fatal("期望存在username字段")
		}
		
		if fuzzyQuery.Fuzziness != "2" {
			t.Errorf("期望Fuzziness为'2', 实际得到: %s", fuzzyQuery.Fuzziness)
		}
		
		if fuzzyQuery.PrefixLength != nil {
			t.Error("期望PrefixLength为nil")
		}
		
		if fuzzyQuery.MaxExpansions != nil {
			t.Error("期望MaxExpansions为nil")
		}
		
		if fuzzyQuery.Transpositions != nil {
			t.Error("期望Transpositions为nil")
		}
		
		if fuzzyQuery.Rewrite != nil {
			t.Error("期望Rewrite为nil")
		}
	})
} 
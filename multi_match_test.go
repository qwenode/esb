package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/textquerytype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// TestMultiMatchSlice 测试MultiMatchSlice查询功能
func TestMultiMatchSlice(t *testing.T) {
	t.Run("测试基本MultiMatchSlice查询", func(t *testing.T) {
		fields := []string{"title", "description"}
		query := NewQuery(MultiMatchSlice("test query", fields))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Query != "test query" {
			t.Errorf("期望Query为'test query', 实际得到: %s", query.MultiMatch.Query)
		}
		
		if len(query.MultiMatch.Fields) != 2 {
			t.Errorf("期望Fields长度为2, 实际得到: %d", len(query.MultiMatch.Fields))
		}
		
		for i, field := range []string{"title", "description"} {
			if query.MultiMatch.Fields[i] != field {
				t.Errorf("期望Fields[%d]为'%s', 实际得到: %s", i, field, query.MultiMatch.Fields[i])
			}
		}
	})
	
	t.Run("测试空切片的MultiMatchSlice查询", func(t *testing.T) {
		var fields []string
		query := NewQuery(MultiMatchSlice("test query", fields))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if len(query.MultiMatch.Fields) != 0 {
			t.Errorf("期望Fields长度为0, 实际得到: %d", len(query.MultiMatch.Fields))
		}
	})
}

// TestMultiMatchBestFieldsSlice 测试MultiMatchBestFieldsSlice查询功能
func TestMultiMatchBestFieldsSlice(t *testing.T) {
	t.Run("测试基本MultiMatchBestFieldsSlice查询", func(t *testing.T) {
		fields := []string{"title", "description"}
		query := NewQuery(MultiMatchBestFieldsSlice("test query", fields))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Query != "test query" {
			t.Errorf("期望Query为'test query', 实际得到: %s", query.MultiMatch.Query)
		}
		
		if query.MultiMatch.Type == nil {
			t.Fatal("Type不应该为nil")
		}
		
		if *query.MultiMatch.Type != textquerytype.Bestfields {
			t.Errorf("期望Type为best_fields, 实际得到: %s", *query.MultiMatch.Type)
		}
	})
}

// TestMultiMatchMostFieldsSlice 测试MultiMatchMostFieldsSlice查询功能
func TestMultiMatchMostFieldsSlice(t *testing.T) {
	t.Run("测试基本MultiMatchMostFieldsSlice查询", func(t *testing.T) {
		fields := []string{"title", "description"}
		query := NewQuery(MultiMatchMostFieldsSlice("test query", fields))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Type == nil {
			t.Fatal("Type不应该为nil")
		}
		
		if *query.MultiMatch.Type != textquerytype.Mostfields {
			t.Errorf("期望Type为most_fields, 实际得到: %s", *query.MultiMatch.Type)
		}
	})
}

// TestMultiMatchCrossFieldsSlice 测试MultiMatchCrossFieldsSlice查询功能
func TestMultiMatchCrossFieldsSlice(t *testing.T) {
	t.Run("测试基本MultiMatchCrossFieldsSlice查询", func(t *testing.T) {
		fields := []string{"title", "description"}
		query := NewQuery(MultiMatchCrossFieldsSlice("test query", fields))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Type == nil {
			t.Fatal("Type不应该为nil")
		}
		
		if *query.MultiMatch.Type != textquerytype.Crossfields {
			t.Errorf("期望Type为cross_fields, 实际得到: %s", *query.MultiMatch.Type)
		}
	})
}

// TestMultiMatchPhraseSlice 测试MultiMatchPhraseSlice查询功能
func TestMultiMatchPhraseSlice(t *testing.T) {
	t.Run("测试基本MultiMatchPhraseSlice查询", func(t *testing.T) {
		fields := []string{"title", "description"}
		query := NewQuery(MultiMatchPhraseSlice("test query", fields))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Type == nil {
			t.Fatal("Type不应该为nil")
		}
		
		if *query.MultiMatch.Type != textquerytype.Phrase {
			t.Errorf("期望Type为phrase, 实际得到: %s", *query.MultiMatch.Type)
		}
	})
}

// TestMultiMatchPhrasePrefixSlice 测试MultiMatchPhrasePrefixSlice查询功能
func TestMultiMatchPhrasePrefixSlice(t *testing.T) {
	t.Run("测试基本MultiMatchPhrasePrefixSlice查询", func(t *testing.T) {
		fields := []string{"title", "description"}
		query := NewQuery(MultiMatchPhrasePrefixSlice("test query", fields))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Type == nil {
			t.Fatal("Type不应该为nil")
		}
		
		if *query.MultiMatch.Type != textquerytype.Phraseprefix {
			t.Errorf("期望Type为phrase_prefix, 实际得到: %s", *query.MultiMatch.Type)
		}
	})
} 

// TestMultiMatch 测试MultiMatch查询功能
func TestMultiMatch(t *testing.T) {
	t.Run("测试基本MultiMatch查询", func(t *testing.T) {
		query := NewQuery(MultiMatch("test query", "title", "description"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Query != "test query" {
			t.Errorf("期望Query为'test query', 实际得到: %s", query.MultiMatch.Query)
		}
		
		if len(query.MultiMatch.Fields) != 2 {
			t.Errorf("期望Fields长度为2, 实际得到: %d", len(query.MultiMatch.Fields))
		}
		
		for i, field := range []string{"title", "description"} {
			if query.MultiMatch.Fields[i] != field {
				t.Errorf("期望Fields[%d]为'%s', 实际得到: %s", i, field, query.MultiMatch.Fields[i])
			}
		}
	})
	
	t.Run("测试无字段的MultiMatch查询", func(t *testing.T) {
		query := NewQuery(MultiMatch("test query"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if len(query.MultiMatch.Fields) != 0 {
			t.Errorf("期望Fields长度为0, 实际得到: %d", len(query.MultiMatch.Fields))
		}
	})
}

// TestMultiMatchBestFields 测试MultiMatchBestFields查询功能
func TestMultiMatchBestFields(t *testing.T) {
	t.Run("测试基本MultiMatchBestFields查询", func(t *testing.T) {
		query := NewQuery(MultiMatchBestFields("test query", "title", "description"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Query != "test query" {
			t.Errorf("期望Query为'test query', 实际得到: %s", query.MultiMatch.Query)
		}
		
		if query.MultiMatch.Type == nil {
			t.Fatal("Type不应该为nil")
		}
		
		if *query.MultiMatch.Type != textquerytype.Bestfields {
			t.Errorf("期望Type为best_fields, 实际得到: %s", *query.MultiMatch.Type)
		}
	})
}

// TestMultiMatchMostFields 测试MultiMatchMostFields查询功能
func TestMultiMatchMostFields(t *testing.T) {
	t.Run("测试基本MultiMatchMostFields查询", func(t *testing.T) {
		query := NewQuery(MultiMatchMostFields("test query", "title", "description"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Type == nil {
			t.Fatal("Type不应该为nil")
		}
		
		if *query.MultiMatch.Type != textquerytype.Mostfields {
			t.Errorf("期望Type为most_fields, 实际得到: %s", *query.MultiMatch.Type)
		}
	})
} 

// TestMultiMatchCrossFields 测试MultiMatchCrossFields查询功能
func TestMultiMatchCrossFields(t *testing.T) {
	t.Run("测试基本MultiMatchCrossFields查询", func(t *testing.T) {
		query := NewQuery(MultiMatchCrossFields("test query", "title", "description"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Type == nil {
			t.Fatal("Type不应该为nil")
		}
		
		if *query.MultiMatch.Type != textquerytype.Crossfields {
			t.Errorf("期望Type为cross_fields, 实际得到: %s", *query.MultiMatch.Type)
		}
	})
}

// TestMultiMatchPhrase 测试MultiMatchPhrase查询功能
func TestMultiMatchPhrase(t *testing.T) {
	t.Run("测试基本MultiMatchPhrase查询", func(t *testing.T) {
		query := NewQuery(MultiMatchPhrase("test query", "title", "description"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Type == nil {
			t.Fatal("Type不应该为nil")
		}
		
		if *query.MultiMatch.Type != textquerytype.Phrase {
			t.Errorf("期望Type为phrase, 实际得到: %s", *query.MultiMatch.Type)
		}
	})
}

// TestMultiMatchPhrasePrefix 测试MultiMatchPhrasePrefix查询功能
func TestMultiMatchPhrasePrefix(t *testing.T) {
	t.Run("测试基本MultiMatchPhrasePrefix查询", func(t *testing.T) {
		query := NewQuery(MultiMatchPhrasePrefix("test query", "title", "description"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Type == nil {
			t.Fatal("Type不应该为nil")
		}
		
		if *query.MultiMatch.Type != textquerytype.Phraseprefix {
			t.Errorf("期望Type为phrase_prefix, 实际得到: %s", *query.MultiMatch.Type)
		}
	})
} 

// TestMultiMatchWithOptionsExtended 测试MultiMatchWithOptions的扩展功能
func TestMultiMatchWithOptionsExtended(t *testing.T) {
	t.Run("测试带分析器和最小匹配数的MultiMatch查询", func(t *testing.T) {
		analyzer := "standard"
		minShouldMatch := "75%"
		op := operator.And
		query := NewQuery(
			MultiMatchWithOptions("test query", []string{"title", "description"}, func(q *types.MultiMatchQuery) {
				q.Analyzer = &analyzer
				q.MinimumShouldMatch = minShouldMatch
				q.Operator = &op
			}),
		)
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Analyzer == nil || *query.MultiMatch.Analyzer != "standard" {
			t.Error("期望分析器为'standard'")
		}
		
		if query.MultiMatch.MinimumShouldMatch != "75%" {
			t.Errorf("期望最小匹配数为'75%%', 实际得到: %s", query.MultiMatch.MinimumShouldMatch)
		}
		
		if query.MultiMatch.Operator == nil || *query.MultiMatch.Operator != operator.And {
			t.Error("期望操作符为'AND'")
		}
	})
	
	t.Run("测试带模糊匹配和切片的MultiMatch查询", func(t *testing.T) {
		fuzziness := types.Fuzziness("AUTO")
		prefixLength := 2
		maxExpansions := 50
		query := NewQuery(
			MultiMatchWithOptions("test query", []string{"title", "description"}, func(q *types.MultiMatchQuery) {
				q.Fuzziness = fuzziness
				q.PrefixLength = &prefixLength
				q.MaxExpansions = &maxExpansions
			}),
		)
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Fuzziness != "AUTO" {
			t.Errorf("期望模糊匹配为'AUTO', 实际得到: %s", query.MultiMatch.Fuzziness)
		}
		
		if query.MultiMatch.PrefixLength == nil || *query.MultiMatch.PrefixLength != 2 {
			t.Error("期望前缀长度为2")
		}
		
		if query.MultiMatch.MaxExpansions == nil || *query.MultiMatch.MaxExpansions != 50 {
			t.Error("期望最大扩展数为50")
		}
	})
	
	t.Run("测试带零项查询和自动生成同义词短语的MultiMatch查询", func(t *testing.T) {
		zeroTerms := zerotermsquery.All
		autoGenerateSynonyms := true
		lenient := true
		query := NewQuery(
			MultiMatchWithOptions("test query", []string{"title", "description"}, func(q *types.MultiMatchQuery) {
				q.ZeroTermsQuery = &zeroTerms
				q.AutoGenerateSynonymsPhraseQuery = &autoGenerateSynonyms
				q.Lenient = &lenient
			}),
		)
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.ZeroTermsQuery == nil || *query.MultiMatch.ZeroTermsQuery != zerotermsquery.All {
			t.Error("期望零项查询为'all'")
		}
		
		if query.MultiMatch.AutoGenerateSynonymsPhraseQuery == nil || !*query.MultiMatch.AutoGenerateSynonymsPhraseQuery {
			t.Error("期望自动生成同义词短语为true")
		}
		
		if query.MultiMatch.Lenient == nil || !*query.MultiMatch.Lenient {
			t.Error("期望lenient为true")
		}
	})
} 

func TestMultiMatchWithOptionsComplete(t *testing.T) {
	t.Run("测试模糊匹配相关选项", func(t *testing.T) {
		fuzzyRewrite := "constant_score"
		fuzzyTranspositions := false
		query := NewQuery(
			MultiMatchWithOptions("test query", []string{"title", "description"}, func(q *types.MultiMatchQuery) {
				q.FuzzyRewrite = &fuzzyRewrite
				q.FuzzyTranspositions = &fuzzyTranspositions
			}),
		)
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.FuzzyRewrite == nil || *query.MultiMatch.FuzzyRewrite != "constant_score" {
			t.Error("期望fuzzy_rewrite为'constant_score'")
		}
		
		if query.MultiMatch.FuzzyTranspositions == nil || *query.MultiMatch.FuzzyTranspositions != false {
			t.Error("期望fuzzy_transpositions为false")
		}
	})
	
	t.Run("测试频率和评分相关选项", func(t *testing.T) {
		cutoffFrequency := types.Float64(0.001)
		tieBreaker := types.Float64(0.3)
		query := NewQuery(
			MultiMatchWithOptions("test query", []string{"title", "description"}, func(q *types.MultiMatchQuery) {
				q.CutoffFrequency = &cutoffFrequency
				q.TieBreaker = &tieBreaker
			}),
		)
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.CutoffFrequency == nil || *query.MultiMatch.CutoffFrequency != 0.001 {
			t.Error("期望cutoff_frequency为0.001")
		}
		
		if query.MultiMatch.TieBreaker == nil || *query.MultiMatch.TieBreaker != 0.3 {
			t.Error("期望tie_breaker为0.3")
		}
	})
	
	t.Run("测试短语匹配和查询名称选项", func(t *testing.T) {
		slop := 2
		queryName := "test_query"
		query := NewQuery(
			MultiMatchWithOptions("test query", []string{"title", "description"}, func(q *types.MultiMatchQuery) {
				q.Slop = &slop
				q.QueryName_ = &queryName
			}),
		)
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Slop == nil || *query.MultiMatch.Slop != 2 {
			t.Error("期望slop为2")
		}
		
		if query.MultiMatch.QueryName_ == nil || *query.MultiMatch.QueryName_ != "test_query" {
			t.Error("期望query_name为'test_query'")
		}
	})
} 

// TestMultiMatchWithOptionsErrorHandling 测试MultiMatchWithOptions的错误处理
func TestMultiMatchWithOptionsErrorHandling(t *testing.T) {
	t.Run("测试无效的模糊匹配值", func(t *testing.T) {
		fuzziness := types.Fuzziness("INVALID")
		query := NewQuery(
			MultiMatchWithOptions("test query", []string{"title"}, func(q *types.MultiMatchQuery) {
				q.Fuzziness = fuzziness
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Fuzziness != "INVALID" {
			t.Errorf("期望Fuzziness为'INVALID', 实际得到: %s", query.MultiMatch.Fuzziness)
		}
	})
	
	t.Run("测试无效的最小匹配数", func(t *testing.T) {
		minShouldMatch := "invalid_value"
		query := NewQuery(
			MultiMatchWithOptions("test query", []string{"title"}, func(q *types.MultiMatchQuery) {
				q.MinimumShouldMatch = minShouldMatch
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.MinimumShouldMatch != "invalid_value" {
			t.Errorf("期望MinimumShouldMatch为'invalid_value', 实际得到: %s", query.MultiMatch.MinimumShouldMatch)
		}
	})
	
	t.Run("测试查询类型", func(t *testing.T) {
		typeVal := textquerytype.Bestfields
		query := NewQuery(
			MultiMatchWithOptions("test query", []string{"title"}, func(q *types.MultiMatchQuery) {
				q.Type = &typeVal
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Type == nil || *query.MultiMatch.Type != textquerytype.Bestfields {
			t.Error("期望Type为best_fields")
		}
	})
	
	t.Run("测试零项查询", func(t *testing.T) {
		zeroTerms := zerotermsquery.None
		query := NewQuery(
			MultiMatchWithOptions("test query", []string{"title"}, func(q *types.MultiMatchQuery) {
				q.ZeroTermsQuery = &zeroTerms
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.ZeroTermsQuery == nil || *query.MultiMatch.ZeroTermsQuery != zerotermsquery.None {
			t.Error("期望ZeroTermsQuery为none")
		}
	})
	
	t.Run("测试操作符", func(t *testing.T) {
		op := operator.And
		query := NewQuery(
			MultiMatchWithOptions("test query", []string{"title"}, func(q *types.MultiMatchQuery) {
				q.Operator = &op
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Operator == nil || *query.MultiMatch.Operator != operator.And {
			t.Error("期望Operator为AND")
		}
	})
}

// TestMultiMatchFieldBoost 测试字段权重功能
func TestMultiMatchFieldBoost(t *testing.T) {
	t.Run("测试字段权重", func(t *testing.T) {
		fields := []string{"title^2", "description"}
		query := NewQuery(MultiMatchSlice("test query", fields))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if len(query.MultiMatch.Fields) != 2 {
			t.Errorf("期望Fields长度为2, 实际得到: %d", len(query.MultiMatch.Fields))
		}
		
		if query.MultiMatch.Fields[0] != "title^2" {
			t.Errorf("期望Fields[0]为'title^2', 实际得到: %s", query.MultiMatch.Fields[0])
		}
		
		if query.MultiMatch.Fields[1] != "description" {
			t.Errorf("期望Fields[1]为'description', 实际得到: %s", query.MultiMatch.Fields[1])
		}
	})
	
	t.Run("测试多个字段权重", func(t *testing.T) {
		fields := []string{"title^3", "description^2", "content"}
		query := NewQuery(MultiMatchSlice("test query", fields))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if len(query.MultiMatch.Fields) != 3 {
			t.Errorf("期望Fields长度为3, 实际得到: %d", len(query.MultiMatch.Fields))
		}
		
		expectedFields := []string{"title^3", "description^2", "content"}
		for i, field := range expectedFields {
			if query.MultiMatch.Fields[i] != field {
				t.Errorf("期望Fields[%d]为'%s', 实际得到: %s", i, field, query.MultiMatch.Fields[i])
			}
		}
	})
}

// TestMultiMatchEdgeCases 测试边界条件
func TestMultiMatchEdgeCases(t *testing.T) {
	t.Run("测试空查询字符串", func(t *testing.T) {
		query := NewQuery(MultiMatch("", "title", "description"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Query != "" {
			t.Errorf("期望Query为空字符串, 实际得到: %s", query.MultiMatch.Query)
		}
	})
	
	t.Run("测试特殊字符字段名", func(t *testing.T) {
		fields := []string{"title.raw", "description.keyword^2"}
		query := NewQuery(MultiMatchSlice("test query", fields))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		expectedFields := []string{"title.raw", "description.keyword^2"}
		for i, field := range expectedFields {
			if query.MultiMatch.Fields[i] != field {
				t.Errorf("期望Fields[%d]为'%s', 实际得到: %s", i, field, query.MultiMatch.Fields[i])
			}
		}
	})
	
	t.Run("测试Unicode字段名和查询", func(t *testing.T) {
		fields := []string{"标题^2", "描述"}
		query := NewQuery(MultiMatchSlice("测试查询", fields))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.MultiMatch == nil {
			t.Fatal("MultiMatch查询不应该为nil")
		}
		
		if query.MultiMatch.Query != "测试查询" {
			t.Errorf("期望Query为'测试查询', 实际得到: %s", query.MultiMatch.Query)
		}
		
		expectedFields := []string{"标题^2", "描述"}
		for i, field := range expectedFields {
			if query.MultiMatch.Fields[i] != field {
				t.Errorf("期望Fields[%d]为'%s', 实际得到: %s", i, field, query.MultiMatch.Fields[i])
			}
		}
	})
} 
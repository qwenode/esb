package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/textquerytype"
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
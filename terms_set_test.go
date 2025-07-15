package esb

import (
	"fmt"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestTermsSet(t *testing.T) {
	query := NewQuery(TermsSet("skills", []string{"java", "golang", "elasticsearch"}))
	
	if query.TermsSet == nil {
		t.Fatal("TermsSet查询不应该为nil")
	}
	
	termsSetQuery, exists := query.TermsSet["skills"]
	if !exists {
		t.Fatal("应该存在skills字段的TermsSet查询")
	}
	
	expectedTerms := []string{"java", "golang", "elasticsearch"}
	if len(termsSetQuery.Terms) != len(expectedTerms) {
		t.Fatalf("期望%d个terms，但得到%d个", len(expectedTerms), len(termsSetQuery.Terms))
	}
	
	for i, term := range expectedTerms {
		if termsSetQuery.Terms[i] != term {
			t.Errorf("期望term[%d]为%s，但得到%s", i, term, termsSetQuery.Terms[i])
		}
	}
}

func TestTermsSetWithOptions_MinimumShouldMatch(t *testing.T) {
	minMatch := 2
	query := NewQuery(TermsSetWithOptions("languages", []string{"go", "python", "java"}, func(q *types.TermsSetQuery) {
		q.MinimumShouldMatch = minMatch
	}))
	
	if query.TermsSet == nil {
		t.Fatal("TermsSet查询不应该为nil")
	}
	
	termsSetQuery, exists := query.TermsSet["languages"]
	if !exists {
		t.Fatal("应该存在languages字段的TermsSet查询")
	}
	
	if termsSetQuery.MinimumShouldMatch == nil {
		t.Fatal("MinimumShouldMatch不应该为nil")
	}
	
	if termsSetQuery.MinimumShouldMatch != minMatch {
		t.Errorf("期望MinimumShouldMatch为%d，但得到%v", minMatch, termsSetQuery.MinimumShouldMatch)
	}
}

func TestTermsSetWithOptions_MinimumShouldMatchField(t *testing.T) {
	minMatchField := "min_required"
	query := NewQuery(TermsSetWithOptions("skills", []string{"elasticsearch", "golang"}, func(q *types.TermsSetQuery) {
		q.MinimumShouldMatchField = &minMatchField
	}))
	
	if query.TermsSet == nil {
		t.Fatal("TermsSet查询不应该为nil")
	}
	
	termsSetQuery, exists := query.TermsSet["skills"]
	if !exists {
		t.Fatal("应该存在skills字段的TermsSet查询")
	}
	
	if termsSetQuery.MinimumShouldMatchField == nil {
		t.Fatal("MinimumShouldMatchField不应该为nil")
	}
	
	if *termsSetQuery.MinimumShouldMatchField != minMatchField {
		t.Errorf("期望MinimumShouldMatchField为%s，但得到%s", minMatchField, *termsSetQuery.MinimumShouldMatchField)
	}
}

func TestTermsSetWithOptions_Boost(t *testing.T) {
	boost := float32(2.5)
	query := NewQuery(TermsSetWithOptions("categories", []string{"books", "electronics"}, func(q *types.TermsSetQuery) {
		q.Boost = &boost
	}))
	
	if query.TermsSet == nil {
		t.Fatal("TermsSet查询不应该为nil")
	}
	
	termsSetQuery, exists := query.TermsSet["categories"]
	if !exists {
		t.Fatal("应该存在categories字段的TermsSet查询")
	}
	
	if termsSetQuery.Boost == nil {
		t.Fatal("Boost不应该为nil")
	}
	
	if *termsSetQuery.Boost != boost {
		t.Errorf("期望Boost为%f，但得到%f", boost, *termsSetQuery.Boost)
	}
}

func TestTermsSetWithOptions_MultipleOptions(t *testing.T) {
	minMatch := 1
	boost := float32(1.5)
	query := NewQuery(TermsSetWithOptions("tags", []string{"tech", "golang", "elasticsearch"}, func(q *types.TermsSetQuery) {
		q.MinimumShouldMatch = minMatch
		q.Boost = &boost
	}))
	
	if query.TermsSet == nil {
		t.Fatal("TermsSet查询不应该为nil")
	}
	
	termsSetQuery, exists := query.TermsSet["tags"]
	if !exists {
		t.Fatal("应该存在tags字段的TermsSet查询")
	}
	
	if termsSetQuery.MinimumShouldMatch == nil {
		t.Fatal("MinimumShouldMatch不应该为nil")
	}
	
	if termsSetQuery.MinimumShouldMatch != minMatch {
		t.Errorf("期望MinimumShouldMatch为%d，但得到%v", minMatch, termsSetQuery.MinimumShouldMatch)
	}
	
	if termsSetQuery.Boost == nil {
		t.Fatal("Boost不应该为nil")
	}
	
	if *termsSetQuery.Boost != boost {
		t.Errorf("期望Boost为%f，但得到%f", boost, *termsSetQuery.Boost)
	}
}

func TestTermsSet_EmptyValues(t *testing.T) {
	query := NewQuery(TermsSet("field", []string{}))
	
	if query.TermsSet == nil {
		t.Fatal("TermsSet查询不应该为nil")
	}
	
	termsSetQuery, exists := query.TermsSet["field"]
	if !exists {
		t.Fatal("应该存在field字段的TermsSet查询")
	}
	
	if len(termsSetQuery.Terms) != 0 {
		t.Errorf("期望空的terms数组，但得到%d个元素", len(termsSetQuery.Terms))
	}
}

func TestTermsSet_SingleValue(t *testing.T) {
	query := NewQuery(TermsSet("status", []string{"active"}))
	
	if query.TermsSet == nil {
		t.Fatal("TermsSet查询不应该为nil")
	}
	
	termsSetQuery, exists := query.TermsSet["status"]
	if !exists {
		t.Fatal("应该存在status字段的TermsSet查询")
	}
	
	if len(termsSetQuery.Terms) != 1 {
		t.Fatalf("期望1个term，但得到%d个", len(termsSetQuery.Terms))
	}
	
	if termsSetQuery.Terms[0] != "active" {
		t.Errorf("期望term为'active'，但得到'%s'", termsSetQuery.Terms[0])
	}
}

func BenchmarkTermsSet(b *testing.B) {
	values := []string{"tech", "blog", "go", "elasticsearch", "programming"}
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_ = NewQuery(TermsSet("tags", values))
	}
}

func BenchmarkTermsSetWithOptions(b *testing.B) {
	values := []string{"java", "golang", "python", "javascript", "rust"}
	minMatch := 2
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_ = NewQuery(TermsSetWithOptions("skills", values, func(q *types.TermsSetQuery) {
			q.MinimumShouldMatch = minMatch
		}))
	}
}

// 测试字符串类型的MinimumShouldMatch
func TestTermsSetWithOptions_MinimumShouldMatchString(t *testing.T) {
	minMatch := "50%"
	query := NewQuery(TermsSetWithOptions("skills", []string{"java", "golang", "python"}, func(q *types.TermsSetQuery) {
		q.MinimumShouldMatch = minMatch
	}))
	
	if query.TermsSet == nil {
		t.Fatal("TermsSet查询不应该为nil")
	}
	
	termsSetQuery, exists := query.TermsSet["skills"]
	if !exists {
		t.Fatal("应该存在skills字段的TermsSet查询")
	}
	
	if termsSetQuery.MinimumShouldMatch == nil {
		t.Fatal("MinimumShouldMatch不应该为nil")
	}
	
	if termsSetQuery.MinimumShouldMatch != minMatch {
		t.Errorf("期望MinimumShouldMatch为%s，但得到%v", minMatch, termsSetQuery.MinimumShouldMatch)
	}
}

// 测试nil选项函数
func TestTermsSetWithOptions_NilOptions(t *testing.T) {
	query := NewQuery(TermsSetWithOptions("tags", []string{"tech", "golang"}, nil))
	
	if query.TermsSet == nil {
		t.Fatal("TermsSet查询不应该为nil")
	}
	
	termsSetQuery, exists := query.TermsSet["tags"]
	if !exists {
		t.Fatal("应该存在tags字段的TermsSet查询")
	}
	
	expectedTerms := []string{"tech", "golang"}
	if len(termsSetQuery.Terms) != len(expectedTerms) {
		t.Fatalf("期望%d个terms，但得到%d个", len(expectedTerms), len(termsSetQuery.Terms))
	}
}

// 测试大量terms的情况
func TestTermsSet_LargeTermsList(t *testing.T) {
	largeTerms := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		largeTerms[i] = fmt.Sprintf("term_%d", i)
	}
	
	query := NewQuery(TermsSet("large_field", largeTerms))
	
	if query.TermsSet == nil {
		t.Fatal("TermsSet查询不应该为nil")
	}
	
	termsSetQuery, exists := query.TermsSet["large_field"]
	if !exists {
		t.Fatal("应该存在large_field字段的TermsSet查询")
	}
	
	if len(termsSetQuery.Terms) != 1000 {
		t.Errorf("期望1000个terms，但得到%d个", len(termsSetQuery.Terms))
	}
}

// 测试特殊字符的字段名和值
func TestTermsSet_SpecialCharacters(t *testing.T) {
	specialField := "field.with-special_chars@domain"
	specialValues := []string{"value with spaces", "value-with-dashes", "value_with_underscores", "value@domain.com"}
	
	query := NewQuery(TermsSet(specialField, specialValues))
	
	if query.TermsSet == nil {
		t.Fatal("TermsSet查询不应该为nil")
	}
	
	termsSetQuery, exists := query.TermsSet[specialField]
	if !exists {
		t.Fatalf("应该存在%s字段的TermsSet查询", specialField)
	}
	
	if len(termsSetQuery.Terms) != len(specialValues) {
		t.Fatalf("期望%d个terms，但得到%d个", len(specialValues), len(termsSetQuery.Terms))
	}
	
	for i, expectedValue := range specialValues {
		if termsSetQuery.Terms[i] != expectedValue {
			t.Errorf("期望term[%d]为%s，但得到%s", i, expectedValue, termsSetQuery.Terms[i])
		}
	}
}
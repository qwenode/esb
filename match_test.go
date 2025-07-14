package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/textquerytype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"
)

func TestMatch(t *testing.T) {
	t.Run("应该创建基本的匹配查询", func(t *testing.T) {
		query := NewQuery(Match("title", "elasticsearch search"))
		if query == nil {
			t.Error("预期查询不为 nil")
		}
		if query.Match == nil {
			t.Error("预期为 Match 查询")
		}
		if query.Match["title"].Query != "elasticsearch search" {
			t.Errorf("预期查询为 'elasticsearch search'，得到 %s", query.Match["title"].Query)
		}
	})

	t.Run("应该支持布尔查询中的多个匹配查询", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Must(
					Match("title", "elasticsearch"),
					Match("content", "search engine"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
		if len(query.Bool.Must) != 2 {
			t.Errorf("预期有 2 个必须查询，得到 %d", len(query.Bool.Must))
		}
		// 检查第一个匹配查询
		if query.Bool.Must[0].Match == nil {
			t.Error("预期第一个查询为 Match")
		}
		// 检查第二个匹配查询
		if query.Bool.Must[1].Match == nil {
			t.Error("预期第二个查询为 Match")
		}
	})
}

func TestMatchWithOptions(t *testing.T) {
	t.Run("应该创建带操作符的匹配查询", func(t *testing.T) {
		op := operator.And
		query := NewQuery(
			MatchWithOptions("title", "elasticsearch search", func(opts *types.MatchQuery) {
				opts.Operator = &op
			}),
		)
		if query.Match == nil {
			t.Error("预期为 Match 查询")
		}
		if query.Match["title"].Operator == nil {
			t.Error("预期操作符已设置")
		}
		if *query.Match["title"].Operator != operator.And {
			t.Errorf("预期操作符为 AND，得到 %v", *query.Match["title"].Operator)
		}
	})

	t.Run("应该创建带模糊匹配的匹配查询", func(t *testing.T) {
		fuzziness := types.Fuzziness("AUTO")
		query := NewQuery(
			MatchWithOptions("title", "elasticsearch", func(opts *types.MatchQuery) {
				opts.Fuzziness = fuzziness
			}),
		)
		if query.Match["title"].Fuzziness == nil {
			t.Error("预期模糊匹配已设置")
		}
		if query.Match["title"].Fuzziness != types.Fuzziness("AUTO") {
			t.Errorf("预期模糊匹配为 AUTO，得到 %v", query.Match["title"].Fuzziness)
		}
	})

	t.Run("应该创建带多个选项的匹配查询", func(t *testing.T) {
		op := operator.And
		analyzer := "standard"
		boost := float32(1.5)
		lenient := true
		
		query := NewQuery(
			MatchWithOptions("title", "elasticsearch search", func(opts *types.MatchQuery) {
				opts.Operator = &op
				opts.Analyzer = &analyzer
				opts.Boost = &boost
				opts.Lenient = &lenient
			}),
		)
		
		matchQuery := query.Match["title"]
		if matchQuery.Operator == nil || *matchQuery.Operator != operator.And {
			t.Error("预期操作符为 AND")
		}
		if matchQuery.Analyzer == nil || *matchQuery.Analyzer != "standard" {
			t.Error("预期分析器为 'standard'")
		}
		if matchQuery.Boost == nil || *matchQuery.Boost != 1.5 {
			t.Error("预期权重为 1.5")
		}
		if matchQuery.Lenient == nil || *matchQuery.Lenient != true {
			t.Error("预期宽松模式为 true")
		}
	})

	t.Run("应该可以处理空选项", func(t *testing.T) {
		query := NewQuery(
			MatchWithOptions("title", "elasticsearch", nil),
		)
		if query.Match == nil {
			t.Error("预期为 Match 查询")
		}
		if query.Match["title"].Query != "elasticsearch" {
			t.Errorf("预期查询为 'elasticsearch'，得到 %s", query.Match["title"].Query)
		}
	})
}

func TestMatchWith(t *testing.T) {
	t.Run("应该创建带回调选项的匹配查询", func(t *testing.T) {
		analyzer := "standard"
		boost := float32(1.5)
		fuzziness := types.Fuzziness("AUTO")
		query := NewQuery(
			MatchWithOptions("title", "elasticsearch guide", func(q *types.MatchQuery) {
				q.Fuzziness = fuzziness
				q.Analyzer = &analyzer
				q.Boost = &boost
			}),
		)
		if query.Match == nil {
			t.Error("预期为 Match 查询")
		}
		matchQuery := query.Match["title"]
		if matchQuery.Fuzziness != fuzziness {
			t.Errorf("预期模糊匹配 %v，得到 %v", fuzziness, matchQuery.Fuzziness)
		}
		if matchQuery.Analyzer == nil || *matchQuery.Analyzer != "standard" {
			t.Error("预期分析器为 'standard'")
		}
		if matchQuery.Boost == nil || *matchQuery.Boost != 1.5 {
			t.Error("预期权重为 1.5")
		}
	})

	t.Run("应该可以处理 nil 回调", func(t *testing.T) {
		query := NewQuery(MatchWithOptions("title", "elasticsearch", nil))
		if query.Match == nil {
			t.Error("预期为 Match 查询")
		}
		if query.Match["title"].Query != "elasticsearch" {
			t.Errorf("预期查询为 'elasticsearch'，得到 %s", query.Match["title"].Query)
		}
	})
}

func TestMatchPhrase(t *testing.T) {
	t.Run("应该创建短语匹配查询", func(t *testing.T) {
		query := NewQuery(MatchPhrase("content", "elasticsearch is awesome"))
		if query == nil {
			t.Error("预期查询不为 nil")
		}
		if query.MatchPhrase == nil {
			t.Error("预期为 MatchPhrase 查询")
		}
		if query.MatchPhrase["content"].Query != "elasticsearch is awesome" {
			t.Errorf("预期查询为 'elasticsearch is awesome'，得到 %s", query.MatchPhrase["content"].Query)
		}
	})


}

func TestMatchPhraseWithOptions(t *testing.T) {
	t.Run("应该创建带 slop 的短语匹配查询", func(t *testing.T) {
		slop := 2
		query := NewQuery(
			MatchPhraseWithOptions("content", "elasticsearch search", func(opts *types.MatchPhraseQuery) {
				opts.Slop = &slop
			}),
		)
		if query.MatchPhrase == nil {
			t.Error("预期为 MatchPhrase 查询")
		}
		if query.MatchPhrase["content"].Slop == nil {
			t.Error("预期 slop 已设置")
		}
		if *query.MatchPhrase["content"].Slop != 2 {
			t.Errorf("预期 slop 为 2，得到 %d", *query.MatchPhrase["content"].Slop)
		}
	})

	t.Run("应该创建带分析器和权重的短语匹配查询", func(t *testing.T) {
		analyzer := "keyword"
		boost := float32(2.0)
		
		query := NewQuery(
			MatchPhraseWithOptions("content", "exact phrase", func(opts *types.MatchPhraseQuery) {
				opts.Analyzer = &analyzer
				opts.Boost = &boost
			}),
		)
		
		matchPhraseQuery := query.MatchPhrase["content"]
		if matchPhraseQuery.Analyzer == nil || *matchPhraseQuery.Analyzer != "keyword" {
			t.Error("预期分析器为 'keyword'")
		}
		if matchPhraseQuery.Boost == nil || *matchPhraseQuery.Boost != 2.0 {
			t.Error("预期权重为 2.0")
		}
	})
}

func TestMatchPhraseWith(t *testing.T) {
	t.Run("应该创建带回调选项的短语匹配查询", func(t *testing.T) {
		analyzer := "keyword"
		boost := float32(2.0)
		slop := 2
		query := NewQuery(
			MatchPhraseWithOptions("content", "elasticsearch is awesome", func(q *types.MatchPhraseQuery) {
				q.Slop = &slop
				q.Analyzer = &analyzer
				q.Boost = &boost
			}),
		)
		if query.MatchPhrase == nil {
			t.Error("预期为 MatchPhrase 查询")
		}
		matchPhraseQuery := query.MatchPhrase["content"]
		if matchPhraseQuery.Slop == nil || *matchPhraseQuery.Slop != 2 {
			t.Error("预期 slop 为 2")
		}
		if matchPhraseQuery.Analyzer == nil || *matchPhraseQuery.Analyzer != "keyword" {
			t.Error("预期分析器为 'keyword'")
		}
		if matchPhraseQuery.Boost == nil || *matchPhraseQuery.Boost != 2.0 {
			t.Error("预期权重为 2.0")
		}
	})

	t.Run("应该可以处理 nil 回调", func(t *testing.T) {
		query := NewQuery(MatchPhraseWithOptions("content", "elasticsearch is awesome", nil))
		if query.MatchPhrase == nil {
			t.Error("预期为 MatchPhrase 查询")
		}
		if query.MatchPhrase["content"].Query != "elasticsearch is awesome" {
			t.Errorf("预期查询为 'elasticsearch is awesome'，得到 %s", query.MatchPhrase["content"].Query)
		}
	})
}

func TestMatchPhrasePrefix(t *testing.T) {
	t.Run("应该创建短语前缀匹配查询", func(t *testing.T) {
		query := NewQuery(MatchPhrasePrefix("title", "elasticsearch sea"))
		if query == nil {
			t.Error("预期查询不为 nil")
		}
		if query.MatchPhrasePrefix == nil {
			t.Error("预期为 MatchPhrasePrefix 查询")
		}
		if query.MatchPhrasePrefix["title"].Query != "elasticsearch sea" {
			t.Errorf("预期查询为 'elasticsearch sea'，得到 %s", query.MatchPhrasePrefix["title"].Query)
		}
	})


}

func TestMatchPhrasePrefixWithOptions(t *testing.T) {
	t.Run("应该创建带回调选项的短语前缀匹配查询", func(t *testing.T) {
		analyzer := "standard"
		boost := float32(2.0)
		slop := 1
		maxExpansions := 10
		zeroTermsQueryVal := zerotermsquery.None
		query := NewQuery(
			MatchPhrasePrefixWithOptions("title", "quick brown f", func(opts *types.MatchPhrasePrefixQuery) {
				opts.Analyzer = &analyzer
				opts.Boost = &boost
				opts.Slop = &slop
				opts.MaxExpansions = &maxExpansions
				opts.ZeroTermsQuery = &zeroTermsQueryVal
			}),
		)
		if query.MatchPhrasePrefix == nil {
			t.Error("预期为 MatchPhrasePrefix 查询")
		}
		matchPhrasePrefixQuery := query.MatchPhrasePrefix["title"]
		if matchPhrasePrefixQuery.Analyzer == nil || *matchPhrasePrefixQuery.Analyzer != "standard" {
			t.Error("预期分析器为 'standard'")
		}
		if matchPhrasePrefixQuery.Boost == nil || *matchPhrasePrefixQuery.Boost != 2.0 {
			t.Error("预期权重为 2.0")
		}
		if matchPhrasePrefixQuery.Slop == nil || *matchPhrasePrefixQuery.Slop != 1 {
			t.Error("预期 slop 为 1")
		}
		if matchPhrasePrefixQuery.MaxExpansions == nil || *matchPhrasePrefixQuery.MaxExpansions != 10 {
			t.Error("预期 maxExpansions 为 10")
		}
		if matchPhrasePrefixQuery.ZeroTermsQuery == nil || *matchPhrasePrefixQuery.ZeroTermsQuery != zerotermsquery.None {
			t.Error("预期 zeroTermsQuery 为 'none'")
		}
	})

	t.Run("应该可以处理 nil 回调", func(t *testing.T) {
		query := NewQuery(MatchPhrasePrefixWithOptions("title", "quick brown f", nil))
		if query.MatchPhrasePrefix == nil {
			t.Error("预期为 MatchPhrasePrefix 查询")
		}
		if query.MatchPhrasePrefix["title"].Query != "quick brown f" {
			t.Errorf("预期查询为 'quick brown f'，得到 %s", query.MatchPhrasePrefix["title"].Query)
		}
	})
}

func TestMatchInBoolQuery(t *testing.T) {
	t.Run("应该可以处理布尔查询", func(t *testing.T) {
		query := NewQuery(
			Bool(
				Must(
					Match("title", "elasticsearch"),
					MatchPhrase("content", "search engine"),
				),
				Should(
					Match("tags", "database"),
					MatchPhrasePrefix("description", "fast sea"),
				),
			),
		)
		if query.Bool == nil {
			t.Error("预期为布尔查询")
		}
		if len(query.Bool.Must) != 2 {
			t.Errorf("预期有 2 个 Must 子句，得到 %d", len(query.Bool.Must))
		}
		if len(query.Bool.Should) != 2 {
			t.Errorf("预期有 2 个 Should 子句，得到 %d", len(query.Bool.Should))
		}
		
		// 检查 Match 查询是否正确嵌套
		mustQuery1 := query.Bool.Must[0]
		if mustQuery1.Match == nil {
			t.Error("预期布尔查询中的 Must 子句包含 Match 查询")
		}
		
		mustQuery2 := query.Bool.Must[1]
		if mustQuery2.MatchPhrase == nil {
			t.Error("预期布尔查询中的 Must 子句包含 MatchPhrase 查询")
		}
		
		shouldQuery1 := query.Bool.Should[0]
		if shouldQuery1.Match == nil {
			t.Error("预期布尔查询中的 Should 子句包含 Match 查询")
		}
		
		shouldQuery2 := query.Bool.Should[1]
		if shouldQuery2.MatchPhrasePrefix == nil {
			t.Error("预期布尔查询中的 Should 子句包含 MatchPhrasePrefix 查询")
		}
	})
}

func TestMatchCompatibility(t *testing.T) {
	t.Run("应该生成兼容的 Match 查询结构", func(t *testing.T) {
		query := NewQuery(Match("title", "elasticsearch search"))
		
		// 验证结构是否与 elasticsearch 期望的结构匹配
		if query.Match == nil {
			t.Error("预期为 Match 查询")
		}
		
		matchQuery := query.Match["title"]
		if matchQuery.Query != "elasticsearch search" {
			t.Errorf("预期查询为 'elasticsearch search'，得到 %s", matchQuery.Query)
		}
	})

	t.Run("应该匹配手动 Match 查询构造", func(t *testing.T) {
		// 我们的构建器方法
		builderQuery := NewQuery(Match("title", "elasticsearch"))

		// 手动构造
		manualQuery := &types.Query{
			Match: map[string]types.MatchQuery{
				"title": {
					Query: "elasticsearch",
				},
			},
		}

		// 比较结构
		if builderQuery.Match == nil || manualQuery.Match == nil {
			t.Error("两个查询都应该有 Match 查询")
		}
		
		if builderQuery.Match["title"].Query != manualQuery.Match["title"].Query {
			t.Errorf("查询不匹配：builder=%s，manual=%s", 
				builderQuery.Match["title"].Query, manualQuery.Match["title"].Query)
		}
	})
}

func TestMultiMatchWith(t *testing.T) {
	t.Run("应该创建带回调选项的多匹配查询", func(t *testing.T) {
		analyzer := "standard"
		boost := float32(1.5)
		typeVal := textquerytype.Bestfields
		query := NewQuery(
			MultiMatchWithOptions("elasticsearch", []string{"title", "content"}, func(q *types.MultiMatchQuery) {
				q.Analyzer = &analyzer
				q.Boost = &boost
				q.Type = &typeVal
			}),
		)
		if query.MultiMatch == nil {
			t.Error("预期为 MultiMatch 查询")
		}
		if query.MultiMatch.Analyzer == nil || *query.MultiMatch.Analyzer != "standard" {
			t.Error("预期分析器为 'standard'")
		}
		if query.MultiMatch.Boost == nil || *query.MultiMatch.Boost != 1.5 {
			t.Error("预期权重为 1.5")
		}
		if query.MultiMatch.Type == nil || *query.MultiMatch.Type != textquerytype.Bestfields {
			t.Error("预期类型为 Bestfields")
		}
	})

	t.Run("应该可以处理 nil 回调", func(t *testing.T) {
		query := NewQuery(MultiMatchWithOptions("elasticsearch", []string{"title"}, nil))
		if query.MultiMatch == nil {
			t.Error("预期为 MultiMatch 查询")
		}
		if query.MultiMatch.Query != "elasticsearch" {
			t.Errorf("预期查询为 'elasticsearch'，得到 %s", query.MultiMatch.Query)
		}
	})
}

func TestWildcardWith(t *testing.T) {
	t.Run("应该创建带回调选项的通配符查询", func(t *testing.T) {
		boost := float32(2.0)
		caseInsensitive := true
		query := NewQuery(
			WildcardWithOptions("username", "john*", func(q *types.WildcardQuery) {
				q.Boost = &boost
				q.CaseInsensitive = &caseInsensitive
			}),
		)
		if query.Wildcard == nil {
			t.Error("预期为 Wildcard 查询")
		}
		wildcardQuery := query.Wildcard["username"]
		if wildcardQuery.Boost == nil || *wildcardQuery.Boost != 2.0 {
			t.Error("预期权重为 2.0")
		}
		if wildcardQuery.CaseInsensitive == nil || *wildcardQuery.CaseInsensitive != true {
			t.Error("预期大小写不敏感为 true")
		}
		if wildcardQuery.Value == nil || *wildcardQuery.Value != "john*" {
			t.Errorf("预期值为 'john*'，得到 %v", wildcardQuery.Value)
		}
	})

	t.Run("应该可以处理 nil 回调", func(t *testing.T) {
		query := NewQuery(WildcardWithOptions("username", "john*", nil))
		if query.Wildcard == nil {
			t.Error("预期为 Wildcard 查询")
		}
		if query.Wildcard["username"].Value == nil || *query.Wildcard["username"].Value != "john*" {
			t.Errorf("预期值为 'john*'，得到 %v", query.Wildcard["username"].Value)
		}
	})
}

func TestFuzzyWith(t *testing.T) {
	t.Run("应该创建带回调选项的模糊查询", func(t *testing.T) {
		boost := float32(1.5)
		fuzziness := types.Fuzziness("AUTO")
		query := NewQuery(
			FuzzyWithOptions("username", "john", func(q *types.FuzzyQuery) {
				q.Fuzziness = fuzziness
				q.Boost = &boost
			}),
		)
		if query.Fuzzy == nil {
			t.Error("预期为 Fuzzy 查询")
		}
		fuzzyQuery := query.Fuzzy["username"]
		if fuzzyQuery.Fuzziness != fuzziness {
			t.Errorf("预期模糊匹配 %v，得到 %v", fuzziness, fuzzyQuery.Fuzziness)
		}
		if fuzzyQuery.Boost == nil || *fuzzyQuery.Boost != 1.5 {
			t.Error("预期权重为 1.5")
		}
	})

	t.Run("应该可以处理 nil 回调", func(t *testing.T) {
		query := NewQuery(FuzzyWithOptions("username", "john", nil))
		if query.Fuzzy == nil {
			t.Error("预期为 Fuzzy 查询")
		}
		if query.Fuzzy["username"].Value != "john" {
			t.Errorf("预期值为 'john'，得到 %s", query.Fuzzy["username"].Value)
		}
	})
}

func TestPrefixWith(t *testing.T) {
	t.Run("应该创建带回调选项的前缀查询", func(t *testing.T) {
		boost := float32(1.5)
		caseInsensitive := true
		query := NewQuery(
			PrefixWithOptions("username", "john", func(q *types.PrefixQuery) {
				q.Boost = &boost
				q.CaseInsensitive = &caseInsensitive
			}),
		)
		if query.Prefix == nil {
			t.Error("预期为 Prefix 查询")
		}
		prefixQuery := query.Prefix["username"]
		if prefixQuery.Boost == nil || *prefixQuery.Boost != 1.5 {
			t.Error("预期权重为 1.5")
		}
		if prefixQuery.CaseInsensitive == nil || *prefixQuery.CaseInsensitive != true {
			t.Error("预期大小写不敏感为 true")
		}
	})

	t.Run("应该可以处理 nil 回调", func(t *testing.T) {
		query := NewQuery(PrefixWithOptions("username", "john", nil))
		if query.Prefix == nil {
			t.Error("预期为 Prefix 查询")
		}
		if query.Prefix["username"].Value != "john" {
			t.Errorf("预期值为 'john'，得到 %s", query.Prefix["username"].Value)
		}
	})
}

func TestMatchWithOptionsFunc(t *testing.T) {
	t.Run("应该创建带回调选项的匹配查询", func(t *testing.T) {
		analyzer := "standard"
		boost := float32(1.5)
		fuzziness := types.Fuzziness("AUTO")
		query := NewQuery(
			MatchWithOptions("title", "elasticsearch guide", func(opts *types.MatchQuery) {
				opts.Fuzziness = fuzziness
				opts.Analyzer = &analyzer
				opts.Boost = &boost
			}),
		)
		if query.Match == nil {
			t.Error("预期为 Match 查询")
		}
		matchQuery := query.Match["title"]
		if matchQuery.Fuzziness != fuzziness {
			t.Errorf("预期模糊匹配 %v，得到 %v", fuzziness, matchQuery.Fuzziness)
		}
		if matchQuery.Analyzer == nil || *matchQuery.Analyzer != "standard" {
			t.Error("预期分析器为 'standard'")
		}
		if matchQuery.Boost == nil || *matchQuery.Boost != 1.5 {
			t.Error("预期权重为 1.5")
		}
	})

	t.Run("应该可以处理 nil 回调", func(t *testing.T) {
		query := NewQuery(MatchWithOptions("title", "elasticsearch", nil))
		if query.Match == nil {
			t.Error("预期为 Match 查询")
		}
		if query.Match["title"].Query != "elasticsearch" {
			t.Errorf("预期查询为 'elasticsearch'，得到 %s", query.Match["title"].Query)
		}
	})
}

func TestMatchPhraseWithOptionsFunc(t *testing.T) {
	t.Run("应该创建带回调选项的短语匹配查询", func(t *testing.T) {
		analyzer := "keyword"
		boost := float32(2.0)
		slop := 2
		query := NewQuery(
			MatchPhraseWithOptions("content", "elasticsearch is awesome", func(opts *types.MatchPhraseQuery) {
				opts.Slop = &slop
				opts.Analyzer = &analyzer
				opts.Boost = &boost
			}),
		)
		if query.MatchPhrase == nil {
			t.Error("预期为 MatchPhrase 查询")
		}
		matchPhraseQuery := query.MatchPhrase["content"]
		if matchPhraseQuery.Slop == nil || *matchPhraseQuery.Slop != 2 {
			t.Error("预期 slop 为 2")
		}
		if matchPhraseQuery.Analyzer == nil || *matchPhraseQuery.Analyzer != "keyword" {
			t.Error("预期分析器为 'keyword'")
		}
		if matchPhraseQuery.Boost == nil || *matchPhraseQuery.Boost != 2.0 {
			t.Error("预期权重为 2.0")
		}
	})

	t.Run("应该可以处理 nil 回调", func(t *testing.T) {
		query := NewQuery(MatchPhraseWithOptions("content", "elasticsearch is awesome", nil))
		if query.MatchPhrase == nil {
			t.Error("预期为 MatchPhrase 查询")
		}
		if query.MatchPhrase["content"].Query != "elasticsearch is awesome" {
			t.Errorf("预期查询为 'elasticsearch is awesome'，得到 %s", query.MatchPhrase["content"].Query)
		}
	})
}

func TestMultiMatchWithOptionsFunc(t *testing.T) {
	t.Run("应该创建带回调选项的多匹配查询", func(t *testing.T) {
		analyzer := "standard"
		boost := float32(1.5)
		typeVal := textquerytype.Bestfields
		query := NewQuery(
			MultiMatchWithOptions("elasticsearch", []string{"title", "content"}, func(opts *types.MultiMatchQuery) {
				opts.Analyzer = &analyzer
				opts.Boost = &boost
				opts.Type = &typeVal
			}),
		)
		if query.MultiMatch == nil {
			t.Error("预期为 MultiMatch 查询")
		}
		if query.MultiMatch.Analyzer == nil || *query.MultiMatch.Analyzer != "standard" {
			t.Error("预期分析器为 'standard'")
		}
		if query.MultiMatch.Boost == nil || *query.MultiMatch.Boost != 1.5 {
			t.Error("预期权重为 1.5")
		}
		if query.MultiMatch.Type == nil || *query.MultiMatch.Type != textquerytype.Bestfields {
			t.Error("预期类型为 Bestfields")
		}
	})

	t.Run("应该可以处理 nil 回调", func(t *testing.T) {
		query := NewQuery(MultiMatchWithOptions("elasticsearch", []string{"title"}, nil))
		if query.MultiMatch == nil {
			t.Error("预期为 MultiMatch 查询")
		}
		if query.MultiMatch.Query != "elasticsearch" {
			t.Errorf("预期查询为 'elasticsearch'，得到 %s", query.MultiMatch.Query)
		}
	})
}

func TestWildcardWithOptionsFunc(t *testing.T) {
	t.Run("应该创建带回调选项的通配符查询", func(t *testing.T) {
		boost := float32(2.0)
		caseInsensitive := true
		query := NewQuery(
			WildcardWithOptions("username", "john*", func(opts *types.WildcardQuery) {
				opts.Boost = &boost
				opts.CaseInsensitive = &caseInsensitive
			}),
		)
		if query.Wildcard == nil {
			t.Error("预期为 Wildcard 查询")
		}
		wildcardQuery := query.Wildcard["username"]
		if wildcardQuery.Boost == nil || *wildcardQuery.Boost != 2.0 {
			t.Error("预期权重为 2.0")
		}
		if wildcardQuery.CaseInsensitive == nil || *wildcardQuery.CaseInsensitive != true {
			t.Error("预期大小写不敏感为 true")
		}
		if wildcardQuery.Value == nil || *wildcardQuery.Value != "john*" {
			t.Errorf("预期值为 'john*'，得到 %v", wildcardQuery.Value)
		}
	})

	t.Run("应该可以处理 nil 回调", func(t *testing.T) {
		query := NewQuery(WildcardWithOptions("username", "john*", nil))
		if query.Wildcard == nil {
			t.Error("预期为 Wildcard 查询")
		}
		if query.Wildcard["username"].Value == nil || *query.Wildcard["username"].Value != "john*" {
			t.Errorf("预期值为 'john*'，得到 %v", query.Wildcard["username"].Value)
		}
	})
}

func TestFuzzyWithOptionsFunc(t *testing.T) {
	t.Run("应该创建带回调选项的模糊查询", func(t *testing.T) {
		boost := float32(1.5)
		fuzziness := types.Fuzziness("AUTO")
		query := NewQuery(
			FuzzyWithOptions("username", "john", func(opts *types.FuzzyQuery) {
				opts.Fuzziness = fuzziness
				opts.Boost = &boost
			}),
		)
		if query.Fuzzy == nil {
			t.Error("预期为 Fuzzy 查询")
		}
		fuzzyQuery := query.Fuzzy["username"]
		if fuzzyQuery.Fuzziness != fuzziness {
			t.Errorf("预期模糊匹配 %v，得到 %v", fuzziness, fuzzyQuery.Fuzziness)
		}
		if fuzzyQuery.Boost == nil || *fuzzyQuery.Boost != 1.5 {
			t.Error("预期权重为 1.5")
		}
	})

	t.Run("应该可以处理 nil 回调", func(t *testing.T) {
		query := NewQuery(FuzzyWithOptions("username", "john", nil))
		if query.Fuzzy == nil {
			t.Error("预期为 Fuzzy 查询")
		}
		if query.Fuzzy["username"].Value != "john" {
			t.Errorf("预期值为 'john'，得到 %s", query.Fuzzy["username"].Value)
		}
	})
}

func TestPrefixWithOptionsFunc(t *testing.T) {
	t.Run("应该创建带回调选项的前缀查询", func(t *testing.T) {
		boost := float32(1.5)
		caseInsensitive := true
		query := NewQuery(
			PrefixWithOptions("username", "john", func(opts *types.PrefixQuery) {
				opts.Boost = &boost
				opts.CaseInsensitive = &caseInsensitive
			}),
		)
		if query.Prefix == nil {
			t.Error("预期为 Prefix 查询")
		}
		prefixQuery := query.Prefix["username"]
		if prefixQuery.Boost == nil || *prefixQuery.Boost != 1.5 {
			t.Error("预期权重为 1.5")
		}
		if prefixQuery.CaseInsensitive == nil || *prefixQuery.CaseInsensitive != true {
			t.Error("预期大小写不敏感为 true")
		}
	})

	t.Run("应该可以处理 nil 回调", func(t *testing.T) {
		query := NewQuery(PrefixWithOptions("username", "john", nil))
		if query.Prefix == nil {
			t.Error("预期为 Prefix 查询")
		}
		if query.Prefix["username"].Value != "john" {
			t.Errorf("预期值为 'john'，得到 %s", query.Prefix["username"].Value)
		}
	})
}

// 匹配查询的基准测试
func BenchmarkMatchQuery(b *testing.B) {
	b.Run("简单匹配", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(Match("title", "elasticsearch search"))
		}
	})

	b.Run("带选项的匹配", func(b *testing.B) {
		op := operator.And
		fuzziness := types.Fuzziness("AUTO")
		for i := 0; i < b.N; i++ {
			_ = NewQuery(
				MatchWithOptions("title", "elasticsearch search", func(opts *types.MatchQuery) {
					opts.Operator = &op
					opts.Fuzziness = fuzziness
				}),
			)
		}
	})

	b.Run("短语匹配", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(MatchPhrase("content", "elasticsearch is awesome"))
		}
	})

	b.Run("短语前缀匹配", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(MatchPhrasePrefix("title", "elasticsearch sea"))
		}
	})

	b.Run("布尔查询中的复杂匹配", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewQuery(
				Bool(
					Must(
						Match("title", "elasticsearch"),
						MatchPhrase("content", "search engine"),
					),
					Should(
						Match("tags", "database"),
						MatchPhrasePrefix("description", "fast sea"),
					),
				),
			)
		}
	})
} 
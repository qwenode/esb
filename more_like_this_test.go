package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestMoreLikeThis(t *testing.T) {
	t.Run("test basic more like this query", func(t *testing.T) {
		query := NewQuery(
			MoreLikeThis([]string{"title", "content"}, "This is a sample text"),
		)

		if query.MoreLikeThis == nil {
			t.Error("expected MoreLikeThis to be set")
		}

		if len(query.MoreLikeThis.Fields) != 2 {
			t.Errorf("expected 2 fields, got %d", len(query.MoreLikeThis.Fields))
		}

		if query.MoreLikeThis.Fields[0] != "title" {
			t.Errorf("expected first field to be 'title', got %s", query.MoreLikeThis.Fields[0])
		}

		if query.MoreLikeThis.Fields[1] != "content" {
			t.Errorf("expected second field to be 'content', got %s", query.MoreLikeThis.Fields[1])
		}

		if len(query.MoreLikeThis.Like) != 1 {
			t.Errorf("expected 1 like item, got %d", len(query.MoreLikeThis.Like))
		}

		if likeText, ok := query.MoreLikeThis.Like[0].(string); !ok || likeText != "This is a sample text" {
			t.Errorf("expected like text to be 'This is a sample text', got %v", query.MoreLikeThis.Like[0])
		}
	})
}

func TestMoreLikeThisWithOptions(t *testing.T) {
	t.Run("test more like this query with options", func(t *testing.T) {
		query := NewQuery(
			MoreLikeThisWithOptions([]string{"title", "content"}, "This is a sample text", func(opts *types.MoreLikeThisQuery) {
				minTermFreq := 2
				opts.MinTermFreq = &minTermFreq
				maxQueryTerms := 25
				opts.MaxQueryTerms = &maxQueryTerms
				minDocFreq := 1
				opts.MinDocFreq = &minDocFreq
			}),
		)

		if query.MoreLikeThis == nil {
			t.Error("expected MoreLikeThis to be set")
		}

		if query.MoreLikeThis.MinTermFreq == nil || *query.MoreLikeThis.MinTermFreq != 2 {
			t.Errorf("expected MinTermFreq to be 2, got %v", query.MoreLikeThis.MinTermFreq)
		}

		if query.MoreLikeThis.MaxQueryTerms == nil || *query.MoreLikeThis.MaxQueryTerms != 25 {
			t.Errorf("expected MaxQueryTerms to be 25, got %v", query.MoreLikeThis.MaxQueryTerms)
		}

		if query.MoreLikeThis.MinDocFreq == nil || *query.MoreLikeThis.MinDocFreq != 1 {
			t.Errorf("expected MinDocFreq to be 1, got %v", query.MoreLikeThis.MinDocFreq)
		}

		if len(query.MoreLikeThis.Fields) != 2 {
			t.Errorf("expected 2 fields, got %d", len(query.MoreLikeThis.Fields))
		}

		if len(query.MoreLikeThis.Like) != 1 {
			t.Errorf("expected 1 like item, got %d", len(query.MoreLikeThis.Like))
		}

		if likeText, ok := query.MoreLikeThis.Like[0].(string); !ok || likeText != "This is a sample text" {
			t.Errorf("expected like text to be 'This is a sample text', got %v", query.MoreLikeThis.Like[0])
		}
	})

	t.Run("test more like this query with nil setOpts", func(t *testing.T) {
		query := NewQuery(
			MoreLikeThisWithOptions([]string{"title"}, "sample text", nil),
		)

		if query.MoreLikeThis == nil {
			t.Error("expected MoreLikeThis to be set")
		}

		if len(query.MoreLikeThis.Fields) != 1 {
			t.Errorf("expected 1 field, got %d", len(query.MoreLikeThis.Fields))
		}

		if len(query.MoreLikeThis.Like) != 1 {
			t.Errorf("expected 1 like item, got %d", len(query.MoreLikeThis.Like))
		}

		// 应该没有设置其他选项
		if query.MoreLikeThis.MinTermFreq != nil {
			t.Error("expected MinTermFreq to be nil")
		}

		if query.MoreLikeThis.MaxQueryTerms != nil {
			t.Error("expected MaxQueryTerms to be nil")
		}

		if query.MoreLikeThis.MinDocFreq != nil {
			t.Error("expected MinDocFreq to be nil")
		}
	})
}

func TestMoreLikeThisWithDocument(t *testing.T) {
	t.Run("test more like this query with document", func(t *testing.T) {
		query := NewQuery(
			MoreLikeThisWithDocument([]string{"title", "content"}, "my-index", "1"),
		)

		if query.MoreLikeThis == nil {
			t.Error("expected MoreLikeThis to be set")
		}

		if len(query.MoreLikeThis.Like) != 1 {
			t.Errorf("expected 1 like item, got %d", len(query.MoreLikeThis.Like))
		}

		if likeDoc, ok := query.MoreLikeThis.Like[0].(types.LikeDocument); ok {
			if likeDoc.Index_ == nil || *likeDoc.Index_ != "my-index" {
				t.Errorf("expected index to be 'my-index', got %v", likeDoc.Index_)
			}
			if likeDoc.Id_ == nil || *likeDoc.Id_ != "1" {
				t.Errorf("expected id to be '1', got %v", likeDoc.Id_)
			}
		} else {
			t.Error("expected like item to be LikeDocument")
		}
	})
}

func TestMoreLikeThisWithMultipleLikes(t *testing.T) {
	t.Run("test more like this query with multiple likes", func(t *testing.T) {
		likes := []types.Like{"text1", "text2"}
		query := NewQuery(
			MoreLikeThisWithMultipleLikes([]string{"title", "content"}, likes),
		)

		if query.MoreLikeThis == nil {
			t.Error("expected MoreLikeThis to be set")
		}

		if len(query.MoreLikeThis.Like) != 2 {
			t.Errorf("expected 2 like items, got %d", len(query.MoreLikeThis.Like))
		}

		if likeText1, ok := query.MoreLikeThis.Like[0].(string); !ok || likeText1 != "text1" {
			t.Errorf("expected first like text to be 'text1', got %v", query.MoreLikeThis.Like[0])
		}

		if likeText2, ok := query.MoreLikeThis.Like[1].(string); !ok || likeText2 != "text2" {
			t.Errorf("expected second like text to be 'text2', got %v", query.MoreLikeThis.Like[1])
		}
	})
}

func TestMoreLikeThisWithUnlike(t *testing.T) {
	t.Run("test more like this query with unlike", func(t *testing.T) {
		query := NewQuery(
			MoreLikeThisWithUnlike([]string{"title", "content"}, "This is a sample text", "unwanted text"),
		)

		if query.MoreLikeThis == nil {
			t.Error("expected MoreLikeThis to be set")
		}

		if len(query.MoreLikeThis.Like) != 1 {
			t.Errorf("expected 1 like item, got %d", len(query.MoreLikeThis.Like))
		}

		if likeText, ok := query.MoreLikeThis.Like[0].(string); !ok || likeText != "This is a sample text" {
			t.Errorf("expected like text to be 'This is a sample text', got %v", query.MoreLikeThis.Like[0])
		}

		if len(query.MoreLikeThis.Unlike) != 1 {
			t.Errorf("expected 1 unlike item, got %d", len(query.MoreLikeThis.Unlike))
		}

		if unlikeText, ok := query.MoreLikeThis.Unlike[0].(string); !ok || unlikeText != "unwanted text" {
			t.Errorf("expected unlike text to be 'unwanted text', got %v", query.MoreLikeThis.Unlike[0])
		}
	})
} 
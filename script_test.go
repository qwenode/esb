package esb

import (
	"encoding/json"
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestScript(t *testing.T) {
	t.Run("test basic script query", func(t *testing.T) {
		query := NewQuery(
			Script("doc['age'].value > 30"),
		)

		if query.Script == nil {
			t.Error("expected Script to be set")
		}

		if *query.Script.Script.Source != "doc['age'].value > 30" {
			t.Errorf("expected Script.Source to be 'doc['age'].value > 30', got %s", *query.Script.Script.Source)
		}
	})
}

func TestScriptWithParams(t *testing.T) {
	t.Run("test script query with parameters", func(t *testing.T) {
		params := map[string]any{"min_age": 30}
		query := NewQuery(
			ScriptWithParams("doc['age'].value > params.min_age", params),
		)

		if query.Script == nil {
			t.Error("expected Script to be set")
		}

		if *query.Script.Script.Source != "doc['age'].value > params.min_age" {
			t.Errorf("expected Script.Source to be 'doc['age'].value > params.min_age', got %s", *query.Script.Script.Source)
		}

		if query.Script.Script.Params == nil {
			t.Error("expected Script.Params to be set")
		}

		if minAgeRaw, ok := query.Script.Script.Params["min_age"]; ok {
			var minAge int
			if err := json.Unmarshal(minAgeRaw, &minAge); err != nil {
				t.Errorf("failed to unmarshal min_age: %v", err)
			}
			if minAge != 30 {
				t.Errorf("expected Script.Params['min_age'] to be 30, got %v", minAge)
			}
		} else {
			t.Error("expected Script.Params['min_age'] to exist")
		}
	})
}

func TestScriptWithLang(t *testing.T) {
	t.Run("test script query with painless language", func(t *testing.T) {
		query := NewQuery(
			ScriptWithLang("doc['age'].value > 30", scriptlanguage.Painless),
		)

		if query.Script == nil {
			t.Error("expected Script to be set")
		}

		if *query.Script.Script.Source != "doc['age'].value > 30" {
			t.Errorf("expected Script.Source to be 'doc['age'].value > 30', got %s", *query.Script.Script.Source)
		}

		if query.Script.Script.Lang == nil {
			t.Error("expected Script.Lang to be set")
		}

		if *query.Script.Script.Lang != scriptlanguage.Painless {
			t.Errorf("expected Script.Lang to be 'painless', got %s", query.Script.Script.Lang.String())
		}
	})

	t.Run("test script query with expression language", func(t *testing.T) {
		query := NewQuery(
			ScriptWithLang("doc['age'].value > 30", scriptlanguage.Expression),
		)

		if *query.Script.Script.Lang != scriptlanguage.Expression {
			t.Errorf("expected Script.Lang to be 'expression', got %s", query.Script.Script.Lang.String())
		}
	})

	t.Run("test script query with mustache language", func(t *testing.T) {
		query := NewQuery(
			ScriptWithLang("doc['age'].value > 30", scriptlanguage.Mustache),
		)

		if *query.Script.Script.Lang != scriptlanguage.Mustache {
			t.Errorf("expected Script.Lang to be 'mustache', got %s", query.Script.Script.Lang.String())
		}
	})

	t.Run("test script query with java language", func(t *testing.T) {
		query := NewQuery(
			ScriptWithLang("doc['age'].value > 30", scriptlanguage.Java),
		)

		if *query.Script.Script.Lang != scriptlanguage.Java {
			t.Errorf("expected Script.Lang to be 'java', got %s", query.Script.Script.Lang.String())
		}
	})

	t.Run("test script query with custom language", func(t *testing.T) {
		customLang := scriptlanguage.ScriptLanguage{Name: "custom"}
		query := NewQuery(
			ScriptWithLang("doc['age'].value > 30", customLang),
		)

		if query.Script.Script.Lang.String() != "custom" {
			t.Errorf("expected Script.Lang to be 'custom', got %s", query.Script.Script.Lang.String())
		}
	})
}

func TestScriptWithOptions(t *testing.T) {
	t.Run("test script query with all options", func(t *testing.T) {
		params := map[string]any{"min_age": 30}
		query := NewQuery(
			ScriptWithOptions("doc['age'].value > params.min_age", func(opts *types.ScriptQuery) {
				lang := scriptlanguage.Painless
				opts.Script.Lang = &lang
				
				// 设置参数
				jsonParams := make(map[string]json.RawMessage)
				for k, v := range params {
					jsonBytes, _ := json.Marshal(v)
					jsonParams[k] = json.RawMessage(jsonBytes)
				}
				opts.Script.Params = jsonParams
			}),
		)

		if query.Script == nil {
			t.Error("expected Script to be set")
		}

		if *query.Script.Script.Source != "doc['age'].value > params.min_age" {
			t.Errorf("expected Script.Source to be 'doc['age'].value > params.min_age', got %s", *query.Script.Script.Source)
		}

		if query.Script.Script.Lang == nil {
			t.Error("expected Script.Lang to be set")
		}

		if *query.Script.Script.Lang != scriptlanguage.Painless {
			t.Errorf("expected Script.Lang to be 'painless', got %s", query.Script.Script.Lang.String())
		}

		if query.Script.Script.Params == nil {
			t.Error("expected Script.Params to be set")
		}

		if minAgeRaw, ok := query.Script.Script.Params["min_age"]; ok {
			var minAge int
			if err := json.Unmarshal(minAgeRaw, &minAge); err != nil {
				t.Errorf("failed to unmarshal min_age: %v", err)
			}
			if minAge != 30 {
				t.Errorf("expected Script.Params['min_age'] to be 30, got %v", minAge)
			}
		} else {
			t.Error("expected Script.Params['min_age'] to exist")
		}
	})

	t.Run("test script query with expression language", func(t *testing.T) {
		query := NewQuery(
			ScriptWithOptions("doc['age'].value > 30", func(opts *types.ScriptQuery) {
				lang := scriptlanguage.Expression
				opts.Script.Lang = &lang
			}),
		)

		if *query.Script.Script.Lang != scriptlanguage.Expression {
			t.Errorf("expected Script.Lang to be 'expression', got %s", query.Script.Script.Lang.String())
		}
	})

	t.Run("test script query with mustache language", func(t *testing.T) {
		query := NewQuery(
			ScriptWithOptions("doc['age'].value > 30", func(opts *types.ScriptQuery) {
				lang := scriptlanguage.Mustache
				opts.Script.Lang = &lang
			}),
		)

		if *query.Script.Script.Lang != scriptlanguage.Mustache {
			t.Errorf("expected Script.Lang to be 'mustache', got %s", query.Script.Script.Lang.String())
		}
	})

	t.Run("test script query with java language", func(t *testing.T) {
		query := NewQuery(
			ScriptWithOptions("doc['age'].value > 30", func(opts *types.ScriptQuery) {
				lang := scriptlanguage.Java
				opts.Script.Lang = &lang
			}),
		)

		if *query.Script.Script.Lang != scriptlanguage.Java {
			t.Errorf("expected Script.Lang to be 'java', got %s", query.Script.Script.Lang.String())
		}
	})

	t.Run("test script query with custom language", func(t *testing.T) {
		query := NewQuery(
			ScriptWithOptions("doc['age'].value > 30", func(opts *types.ScriptQuery) {
				customLang := scriptlanguage.ScriptLanguage{Name: "custom"}
				opts.Script.Lang = &customLang
			}),
		)

		if query.Script.Script.Lang.String() != "custom" {
			t.Errorf("expected Script.Lang to be 'custom', got %s", query.Script.Script.Lang.String())
		}
	})

	t.Run("test script query with boost", func(t *testing.T) {
		query := NewQuery(
			ScriptWithOptions("doc['age'].value > 30", func(opts *types.ScriptQuery) {
				boost := float32(1.5)
				opts.Boost = &boost
			}),
		)

		if query.Script.Boost == nil {
			t.Error("expected Script.Boost to be set")
		}

		if *query.Script.Boost != 1.5 {
			t.Errorf("expected Script.Boost to be 1.5, got %f", *query.Script.Boost)
		}
	})

	t.Run("test script query with query name", func(t *testing.T) {
		query := NewQuery(
			ScriptWithOptions("doc['age'].value > 30", func(opts *types.ScriptQuery) {
				name := "test_script"
				opts.QueryName_ = &name
			}),
		)

		if query.Script.QueryName_ == nil {
			t.Error("expected Script.QueryName_ to be set")
		}

		if *query.Script.QueryName_ != "test_script" {
			t.Errorf("expected Script.QueryName_ to be 'test_script', got %s", *query.Script.QueryName_)
		}
	})

	t.Run("test script query with nil setOpts", func(t *testing.T) {
		query := NewQuery(
			ScriptWithOptions("doc['age'].value > 30", nil),
		)

		if query.Script == nil {
			t.Error("expected Script to be set")
		}

		if *query.Script.Script.Source != "doc['age'].value > 30" {
			t.Errorf("expected Script.Source to be 'doc['age'].value > 30', got %s", *query.Script.Script.Source)
		}

		// 应该没有设置其他选项
		if query.Script.Script.Lang != nil {
			t.Error("expected Script.Lang to be nil")
		}

		if query.Script.Script.Params != nil {
			t.Error("expected Script.Params to be nil")
		}

		if query.Script.Boost != nil {
			t.Error("expected Script.Boost to be nil")
		}

		if query.Script.QueryName_ != nil {
			t.Error("expected Script.QueryName_ to be nil")
		}
	})
} 
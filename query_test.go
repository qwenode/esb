package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestNewQuery(t *testing.T) {
	t.Run("当没有提供选项时应该创建空查询", func(t *testing.T) {
		query := NewQuery()
		if query == nil {
			t.Error("预期查询不为 nil")
		}
	})

	t.Run("应该使用有效选项创建查询", func(t *testing.T) {
		query := NewQuery(Term("status", "published"))
		if query == nil {
			t.Error("预期查询不为 nil")
		}
	})

	t.Run("应该使用自定义选项创建查询", func(t *testing.T) {
		validOption := func(q *types.Query) {
			// 简单的选项，设置一个字段（我们稍后会实现实际的选项）
		}

		query := NewQuery(validOption)
		if query == nil {
			t.Error("预期查询不为 nil")
		}
	})

	t.Run("应该按顺序应用多个选项", func(t *testing.T) {
		callOrder := []int{}
		
		option1 := func(q *types.Query) {
			callOrder = append(callOrder, 1)
		}
		
		option2 := func(q *types.Query) {
			callOrder = append(callOrder, 2)
		}

		query := NewQuery(option1, option2)
		if query == nil {
			t.Error("预期查询不为 nil")
		}
		
		if len(callOrder) != 2 || callOrder[0] != 1 || callOrder[1] != 2 {
			t.Errorf("预期调用顺序为 [1, 2]，得到 %v", callOrder)
		}
	})
}

// 基准测试以确保性能可接受
func BenchmarkNewQuery(b *testing.B) {
	option := func(q *types.Query) {}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewQuery(option)
	}
}

func BenchmarkNewQueryMultipleOptions(b *testing.B) {
	option1 := func(q *types.Query) {}
	option2 := func(q *types.Query) {}
	option3 := func(q *types.Query) {}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewQuery(option1, option2, option3)
	}
} 
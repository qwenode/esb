package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// TestIDs 测试基本的IDs查询功能
func TestIDs(t *testing.T) {
	t.Run("测试基本IDs查询", func(t *testing.T) {
		query := NewQuery(IDs("1", "2", "3"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Ids == nil {
			t.Fatal("Ids查询不应该为nil")
		}
		
		expectedIds := []string{"1", "2", "3"}
		if len(query.Ids.Values) != len(expectedIds) {
			t.Fatalf("期望IDs长度为%d, 实际得到: %d", len(expectedIds), len(query.Ids.Values))
		}
		
		for i, id := range expectedIds {
			if query.Ids.Values[i] != id {
				t.Errorf("期望ID[%d]为'%s', 实际得到: '%s'", i, id, query.Ids.Values[i])
			}
		}
	})
	
	t.Run("测试单个ID", func(t *testing.T) {
		query := NewQuery(IDs("single-id"))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Ids == nil {
			t.Fatal("Ids查询不应该为nil")
		}
		
		if len(query.Ids.Values) != 1 {
			t.Fatalf("期望IDs长度为1, 实际得到: %d", len(query.Ids.Values))
		}
		
		if query.Ids.Values[0] != "single-id" {
			t.Errorf("期望ID为'single-id', 实际得到: '%s'", query.Ids.Values[0])
		}
	})
	
	t.Run("测试空ID列表", func(t *testing.T) {
		query := NewQuery(IDs())
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Ids == nil {
			t.Fatal("Ids查询不应该为nil")
		}
		
		if len(query.Ids.Values) != 0 {
			t.Fatalf("期望IDs长度为0, 实际得到: %d", len(query.Ids.Values))
		}
	})
}

// TestIDsWithOptions 测试带选项的IDs查询
func TestIDsWithOptions(t *testing.T) {
	t.Run("测试带boost选项的IDs查询", func(t *testing.T) {
		boost := float32(2.0)
		queryName := "test-ids-query"
		
		query := NewQuery(
			IDsWithOptions([]string{"1", "2", "3"}, func(opts *types.IdsQuery) {
				opts.Boost = &boost
				opts.QueryName_ = &queryName
			}),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Ids == nil {
			t.Fatal("Ids查询不应该为nil")
		}
		
		// 验证IDs值
		expectedIds := []string{"1", "2", "3"}
		if len(query.Ids.Values) != len(expectedIds) {
			t.Fatalf("期望IDs长度为%d, 实际得到: %d", len(expectedIds), len(query.Ids.Values))
		}
		
		for i, id := range expectedIds {
			if query.Ids.Values[i] != id {
				t.Errorf("期望ID[%d]为'%s', 实际得到: '%s'", i, id, query.Ids.Values[i])
			}
		}
		
		// 验证boost选项
		if query.Ids.Boost == nil || *query.Ids.Boost != 2.0 {
			t.Errorf("期望Boost为2.0, 实际得到: %v", query.Ids.Boost)
		}
		
		// 验证queryName选项
		if query.Ids.QueryName_ == nil || *query.Ids.QueryName_ != "test-ids-query" {
			t.Errorf("期望QueryName为'test-ids-query', 实际得到: %v", query.Ids.QueryName_)
		}
	})
	
	t.Run("测试不带选项的IDs查询", func(t *testing.T) {
		query := NewQuery(
			IDsWithOptions([]string{"1", "2"}, nil),
		)
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Ids == nil {
			t.Fatal("Ids查询不应该为nil")
		}
		
		// 验证默认选项
		if query.Ids.Boost != nil {
			t.Errorf("期望Boost为nil, 实际得到: %v", query.Ids.Boost)
		}
		
		if query.Ids.QueryName_ != nil {
			t.Errorf("期望QueryName为nil, 实际得到: %v", query.Ids.QueryName_)
		}
	})
}

// TestIDsFromSlice 测试从切片创建IDs查询
func TestIDsFromSlice(t *testing.T) {
	t.Run("测试从切片创建IDs查询", func(t *testing.T) {
		userIds := []string{"user-1", "user-2", "user-3"}
		query := NewQuery(IDsFromSlice(userIds))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Ids == nil {
			t.Fatal("Ids查询不应该为nil")
		}
		
		if len(query.Ids.Values) != len(userIds) {
			t.Fatalf("期望IDs长度为%d, 实际得到: %d", len(userIds), len(query.Ids.Values))
		}
		
		for i, id := range userIds {
			if query.Ids.Values[i] != id {
				t.Errorf("期望ID[%d]为'%s', 实际得到: '%s'", i, id, query.Ids.Values[i])
			}
		}
	})
	
	t.Run("测试空切片", func(t *testing.T) {
		emptyIds := []string{}
		query := NewQuery(IDsFromSlice(emptyIds))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Ids == nil {
			t.Fatal("Ids查询不应该为nil")
		}
		
		if len(query.Ids.Values) != 0 {
			t.Fatalf("期望IDs长度为0, 实际得到: %d", len(query.Ids.Values))
		}
	})
} 

// TestIDsSlice 测试IDsSlice查询功能
func TestIDsSlice(t *testing.T) {
	t.Run("测试基本IDsSlice查询", func(t *testing.T) {
		ids := []string{"1", "2", "3"}
		query := NewQuery(IDsSlice(ids))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Ids == nil {
			t.Fatal("Ids查询不应该为nil")
		}
		
		if len(query.Ids.Values) != 3 {
			t.Errorf("期望Values长度为3, 实际得到: %d", len(query.Ids.Values))
		}
		
		for i, id := range []string{"1", "2", "3"} {
			if query.Ids.Values[i] != id {
				t.Errorf("期望Values[%d]为'%s', 实际得到: %s", i, id, query.Ids.Values[i])
			}
		}
	})
	
	t.Run("测试单个值的IDsSlice查询", func(t *testing.T) {
		ids := []string{"1"}
		query := NewQuery(IDsSlice(ids))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Ids == nil {
			t.Fatal("Ids查询不应该为nil")
		}
		
		if len(query.Ids.Values) != 1 {
			t.Errorf("期望Values长度为1, 实际得到: %d", len(query.Ids.Values))
		}
		
		if query.Ids.Values[0] != "1" {
			t.Errorf("期望Values[0]为'1', 实际得到: %s", query.Ids.Values[0])
		}
	})
	
	t.Run("测试空切片的IDsSlice查询", func(t *testing.T) {
		var ids []string
		query := NewQuery(IDsSlice(ids))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Ids == nil {
			t.Fatal("Ids查询不应该为nil")
		}
		
		if len(query.Ids.Values) != 0 {
			t.Errorf("期望Values长度为0, 实际得到: %d", len(query.Ids.Values))
		}
	})
	
	t.Run("测试nil切片的IDsSlice查询", func(t *testing.T) {
		var ids []string = nil
		query := NewQuery(IDsSlice(ids))
		
		if query == nil {
			t.Fatal("query不应该为nil")
		}
		
		if query.Ids == nil {
			t.Fatal("Ids查询不应该为nil")
		}
		
		if len(query.Ids.Values) != 0 {
			t.Errorf("期望Values长度为0, 实际得到: %d", len(query.Ids.Values))
		}
	})
} 
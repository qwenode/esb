package esb

import (
	"testing"
)

func TestExists(t *testing.T) {
	tests := []struct {
		name    string
		field   string
		wantErr bool
	}{
		{
			name:    "有效字段名",
			field:   "user.name",
			wantErr: false,
		},
		{
			name:    "简单字段名",
			field:   "status",
			wantErr: false,
		},
		{
			name:    "嵌套字段名",
			field:   "metadata.timestamp",
			wantErr: false,
		},
		{
			name:    "空字段名",
			field:   "",
			wantErr: false,
		},
		{
			name:    "空白字段名",
			field:   "   ",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := NewQuery(Exists(tt.field))
			
			if tt.wantErr {
				if query.Exists == nil {
					t.Errorf("Exists() 预期出错，但得到 nil")
				}
				return
			}
			
			if query.Exists == nil {
				t.Errorf("Exists() query.Exists 为 nil")
				return
			}
			
			if query.Exists.Field != tt.field {
				t.Errorf("Exists() 字段 = %v，期望 %v", query.Exists.Field, tt.field)
			}
		})
	}
}

func TestExistsWithOtherQueries(t *testing.T) {
	// 测试将 Exists 与布尔查询组合
	query := NewQuery(
		Bool(
			Must(
				Exists("user.name"),
				Term("status", "active"),
			),
		),
	)
	
	if query.Bool == nil {
		t.Errorf("布尔查询为 nil")
		return
	}
	
	if len(query.Bool.Must) != 2 {
		t.Errorf("Bool.Must 长度 = %d，期望 2", len(query.Bool.Must))
		return
	}
	
	// 检查第一个 Must 子句是否为 Exists
	if query.Bool.Must[0].Exists == nil {
		t.Errorf("第一个 Must 子句应该是 Exists 查询")
		return
	}
	
	if query.Bool.Must[0].Exists.Field != "user.name" {
		t.Errorf("Exists 字段 = %v，期望 user.name", query.Bool.Must[0].Exists.Field)
	}
	
	// 检查第二个 Must 子句是否为 Term
	if query.Bool.Must[1].Term == nil {
		t.Errorf("第二个 Must 子句应该是 Term 查询")
		return
	}
}

// Exists 查询的基准测试
func BenchmarkExists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewQuery(Exists("user.name"))
	}
}

func BenchmarkExistsWithBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewQuery(
			Bool(
				Must(
					Exists("user.name"),
					Term("status", "active"),
				),
			),
		)
	}
} 
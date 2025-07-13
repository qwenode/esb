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
			name:    "valid field name",
			field:   "user.name",
			wantErr: false,
		},
		{
			name:    "simple field name",
			field:   "status",
			wantErr: false,
		},
		{
			name:    "nested field name",
			field:   "metadata.timestamp",
			wantErr: false,
		},
		{
			name:    "empty field name",
			field:   "",
			wantErr: false,
		},
		{
			name:    "whitespace field name",
			field:   "   ",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := NewQuery(Exists(tt.field))
			
			if tt.wantErr {
				if query.Exists == nil {
					t.Errorf("Exists() expected error, got nil")
				}
				return
			}
			
			if query.Exists == nil {
				t.Errorf("Exists() query.Exists is nil")
				return
			}
			
			if query.Exists.Field != tt.field {
				t.Errorf("Exists() field = %v, want %v", query.Exists.Field, tt.field)
			}
		})
	}
}

func TestExistsWithOtherQueries(t *testing.T) {
	// Test combining Exists with Bool query
	query := NewQuery(
		Bool(
			Must(
				Exists("user.name"),
				Term("status", "active"),
			),
		),
	)
	
	if query.Bool == nil {
		t.Errorf("Bool query is nil")
		return
	}
	
	if len(query.Bool.Must) != 2 {
		t.Errorf("Bool.Must length = %d, want 2", len(query.Bool.Must))
		return
	}
	
	// Check first Must clause is Exists
	if query.Bool.Must[0].Exists == nil {
		t.Errorf("First Must clause should be Exists query")
		return
	}
	
	if query.Bool.Must[0].Exists.Field != "user.name" {
		t.Errorf("Exists field = %v, want user.name", query.Bool.Must[0].Exists.Field)
	}
	
	// Check second Must clause is Term
	if query.Bool.Must[1].Term == nil {
		t.Errorf("Second Must clause should be Term query")
		return
	}
}



// Benchmark tests for Exists query
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
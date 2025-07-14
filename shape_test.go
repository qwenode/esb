package esb

import (
	"testing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geoshaperelation"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestShape(t *testing.T) {
	t.Run("test basic shape query", func(t *testing.T) {
		query := NewQuery(
			Shape("location", "POINT(13.0 53.0)"),
		)

		if query.Shape == nil {
			t.Error("expected Shape to be set")
		}

		shapeQuery, exists := query.Shape.ShapeQuery["location"]
		if !exists {
			t.Error("expected location field to exist")
		}

		if string(shapeQuery.Shape) != "POINT(13.0 53.0)" {
			t.Errorf("expected shape to be 'POINT(13.0 53.0)', got %s", string(shapeQuery.Shape))
		}
	})
}

func TestShapeWithRelation(t *testing.T) {
	t.Run("test shape query with relation", func(t *testing.T) {
		query := NewQuery(
			ShapeWithRelation("location", "POINT(13.0 53.0)", geoshaperelation.Intersects),
		)

		if query.Shape == nil {
			t.Error("expected Shape to be set")
		}

		shapeQuery, exists := query.Shape.ShapeQuery["location"]
		if !exists {
			t.Error("expected location field to exist")
		}

		if string(shapeQuery.Shape) != "POINT(13.0 53.0)" {
			t.Errorf("expected shape to be 'POINT(13.0 53.0)', got %s", string(shapeQuery.Shape))
		}

		if shapeQuery.Relation == nil {
			t.Error("expected Relation to be set")
		}

		if *shapeQuery.Relation != geoshaperelation.Intersects {
			t.Errorf("expected relation to be 'intersects', got %s", shapeQuery.Relation.String())
		}
	})

	t.Run("test shape query with disjoint relation", func(t *testing.T) {
		query := NewQuery(
			ShapeWithRelation("location", "POINT(13.0 53.0)", geoshaperelation.Disjoint),
		)

		shapeQuery, exists := query.Shape.ShapeQuery["location"]
		if !exists {
			t.Error("expected location field to exist")
		}

		if *shapeQuery.Relation != geoshaperelation.Disjoint {
			t.Errorf("expected relation to be 'disjoint', got %s", shapeQuery.Relation.String())
		}
	})

	t.Run("test shape query with within relation", func(t *testing.T) {
		query := NewQuery(
			ShapeWithRelation("location", "POINT(13.0 53.0)", geoshaperelation.Within),
		)

		shapeQuery, exists := query.Shape.ShapeQuery["location"]
		if !exists {
			t.Error("expected location field to exist")
		}

		if *shapeQuery.Relation != geoshaperelation.Within {
			t.Errorf("expected relation to be 'within', got %s", shapeQuery.Relation.String())
		}
	})

	t.Run("test shape query with contains relation", func(t *testing.T) {
		query := NewQuery(
			ShapeWithRelation("location", "POINT(13.0 53.0)", geoshaperelation.Contains),
		)

		shapeQuery, exists := query.Shape.ShapeQuery["location"]
		if !exists {
			t.Error("expected location field to exist")
		}

		if *shapeQuery.Relation != geoshaperelation.Contains {
			t.Errorf("expected relation to be 'contains', got %s", shapeQuery.Relation.String())
		}
	})

	// 可选：如需测试自定义 relation，可直接构造 geoshaperelation.GeoShapeRelation{Name: "custom"}
	// t.Run("test shape query with custom relation", func(t *testing.T) {
	// 	customRel := geoshaperelation.GeoShapeRelation{Name: "custom"}
	// 	query := NewQuery(
	// 		ShapeWithRelation("location", "POINT(13.0 53.0)", customRel),
	// 	)
	// 	shapeQuery, exists := query.Shape.ShapeQuery["location"]
	// 	if !exists {
	// 		t.Error("expected location field to exist")
	// 	}
	// 	if shapeQuery.Relation.String() != "custom" {
	// 		t.Errorf("expected relation to be 'custom', got %s", shapeQuery.Relation.String())
	// 	}
	// })
}

func TestShapeWithIndexedShape(t *testing.T) {
	t.Run("test shape query with indexed shape", func(t *testing.T) {
		query := NewQuery(
			ShapeWithIndexedShape("location", "shapes", "deu", "location"),
		)

		if query.Shape == nil {
			t.Error("expected Shape to be set")
		}

		shapeQuery, exists := query.Shape.ShapeQuery["location"]
		if !exists {
			t.Error("expected location field to exist")
		}

		if shapeQuery.IndexedShape == nil {
			t.Error("expected IndexedShape to be set")
		}

		if shapeQuery.IndexedShape.Index == nil || *shapeQuery.IndexedShape.Index != "shapes" {
			t.Errorf("expected index to be 'shapes', got %v", shapeQuery.IndexedShape.Index)
		}

		if shapeQuery.IndexedShape.Id != "deu" {
			t.Errorf("expected id to be 'deu', got %s", shapeQuery.IndexedShape.Id)
		}

		if shapeQuery.IndexedShape.Path == nil || *shapeQuery.IndexedShape.Path != "location" {
			t.Errorf("expected path to be 'location', got %v", shapeQuery.IndexedShape.Path)
		}
	})
}

func TestShapeWithOptions(t *testing.T) {
	t.Run("test shape query with options", func(t *testing.T) {
		query := NewQuery(
			ShapeWithOptions("location", "POINT(13.0 53.0)", func(shapeQuery *types.ShapeQuery, fieldQuery *types.ShapeFieldQuery) {
				relation := geoshaperelation.Intersects
				fieldQuery.Relation = &relation
				ignoreUnmapped := true
				shapeQuery.IgnoreUnmapped = &ignoreUnmapped
			}),
		)

		if query.Shape == nil {
			t.Error("expected Shape to be set")
		}

		shapeFieldQuery, exists := query.Shape.ShapeQuery["location"]
		if !exists {
			t.Error("expected location field to exist")
		}

		if string(shapeFieldQuery.Shape) != "POINT(13.0 53.0)" {
			t.Errorf("expected shape to be 'POINT(13.0 53.0)', got %s", string(shapeFieldQuery.Shape))
		}

		if shapeFieldQuery.Relation == nil {
			t.Error("expected Relation to be set")
		}

		if *shapeFieldQuery.Relation != geoshaperelation.Intersects {
			t.Errorf("expected relation to be 'intersects', got %s", shapeFieldQuery.Relation.String())
		}

		if query.Shape.IgnoreUnmapped == nil || *query.Shape.IgnoreUnmapped != true {
			t.Errorf("expected IgnoreUnmapped to be true, got %v", query.Shape.IgnoreUnmapped)
		}
	})

	t.Run("test shape query with disjoint relation and options", func(t *testing.T) {
		query := NewQuery(
			ShapeWithOptions("location", "POINT(13.0 53.0)", func(shapeQuery *types.ShapeQuery, fieldQuery *types.ShapeFieldQuery) {
				relation := geoshaperelation.Disjoint
				fieldQuery.Relation = &relation
				ignoreUnmapped := false
				shapeQuery.IgnoreUnmapped = &ignoreUnmapped
			}),
		)

		shapeFieldQuery, exists := query.Shape.ShapeQuery["location"]
		if !exists {
			t.Error("expected location field to exist")
		}

		if *shapeFieldQuery.Relation != geoshaperelation.Disjoint {
			t.Errorf("expected relation to be 'disjoint', got %s", shapeFieldQuery.Relation.String())
		}

		if query.Shape.IgnoreUnmapped == nil || *query.Shape.IgnoreUnmapped != false {
			t.Errorf("expected IgnoreUnmapped to be false, got %v", query.Shape.IgnoreUnmapped)
		}
	})

	t.Run("test shape query with within relation and options", func(t *testing.T) {
		query := NewQuery(
			ShapeWithOptions("location", "POINT(13.0 53.0)", func(shapeQuery *types.ShapeQuery, fieldQuery *types.ShapeFieldQuery) {
				relation := geoshaperelation.Within
				fieldQuery.Relation = &relation
			}),
		)

		shapeFieldQuery, exists := query.Shape.ShapeQuery["location"]
		if !exists {
			t.Error("expected location field to exist")
		}

		if *shapeFieldQuery.Relation != geoshaperelation.Within {
			t.Errorf("expected relation to be 'within', got %s", shapeFieldQuery.Relation.String())
		}
	})

	t.Run("test shape query with contains relation and options", func(t *testing.T) {
		query := NewQuery(
			ShapeWithOptions("location", "POINT(13.0 53.0)", func(shapeQuery *types.ShapeQuery, fieldQuery *types.ShapeFieldQuery) {
				relation := geoshaperelation.Contains
				fieldQuery.Relation = &relation
			}),
		)

		shapeFieldQuery, exists := query.Shape.ShapeQuery["location"]
		if !exists {
			t.Error("expected location field to exist")
		}

		if *shapeFieldQuery.Relation != geoshaperelation.Contains {
			t.Errorf("expected relation to be 'contains', got %s", shapeFieldQuery.Relation.String())
		}
	})

	t.Run("test shape query with custom relation and options", func(t *testing.T) {
		query := NewQuery(
			ShapeWithOptions("location", "POINT(13.0 53.0)", func(shapeQuery *types.ShapeQuery, fieldQuery *types.ShapeFieldQuery) {
				customRelation := geoshaperelation.GeoShapeRelation{Name: "custom"}
				fieldQuery.Relation = &customRelation
			}),
		)

		shapeFieldQuery, exists := query.Shape.ShapeQuery["location"]
		if !exists {
			t.Error("expected location field to exist")
		}

		if shapeFieldQuery.Relation.String() != "custom" {
			t.Errorf("expected relation to be 'custom', got %s", shapeFieldQuery.Relation.String())
		}
	})

	t.Run("test shape query with nil setOpts", func(t *testing.T) {
		query := NewQuery(
			ShapeWithOptions("location", "POINT(13.0 53.0)", nil),
		)

		if query.Shape == nil {
			t.Error("expected Shape to be set")
		}

		shapeFieldQuery, exists := query.Shape.ShapeQuery["location"]
		if !exists {
			t.Error("expected location field to exist")
		}

		if string(shapeFieldQuery.Shape) != "POINT(13.0 53.0)" {
			t.Errorf("expected shape to be 'POINT(13.0 53.0)', got %s", string(shapeFieldQuery.Shape))
		}

		// 应该没有设置其他选项
		if shapeFieldQuery.Relation != nil {
			t.Error("expected Relation to be nil")
		}

		if query.Shape.IgnoreUnmapped != nil {
			t.Error("expected IgnoreUnmapped to be nil")
		}
	})
} 
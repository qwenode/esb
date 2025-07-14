package esb

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geoshaperelation"
)

// Shape 创建一个形状查询，用于查找与指定形状相交的几何形状。
//
// 示例：
//   esb.Shape("location", "POINT(13.0 53.0)")
func Shape(field string, shape string) QueryOption {
	return func(q *types.Query) {
		query := types.NewShapeQuery()
		query.ShapeQuery[field] = types.ShapeFieldQuery{
			Shape: json.RawMessage(shape),
		}
		q.Shape = query
	}
}

// ShapeWithRelation 创建一个带有空间关系的形状查询，relation 参数为强类型。
//
// 示例：
//   esb.ShapeWithRelation("location", "POINT(13.0 53.0)", geoshaperelation.Intersects)
func ShapeWithRelation(field string, shape string, relation geoshaperelation.GeoShapeRelation) QueryOption {
	return func(q *types.Query) {
		query := types.NewShapeQuery()
		query.ShapeQuery[field] = types.ShapeFieldQuery{
			Shape:    json.RawMessage(shape),
			Relation: &relation,
		}
		q.Shape = query
	}
}

// ShapeWithIndexedShape 创建一个基于索引中存储形状的查询。
//
// 示例：
//   esb.ShapeWithIndexedShape("location", "shapes", "deu", "location")
func ShapeWithIndexedShape(field string, index, id, path string) QueryOption {
	return func(q *types.Query) {
		query := types.NewShapeQuery()
		query.ShapeQuery[field] = types.ShapeFieldQuery{
			IndexedShape: &types.FieldLookup{
				Index: &index,
				Id:    id,
				Path:  &path,
			},
		}
		q.Shape = query
	}
}

// ShapeWithOptions 提供回调函数式的形状查询配置。
//
// 示例：
//   esb.ShapeWithOptions("location", "POINT(13.0 53.0)", func(shapeQuery *types.ShapeQuery, fieldQuery *types.ShapeFieldQuery) {
//       relation := geoshaperelation.Intersects
//       fieldQuery.Relation = &relation
//       ignoreUnmapped := true
//       shapeQuery.IgnoreUnmapped = &ignoreUnmapped
//   })
func ShapeWithOptions(field string, shape string, setOpts func(shapeQuery *types.ShapeQuery, fieldQuery *types.ShapeFieldQuery)) QueryOption {
	return func(q *types.Query) {
		shapeQuery := types.NewShapeQuery()
		fieldQuery := types.ShapeFieldQuery{
			Shape: json.RawMessage(shape),
		}
		
		if setOpts != nil {
			setOpts(shapeQuery, &fieldQuery)
		}
		
		shapeQuery.ShapeQuery[field] = fieldQuery
		q.Shape = shapeQuery
	}
} 
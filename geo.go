// Package esb provides Elasticsearch query builder functionality
package esb

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// GeoBoundingBox 创建一个地理边界框查询，用于查找位于指定矩形区域内的地理点。
// 
// 示例：
//   esb.GeoBoundingBox("location", 40.73, -74.1, 40.01, -71.12)
func GeoBoundingBox(field string, topLeftLat, topLeftLon, bottomRightLat, bottomRightLon float64) QueryOption {
	return func(q *types.Query) {
		query := types.NewGeoBoundingBoxQuery()
		query.GeoBoundingBoxQuery[field] = types.TopLeftBottomRightGeoBounds{
			TopLeft: types.LatLonGeoLocation{
				Lat: types.Float64(topLeftLat),
				Lon: types.Float64(topLeftLon),
			},
			BottomRight: types.LatLonGeoLocation{
				Lat: types.Float64(bottomRightLat),
				Lon: types.Float64(bottomRightLon),
			},
		}
		q.GeoBoundingBox = query
	}
}

// GeoDistance 创建一个地理距离查询，用于查找距离指定点在给定距离内的地理点。
//
// 示例：
//   esb.GeoDistance("location", 40.0, -70.0, "200km")
func GeoDistance(field string, lat, lon float64, distance string) QueryOption {
	return func(q *types.Query) {
		query := types.NewGeoDistanceQuery()
		query.Distance = distance
		query.GeoDistanceQuery[field] = types.LatLonGeoLocation{
			Lat: types.Float64(lat),
			Lon: types.Float64(lon),
		}
		q.GeoDistance = query
	}
}

// GeoPolygon 创建一个地理多边形查询，用于查找位于指定多边形内的地理点。
//
// 示例：
//   points := [][]float64{{40, -70}, {30, -80}, {20, -90}}
//   esb.GeoPolygon("location", points)
func GeoPolygon(field string, points [][]float64) QueryOption {
	return func(q *types.Query) {
		query := types.NewGeoPolygonQuery()
		geoPoints := make([]types.GeoLocation, len(points))
		for i, point := range points {
			if len(point) >= 2 {
				geoPoints[i] = types.LatLonGeoLocation{
					Lat: types.Float64(point[0]),
					Lon: types.Float64(point[1]),
				}
			}
		}
		
		query.GeoPolygonQuery[field] = types.GeoPolygonPoints{
			Points: geoPoints,
		}
		q.GeoPolygon = query
	}
}

// GeoShape 创建一个地理形状查询，用于查找与指定形状相交的地理形状。
//
// 示例：
//   esb.GeoShape("location", "POINT(13.0 53.0)")
func GeoShape(field string, shape string) QueryOption {
	return func(q *types.Query) {
		query := types.NewGeoShapeQuery()
		query.GeoShapeQuery[field] = types.GeoShapeFieldQuery{
			Shape: []byte(shape),
		}
		q.GeoShape = query
	}
} 
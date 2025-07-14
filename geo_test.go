package esb

import (
	"testing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestGeoBoundingBox(t *testing.T) {
	t.Run("test geo bounding box query", func(t *testing.T) {
		query := NewQuery(
			GeoBoundingBox("location", 40.73, -74.1, 40.01, -71.12),
		)

		if query.GeoBoundingBox == nil {
			t.Error("expected GeoBoundingBox to be set")
		}

		bounds, exists := query.GeoBoundingBox.GeoBoundingBoxQuery["location"]
		if !exists {
			t.Error("expected location field to exist")
		}

		// 类型断言为 TopLeftBottomRightGeoBounds
		if topLeftBottomRight, ok := bounds.(types.TopLeftBottomRightGeoBounds); ok {
			if topLeftLocation, ok := topLeftBottomRight.TopLeft.(types.LatLonGeoLocation); ok {
				if float64(topLeftLocation.Lat) != 40.73 {
					t.Errorf("expected TopLeft.Lat to be 40.73, got %f", float64(topLeftLocation.Lat))
				}
				if float64(topLeftLocation.Lon) != -74.1 {
					t.Errorf("expected TopLeft.Lon to be -74.1, got %f", float64(topLeftLocation.Lon))
				}
			} else {
				t.Error("expected TopLeft to be LatLonGeoLocation")
			}

			if bottomRightLocation, ok := topLeftBottomRight.BottomRight.(types.LatLonGeoLocation); ok {
				if float64(bottomRightLocation.Lat) != 40.01 {
					t.Errorf("expected BottomRight.Lat to be 40.01, got %f", float64(bottomRightLocation.Lat))
				}
				if float64(bottomRightLocation.Lon) != -71.12 {
					t.Errorf("expected BottomRight.Lon to be -71.12, got %f", float64(bottomRightLocation.Lon))
				}
			} else {
				t.Error("expected BottomRight to be LatLonGeoLocation")
			}
		} else {
			t.Error("expected bounds to be TopLeftBottomRightGeoBounds")
		}
	})
}

func TestGeoDistance(t *testing.T) {
	t.Run("test geo distance query", func(t *testing.T) {
		query := NewQuery(
			GeoDistance("location", 40.0, -70.0, "200km"),
		)

		if query.GeoDistance == nil {
			t.Error("expected GeoDistance to be set")
		}

		locationField, exists := query.GeoDistance.GeoDistanceQuery["location"]
		if !exists {
			t.Error("expected location field to exist")
		}

		if latLonLocation, ok := locationField.(types.LatLonGeoLocation); ok {
			if float64(latLonLocation.Lat) != 40.0 {
				t.Errorf("expected Lat to be 40.0, got %f", float64(latLonLocation.Lat))
			}
			if float64(latLonLocation.Lon) != -70.0 {
				t.Errorf("expected Lon to be -70.0, got %f", float64(latLonLocation.Lon))
			}
		} else {
			t.Error("expected location to be LatLonGeoLocation")
		}

		if query.GeoDistance.Distance != "200km" {
			t.Errorf("expected Distance to be '200km', got %s", query.GeoDistance.Distance)
		}
	})
}

func TestGeoPolygon(t *testing.T) {
	t.Run("test geo polygon query", func(t *testing.T) {
		points := [][]float64{{40, -70}, {30, -80}, {20, -90}}
		query := NewQuery(
			GeoPolygon("location", points),
		)

		if query.GeoPolygon == nil {
			t.Error("expected GeoPolygon to be set")
		}

		polygonPoints, exists := query.GeoPolygon.GeoPolygonQuery["location"]
		if !exists {
			t.Error("expected location field to exist")
		}

		if len(polygonPoints.Points) != 3 {
			t.Errorf("expected 3 points, got %d", len(polygonPoints.Points))
		}

		if firstPoint, ok := polygonPoints.Points[0].(types.LatLonGeoLocation); ok {
			if float64(firstPoint.Lat) != 40 {
				t.Errorf("expected first point Lat to be 40, got %f", float64(firstPoint.Lat))
			}
			if float64(firstPoint.Lon) != -70 {
				t.Errorf("expected first point Lon to be -70, got %f", float64(firstPoint.Lon))
			}
		} else {
			t.Error("expected first point to be LatLonGeoLocation")
		}
	})
}

func TestGeoShape(t *testing.T) {
	t.Run("test geo shape query", func(t *testing.T) {
		query := NewQuery(
			GeoShape("location", "POINT(13.0 53.0)"),
		)

		if query.GeoShape == nil {
			t.Error("expected GeoShape to be set")
		}

		shapeQuery, exists := query.GeoShape.GeoShapeQuery["location"]
		if !exists {
			t.Error("expected location field to exist")
		}

		if string(shapeQuery.Shape) != "POINT(13.0 53.0)" {
			t.Errorf("expected Shape to be 'POINT(13.0 53.0)', got %s", string(shapeQuery.Shape))
		}
	})
} 
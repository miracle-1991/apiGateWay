package service

import (
	"context"
	"errors"
	"github.com/Willyham/hashfill"
	echo "github.com/miracle-1991/apiGateWay/server/echo/proto"
	geom "github.com/twpayne/go-geom"
)

type Service interface {
	FillGeoHash(context.Context, *echo.FillGeoHashRequest) (*echo.FillGeoHashResponse, error)
}

type GeosService struct{}

func (g *GeosService) FillGeoHash(ctx context.Context, request *echo.FillGeoHashRequest) (*echo.FillGeoHashResponse, error) {
	boundary := request.GetBoundary()
	if boundary == nil {
		return nil, errors.New("Invalid Boundary")
	}
	geoFence := convertBoundaryToGeoFence(boundary)
	precision := request.GetPrecision()
	filler := hashfill.NewRecursiveFiller(hashfill.WithMaxPrecision(int(precision)))
	hashes, err := filler.Fill(geoFence, hashfill.FillIntersects)
	if err != nil {
		return nil, errors.New("Fill Fail")
	}
	return &echo.FillGeoHashResponse{GeoHash: hashes}, nil
}

func convertBoundaryToGeoFence(boundary *echo.MultiPolygon) *geom.Polygon {
	var outerRing []geom.Coord
	var innerRing [][]geom.Coord
	for i, polygon := range boundary.Polygons {
		coord := []geom.Coord{}
		for _, point := range polygon.Vertices {
			coord = append(coord, geom.Coord{point.Lon, point.Lat})
		}
		coord = append(coord, coord[0])
		if i == 0 {
			outerRing = coord
		} else {
			innerRing = append(innerRing, coord)
		}
	}
	geofence := geom.NewPolygon(geom.XY).MustSetCoords(append([][]geom.Coord{outerRing}, innerRing...))
	return geofence
}

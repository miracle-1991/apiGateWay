package service

import (
	"context"
	"errors"
	"fmt"
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
	geoFence, err := convertBoundaryToGeoFence(boundary)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Parse Boundary Fail: %v", err))
	}

	precision := request.GetPrecision()
	filler := hashfill.NewRecursiveFiller(hashfill.WithMaxPrecision(int(precision)))
	hashes, err := filler.Fill(geoFence, hashfill.FillIntersects)
	if err != nil {
		return nil, errors.New("Fill Fail")
	}
	return &echo.FillGeoHashResponse{GeoHash: hashes}, nil
}

func convertBoundaryToGeoFence(boundary *echo.MultiPolygon) (*geom.Polygon, error) {
	geofence := &geom.Polygon{}
	coords := [][]geom.Coord{}
	for _, polygon := range boundary.Polygons {
		coord := []geom.Coord{}
		for _, point := range polygon.Vertices {
			coord = append(coord, geom.Coord{point.Lon, point.Lat})
		}
		coords = append(coords, coord)
	}
	p, err := geofence.SetCoords(coords)
	if err != nil {
		return nil, err
	}
	return p, nil
}

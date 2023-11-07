package service

import (
	"context"
	"errors"
	"github.com/Willyham/hashfill"
	echo "github.com/miracle-1991/apiGateWay/server/echo/proto"
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
		return nil, errors.New("Parse Boundary Fail")
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
	for _, polygon := range boundary.Polygons {
		coords := []geom.Coord{}
		for _, point := range polygon.Vertices {
			coords = append(coords, geom.Coord{point.Lon, point.Lat})
		}
		ring, err := geom.NewLinearRing(geom.XY).SetCoords(coords)
		if err != nil {
			return nil, err
		}
		err = geofence.Push(ring)
		if err != nil {
			return nil, err
		}
	}
	return geofence, nil
}

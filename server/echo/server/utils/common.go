package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/Willyham/hashfill"
	"github.com/miracle-1991/apiGateWay/server/echo/observable/trace"
	echo "github.com/miracle-1991/apiGateWay/server/echo/proto"
	gogeom "github.com/twpayne/go-geom"
)

func convertBoundaryToGeoFence(boundary *echo.MultiPolygon) *gogeom.Polygon {
	var outerRing []gogeom.Coord
	var innerRing [][]gogeom.Coord
	for i, polygon := range boundary.Polygons {
		coord := []gogeom.Coord{}
		for _, point := range polygon.Vertices {
			coord = append(coord, gogeom.Coord{point.Lon, point.Lat})
		}
		coord = append(coord, coord[0])
		if i == 0 {
			outerRing = coord
		} else {
			innerRing = append(innerRing, coord)
		}
	}
	geofence := gogeom.NewPolygon(gogeom.XY).MustSetCoords(append([][]gogeom.Coord{outerRing}, innerRing...))
	return geofence
}

func FillGeoHash(ctx context.Context, request *echo.FillGeoHashRequest) (*echo.FillGeoHashResponse, error) {
	ctx, span := trace.Tracer.Start(ctx, "utils-fillGeoHash")
	defer span.End()

	boundary := request.GetBoundary()
	if boundary == nil {
		err := errors.New("Invalid Boundary")
		span.RecordError(err)
		return nil, err
	}

	geoFence := convertBoundaryToGeoFence(boundary)
	precision := request.GetPrecision()
	filler := hashfill.NewRecursiveFiller(hashfill.WithMaxPrecision(int(precision)))
	hashes, err := filler.Fill(geoFence, hashfill.FillIntersects)
	if err != nil {
		err = errors.New(fmt.Sprintf("FillFail: %v", err))
		span.RecordError(err)
		return nil, err
	}

	return &echo.FillGeoHashResponse{GeoHash: hashes}, nil
}

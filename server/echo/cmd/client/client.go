package main

import (
	"context"
	"fmt"
	echo "github.com/miracle-1991/apiGateWay/server/echo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("gRPC dial err:", err)
	}
	defer conn.Close()

	client := echo.NewGeoServiceClient(conn)
	resp, err := client.FillGeoHash(context.Background(), &echo.FillGeoHashRequest{
		BoundaryName: "beijing",
		Precision:    6,
		Boundary: &echo.MultiPolygon{
			Polygons: []*echo.Polygon{
				&echo.Polygon{
					Vertices: []*echo.Point{
						&echo.Point{Lat: 30.725932779733828, Lon: 104.04810399201176},
						&echo.Point{Lat: 30.718933590173382, Lon: 104.10548984714013},
						&echo.Point{Lat: 30.712244330215963, Lon: 104.13622682435815},
						&echo.Point{Lat: 30.702910389863735, Lon: 104.15873257503482},
						&echo.Point{Lat: 30.66954989275946, Lon: 104.16300567773708},
						&echo.Point{Lat: 30.611560542033516, Lon: 104.14115260383586},
						&echo.Point{Lat: 30.60368261457832, Lon: 104.13460495126948},
						&echo.Point{Lat: 30.597337315209458, Lon: 104.12197165037554},
						&echo.Point{Lat: 30.594520110636584, Lon: 104.09585145457078},
						&echo.Point{Lat: 30.597848397395246, Lon: 104.08054653476084},
						&echo.Point{Lat: 30.59897710829587, Lon: 104.04449447431418},
						&echo.Point{Lat: 30.626768854130056, Lon: 104.0006597543566},
						&echo.Point{Lat: 30.652212454847685, Lon: 103.98381719828197},
						&echo.Point{Lat: 30.67412880125752, Lon: 103.99050661667191},
						&echo.Point{Lat: 30.701291076077624, Lon: 104.00698694578595},
						&echo.Point{Lat: 30.70704013040491, Lon: 104.02065715840023},
						&echo.Point{Lat: 30.725932779733828, Lon: 104.04810399201176}},
				},
			},
		},
	})
	if err != nil {
		fmt.Printf("fail, error:%v\n", err)
		return
	}
	fmt.Printf("success: %v\n", resp)
	return
}

package main

import (
    "context"
    "fmt"
    grpcclient "techdash/internal/transport/grpc"
    "techdash/internal/domain"
)

func main() {
    client, err := grpcclient.NewAnalyticsClient("localhost:50051")
    if err != nil {
        panic(err)
    }
    points := []domain.GeoPoint{
        {Lat: 51.169392, Lon: 71.449074},
        {Lat: 51.17, Lon: 71.45},
        {Lat: 51.168, Lon: 71.448},
    }

    result, err := client.GetHeatmap(context.Background(), points)
    if err != nil {
        fmt.Println("error:", err)
        return
    }
    fmt.Printf("result=%#v\n", result)
}

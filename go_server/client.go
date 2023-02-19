package main

import (
    "go_server/proto/aiservice"
    "google.golang.org/grpc"
    "golang.org/x/net/context"
    "google.golang.org/grpc/grpclog"
    "fmt"
)

const (
    // 远端gRPC服务地址
    Address = "127.0.0.1:50002"
)


func main() {
    conn, err := grpc.Dial(Address, grpc.WithInsecure())
    if err != nil {
        grpclog.Fatalln(err)
    }
    defer conn.Close()

    c := aiservice.NewAiServiceClient(conn)

    req := &aiservice.PredictRequest{Image:"This is base64Image!"}
    res, err := c.Predict(context.Background(), req)
    if err != nil {
        grpclog.Fatalln(err)
    }

    fmt.Println(res.Result)
}



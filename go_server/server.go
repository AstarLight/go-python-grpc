package main

// server

import (
    "fmt"
    "net"
    "go_server/proto/aiservice"
    "google.golang.org/grpc"
    "golang.org/x/net/context"
    "google.golang.org/grpc/grpclog"
)

const (
    // 本服务地址
    Address = "127.0.0.1:9988"
)

type gateService struct {
    aiservice.UnimplementedGateServiceServer
}

func (h gateService) Report(ctx context.Context, req *aiservice.ReportRequest) (*aiservice.ReportResponse, error) {
	
    resp := new(aiservice.ReportResponse)
    resp.Status = fmt.Sprintf("recv %s %s %s. Go server reply ok", req.Address, req.CpuLoad, req.MemLoad)

    return resp, nil
}


func main() {
    listen, err := net.Listen("tcp", Address)
    if err != nil {
        grpclog.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()

    aiservice.RegisterGateServiceServer(s, &gateService{})
    fmt.Println("Listen on " + Address)
    grpclog.Println("Listen on " + Address)
    s.Serve(listen)
}

package service

import (
	"context"

	"sample.com/grpc-server-with-go/pb"
)

type SampleServiceServer struct {
}

func NewSampleServiceServer() *SampleServiceServer {
	return &SampleServiceServer{}
}
func (c *SampleServiceServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	var sum int32 = 0
	for _, value := range req.Values {
		sum += value
	}
	return &pb.AddResponse{
		Sum: sum,
	}, nil
}

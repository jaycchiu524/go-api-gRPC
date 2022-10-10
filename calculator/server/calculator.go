package main

import (
	"context"

	pb "github.com/jaycchiu524/go-api-grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{
		Sum: in.A + in.B,
	}, nil
}
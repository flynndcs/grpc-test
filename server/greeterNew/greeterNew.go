package greeterNew

import (
	"context"

	pb "grpc-test/proto"
)

type GreeterNew struct {
	pb.UnimplementedGreeterNewServer
}

func (s *GreeterNew) SayHello(ctx context.Context, in *pb.HelloRequestNew) (*pb.HelloReplyNew, error) {
	return &pb.HelloReplyNew{Message: "Hello new " + in.GetName()}, nil
}

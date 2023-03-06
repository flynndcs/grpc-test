package greeter

import (
	"context"
	"log"

	pb "grpc-test/proto"
)

// Greeter is used to implement helloworld.GreeterServer.
type Greeter struct {
	New pb.GreeterNewServer
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *Greeter) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	inNew := pb.HelloRequestNew{Name: in.Name}
	response, err := s.New.SayHello(ctx, &inNew)
	if err != nil {
		log.Printf("error from external call to new greeter")
	}
	return &pb.HelloReply{Message: response.Message}, nil

}

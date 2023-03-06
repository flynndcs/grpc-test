package greeter

import (
	"context"
	"log"

	gen "grpc-test/gen"
)

// Greeter is used to implement helloworld.GreeterServer.
type Greeter struct {
	New gen.GreeterNewServer
	gen.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *Greeter) SayHello(ctx context.Context, in *gen.HelloRequest) (*gen.HelloReply, error) {
	inNew := gen.HelloRequestNew{Name: in.Name}
	response, err := s.New.SayHello(ctx, &inNew)
	if err != nil {
		log.Printf("error from external call to new greeter")
	}
	return &gen.HelloReply{Message: response.Message}, nil

}

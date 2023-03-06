package greeterNew

import (
	"context"

	gen "grpc-test/gen"
)

type GreeterNew struct {
	gen.UnimplementedGreeterNewServer
}

func (s *GreeterNew) SayHello(ctx context.Context, in *gen.HelloRequestNew) (*gen.HelloReplyNew, error) {
	return &gen.HelloReplyNew{Message: "Hello new " + in.GetName()}, nil
}

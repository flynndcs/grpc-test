/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	gen "grpc-test/gen"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"grpc-test/server/greeter"
	"grpc-test/server/greeterNew"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	new := greeterNew.GreeterNew{}

	gen.RegisterGreeterServer(s, &greeter.Greeter{New: &new})
	gen.RegisterGreeterNewServer(s, &new)

	log.Printf("server listening at %v", lis.Addr())
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	log.Println("here")

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:50051",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	log.Println("dialed")

	if err != nil {
		log.Fatalln("failed to dial grpc")
	}

	gwmux := runtime.NewServeMux()

	err = gen.RegisterGreeterHandler(context.Background(), gwmux, conn)
	err = gen.RegisterGreeterNewHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("failed to register")
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gateway on 8090")
	log.Fatalln(gwServer.ListenAndServe())

}

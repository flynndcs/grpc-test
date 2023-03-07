
# GRPC Test
### Adapted from [this tutorial](https://grpc.io/docs/languages/go/quickstart)
---
## Installation

```
go mod download
```
- Install [protoc](https://grpc.io/docs/protoc-installation/), [protoc-gen-go](https://grpc.io/docs/languages/go/quickstart/#prerequisites), [protoc-gen-doc](https://github.com/pseudomuto/protoc-gen-doc), [protoc-gen-grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway#installation), [protoc-gen-openapiv2](https://github.com/grpc-ecosystem/grpc-gateway#installation)
---
## Description
- This repo defines a gRPC client/server structure, leveraging code generation via protoc and protoc-gen-go and documentation generation via protoc-gen-doc.
- gRPC is an evolution of RPC (remote procedure call) concepts existing pre-REST. Behavior is categorized by *services* that are comprised of RPCs which involve *messages* which are comprised of *types*.
- The `protoc-gen-go` plugin for `protoc` and `protoc` itself generate the service/message/type code defined by the Protocol Buffer definitions (.proto files) which define the services, messages, and types.  
- Among the generated code exist "stubs" for both client and server interactions. Implementing a server stub enables a server to be run.
- gRPC communicates over HTTP/2 and with protocol buffers, both of which provide notable performance improvements.


---
## Code Generation
### Golang - gRPC client/server, gRPC Gateway, OpenAPI spec/docs
```
protoc -I ./proto --grpc-gateway_out . \
    --grpc-gateway_opt paths=source_relative \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --openapiv2_out ./gen/openapiv2 \
    ./proto/gen/*.proto
```
- the gRPC gateway code generation requires the Google API protobufs (used for annotations) to be importable - they are included in this repo for convenience, but may need to be updated.
### Python gRPC client/server
```
python3 -m grpc_tools.protoc -I proto --python_out=. --pyi_out=. --grpc_python_out=. ./proto/gen/*.proto
```

### HTML Protobuf docs
```
protoc --doc_out=./docs --doc_opt=html,index.html proto/gen/*.proto
```
---
## Local Development
- First, run code generation for Golang and Python if you haven't already.
- Run server:
```
go run server/main.go
```
- Run test scripts, which each use the client generated for that language
```
go run test.go
```
```
python3 test.py
```
---
## JSON/REST gRPC Gateway
- This provides a HTTP/JSON reverse proxy to gRPC services - allowing this server to be accessed/tested via familiar REST calls.
- The `protoc` invocation for Golang above generates the gRPC Gateway code.
- The `server/main.go` file starts the gRPC server on one port, then starts gRPC Gateway on another that proxies to the first. 
    - gRPC: localhost:50051
    - gRPC gateway: localhost:8090
- We also generate OpenAPI definitions with the invocation above.
### In progress: client generation from OpenAPI definitions (likely using deepmap/oapi-codegen for go), OpenAPI v3
---
##
_The jank folder structure (/gen, /proto/gen, /gen/openapiv2/gen) is a 1. hack to mitigate https://github.com/protocolbuffers/protobuf/issues/1490 for python and 2. the supported behavior of the openapiv2 generator in using the location of the protocol buffers to determine the location of the output of the OpenAPI spec_

# GRPC 
## Go Server, Go/Python client
### Adapted from [this tutorial](https://grpc.io/docs/languages/go/quickstart)
---
## Installation

```
go mod download
```
- Install [protoc](https://grpc.io/docs/protoc-installation/), [protoc-gen-go](https://grpc.io/docs/languages/go/quickstart/#prerequisites), [protoc-gen-doc](https://github.com/pseudomuto/protoc-gen-doc)
---
## Description
- This repo defines a gRPC client/server structure, leveraging code generation via protoc and protoc-gen-go and documentation generation via protoc-gen-doc.
- gRPC is an evolution of RPC (remote procedure call) concepts existing pre-REST. Behavior is categorized by *services* that are comprised of RPCs which involve *messages* which are comprised of *types*.
- The `protoc-gen-go` plugin for `protoc` and `protoc` itself generate the service/message/type code defined by the Protocol Buffer definitions (.proto files) which define the services, messages, and types.  
- Among the generated code exist "stubs" for both client and server interactions. Implementing a server stub enables a server to be run.
- gRPC communicates over HTTP/2 and with protocol buffers, both of which provide notable performance improvements.


---
## Client/Server/Docs Generation
### Golang 
```
protoc --go_out=./gen --go-grpc_out=./gen ./proto/gen/*.proto
```
### Python 
```
python3 -m grpc_tools.protoc -I proto --python_out=. --pyi_out=. --grpc_python_out=. ./proto/gen/*.proto
```

### HTML docs
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
### In progress with [gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- This provides a reverse proxy to gRPC services - you give it JSON via HTTP, it passes it through to gRPC.
---
##
_The jank folder structure (/gen, /proto/gen) is a hack to mitigate https://github.com/protocolbuffers/protobuf/issues/1490 for python_
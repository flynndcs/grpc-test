import grpc
import greeter_pb2
import greeter_pb2_grpc
import greeterNew_pb2
import greeterNew_pb2_grpc

def run():
    # connect
    with grpc.insecure_channel('localhost:50051') as channel:
        # request to greeter
        stub = greeter_pb2_grpc.GreeterStub(channel)
        response = stub.SayHello(greeter_pb2.HelloRequest(name="my name"))
        print(response)

        #request to greeter new
        stubNew = greeterNew_pb2_grpc.GreeterNewStub(channel)
        response = stubNew.SayHello(greeterNew_pb2.HelloRequestNew(name="my name"))
        print(response)
if __name__ == '__main__':
    run()
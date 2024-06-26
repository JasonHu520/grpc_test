package main

import (
	"context"
	"log"
	"net"

	"github.com/cubefs/cubefs/blobstore/common/trace"
	"google.golang.org/grpc"

	test "github.com/grpc_test"
)

type Server struct {
	test.UnimplementedGreeterServer

	*grpc.Server
}

func (s *Server) SayHello(ctx context.Context, in *test.HelloRequest) (*test.HelloReply, error) {
	span := trace.SpanFromContextSafe(ctx)
	span.Infof("SayHello: %s", in)
	return &test.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		return
	}
	server := grpc.NewServer(grpc.ChainUnaryInterceptor())
	s := Server{Server: server}
	s.RegisterService(&test.Greeter_ServiceDesc, &s)
	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
		return
	}
}

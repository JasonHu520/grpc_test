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

	infos map[string]*test.Info
}

func (s *Server) SayHello(ctx context.Context, in *test.HelloRequest) (*test.HelloReply, error) {
	span := trace.SpanFromContextSafe(ctx)
	span.Infof("SayHello: %s", in)
	return &test.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *Server) GetInfo(ctx context.Context, in *test.HelloRequest) (*test.Info, error) {
	span := trace.SpanFromContextSafe(ctx)
	span.Infof("GetInfo: %s", in)
	return s.infos[in.Name], nil
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		return
	}
	server := grpc.NewServer(grpc.ChainUnaryInterceptor())
	m := make(map[string]*test.Info, 1)
	m["test"] = &test.Info{Name: "test", Age: 25}
	s := Server{Server: server, infos: m}
	s.RegisterService(&test.Greeter_ServiceDesc, &s)
	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
		return
	}
}

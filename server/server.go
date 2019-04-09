package server

import (
	"context"
	"fmt"

	"github.com/akkyie/grpc-echo/echo"
)

type Server struct {
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}

func (server *Server) UnaryEcho(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{Message: req.GetMessage()}, nil
}

func (server *Server) ServerStreamingEcho(*echo.EchoRequest, echo.Echo_ServerStreamingEchoServer) error {
	return fmt.Errorf("unimplemented")
}

func (server *Server) ClientStreamingEcho(echo.Echo_ClientStreamingEchoServer) error {
	return fmt.Errorf("unimplemented")
}

func (server *Server) BidirectionalStreamingEcho(echo.Echo_BidirectionalStreamingEchoServer) error {
	return fmt.Errorf("unimplemented")
}

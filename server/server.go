package server

import (
	"context"
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/akkyie/grpc-echo/echo"
)

type Server struct {
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}

func (server *Server) UnaryEcho(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	message := fmt.Sprintf("[%s] %s", hostname, req.GetMessage())
	return &echo.EchoResponse{Message: message}, nil
}

func (server *Server) ServerStreamingEcho(req *echo.EchoRequest, res echo.Echo_ServerStreamingEchoServer) error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	message := req.GetMessage()
	response := &echo.EchoResponse{}
	for i := 0; i <= utf8.RuneCountInString(message); i++ {
		response.Message = fmt.Sprintf("[%s] %s", hostname, string([]rune(message)[:i]))
		res.Send(response)
	}
	return nil
}

func (server *Server) ClientStreamingEcho(echo.Echo_ClientStreamingEchoServer) error {
	return fmt.Errorf("unimplemented")
}

func (server *Server) BidirectionalStreamingEcho(echo.Echo_BidirectionalStreamingEchoServer) error {
	return fmt.Errorf("unimplemented")
}

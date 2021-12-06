package actuator

import (
	context "context"

	empty "github.com/golang/protobuf/ptypes/empty"
	core "github.com/imhshekhar47/go-grpc-api/core"
)

type Server struct {
	config core.AppConfig
}

func NewServer(appConfig core.AppConfig) *Server {
	return &Server{
		config: appConfig,
	}
}

func (s *Server) GetHealth(context.Context, *empty.Empty) (*Health, error) {
	return &Health{
		Status: "UP",
	}, nil
}

func (s *Server) GetInfo(context.Context, *empty.Empty) (*Info, error) {
	return &Info{
		Name:    s.config.Application.Name,
		Version: s.config.Application.Version,
	}, nil
}

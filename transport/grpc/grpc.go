package grpc

import (
	"context"
	"net"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/YxTiBlya/ci-api/pkg/executor"
)

func New(cfg Config, svc Service) *Server {
	s := &Server{
		svc: svc,
		cfg: cfg,
		log: zap.Must(zap.NewDevelopment()).Sugar(),
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
				if r, ok := req.(*executor.ExecuteRequest); ok {
					s.log.Infof("grpc request %s repository: %s", info.FullMethod, r.Repo)
				}

				return handler(ctx, req)
			},
		),
	}
	s.srv = grpc.NewServer(opts...)

	executor.RegisterExecutorAPIServer(s.srv, s)
	return s
}

type Server struct {
	executor.UnimplementedExecutorAPIServer
	svc Service
	srv *grpc.Server
	cfg Config
	log *zap.SugaredLogger
}

func (s *Server) Start(ctx context.Context) error {
	conn, err := net.Listen("tcp", s.cfg.Address)
	if err != nil {
		return errors.Wrapf(err, "cannot listen %q", s.cfg.Address)
	}

	errCh := make(chan error)
	go func() {
		s.log.Infof("grpc start listening on %q", s.cfg.Address)
		if err := s.srv.Serve(conn); err != nil {
			errCh <- errors.Wrap(err, "cannot server connection")
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-time.After(s.cfg.StartTimeout):
		return nil
	}
}

func (s *Server) Stop(ctx context.Context) error {
	s.srv.GracefulStop()
	return nil
}

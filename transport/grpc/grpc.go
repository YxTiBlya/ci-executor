package grpc

import (
	"context"
	"net"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/YxTiBlya/ci-api/pkg/executor"
	"github.com/YxTiBlya/ci-core/logger"
)

func New(cfg Config, svc Service) *Server {
	s := &Server{
		svc: svc,
		cfg: cfg,
		log: logger.New("grpc"),
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
				if r, ok := req.(*executor.ExecuteRequest); ok {
					s.log.Info().Str("method", info.FullMethod).Str("repo", r.Repo).Msg("grpc request")
				}

				return handler(ctx, req)
			},
		),
	}
	s.srv = grpc.NewServer(opts...)

	executor.RegisterExecutorAPIServer(s.srv, s)
	reflection.Register(s.srv)

	return s
}

type Server struct {
	executor.UnimplementedExecutorAPIServer
	svc Service
	srv *grpc.Server
	cfg Config
	log *logger.Logger
}

func (s *Server) Start(ctx context.Context) error {
	conn, err := net.Listen("tcp", s.cfg.Address)
	if err != nil {
		return errors.Wrapf(err, "cannot listen %q", s.cfg.Address)
	}

	errCh := make(chan error)
	go func() {
		s.log.Info().Str("address", s.cfg.Address).Msg("starting grpc server")
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

package grpc

import (
	"context"

	"github.com/YxTiBlya/ci-api/pkg/executor"

	"github.com/YxTiBlya/ci-executor/internal/service/dto"
)

func (s *Server) ExecuteTask(ctx context.Context, in *executor.ExecuteRequest) (*executor.ExecuteResponse, error) {
	if err := in.Validate(); err != nil {
		s.log.Error().Err(err).Msg("failed to validate request")
		return nil, err
	}

	result := s.svc.ExecuteTask(ctx, dto.ExecuteTaskRequest{
		Repository: in.Repo,
		Command:    in.Cmd,
	})

	return result.ToGRPC(), nil
}

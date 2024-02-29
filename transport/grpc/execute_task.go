package grpc

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/YxTiBlya/ci-api/pkg/executor"

	"github.com/YxTiBlya/ci-executor/internal/service/dto"
)

func (s *Server) ExecuteTask(ctx context.Context, in *executor.ExecuteRequest) (*executor.ExecuteResponse, error) {
	if err := in.Validate(); err != nil {
		s.log.Error("failed to validate request", zap.Error(err))
		return nil, errors.Wrap(err, "failed to validate request")
	}

	result := s.svc.ExecuteTask(ctx, dto.ExecuteTaskRequest{
		Repository: in.Repo,
		Command:    in.Cmd,
	})

	return result.ToGRPC(), nil
}

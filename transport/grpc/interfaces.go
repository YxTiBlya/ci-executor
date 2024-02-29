package grpc

import (
	"context"

	"github.com/YxTiBlya/ci-executor/internal/service/dto"
)

type Service interface {
	ExecuteTask(ctx context.Context, req dto.ExecuteTaskRequest) *dto.ExecuteTaskResponse
}

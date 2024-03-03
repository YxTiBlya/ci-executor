package service

import (
	"context"
	"os/exec"
	"strings"
	"time"

	"github.com/YxTiBlya/ci-executor/internal/service/dto"
)

func (svc *Service) ExecuteTask(ctx context.Context, req dto.ExecuteTaskRequest) *dto.ExecuteTaskResponse {
	command := strings.Split(req.Command, " ")

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Dir = req.Repository

	start := time.Now()
	b, err := cmd.Output()
	if err != nil {
		if exitCode, ok := err.(*exec.ExitError); ok {
			svc.log.Error().
				Str("command", req.Command).
				Str("output", string(b)).
				Int("exitCode", exitCode.ExitCode()).
				Msg("failed to execute command")
			return dto.NewExecuteTaskResponse(dto.ExecuteFailedStatus, exitCode.ExitCode(), &b, &start)
		}

		svc.log.Error().Err(err).Msg("failed to execute command")
		return dto.NewExecuteTaskResponse(dto.ExecuteFailedStatus, 1, &b, &start)
	}

	return dto.NewExecuteTaskResponse(dto.ExecuteSuccessStatus, 0, &b, &start)
}

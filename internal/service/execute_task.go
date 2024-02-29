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
		exitCode := err.(*exec.ExitError)
		svc.log.Errorf("failed to execute command: %s\noutput: %s\nexit status: %d", req.Command, b, exitCode.ExitCode())
		return dto.NewExecuteTaskResponse(dto.ExecuteFailedStatus, exitCode.ExitCode(), &b, &start)
	}

	return dto.NewExecuteTaskResponse(dto.ExecuteSuccessStatus, 0, &b, &start)
}

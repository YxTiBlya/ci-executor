package dto

import (
	"time"

	"github.com/YxTiBlya/ci-api/pkg/executor"
)

const (
	ExecuteSuccessStatus = "success"
	ExecuteFailedStatus  = "failed"
)

type ExecuteTaskRequest struct {
	Repository string
	Command    string
}

type ExecuteTaskResponse struct {
	Status   string
	ExitCode int
	Output   string
	Time     float64
}

func NewExecuteTaskResponse(status string, code int, b *[]byte, startTime *time.Time) *ExecuteTaskResponse {
	return &ExecuteTaskResponse{
		Status:   status,
		ExitCode: code,
		Output:   string(*b),
		Time:     time.Since(*startTime).Seconds(),
	}
}

func (eResp *ExecuteTaskResponse) ToGRPC() *executor.ExecuteResponse {
	return &executor.ExecuteResponse{
		Status:   compareStatus(eResp.Status),
		ExitCode: int32(eResp.ExitCode),
		Output:   eResp.Output,
		Time:     eResp.Time,
	}
}

func compareStatus(status string) executor.ExecuteStatus {
	switch status {
	case ExecuteSuccessStatus:
		return executor.ExecuteStatus_SUCCESS
	case ExecuteFailedStatus:
		return executor.ExecuteStatus_FAILED
	}

	return -1
}

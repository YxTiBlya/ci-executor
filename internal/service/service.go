package service

import (
	"context"

	"github.com/YxTiBlya/ci-core/logger"
)

func New(cfg Config) *Service {
	return &Service{
		cfg: cfg,
		log: logger.New("service"),
	}
}

type Service struct {
	cfg Config
	log *logger.Logger
}

func (svc *Service) Start(ctx context.Context) error {
	return nil
}

func (svc *Service) Stop(ctx context.Context) error {
	svc.log.Sync()
	return nil
}

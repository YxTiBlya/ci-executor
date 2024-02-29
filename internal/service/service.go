package service

import (
	"context"

	"go.uber.org/zap"
)

type Relations struct {
}

func New(cfg Config, log *zap.SugaredLogger, rel Relations) *Service {
	return &Service{
		cfg:       cfg,
		log:       log,
		Relations: rel,
	}
}

type Service struct {
	cfg Config
	log *zap.SugaredLogger
	Relations
}

func (svc *Service) Start(ctx context.Context) error {

	return nil
}

func (svc *Service) Stop(ctx context.Context) error {
	return nil
}

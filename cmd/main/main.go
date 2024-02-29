package main

import (
	"flag"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"

	"github.com/YxTiBlya/ci-core/scheduler"

	"github.com/YxTiBlya/ci-executor/internal/service"
	"github.com/YxTiBlya/ci-executor/transport/grpc"
)

type Config struct {
	Executor struct {
		Service service.Config `yaml:"service"`
		GRPC    grpc.Config    `yaml:"grpc"`
	} `yaml:"executor"`
}

var cfgPath string

func init() {
	flag.StringVar(&cfgPath, "cfg", "config.yaml", "app cfg path")
	flag.Parse()
}
func main() {
	logger := zap.Must(zap.NewDevelopment()).Sugar()

	yamlFile, err := os.ReadFile(cfgPath)
	if err != nil {
		logger.Fatal("failed to open config file", zap.Error(err))
	}

	var cfg Config
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		logger.Fatal("failed to unmarshal config file", zap.Error(err))
	}

	svc := service.New(cfg.Executor.Service, logger, service.Relations{})

	grpc := grpc.New(cfg.Executor.GRPC, svc)

	sch := scheduler.NewScheduler(
		zap.Must(zap.NewDevelopment()).Sugar(),
		scheduler.NewComponent("service", svc),
		scheduler.NewComponent("grpc", grpc),
	)
	sch.Start()
}

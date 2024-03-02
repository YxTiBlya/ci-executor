package main

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/YxTiBlya/ci-core/logger"
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
	logger.Init(logger.DevelopmentConfig)
	flag.StringVar(&cfgPath, "cfg", "config.yaml", "app cfg path")
	flag.Parse()
}
func main() {
	yamlFile, err := os.ReadFile(cfgPath)
	if err != nil {
		log.Fatal("failed to open config file", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		log.Fatal("failed to unmarshal config file", err)
	}

	svc := service.New(cfg.Executor.Service)
	grpc := grpc.New(cfg.Executor.GRPC, svc)

	sch := scheduler.NewScheduler(
		scheduler.NewComponent("service", svc),
		scheduler.NewComponent("grpc", grpc),
	)
	sch.Start()
}

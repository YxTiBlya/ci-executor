package grpc

import "time"

type Config struct {
	Address      string        `yaml:"address" default:"0.0.0.0:8000"`
	StartTimeout time.Duration `yaml:"start_timeout" default:"5s"`
}

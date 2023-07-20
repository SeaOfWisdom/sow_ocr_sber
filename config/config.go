package config

import (
	"fmt"

	"github.com/namsral/flag"
)

const defaultRpcPort = 50051

type Config struct {
	GrpcPort uint64

	OpenAIApiKey   string
	VisionFilepath string
}

func NewConfig() *Config {
	config := &Config{}
	/* gRPC */
	flag.Uint64Var(&config.GrpcPort, "grpc-address", defaultRpcPort, "gRPC port for inter-service communications")
	/* Cron */
	flag.StringVar(&config.OpenAIApiKey, "openai-api-key", "", "")
	flag.StringVar(&config.VisionFilepath, "vision-credentials", "", "")

	/* parse config from envs or config files */
	flag.Parse()
	// config.verify()

	return config
}

func (c *Config) verify() {
	if c.OpenAIApiKey == "" {
		panic(fmt.Errorf("OpenAIApiKey is empty"))
	}
	if c.VisionFilepath == "" {
		panic(fmt.Errorf("VisionFilepath is empty"))
	}
}

package config

import (
	"encoding/json"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Server     ServerConfig
	BaconIpsum BaconIpsumConfig
}

type ServerConfig struct {
	GRPCPort string `envconfig:"GRPC_PORT" default:"50051"`
	RESTPort string `envconfig:"REST_PORT" default:"8080"`
}

type BaconIpsumConfig struct {
	URL string `envconfig:"BACON_IPSUM_URL"`
}

var config AppConfig

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load .env file: %v", err)
	}

	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("failed to process env var: %v", err)
	}

	log.Printf("[Config]\n%s\n", config.pretty())
}

func Get() *AppConfig {
	return &config
}

func (c *AppConfig) pretty() string {
	s, _ := json.MarshalIndent(c, "", "\t")
	return string(s)
}

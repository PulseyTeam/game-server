package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Config struct {
	AppVersion string
	Server     Server
	MongoDB    MongoDB
}

type Server struct {
	Port              string
	Development       bool
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	MaxConnectionIdle time.Duration
	MaxConnectionAge  time.Duration
}

type MongoDB struct {
	URI      string
	User     string
	Password string
	DB       string
}

func exportConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if os.Getenv("DOCKER_ENVIRONMENT") == "1" {
		viper.SetConfigName("config-docker.yml")
	} else {
		viper.SetConfigName("config.yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func ParseConfig() (*Config, error) {
	if err := exportConfig(); err != nil {
		return nil, err
	}

	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	gRPCPort := os.Getenv("GRPC_PORT")
	if gRPCPort != "" {
		c.Server.Port = gRPCPort
	}

	return &c, nil
}

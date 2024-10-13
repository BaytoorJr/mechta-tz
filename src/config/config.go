package config

import (
	"errors"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var MainConfig *Config

type Config struct {
	ExampleFilePath string `envconfig:"json_file_path" required:"true" default:"./example.json"`
	WorkersCount    int    `envconfig:"workers_count" required:"true" default:"4"`
}

func InitConfig() error {
	err := godotenv.Load("env_local.env")
	if err != nil {
		return err
	}

	var cfg Config

	err = envconfig.Process("", &cfg)
	if err != nil {
		return errors.New("could not process env variables")
	}

	MainConfig = &cfg
	return nil
}

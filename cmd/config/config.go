package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App         App
		Logger      Logger
		FeatureFlag FeatureFlag
		Datasource  Datasource
		Debug       bool `env:"DEBUG" env-default:"false"`
	}

	App struct {
		Environment string `env:"ENVIRONMENT" env-default:"local"`
		Name        string `env:"APP_NAME" env-default:"interview-crud-files"`
	}

	Logger struct {
		Level string `env:"LOGGING_LEVEL" env-default:"debug"`
	}

	Datasource struct {
		Path string `env:"DATASOURCE_PATH" env-default:"/datasource"`
	}

	FeatureFlag struct {
		DumpConfig bool `env:"FEATURE_FLAG_DUMP_CONFIG" env-default:"true"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	return cfg, err
}

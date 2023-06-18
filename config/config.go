package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App   App
		Debug bool `env:"DEBUG" env-default:"false"`

		HTTP       HTTP
		Postgresql Postgresql
		Datasource Datasource

		Trace       Trace
		Logger      Logger
		FeatureFlag FeatureFlag
	}

	App struct {
		Environment string `env:"ENVIRONMENT" env-default:"local"`
		Name        string `env:"APP_NAME" env-default:"interview-crud-files"`
	}

	HTTP struct {
		Port string `env:"HTTP_PORT" env-default:"8080"`
	}

	Postgresql struct {
		URL string `env:"POSTGRESQL_URL" env-default:"postgres://user:pass@localhost:5432/postgres?pool_max_conn_idle_time=30s"`
	}

	Datasource struct {
		Path string `env:"DATASOURCE_PATH" env-default:"/datasource"`
	}

	Trace struct {
		URL string `env:"TRACE_URL" env-default:"http://localhost:14268/api/traces"`
	}

	Logger struct {
		Level string `env:"LOGGING_LEVEL" env-default:"debug"`
	}

	FeatureFlag struct {
		DumpConfig   bool `env:"FEATURE_FLAG_DUMP_CONFIG" env-default:"true"`
		TraceEnabled bool `env:"FEATURE_FLAG_TRACE_ENABLED" env-default:"false"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	return cfg, err
}

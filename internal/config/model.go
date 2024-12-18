package config

import (
	"time"

	"github.com/Kroning/mytheresa/internal/transport/http"
)

type Config struct {
	App    App    `mapstructure:"app"`
	Server Server `mapstructure:"server"`
	DB     DB     `mapstructure:"db"`
}

type App struct {
	Env      string `mapstructure:"env"`
	Name     string `mapstructure:"name"`
	LogLevel string `mapstructure:"log_level" default:"debug"`
}

type DB struct {
	Master DBConfig `mapstructure:"master"`
}

type DBConfig struct {
	Host           string        `mapstructure:"host"`
	Port           string        `mapstructure:"port"`
	User           string        `mapstructure:"user"`
	Password       string        `mapstructure:"password"`
	Database       string        `mapstructure:"database"`
	MaxOpen        uint          `mapstructure:"max_open"`
	Timeout        time.Duration `mapstructure:"timeout"`
	MigrationsPath string        `mapstructure:"migrations_path"`
}

type Server struct {
	HTTP *http.Config `mapstructure:"http"`
}

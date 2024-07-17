package config

import (
	"fmt"
	"time"

	"github.com/jessevdk/go-flags"
)

const (
	// DefEnvProd define production mode
	DefEnvProd = "prod"
	// DefEnvDev define dev mode
	DefEnvDev = "dev"
	// DefEnvTest define test mode
	DefEnvTest = "test"
)

// Config is a main configuration struct for application
type Config struct {
	Environment       string             `long:"environment" env:"PLATFORM_ENVIRONMENT" default:"test"`
	Debug             bool               `long:"debug" env:"PLATFORM_DEBUG"`
	Timeout           int                `long:"timeout" env:"PLATFORM_TIMEOUT" default:"1000000"`
	MainBackendConfig *MainBackendConfig `group:"Main backend args" namespace:"mainbackend" env-namespace:"PLATFORM_MAIN_BACKEND"`
	Log               *LogConfig         `group:"Logger args" namespace:"logger" env-namespace:"PLATFORM_LOGGER"`
	Db                *Db                `group:"database args" namespace:"db" env-namespace:"PLATFORM_DATABASE"`
}

// Db struct contains database configuration
type Db struct {
	Host            string        `long:"host" env:"HOST" description:"Postgres host" required:"yes"`
	Port            string        `long:"port" env:"PORT" description:"Postgres port" required:"yes"`
	User            string        `long:"user" env:"USER" description:"Postgres user" required:"yes"`
	Password        string        `long:"password" env:"PASSWORD" description:"Postgres password" required:"yes"`
	Name            string        `long:"name" env:"NAME" description:"Postgres name" required:"yes"`
	MaxOpenConns    int           `long:"max-open-conns" env:"MAX_OPEN_CONNS" default:"25" description:"maximum of open database connections"`
	MaxIdleConns    int           `long:"max-idle-conns" env:"MAX_IDLE_CONNS" default:"10" description:"maximum of idle database connections"`
	ConnMaxLifeTime time.Duration `long:"conn-max-life-time" env:"CONN_MAX_LIFE_TIME" default:"5m" description:"database max connection life time"`

	MigrationsSourceURL string `long:"migrations-source-url" env:"MIGRATIONS_SOURCE_URL" default:"migrations"`
}

type LogConfig struct {
	Level string `short:"l" long:"level" env:"LEVEL" description:"Logger level" required:"yes" default:"error"` // std: trace, debug, info, warning, error, fatal, panic
}

type MainBackendConfig struct {
	Host string `long:"host" env:"HOST" description:"Main backend host" required:"yes"`
	Port string `long:"port" env:"PORT" description:"Main backend port" required:"yes"`
	Path string `long:"path" env:"PATH" description:"Main backend path for user info endpoint" required:"yes"`
}

func Parse() (*Config, error) {
	var config Config
	p := flags.NewParser(&config, flags.HelpFlag|flags.PassDoubleDash)

	_, err := p.ParseArgs([]string{})
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// IsProduction check that application in production mode
func (c *Config) IsProduction() bool {
	return c.Environment == DefEnvProd
}

// IsDevelopment check that application in development mode
func (c *Config) IsDevelopment() bool {
	return c.Environment == DefEnvDev
}

// IsTest check that application in test mode
func (c *Config) IsTest() bool {
	return c.Environment == DefEnvTest
}

func (c *Db) ConnectionString() string {
	uri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable default_query_exec_mode=cache_describe",
		c.Host, c.Port,
		c.User, c.Name,
		c.Password,
	)

	return uri
}

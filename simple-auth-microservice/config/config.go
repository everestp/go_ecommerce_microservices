package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

// App config struct
type Config struct {
	Server ServerConfig
	Postgres PostgresConfig
	Redis RedisConfig
	Logger Logger
	Jaeger Jaeger
	Session Session
	Metrics Metrics
	RabbitMQ RabbitMQ
	Resend Resend
}


// Server config struct
type ServerConfig struct {
	AppVersion string
	Port string
	PprofPort string
	Mode string
	JwtSecretKey string
	CookieName string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	SSL bool
	CtxDefaultTimeout time.Duration
	CSRF bool
	Debug bool
	MaxConnectionIdle time.Duration
	Timeout time.Duration
	MaxConnectionAge time.Duration
	Time time.Duration
}

// Postgresql config
type PostgresConfig struct {
	PostgresqlHost string
	PostgresqlPort string
	PostgresqlUser string
	PostgresqlPassword string
	PostgresqlDbName string
	PostgresqlSSLMode bool
	PgDriver string
}

// Redis config
type RedisConfig struct {
	RedisAddr string
	RedisPassword string
	RedisDB string
	RedisDefaultdb string
	MinIdleConns int
	PoolSize int
	PoolTimeout int
	Password string
	DB int
}

//Logger config
type Logger struct {
	Development bool
	DisableCaller bool
	DisableStacktrace bool
	Encoding string
	Level string
}

// Jaeger config
type Jaeger struct {
	Host string
	ServiceName string
	LogSpans bool
}

// Session config
type Session struct {
	Name string
	Prefix string
	Expire int
}

// Metrics config
type Metrics struct {
	URL string
	ServiceName string
}

// RabbitMQ config
type RabbitMQ struct {
	Host string
	Port string
	User string
	Password string
	Exchange string
	Queue string
	RoutingKey string
	ConsumerTag string
	WorkerPoolSize int
}

// Resend config
type Resend struct {
	ApiKey string
}


// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigFile(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
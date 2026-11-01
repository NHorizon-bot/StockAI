package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	defaultServerAddr         = ":8080"
	defaultGinMode            = "release"
	defaultMySQLHost          = "127.0.0.1"
	defaultMySQLPort          = "3307"
	defaultMySQLUser          = "stockai"
	defaultMySQLPassword      = "stockai_pass"
	defaultMySQLDatabase      = "stockai"
	defaultMySQLMigrationsDir = "./migrations"
	defaultRedisAddr          = "127.0.0.1:6380"
	defaultRedisPassword      = "stockai_redis"
	defaultRedisDB            = 0
)

type Config struct {
	Server ServerConfig
	MySQL  MySQLConfig
	Redis  RedisConfig
}

type ServerConfig struct {
	Addr    string
	GinMode string
}

type MySQLConfig struct {
	Host          string
	Port          string
	User          string
	Password      string
	Database      string
	MigrationsDir string
}

func (c MySQLConfig) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func Load() (Config, error) {
	redisDB, err := strconv.Atoi(envOrDefault("REDIS_DB", strconv.Itoa(defaultRedisDB)))
	if err != nil {
		return Config{}, fmt.Errorf("parse REDIS_DB: %w", err)
	}

	return Config{
		Server: ServerConfig{
			Addr:    envOrDefault("SERVER_ADDR", defaultServerAddr),
			GinMode: envOrDefault("GIN_MODE", defaultGinMode),
		},
		MySQL: MySQLConfig{
			Host:          envOrDefault("MYSQL_HOST", defaultMySQLHost),
			Port:          envOrDefault("MYSQL_PORT", defaultMySQLPort),
			User:          envOrDefault("MYSQL_USER", defaultMySQLUser),
			Password:      envOrDefault("MYSQL_PASSWORD", defaultMySQLPassword),
			Database:      envOrDefault("MYSQL_DATABASE", defaultMySQLDatabase),
			MigrationsDir: envOrDefault("MYSQL_MIGRATIONS_DIR", defaultMySQLMigrationsDir),
		},
		Redis: RedisConfig{
			Addr:     envOrDefault("REDIS_ADDR", defaultRedisAddr),
			Password: envOrDefault("REDIS_PASSWORD", defaultRedisPassword),
			DB:       redisDB,
		},
	}, nil
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

package settings

import (
	"github.com/spf13/viper"
	"log"
)

const (
	errorReadingConfigFile = "Error reading config file: %s"
)

var (
	viperConfig = viper.New()
	config      = Init()
)

func Init() *serviceConfig {
	viperConfig.AddConfigPath(".")
	viperConfig.AddConfigPath("./settings/")
	viperConfig.AddConfigPath("./internal/settings/")
	viperConfig.SetConfigName("settings")
	viperConfig.SetConfigType("yaml")

	setUpDefaults()

	if err := viperConfig.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalf("%s: %s", errorReadingConfigFile, err.Error())
		}
	}

	serverConfig := &ServerConfig{
		Host: viperConfig.GetString("server.host"),
	}
	postgresConfig := &PostgresConfig{
		DbName:     viperConfig.GetString("db.environment.POSTGRES_DB"),
		DbUser:     viperConfig.GetString("db.environment.POSTGRES_USER"),
		DbHost:     viperConfig.GetString("db.environment.POSTGRES_HOST"),
		DbPassword: viperConfig.GetString("db.environment.POSTGRES_PASSWORD"),
		DbPort:     viperConfig.GetInt("db.environment.POSTGRES_PORT"),
	}

	databaseConfig := &DatabaseConfig{Postgres: postgresConfig}

	return &serviceConfig{
		Server:   serverConfig,
		Database: databaseConfig,
	}
}

type serviceConfig struct {
	Server   *ServerConfig
	Database *DatabaseConfig
}

type ServerConfig struct {
	Host string
}

type DatabaseConfig struct {
	Postgres *PostgresConfig
}

type PostgresConfig struct {
	DbHost     string
	DbName     string
	DbUser     string
	DbPassword string
	DbPort     int
}

func setUpDefaults() {
	viperConfig.SetDefault("db.environment.POSTGRES_HOST", "postgres")
	viperConfig.SetDefault("db.environment.POSTGRES_PORT", 5432)
	viperConfig.SetDefault("db.environment.POSTGRES_USER", "driver-register")
	viperConfig.SetDefault("db.environment.POSTGRES_PASSWORD", "driver-register")
	viperConfig.SetDefault("db.environment.POSTGRES_DB", "driver-register")
	viperConfig.SetDefault("db.environment.POSTGRES_DB", "driver-register")
	viperConfig.SetDefault("server.host", ":8000")
	viperConfig.AutomaticEnv()
}

func GetDatabaseConfig() *DatabaseConfig {
	return config.Database
}

func GetServerConfig() *ServerConfig {
	return config.Server
}

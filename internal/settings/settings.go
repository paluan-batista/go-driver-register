package settings

import (
	"github.com/spf13/viper"
	"go-driver-register/internal/container"
)

const (
	errorReadingConfigFile = "Error reading config file: %s"
)

var (
	viperConfig = viper.New()
	config      = Init()
	logger      = container.GetLogger()
)

func Init() *serviceConfig {
	viperConfig.AddConfigPath(".")
	viperConfig.AddConfigPath("./settings/")
	viperConfig.AddConfigPath("./internal/settings/")
	viperConfig.SetConfigName("settings")
	viperConfig.SetConfigType("yaml")

	if err := viperConfig.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logger.Fatalf("%s: %s", errorReadingConfigFile, err.Error())
		}
	}

	serverConfig := &ServerConfig{
		Host: viperConfig.GetString("server.host"),
	}
	postgresConfig := &PostgresConfig{
		DbName:     viperConfig.GetString("postgres.db_name"),
		DbUser:     viperConfig.GetString("postgres.db_user"),
		DbHost:     viperConfig.GetString("postgres.db_host"),
		DbPassword: viperConfig.GetString("postgres.db_password"),
		DbPort:     viperConfig.GetInt("postgres.db_port"),
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

func GetPostgresConfig() *PostgresConfig {
	return config.Database.Postgres
}

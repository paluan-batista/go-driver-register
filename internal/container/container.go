package container

import (
	"fmt"
	"go-driver-register/internal/domain/entities"
	"go-driver-register/internal/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

var (
	database         *gorm.DB
	localContainer   *Container
	runOnceContainer sync.Once
)

type PostgresDB struct {
	*gorm.DB
}

type Container struct {
	Log *log.Logger
	Db  *gorm.DB
}

func InitContainer() *Container {
	runOnceContainer.Do(func() {
		localContainer = NewContainer()
	})
	return localContainer
}

func NewContainer() *Container {
	settings.Init()
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	postgresSQLConnect(logger, settings.GetDatabaseConfig())
	return &Container{
		logger,
		database,
	}
}

func postgresSQLConnect(logger *log.Logger, envs *settings.DatabaseConfig) {
	var err error
	dbUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		envs.Postgres.DbHost, envs.Postgres.DbUser, envs.Postgres.DbPassword, envs.Postgres.DbName, envs.Postgres.DbPort)
	database, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		logger.Fatal("This is the error:", err)
	}

	logger.Printf("We are connected to the %s database", envs.Postgres.DbName)

	if err = database.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		logger.Fatal("Error creating extension:", err)
		os.Exit(1)
	}

	if err := database.Debug().AutoMigrate(&entities.Driver{}, &entities.Vehicle{}); err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}
}

func GetLogger() *log.Logger {
	return localContainer.Log
}

func GetContainer() *Container {
	return localContainer
}

func GetDB() *gorm.DB {
	return localContainer.Db
}

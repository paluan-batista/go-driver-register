package container

import (
	"fmt"
	"go-driver-register/internal/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type PostgresDB struct {
	*gorm.DB
}

type Container struct {
	Log        *log.Logger
	PostgresDB *PostgresDB
}

func NewContainer() *Container {
	return startContainer()

}

func startContainer() *Container {
	settings.Init()
	container := &Container{Log: log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		PostgresDB: startPostgresDB()}
	return container

}

func startPostgresDB() *PostgresDB {
	host := settings.GetPostgresConfig().DbHost
	user := settings.GetPostgresConfig().DbUser
	password := settings.GetPostgresConfig().DbPassword
	dbname := settings.GetPostgresConfig().DbName
	port := settings.GetPostgresConfig().DbPort

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		NewContainer().Log.Fatalf("Error on start postgres :%s", err.Error())
	}
	return &PostgresDB{db}
}

func GetLogger() *log.Logger {
	return NewContainer().Log
}

func GetContainer() *Container {
	return NewContainer()
}

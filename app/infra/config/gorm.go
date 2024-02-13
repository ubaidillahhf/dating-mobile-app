package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func NewGormPostgres(config IConfig) *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Get("PSQL_HOST"),
		config.Get("PSQL_USERNAME"),
		config.Get("PSQL_PASSWORD"),
		config.Get("PSQL_NAME"),
		config.Get("PSQL_PORT"),
	)
	newLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  gormLogger.Info,        // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)
	gormConn, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("start postgreDB connection %s success", gormConn.Name())

	return gormConn
}

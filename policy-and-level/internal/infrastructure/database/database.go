package config

import (
	"github/wstiehler/clean-architecture-concept/policy-and-level/internal/environment"
)

type DBConfig struct {
	databaseHost     string
	databasePort     int64
	databaseUser     string
	databaseName     string
	databasePassword string
}

func buildDBConfig() *DBConfig {
	env := environment.GetInstance()

	dbConfig := DBConfig{
		databaseHost:     env.DATABASE_HOST,
		databasePort:     env.DATABASE_PORT,
		databaseUser:     env.DATABASE_USER,
		databasePassword: env.DATABASE_PASSWORD,
		databaseName:     env.DATABASE_NAME,
	}
	return &dbConfig
}

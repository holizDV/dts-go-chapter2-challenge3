package config

import (
	"os"
	"strings"

	"github.com/holizDV/dts-go-chapter2-challenge3/pkg/helper"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
)

func postgreConnection() string {
	dbConfigs := []string{
		"host=" + os.Getenv("DB_HOST"),
		"user=" + os.Getenv("DB_USER"),
		"password=" + os.Getenv("DB_PASSWORD"),
		"dbname=" + os.Getenv("DB_NAME"),
		"port=" + os.Getenv("PORT"),
		"sslmode=" + os.Getenv("SSL_MODE"),
	}
	return strings.Join(dbConfigs, " ")
}

func PostgreConfig() postgres.Config {
	errEnv := godotenv.Load()
	helper.PanicIfError(errEnv)

	configPostgre := postgres.Config{PreferSimpleProtocol: true}

	configPostgre.DSN = postgreConnection()

	return configPostgre
}

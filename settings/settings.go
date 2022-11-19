package settings

import (
	"github.com/spf13/viper"
)

const dbHost string = "DB_HOST"
const dbPort string = "DB_PORT"
const dbUser string = "DB_USER"
const dbPwd string = "DB_PASSWORD"
const dbName string = "DB_NAME"
const tokenSecretSeed string = "TOKEN_SECRET_SEED"
const AdminRole string = "admin"

func LoadEnv() {
	viper.AutomaticEnv()
}

func TokenSecretSeed() string {
	return viper.GetString(tokenSecretSeed)
}

func DbHost() string {
	return dbHost
}

func DbPort() string {
	return dbPort
}

func DbUser() string {
	return dbUser
}

func DbPwd() string {
	return dbPwd
}

func DbName() string {
	return dbName
}

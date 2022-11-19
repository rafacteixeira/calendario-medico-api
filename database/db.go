package database

import (
	"fmt"
	"github.com/rafacteixeira/calendario-medico-api/model"
	"github.com/rafacteixeira/calendario-medico-api/settings"

	"gorm.io/gorm/logger"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DBStart() {
	user, pwd, host, port, schema := loadEnvVariables()
	dbConnect(user, pwd, host, port, schema)
	createDbStructure()
}

func loadEnvVariables() (string, string, string, int, string) {

	var host, user, pwd, schema string
	var port int
	host = viper.GetString(settings.DbHost())
	port = viper.GetInt(settings.DbPort())
	user = viper.GetString(settings.DbUser())
	pwd = viper.GetString(settings.DbPwd())
	schema = viper.GetString(settings.DbName())

	return user, pwd, host, port, schema
}

func dbConnect(user string, pwd string, host string, port int, schema string) {
	con := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pwd, host, port, schema)
	fmt.Println("String de Conexão: ", con)
	DB, err = gorm.Open(mysql.Open(con), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		log.Panic("Erro ao conectar com o Banco de Dados")
	}
}

func createDbStructure() {
	err := DB.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Note{},
		&model.Event{},
	)

	if err != nil {
		log.Panic("Erro ao construir tabelas")
	}
}

type ChangeConnRequest struct {
	Host     string
	Port     int
	Schema   string
	User     string
	Password string
}

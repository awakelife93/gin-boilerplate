package database

import (
	"fmt"
	"log"
	"os"

	"github.com/awakelife93/gin-boilerplate/src/lib/structures"
	"github.com/awakelife93/gin-boilerplate/src/models"
	"github.com/awakelife93/gin-boilerplate/src/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var database *gorm.DB = nil

func connection(config *structures.MysqlConfig) {
	var error error
	database, error = gorm.Open(
		mysql.Open(
			fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				config.Username,
				config.Password,
				config.Host,
				config.Port,
				config.Database,
			)),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

	if error != nil {
		log.Fatal(error.Error())
	}
}

func migration() {
	if database == nil {
		panic("Database is Nil")
	}

	database.AutoMigrate(
		&models.Product{},
	)
}

func Initialize() {
	config := structures.MysqlConfig{
		Host:     utils.GetStringValue(os.Getenv("mysqlHost"), "localhost"),
		Port:     utils.GetStringToInt(os.Getenv("mysqlPort"), 3306),
		Username: utils.GetStringValue(os.Getenv("mysqlUserName"), "root"),
		Database: os.Getenv("mysqlDatabase"),
		Password: os.Getenv("mysqlPassword"),
	}

	connection(&config)
	migration()
}

func GetDatabase() *gorm.DB {
	if database == nil {
		panic("Database is Nil")
	}

	return database
}

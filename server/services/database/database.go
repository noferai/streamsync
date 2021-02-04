package database

import (
	"fmt"
	"server/core/models"
	"server/services/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB    *gorm.DB
	err   error
	DBErr error
)

type Database struct {
	*gorm.DB
}

// Opens a database and saves the reference to `Database` struct.
func Setup() {
	var db = DB

	configuration := config.GetConfig()

	driver := configuration.Database.Driver
	database := configuration.Database.Dbname
	username := configuration.Database.Username
	password := configuration.Database.Password
	host := configuration.Database.Host
	port := configuration.Database.Port

	if driver == "sqlite" {

		db, err = gorm.Open("sqlite3", "./sqlite3_dev.db")

		if err != nil {
			DBErr = err
			fmt.Println("db err: ", err)
		}

	} else if driver == "postgres" {

		db, err = gorm.Open("postgres", "host="+host+" port="+port+" user="+username+" dbname="+database+"  sslmode=disable password="+password)
		if err != nil {
			DBErr = err
			fmt.Println("db err: ", err)
		}

	}
	db.LogMode(configuration.Database.LogMode)

	db.AutoMigrate(&models.Room{}, &models.Message{})
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

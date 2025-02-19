package db

import (
	"fmt"
	"log"
	"os"

	"github.com/sourabhsd87/URL_Shortner/config"
	"github.com/sourabhsd87/URL_Shortner/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbInit() {
	dbUri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s connect_timeout=%s TimeZone=%s", config.Db_Host, config.Db_User, config.Db_Password, config.Db_Name, config.Db_Port, config.DBConnectionTimeOut, config.Db_TimeZone)
	// Logger.WithFields(logrus.Fields{"name": Db_Name, "user": Db_User}).Info("creating DB connection")
	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		// Logger.WithFields(logrus.Fields{}).Error("error creating db session")
		log.Fatalf("Error connecting to database %v", err)
		os.Exit(1)
	}
	DB = db
	DB.AutoMigrate(&models.URL{})
}

package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

func LoggerInit(dir string) {

	err := CreateFolderIfNotExist(dir, "logs")
	if err != nil {
		log.Println("Error creating logs directory")
	}
	log.Printf("Log Level is: %s", Log_level)
	// Create the info logger
	Logger = logrus.New()
	level, err := logrus.ParseLevel(Log_level)
	if err != nil {
		log.Printf("\nError parsing log level: %v", err.Error())
		log.Println("Setting log level to INFO")
		level = logrus.InfoLevel
	}
	Logger.SetLevel(level)
	infoFile, err := os.OpenFile(dir+"/logs/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.Out = infoFile
	} else {
		Logger.Info("Failed to log to file, using default stderr")
	}
	Logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	Logger.SetReportCaller(true)

}

func CreateFolderIfNotExist(dirPath, folderName string) error {
	// Combine the directory path and folder name to get the full path
	fullPath := filepath.Join(dirPath, folderName)

	// Check if the folder already exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		// Folder does not exist, so create it
		err := os.Mkdir(fullPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating log folder: %v", err)
		}
		fmt.Printf("\nLog folder created: %v", fullPath)
	} else if err != nil {
		// There was an error other than the folder not existing
		return fmt.Errorf("error checking folder existence: %v", err)
	}
	return nil
}

package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Host                 string
	Port                 string
	Db_Host              string
	Db_Port              string
	Db_User              string
	Db_Password          string
	Db_Name              string
	Db_TimeZone          string
	DBConnectionTimeOut  string
	Redis_Host           string
	Redis_Port           string
	Redis_Db             int
	Log_level            string
	Google_Client_ID     string
	Google_Client_Secret string
	Google_Redirect_URL  string
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading it, falling back to OS environment variables")
	}

	envVars := map[string]interface{}{
		"DB_HOST":               &Db_Host,
		"DB_PORT":               &Db_Port,
		"DB_USER":               &Db_User,
		"DB_PASSWORD":           &Db_Password,
		"DB_NAME":               &Db_Name,
		"DB_TIMEZONE":           &Db_TimeZone,
		"DB_CONNECTION_TIMEOUT": &DBConnectionTimeOut,
		"REDIS_HOST":            &Redis_Host,
		"REDIS_PORT":            &Redis_Port,
		"REDIS_DB":              &Redis_Db,
		"HOST":                  &Host,
		"PORT":                  &Port,
		"LOG_LEVEL":             &Log_level,
		"GOOGLE_CLIENT_ID":      &Google_Client_ID,
		"GOOGLE_CLIENT_SECRET":  &Google_Client_Secret,
		"GOOGLE_REDIRECT_URL":   &Google_Redirect_URL,
	}
	for key, value := range envVars {
		val := os.Getenv(key)
		if val == "" {
			log.Printf("Warning: %s is not set in either .env or OS environment", key)
		}

		switch v := value.(type) {
		case *string:
			*v = val
		case *int:
			parsed, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("failed to parse %s as int: %w", key, err)
			}
			*v = parsed
		default:
			return fmt.Errorf("unsupported variable type for %s", key)
		}
	}
	return nil
}

package connection

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Port       string
	DBPort     string
	DBUser     string
	DBPass     string
	DBName     string
	DBHost     string
	DBType     string
	ConnString string
	SslMode    string
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Print("No .env file found")
		os.Exit(3)
	}

	Port = os.Getenv("api_port")
	DBHost = os.Getenv("db_host")
	DBPort = os.Getenv("db_port")
	DBUser = os.Getenv("db_user")
	DBPass = os.Getenv("db_pass")
	DBName = os.Getenv("db_name")
	DBType = os.Getenv("db_type")
	SslMode = os.Getenv("ssl_mode")

	ConnString = "host=" + DBHost + " user=" + DBUser + " dbname=" + DBName + " sslmode=" + SslMode + " password=" + DBPass + " port=" + DBPort
}

func ConnectDatabase() (*gorm.DB, error) {
	return gorm.Open(DBType, ConnString)
}

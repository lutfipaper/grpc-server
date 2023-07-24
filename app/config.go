package app

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var config *Config
var MysqlConfig map[string]interface{}
var connection *gorm.DB

func init() {
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}
	MysqlConfig = map[string]interface{}{
		"host":     os.Getenv("MYSQL_HOST"),
		"password": os.Getenv("MYSQL_PASSWORD"),
		"username": os.Getenv("MYSQL_USERNAME"),
		"port":     os.Getenv("MYSQL_PORT"),
		"database": os.Getenv("MYSQL_DATABASE"),
	}
	// config = &Config{
	// 	Config: libs.NewConfig(),
	// }
}

func GetConfig() (*gorm.DB, error) {
	dbHost := MysqlConfig["host"]
	dbPort := MysqlConfig["port"]
	dbUser := MysqlConfig["username"]
	dbPass := MysqlConfig["password"]
	dbName := MysqlConfig["database"]

	defaultTimezone := os.Getenv("SERVER_TIMEZONE")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=1&loc=%s",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
		url.QueryEscape(defaultTimezone),
	)

	var err error

	connection, err = gorm.Open(mysql.Open(connectionString))
	if nil != err {
		fmt.Println("err", err)

	} else {

		fmt.Println("success", connectionString)

	}
	fmt.Println("Connection is created")
	return connection, nil
}

func GetMysqlConnection() *gorm.DB {
	if connection == nil {
		fmt.Println("Initialize database")
		connection, _ = GetConfig()
	} else {
		fmt.Println("Get connection database")
	}
	return connection
}

// type Config struct {
// 	*libs.Config `yaml:",inline"`
// 	Database     struct {
// 		// Caching interfaces.RedisProviderConfig `yaml:"caching"`
// 		Timeout int                  `yaml:"timeout" default:"30000"`
// 		Adapter interfaces.SQLConfig `yaml:"adapter"`
// 	} `yaml:"database"`
// }

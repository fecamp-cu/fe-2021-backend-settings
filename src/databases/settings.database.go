package databases

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var lock sync.Once

func GetDB() *gorm.DB {
	lock.Do(func() {
		initDB()
	})
	return db
}

func initDB() {
	dsn := getDSN()
	tmpDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB
	db.AutoMigrate(models.Footer{}, models.About{}, models.Setting{}, models.Qualification{}, models.PhotoPreview{}, models.Sponcer{}, models.Timeline{})
}

func getDSN() string {
	host := os.Getenv("DB_HOST")
	if host == "" {
		log.Fatal("DB_HOST is not set")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		log.Fatal("DB_USER is not set")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		log.Fatal("DB_PASSWORD is not set")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME is not set")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	timeZone := os.Getenv("DB_TIMEZONE")
	if timeZone == "" {
		timeZone = "Asia/bangkok"
	}

	sslMode := os.Getenv("DB_SSL_MODE")
	if sslMode == "" {
		sslMode = "disable"
	}
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbName, port, sslMode, timeZone)
}

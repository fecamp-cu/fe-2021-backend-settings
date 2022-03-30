package databases

import (
	"fmt"
	"log"
	"sync"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/configs"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var lock sync.Once

func initGormDB() {
	dsn := getDSN()
	tmpDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	tmpDB.AutoMigrate(models.Footer{}, models.Setting{}, models.About{}, models.Qualification{}, models.PhotoPreview{}, models.Sponcer{}, models.Timeline{})
	db = tmpDB
}

func GetDB() *gorm.DB {
	lock.Do(initGormDB)
	return db
}

func getDSN() string {
	databaseConfig := configs.GetConfigs().Postgres
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", databaseConfig.Host, databaseConfig.User, databaseConfig.Password, databaseConfig.Database, databaseConfig.Port, databaseConfig.SSLMode, databaseConfig.Timezone)
}

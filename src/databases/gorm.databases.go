package databases

import (
	"fmt"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"log"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitGormDB(configs *configs.Configuration) {
	dsn := getDSN(configs)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(models.Footer{}, models.Setting{}, models.About{}, models.Qualification{}, models.PhotoPreview{}, models.Sponcer{}, models.Timeline{}); err != nil {
		log.Fatal(err)
	}
	DB = db
}

func getDSN(configs *configs.Configuration) string {
	databaseConfig := configs.Postgres
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", databaseConfig.Host, databaseConfig.User, databaseConfig.Password, databaseConfig.Database, databaseConfig.Port, databaseConfig.SSLMode, databaseConfig.Timezone)
}

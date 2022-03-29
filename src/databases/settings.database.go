package databases

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db GormDB
var lock sync.Once

type GormDB struct {
	db *gorm.DB
}

func GetDB() GormDB {
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
	tmpDB.AutoMigrate(models.Footer{}, models.Setting{}, models.About{}, models.Qualification{}, models.PhotoPreview{}, models.Sponcer{}, models.Timeline{})
	db = GormDB{tmpDB}
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

// ======= Footer =======

func (s *GormDB) GetFooter(footer *models.Footer) error {
	return s.db.First(footer).Error
}

func (s *GormDB) UpdateFooter(footer *models.Footer) error {
	tmp := models.Footer{}
	err := s.db.First(&tmp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	// if not found, create new
	if err == gorm.ErrRecordNotFound {
		return s.db.Create(footer).Error
	}

	newFooter := *footer
	newFooter.ID = tmp.ID
	return s.db.Save(newFooter).Error
}

// ======= Settings =======
func (s *GormDB) GetAllSettings(settings []*models.Setting) error {
	return s.db.Find(&settings).Error
}

func (s *GormDB) GetAllActiveSettings(settings *[]models.Setting) error {
	return s.db.Preload("Abouts", func(db *gorm.DB) *gorm.DB {
		return db.Order("abouts.order ASC")
	}).Preload("Timelines", func(db *gorm.DB) *gorm.DB {
		return db.Order("timelines.event_start_date ASC")
	}).Preload("Sponcers", func(db *gorm.DB) *gorm.DB {
		return db.Order("sponcers.order ASC")
	}).Preload("Qualifications", func(db *gorm.DB) *gorm.DB {
		return db.Order("qualifications.order ASC")
	}).Where("is_active = ?", true).Find(settings).Error
}

func (s *GormDB) GetSetting(id uint, setting *models.Setting) error {
	return s.db.Where("id = ?", id).First(setting).Error
}

func (s *GormDB) CreateSetting(setting *models.Setting) error {
	return s.db.Create(setting).Error
}

func (s *GormDB) UpdateSetting(setting *models.Setting) error {
	return s.db.Save(setting).Error
}

func (s *GormDB) ActivateSetting(id uint) error {
	return s.db.Model(&models.Setting{}).Where("id = ?", id).Update("is_active", true).Error
}

func (s *GormDB) DeleteSetting(id uint) error {
	return s.db.Where("id = ?", id).Select(clause.Associations).Delete(&models.Setting{ID: id}).Error
}

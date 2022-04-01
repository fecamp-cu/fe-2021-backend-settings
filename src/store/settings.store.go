package store

import (
	"sync"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/databases"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var settingsDB SettingsStore
var lockSettings sync.Once

type SettingsStore struct {
	db    *gorm.DB
	redis databases.RedisClient
}

func GetSetiingsStore() SettingsStore {
	lockSettings.Do(initsettingsStore)
	return settingsDB
}

func initsettingsStore() {
	settingsDB = SettingsStore{
		db:    databases.DB,
		redis: databases.RC,
	}
}

func (s *SettingsStore) GetAllSettings(settings *[]models.Setting) error {
	err := s.redis.Get("settings", settings)
	if err == redis.Nil {
		if err := s.db.Find(settings).Error; err != nil {
			return err
		}
		if err := s.redis.Set("settings", settings); err != nil {
			return err
		}
		return nil
	}
	return err
}

func (s *SettingsStore) GetAllActiveSettings(settings *[]models.Setting) error {
	err := s.redis.Get("active_settings", settings)
	if err == redis.Nil {
		if err := s.db.Preload("Abouts", func(db *gorm.DB) *gorm.DB {
			return db.Order("abouts.order ASC")
		}).Preload("Timelines", func(db *gorm.DB) *gorm.DB {
			return db.Order("timelines.event_start_date ASC")
		}).Preload("Sponcers", func(db *gorm.DB) *gorm.DB {
			return db.Order("sponcers.order ASC")
		}).Preload("Qualifications", func(db *gorm.DB) *gorm.DB {
			return db.Order("qualifications.order ASC")
		}).Where("is_active = ?", true).Find(settings).Error; err != nil {
			return err
		}
		return s.redis.Set("active_settings", settings)
	}
	return err
}

func (s *SettingsStore) GetSetting(id uint, setting *models.Setting) error {
	return s.db.First(setting, id).Error
}

func (s *SettingsStore) CreateSetting(setting *models.Setting) error {
	return s.db.Create(setting).Error
}

func (s *SettingsStore) UpdateSetting(setting *models.Setting) error {
	if err := s.db.Save(setting).Error; err != nil {
		return err
	}

	if err := s.redis.Delete("settings"); err != nil {
		return err
	}

	return s.redis.Delete("active_settings")
}

func (s *SettingsStore) ActivateSetting(id uint) error {
	if err := s.db.Model(&models.Setting{}).Where("id = ?", id).Update("is_active", true).Error; err != nil {
		return err
	}

	if err := s.redis.Delete("settings"); err != nil {
		return err
	}

	return s.redis.Delete("active_settings")
}

func (s *SettingsStore) DeleteSetting(id uint) error {
	if err := s.db.Where("id = ?", id).Select(clause.Associations).Delete(&models.Setting{ID: id}).Error; err != nil {
		return err
	}

	if err := s.redis.Delete("settings"); err != nil {
		return err
	}

	return s.redis.Delete("active_settings")
}

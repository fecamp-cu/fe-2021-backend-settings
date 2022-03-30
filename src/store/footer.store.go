package store

import (
	"sync"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/databases"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var footerDB footerStore
var lockFooter sync.Once

type footerStore struct {
	db    *gorm.DB
	redis databases.RedisClient
}

func GetFooterDB() footerStore {
	lockSettings.Do(initFooterDB)
	return footerDB
}

func initFooterDB() {
	footerDB = footerStore{
		db:    databases.GetDB(),
		redis: databases.GetRedis(),
	}
}

func (s *footerStore) GetFooter(footer *models.Footer) error {
	err := s.redis.Get("footer", footer)
	if err == redis.Nil {
		if err := s.db.First(footer).Error; err != nil {
			return err
		}
		s.redis.Set("footer", footer)
		return nil
	}
	return err
}

func (s *footerStore) UpdateFooter(footer *models.Footer) error {
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
	if err := s.db.Save(newFooter).Error; err != nil {
		return err
	}
	return s.redis.Set("footer", footer)
}

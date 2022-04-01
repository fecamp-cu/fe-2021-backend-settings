package store

import (
	"sync"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/databases"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var sponcerDB SponcerStore
var lockSponcer sync.Once

type SponcerStore struct {
	db    *gorm.DB
	redis databases.RedisClient
}

func initSponcer() {
	sponcerDB = SponcerStore{
		db:    databases.DB,
		redis: databases.RC,
	}
}

func GetSponcerStore() SponcerStore {
	lockSponcer.Do(initSponcer)
	return sponcerDB
}

func (s *SponcerStore) CreateSponcer(sponcer *models.Sponcer) error {
	if err := s.db.Create(sponcer).Error; err != nil {
		return err
	}
	return s.redis.Delete("sponcer")
}

func (s *SponcerStore) GetAllSponcer(sponcer *[]models.Sponcer) error {
	err := s.redis.Get("sponcer", sponcer)
	if err == redis.Nil {
		if err := s.db.Find(sponcer).Error; err != nil {
			return err
		}

		if err := s.redis.Set("sponcer", sponcer); err != nil {
			return err
		}
		return nil
	}
	return err
}

func (s *SponcerStore) GetSponcer(id uint, sponcer *models.Sponcer) error {
	return s.db.Where("id = ?", id).First(sponcer).Error
}

func (s *SponcerStore) UpdateSponcer(sponcer *models.Sponcer) error {
	if err := s.db.Save(sponcer).Error; err != nil {
		return err
	}
	return s.redis.Delete("sponcer")
}

func (s *SponcerStore) DeleteSponcer(id uint) error {
	if err := s.db.Where("id = ?", id).Delete(models.Sponcer{}).Error; err != nil {
		return err
	}
	return s.redis.Delete("sponcer")
}

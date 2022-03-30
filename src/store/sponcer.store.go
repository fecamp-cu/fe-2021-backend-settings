package store

import (
	"sync"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/databases"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var sponcerDB sponcerStore
var lockSponcer sync.Once

type sponcerStore struct {
	db    *gorm.DB
	redis databases.RedisClient
}

func initSponcer() {
	sponcerDB = sponcerStore{
		db:    databases.GetDB(),
		redis: databases.GetRedis(),
	}
}

func GetSponcerStore() sponcerStore {
	lockSponcer.Do(initSponcer)
	return sponcerDB
}

func (s *sponcerStore) CreateSponcer(sponcer *models.Sponcer) error {
	if err := s.db.Create(sponcer).Error; err != nil {
		return err
	}
	return s.redis.Delete("sponcer")
}

func (s *sponcerStore) GetAllSponcer(sponcer *[]models.Sponcer) error {
	err := s.redis.Get("sponcer", sponcer)
	if err == redis.Nil {
		if err := s.db.Find(sponcer).Error; err != nil {
			return err
		}
		s.redis.Set("sponcer", sponcer)
		return nil
	}
	return err
}

func (s *sponcerStore) GetSponcer(id uint, sponcer *models.Sponcer) error {
	return s.db.Where("id = ?", id).First(sponcer).Error
}

func (s *sponcerStore) UpdateSponcer(sponcer *models.Sponcer) error {
	if err := s.db.Save(sponcer).Error; err != nil {
		return err
	}
	return s.redis.Delete("sponcer")
}

func (s *sponcerStore) DeleteSponcer(id uint) error {
	if err := s.db.Where("id = ?", id).Delete(models.Sponcer{}).Error; err != nil {
		return err
	}
	return s.redis.Delete("sponcer")
}

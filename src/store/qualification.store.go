package store

import (
	"sync"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/databases"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var qualificationDB qualificationStore
var lockQualification sync.Once

type qualificationStore struct {
	db    *gorm.DB
	redis databases.RedisClient
}

func initQualification() {
	qualificationDB = qualificationStore{
		db:    databases.GetDB(),
		redis: databases.GetRedis(),
	}
}

func GetQualificationStore() qualificationStore {
	lockQualification.Do(initQualification)
	return qualificationDB
}

func (s *qualificationStore) CreateQualification(qualification *models.Qualification) error {
	if err := s.db.Create(qualification).Error; err != nil {
		return err
	}
	return s.redis.Delete("qualification")
}

func (s *qualificationStore) GetAllQualification(qualification *[]models.Qualification) error {
	err := s.redis.Get("qualification", qualification)
	if err == redis.Nil {
		if err := s.db.Find(qualification).Error; err != nil {
			return err
		}
		s.redis.Set("qualification", qualification)
		return nil
	}
	return err
}

func (s *qualificationStore) GetQualification(id uint, qualification *models.Qualification) error {
	return s.db.Where("id = ?", id).First(qualification).Error
}

func (s *qualificationStore) UpdateQualification(qualification *models.Qualification) error {
	if err := s.db.Save(qualification).Error; err != nil {
		return err
	}
	return s.redis.Delete("qualification")
}

func (s *qualificationStore) DeleteQualification(id uint) error {
	if err := s.db.Where("id = ?", id).Delete(models.Qualification{}).Error; err != nil {
		return err
	}
	return s.redis.Delete("qualification")
}

package store

import (
	"sync"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/databases"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var qualificationDB QualificationStore
var lockQualification sync.Once

type QualificationStore struct {
	db    *gorm.DB
	redis databases.RedisClient
}

func initQualification() {
	qualificationDB = QualificationStore{
		db:    databases.DB,
		redis: databases.RC,
	}
}

func GetQualificationStore() *QualificationStore {
	lockQualification.Do(initQualification)
	return &qualificationDB
}

func (s *QualificationStore) CreateQualification(qualification *models.Qualification) error {
	if err := s.db.Create(qualification).Error; err != nil {
		return err
	}
	return s.redis.Delete("qualification")
}

func (s *QualificationStore) GetAllQualification(qualification *[]models.Qualification) error {
	err := s.redis.Get("qualification", qualification)
	if err == redis.Nil {
		if err := s.db.Find(qualification).Error; err != nil {
			return err
		}
		if err := s.redis.Set("qualification", qualification); err != nil {
			return err
		}
		return nil
	}
	return err
}

func (s *QualificationStore) GetQualification(id uint, qualification *models.Qualification) error {
	return s.db.Where("id = ?", id).First(qualification).Error
}

func (s *QualificationStore) UpdateQualification(qualification *models.Qualification) error {
	if err := s.db.Save(qualification).Error; err != nil {
		return err
	}
	return s.redis.Delete("qualification")
}

func (s *QualificationStore) DeleteQualification(id uint) error {
	if err := s.db.Where("id = ?", id).Delete(models.Qualification{}).Error; err != nil {
		return err
	}
	return s.redis.Delete("qualification")
}

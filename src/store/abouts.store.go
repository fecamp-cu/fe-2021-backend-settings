package store

import (
	"sync"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/databases"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var aboutsDB AboutsStore
var lockAbouts sync.Once

type AboutsStore struct {
	db    *gorm.DB
	redis databases.RedisClient
}

func initAbouts() {
	aboutsDB = AboutsStore{
		db:    databases.DB,
		redis: databases.RC,
	}
}

func GetAboutsStore() AboutsStore {
	lockAbouts.Do(initAbouts)
	return aboutsDB
}

func (s *AboutsStore) CreateAbout(about *models.About) error {
	if err := s.db.Create(about).Error; err != nil {
		return err
	}
	return s.redis.Delete("abouts")
}

func (s *AboutsStore) GetAllAbouts(abouts *[]models.About) error {
	err := s.redis.Get("abouts", abouts)
	if err == redis.Nil {
		if err := s.db.Find(abouts).Error; err != nil {
			return err
		}
		if err := s.redis.Set("abouts", abouts); err != nil {
			return err
		}
		return nil
	}
	return err
}

func (s *AboutsStore) GetAbout(id uint, about *models.About) error {
	return s.db.Where("id = ?", id).First(about).Error
}

func (s *AboutsStore) UpdateAbout(about *models.About) error {
	if err := s.db.Save(about).Error; err != nil {
		return err
	}
	return s.redis.Delete("abouts")
}

func (s *AboutsStore) DeleteAbout(id uint) error {
	if err := s.db.Where("id = ?", id).Delete(models.About{}).Error; err != nil {
		return err
	}
	return s.redis.Delete("abouts")
}

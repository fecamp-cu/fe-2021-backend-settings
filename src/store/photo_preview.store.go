package store

import (
	"sync"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/databases"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var photoPreviewDB photoPreviewStore
var lockPhotoPreview sync.Once

type photoPreviewStore struct {
	db    *gorm.DB
	redis databases.RedisClient
}

func initPhotoPreview() {
	photoPreviewDB = photoPreviewStore{
		db:    databases.GetDB(),
		redis: databases.GetRedis(),
	}
}

func GetPhotoPreviewStore() photoPreviewStore {
	lockPhotoPreview.Do(initPhotoPreview)
	return photoPreviewDB
}

func (s *photoPreviewStore) CreatePhotoPreview(photoPreview *models.PhotoPreview) error {
	if err := s.db.Create(photoPreview).Error; err != nil {
		return err
	}
	return s.redis.Delete("PhotoPreview")
}

func (s *photoPreviewStore) GetAllPhotoPreview(photoPreview *[]models.PhotoPreview) error {
	err := s.redis.Get("PhotoPreview", photoPreview)
	if err == redis.Nil {
		if err := s.db.Find(photoPreview).Error; err != nil {
			return err
		}
		s.redis.Set("PhotoPreview", photoPreview)
		return nil
	}
	return err
}

func (s *photoPreviewStore) GetPhotoPreview(id uint, photoPreview *models.PhotoPreview) error {
	return s.db.Where("id = ?", id).First(photoPreview).Error
}

func (s *photoPreviewStore) UpdatePhotoPreview(photoPreview *models.PhotoPreview) error {
	if err := s.db.Save(photoPreview).Error; err != nil {
		return err
	}
	return s.redis.Delete("PhotoPreview")
}

func (s *photoPreviewStore) DeletePhotoPreview(id uint) error {
	if err := s.db.Where("id = ?", id).Delete(models.PhotoPreview{}).Error; err != nil {
		return err
	}
	return s.redis.Delete("PhotoPreview")
}

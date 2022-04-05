package store

import (
	"sync"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/databases"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var photoPreviewDB PhotoPreviewStore
var lockPhotoPreview sync.Once

type PhotoPreviewStore struct {
	db    *gorm.DB
	redis databases.RedisClient
}

func initPhotoPreview() {
	photoPreviewDB = PhotoPreviewStore{
		db:    databases.DB,
		redis: databases.RC,
	}
}

func GetPhotoPreviewStore() *PhotoPreviewStore {
	lockPhotoPreview.Do(initPhotoPreview)
	return &photoPreviewDB
}

func (s *PhotoPreviewStore) CreatePhotoPreview(photoPreview *models.PhotoPreview) error {
	if err := s.db.Create(photoPreview).Error; err != nil {
		return err
	}
	return s.redis.Delete("PhotoPreview")
}

func (s *PhotoPreviewStore) GetAllPhotoPreview(photoPreview *[]models.PhotoPreview) error {
	err := s.redis.Get("PhotoPreview", photoPreview)
	if err == redis.Nil {
		if err := s.db.Find(photoPreview).Error; err != nil {
			return err
		}
		if err := s.redis.Set("PhotoPreview", photoPreview); err != nil {
			return err
		}
		return nil
	}
	return err
}

func (s *PhotoPreviewStore) GetPhotoPreview(id uint, photoPreview *models.PhotoPreview) error {
	return s.db.Where("id = ?", id).First(photoPreview).Error
}

func (s *PhotoPreviewStore) UpdatePhotoPreview(photoPreview *models.PhotoPreview) error {
	if err := s.db.Save(photoPreview).Error; err != nil {
		return err
	}
	return s.redis.Delete("PhotoPreview")
}

func (s *PhotoPreviewStore) DeletePhotoPreview(id uint) error {
	if err := s.db.Where("id = ?", id).Delete(models.PhotoPreview{}).Error; err != nil {
		return err
	}
	return s.redis.Delete("PhotoPreview")
}

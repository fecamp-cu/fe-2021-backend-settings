package store

import (
	"sync"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/databases"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var timelineDB timelineStore
var lockTimeline sync.Once

type timelineStore struct {
	db    *gorm.DB
	redis databases.RedisClient
}

func initTimeline() {
	timelineDB = timelineStore{
		db:    databases.GetDB(),
		redis: databases.GetRedis(),
	}
}

func GetTimelineStore() timelineStore {
	lockTimeline.Do(initTimeline)
	return timelineDB
}

func (s *timelineStore) CreateTimeline(timeline *models.Timeline) error {
	if err := s.db.Create(timeline).Error; err != nil {
		return err
	}
	return s.redis.Delete("timeline")
}

func (s *timelineStore) GetAllTimeline(timeline *[]models.Timeline) error {
	err := s.redis.Get("timeline", timeline)
	if err == redis.Nil {
		if err := s.db.Find(timeline).Error; err != nil {
			return err
		}
		s.redis.Set("timeline", timeline)
		return nil
	}
	return err
}

func (s *timelineStore) GetTimeline(id uint, timeline *models.Timeline) error {
	return s.db.Where("id = ?", id).First(timeline).Error
}

func (s *timelineStore) UpdateTimeline(timeline *models.Timeline) error {
	if err := s.db.Save(timeline).Error; err != nil {
		return err
	}
	return s.redis.Delete("timeline")
}

func (s *timelineStore) DeleteTimeline(id uint) error {
	if err := s.db.Where("id = ?", id).Delete(models.Timeline{}).Error; err != nil {
		return err
	}
	return s.redis.Delete("timeline")
}

package store

import (
	"sync"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/databases"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var timelineDB TimelineStore
var lockTimeline sync.Once

type TimelineStore struct {
	db    *gorm.DB
	redis databases.RedisClient
}

func initTimeline() {
	timelineDB = TimelineStore{
		db:    databases.DB,
		redis: databases.RC,
	}
}

func GetTimelineStore() *TimelineStore {
	lockTimeline.Do(initTimeline)
	return &timelineDB
}

func (s *TimelineStore) CreateTimeline(timeline *models.Timeline) error {
	if err := s.db.Create(timeline).Error; err != nil {
		return err
	}
	return s.redis.Delete("timeline")
}

func (s *TimelineStore) GetAllTimeline(timeline *[]models.Timeline) error {
	err := s.redis.Get("timeline", timeline)
	if err == redis.Nil {
		if err := s.db.Find(timeline).Error; err != nil {
			return err
		}
		if err := s.redis.Set("timeline", timeline); err != nil {
			return err
		}
		return nil
	}
	return err
}

func (s *TimelineStore) GetTimeline(id uint, timeline *models.Timeline) error {
	return s.db.Where("id = ?", id).First(timeline).Error
}

func (s *TimelineStore) UpdateTimeline(timeline *models.Timeline) error {
	if err := s.db.Save(timeline).Error; err != nil {
		return err
	}
	return s.redis.Delete("timeline")
}

func (s *TimelineStore) DeleteTimeline(id uint) error {
	if err := s.db.Where("id = ?", id).Delete(models.Timeline{}).Error; err != nil {
		return err
	}
	return s.redis.Delete("timeline")
}

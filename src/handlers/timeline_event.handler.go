package handlers

import (
	"net/http"
	"strconv"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/routers"
	"github.com/gofiber/fiber/v2"
)

type timelineStorer interface {
	CreateTimeline(timeline *models.Timeline) error
	GetAllTimeline(timelines *[]models.Timeline) error
	GetTimeline(id uint, timeline *models.Timeline) error
	UpdateTimeline(timeline *models.Timeline) error
	DeleteTimeline(id uint) error
}

type TimelineHandler struct {
	timelineStorer timelineStorer
}

func NewDefaultTimelineHandler(s timelineStorer) *TimelineHandler {
	return &TimelineHandler{s}
}

func (s *TimelineHandler) CreateTimeline(ctx routers.Context) {
	timeline := models.Timeline{}
	id, err := strconv.Atoi(ctx.Params("settingid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	if err := ctx.Bind(&timeline); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	timeline.SettingID = uint(id)
	if err := s.timelineStorer.CreateTimeline(&timeline); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
}

func (s *TimelineHandler) GetAllTimelines(ctx routers.Context) {
	timelines := []models.Timeline{}
	if err := s.timelineStorer.GetAllTimeline(&timelines); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, timelines)
}

func (s *TimelineHandler) GetTimeline(ctx routers.Context) {
	timeline := models.Timeline{}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	if err := s.timelineStorer.GetTimeline(uint(id), &timeline); err != nil {
		ctx.JSON(http.StatusNotFound, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, timeline)
}

func (s *TimelineHandler) UpdateTimeline(ctx routers.Context) {
	timeline := models.Timeline{}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	if err := ctx.Bind(&timeline); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	timeline.ID = uint(id)
	if err := s.timelineStorer.UpdateTimeline(&timeline); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
}

func (s *TimelineHandler) DeleteTimeline(ctx routers.Context) {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	if err := s.timelineStorer.DeleteTimeline(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
}

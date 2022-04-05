package handlers

import (
	"net/http"
	"strconv"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/routers"
	"github.com/gofiber/fiber/v2"
)

type aboutStorer interface {
	CreateAbout(about *models.About) error
	GetAllAbouts(abouts *[]models.About) error
	GetAbout(id uint, about *models.About) error
	UpdateAbout(about *models.About) error
	DeleteAbout(id uint) error
}

type AboutHandler struct {
	aboutStorer aboutStorer
}

func NewDefaultAboutHandler(s aboutStorer) *AboutHandler {
	return &AboutHandler{s}
}

func (s *AboutHandler) CreateAbout(ctx routers.Context) {
	about := models.About{}
	id, err := strconv.Atoi(ctx.Params("settingid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	if err := ctx.Bind(&about); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	about.SettingID = uint(id)
	if err := s.aboutStorer.CreateAbout(&about); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
}

func (s *AboutHandler) GetAllAbouts(ctx routers.Context) {
	abouts := []models.About{}
	if err := s.aboutStorer.GetAllAbouts(&abouts); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, abouts)
}

func (s *AboutHandler) GetAbout(ctx routers.Context) {
	about := models.About{}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	if err := s.aboutStorer.GetAbout(uint(id), &about); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, about)
}

func (s *AboutHandler) UpdateAbout(ctx routers.Context) {
	about := models.About{}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	if err := ctx.Bind(&about); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	about.ID = uint(id)
	if err := s.aboutStorer.UpdateAbout(&about); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
}

func (s *AboutHandler) DeleteAbout(ctx routers.Context) {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	if err := s.aboutStorer.DeleteAbout(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
}

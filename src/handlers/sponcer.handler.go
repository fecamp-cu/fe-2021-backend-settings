package handlers

import (
	"net/http"
	"strconv"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/routers"
	"github.com/gofiber/fiber/v2"
)

type sponcerStorer interface {
	CreateSponcer(sponcer *models.Sponcer) error
	GetAllSponcer(sponcers *[]models.Sponcer) error
	GetSponcer(id uint, sponcer *models.Sponcer) error
	UpdateSponcer(sponcer *models.Sponcer) error
	DeleteSponcer(id uint) error
}

type SponcerHandler struct {
	sponcerStorer sponcerStorer
}

func NewDefaultSponcerHandler(s sponcerStorer) *SponcerHandler {
	return &SponcerHandler{s}
}

func (s *SponcerHandler) CreateSponcer(ctx routers.Context) {
	sponcer := models.Sponcer{}
	id, err := strconv.Atoi(ctx.Params("settingid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	if err := ctx.Bind(&sponcer); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	sponcer.SettingID = uint(id)
	if err := s.sponcerStorer.CreateSponcer(&sponcer); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
}

func (s *SponcerHandler) GetAllSponcers(ctx routers.Context) {
	sponcers := []models.Sponcer{}
	if err := s.sponcerStorer.GetAllSponcer(&sponcers); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, sponcers)
}

func (s *SponcerHandler) GetSponcer(ctx routers.Context) {
	sponcer := models.Sponcer{}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	if err := s.sponcerStorer.GetSponcer(uint(id), &sponcer); err != nil {
		ctx.JSON(http.StatusNotFound, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, sponcer)
}

func (s *SponcerHandler) UpdateSponcer(ctx routers.Context) {
	sponcer := models.Sponcer{}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	if err := ctx.Bind(&sponcer); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	sponcer.ID = uint(id)
	if err := s.sponcerStorer.UpdateSponcer(&sponcer); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
}

func (s *SponcerHandler) DeleteSponcer(ctx routers.Context) {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	if err := s.sponcerStorer.DeleteSponcer(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
}

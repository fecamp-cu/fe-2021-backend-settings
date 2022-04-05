package handlers

import (
	"net/http"
	"strconv"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/routers"
	"github.com/gofiber/fiber/v2"
)

type settingsStorer interface {
	GetAllSettings(settings *[]models.Setting) error
	GetAllActiveSettings(settings *[]models.Setting) error
	GetSetting(id uint, setting *models.Setting) error
	CreateSetting(setting *models.Setting) error
	UpdateSetting(setting *models.Setting) error
	ActivateSetting(id uint) error
	DeleteSetting(id uint) error
}

type SettingsHandler struct {
	settingsStorer settingsStorer
}

func NewDefaultSettingsHandler(s settingsStorer) *SettingsHandler {
	return &SettingsHandler{s}
}

func (s *SettingsHandler) CreateSetting(ctx routers.Context) {
	setting := models.Setting{}
	if err := ctx.Bind(&setting); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	if err := s.settingsStorer.CreateSetting(&setting); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
}

func (s *SettingsHandler) GetAllSettings(ctx routers.Context) {
	settings := []models.Setting{}
	if err := s.settingsStorer.GetAllSettings(&settings); err != nil {
		ctx.JSON(http.StatusNotFound, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, settings)
}

func (s *SettingsHandler) GetAllActiveSettings(ctx routers.Context) {
	settings := []models.Setting{}
	if err := s.settingsStorer.GetAllActiveSettings(&settings); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, settings)
}

func (s *SettingsHandler) GetSetting(ctx routers.Context) {
	i_id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	setting := models.Setting{}
	if err := s.settingsStorer.GetSetting(uint(i_id), &setting); err != nil {
		ctx.JSON(http.StatusNotFound, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, setting)
}

func (s *SettingsHandler) UpdateSetting(ctx routers.Context) {
	i_id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	setting := models.Setting{}
	if err := ctx.Bind(&setting); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	setting.ID = uint(i_id)
	if err := s.settingsStorer.UpdateSetting(&setting); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
}

func (s *SettingsHandler) RemoveSetting(ctx routers.Context) {
	i_id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
	if err := s.settingsStorer.DeleteSetting(uint(i_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
		return
	}
}

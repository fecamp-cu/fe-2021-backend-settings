package handlers

import (
	"net/http"
	"strconv"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/routers"
	"github.com/gofiber/fiber/v2"
)

type photoPreviewStorer interface {
	CreatePhotoPreview(photoPreview *models.PhotoPreview) error
	GetAllPhotoPreview(photoPreviews *[]models.PhotoPreview) error
	GetPhotoPreview(id uint, photoPreview *models.PhotoPreview) error
	UpdatePhotoPreview(photoPreview *models.PhotoPreview) error
	DeletePhotoPreview(id uint) error
}

type PhotoPreviewHandler struct {
	photoPreviewStorer photoPreviewStorer
}

func NewDefaultPhotoPreviewHandler(s photoPreviewStorer) *PhotoPreviewHandler {
	return &PhotoPreviewHandler{s}
}

func (s *PhotoPreviewHandler) CreatePhotoPreview(ctx routers.Context) {
	photoPreview := models.PhotoPreview{}
	id, err := strconv.Atoi(ctx.Params("settingid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	if err := ctx.Bind(&photoPreview); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	photoPreview.SettingID = uint(id)
	if err := s.photoPreviewStorer.CreatePhotoPreview(&photoPreview); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
}

func (s *PhotoPreviewHandler) GetAllPhotoPreviews(ctx routers.Context) {
	photoPreviews := []models.PhotoPreview{}
	if err := s.photoPreviewStorer.GetAllPhotoPreview(&photoPreviews); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, photoPreviews)
}

func (s *PhotoPreviewHandler) GetPhotoPreview(ctx routers.Context) {
	photoPreview := models.PhotoPreview{}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	if err := s.photoPreviewStorer.GetPhotoPreview(uint(id), &photoPreview); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, photoPreview)
}

func (s *PhotoPreviewHandler) UpdatePhotoPreview(ctx routers.Context) {
	photoPreview := models.PhotoPreview{}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	if err := ctx.Bind(&photoPreview); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	photoPreview.ID = uint(id)
	if err := s.photoPreviewStorer.UpdatePhotoPreview(&photoPreview); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
}

func (s *PhotoPreviewHandler) DeletePhotoPreview(ctx routers.Context) {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	if err := s.photoPreviewStorer.DeletePhotoPreview(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
}

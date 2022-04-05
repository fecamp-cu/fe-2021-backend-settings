package handlers

import (
	"net/http"
	"strconv"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/routers"
	"github.com/gofiber/fiber/v2"
)

type qualificationStorer interface {
	CreateQualification(qualification *models.Qualification) error
	GetAllQualification(qualifications *[]models.Qualification) error
	GetQualification(id uint, qualification *models.Qualification) error
	UpdateQualification(qualification *models.Qualification) error
	DeleteQualification(id uint) error
}

type QualificationHandler struct {
	qualificationStorer qualificationStorer
}

func NewDefaultQualificationHandler(s qualificationStorer) *QualificationHandler {
	return &QualificationHandler{s}
}

func (s *QualificationHandler) CreateQualification(ctx routers.Context) {
	qualification := models.Qualification{}
	id, err := strconv.Atoi(ctx.Params("settingid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	if err := ctx.Bind(&qualification); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	qualification.SettingID = uint(id)
	if err := s.qualificationStorer.CreateQualification(&qualification); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
}

func (s *QualificationHandler) GetAllQualifications(ctx routers.Context) {
	qualifications := []models.Qualification{}
	if err := s.qualificationStorer.GetAllQualification(&qualifications); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, qualifications)
}

func (s *QualificationHandler) GetQualification(ctx routers.Context) {
	qualification := models.Qualification{}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	if err := s.qualificationStorer.GetQualification(uint(id), &qualification); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, qualification)
}

func (s *QualificationHandler) UpdateQualification(ctx routers.Context) {
	qualification := models.Qualification{}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	if err := ctx.Bind(&qualification); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	qualification.ID = uint(id)
	if err := s.qualificationStorer.UpdateQualification(&qualification); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
}

func (s *QualificationHandler) DeleteQualification(ctx routers.Context) {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	if err := s.qualificationStorer.DeleteQualification(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
}

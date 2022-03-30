package handlers

import (
	"net/http"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/gofiber/fiber/v2"
)

type footerStorer interface {
	GetFooter(footer *models.Footer) error
	UpdateFooter(footer *models.Footer) (bool, error)
}

type FooterHandler struct {
	footerStorer footerStorer
}

func NewDefaultFooterHandler(s footerStorer) *FooterHandler {
	return &FooterHandler{s}
}

func (h *FooterHandler) GetFooter(ctx *fiber.Ctx) error {
	footer := models.Footer{}
	if err := h.footerStorer.GetFooter(&footer); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(footer)
}

func (h *FooterHandler) UpdateFooter(ctx *fiber.Ctx) error {
	footer := models.Footer{}
	if err := ctx.BodyParser(&footer); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	isNew, err := h.footerStorer.UpdateFooter(&footer)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	status := http.StatusOK
	if isNew {
		status = http.StatusCreated
	}

	return ctx.Status(status).JSON(footer)
}

package handlers

import (
	"net/http"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/models"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/routers"
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

func (h *FooterHandler) GetFooter(ctx routers.Context) {
	footer := models.Footer{}
	if err := h.footerStorer.GetFooter(&footer); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, footer)
}

func (h *FooterHandler) UpdateFooter(ctx routers.Context) {
	footer := models.Footer{}
	if err := ctx.Bind(&footer); err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}
	isNew, err := h.footerStorer.UpdateFooter(&footer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}

	status := http.StatusOK
	if isNew {
		status = http.StatusCreated
	}

	ctx.JSON(status, footer)
}

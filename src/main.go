package main

import (
	"github.com/fecamp-cu/fe-2021-backend-settings/src/configs"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/handlers"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/store"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	r := fiber.New()
	r.Use(logger.New())
	r.Use(cors.New())

	app := r.Group("/api/v1")

	footerRoute := app.Group("/footer")
	footerHandler := handlers.NewDefaultFooterHandler(store.GetFooterStore())
	footerRoute.Get("/", footerHandler.GetFooter)
	footerRoute.Patch("/", footerHandler.UpdateFooter)

	port := configs.GetConfigs().App.Port
	r.Listen(":" + port)
}

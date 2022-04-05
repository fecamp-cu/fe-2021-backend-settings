package main

import (
	"github.com/fecamp-cu/fe-2021-backend-settings/src/configs"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/databases"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/handlers"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/routers"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/store"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config := configs.LoadConfigs()
	databases.InitRedis(config)
	databases.InitGormDB(config)

	// r := fiber.New()
	// r.Use(logger.New())
	// r.Use(cors.New())

	// app := r.Group("/api/v1")

	// footerRoute := app.Group("/footer")
	// footerRoute.Get("/", footerHandler.GetFooter)
	// footerRoute.Patch("/", footerHandler.UpdateFooter)

	r := routers.NewFiberRouter()
	r.Use(logger.New())
	r.Use(cors.New())

	footerHandler := handlers.NewDefaultFooterHandler(store.GetFooterStore())
	r.GET("/api/v1/footer/", footerHandler.GetFooter)
	r.PATCH("/api/v1/footer/", footerHandler.UpdateFooter)

	settingsHandler := handlers.NewDefaultSettingsHandler(store.GetSetiingsStore())
	r.POST("/api/v1/settings/", settingsHandler.CreateSetting)
	r.GET("/api/v1/settings/", settingsHandler.GetAllSettings)
	r.GET("/api/v1/settings/active", settingsHandler.GetAllActiveSettings)
	r.GET("/api/v1/settings/:id", settingsHandler.GetSetting)
	r.PATCH("/api/v1/settings/:id", settingsHandler.UpdateSetting)
	r.DELETE("/api/v1/settings/:id", settingsHandler.RemoveSetting)
	
	if err := r.Listen(":" + config.App.Port); err != nil {
		panic(err)
	}
}

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

	r := routers.NewFiberRouter()
	r.Use(logger.New())
	r.Use(cors.New())

	footerHandler := handlers.NewDefaultFooterHandler(store.GetFooterStore())
	r.GET("/footer/", footerHandler.GetFooter)
	r.PATCH("/footer/", footerHandler.UpdateFooter)

	settingsHandler := handlers.NewDefaultSettingsHandler(store.GetSetiingsStore())
	r.GET("/settings/", settingsHandler.GetAllSettings)
	r.GET("/settings/active", settingsHandler.GetAllActiveSettings)
	r.GET("/settings/:id", settingsHandler.GetSetting)
	r.POST("/settings/", settingsHandler.CreateSetting)
	r.PATCH("/settings/:id", settingsHandler.UpdateSetting)
	r.DELETE("/settings/:id", settingsHandler.RemoveSetting)

	aboutHandler := handlers.NewDefaultAboutHandler(store.GetAboutsStore())
	r.POST("/setting/about/:settingid", aboutHandler.CreateAbout)
	r.GET("/setting/about/", aboutHandler.GetAllAbouts)
	r.GET("/setting/about/:id", aboutHandler.GetAbout)
	r.PATCH("/setting/about/:id", aboutHandler.UpdateAbout)
	r.DELETE("/setting/about/:id", aboutHandler.DeleteAbout)

	photoHandler := handlers.NewDefaultPhotoPreviewHandler(store.GetPhotoPreviewStore())
	r.POST("/setting/photo/:settingid", photoHandler.CreatePhotoPreview)
	r.GET("/setting/photo/", photoHandler.GetAllPhotoPreviews)
	r.GET("/setting/photo/:id", photoHandler.GetPhotoPreview)
	r.PATCH("/setting/photo/:id", photoHandler.UpdatePhotoPreview)
	r.DELETE("/setting/photo/:id", photoHandler.DeletePhotoPreview)

	qualificationHandler := handlers.NewDefaultQualificationHandler(store.GetQualificationStore())
	r.POST("/setting/qualification/:settingid", qualificationHandler.CreateQualification)
	r.GET("/setting/qualification/", qualificationHandler.GetAllQualifications)
	r.GET("/setting/qualification/:id", qualificationHandler.GetQualification)
	r.PATCH("/setting/qualification/:id", qualificationHandler.UpdateQualification)
	r.DELETE("/setting/qualification/:id", qualificationHandler.DeleteQualification)

	sponcerHandler := handlers.NewDefaultSponcerHandler(store.GetSponcerStore())
	r.POST("/setting/sponcer/:settingid", sponcerHandler.CreateSponcer)
	r.GET("/setting/sponcer/", sponcerHandler.GetAllSponcers)
	r.GET("/setting/sponcer/:id", sponcerHandler.GetSponcer)
	r.PATCH("/setting/sponcer/:id", sponcerHandler.UpdateSponcer)
	r.DELETE("/setting/sponcer/:id", sponcerHandler.DeleteSponcer)

	timelineHandler := handlers.NewDefaultTimelineHandler(store.GetTimelineStore())
	r.POST("/setting/timeline/:settingid", timelineHandler.CreateTimeline)
	r.GET("/setting/timeline/", timelineHandler.GetAllTimelines)
	r.GET("/setting/timeline/:id", timelineHandler.GetTimeline)
	r.PATCH("/setting/timeline/:id", timelineHandler.UpdateTimeline)
	r.DELETE("/setting/timeline/:id", timelineHandler.DeleteTimeline)

	if err := r.Listen(":" + config.App.Port); err != nil {
		panic(err)
	}
}

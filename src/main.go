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
	r.GET("/api/v1/footer/", footerHandler.GetFooter)
	r.PATCH("/api/v1/footer/", footerHandler.UpdateFooter)

	settingsHandler := handlers.NewDefaultSettingsHandler(store.GetSetiingsStore())
	r.GET("/api/v1/setting/", settingsHandler.GetAllSettings)
	r.GET("/api/v1/setting/active", settingsHandler.GetAllActiveSettings)
	r.GET("/api/v1/setting/:id", settingsHandler.GetSetting)
	r.POST("/api/v1/setting/", settingsHandler.CreateSetting)
	r.PATCH("/api/v1/setting/:id", settingsHandler.UpdateSetting)
	r.DELETE("/api/v1/setting/:id", settingsHandler.RemoveSetting)

	aboutHandler := handlers.NewDefaultAboutHandler(store.GetAboutsStore())
	r.POST("/api/v1/setting/about/:settingid", aboutHandler.CreateAbout)
	r.GET("/api/v1/setting/about/", aboutHandler.GetAllAbouts)
	r.GET("/api/v1/setting/about/:id", aboutHandler.GetAbout)
	r.PATCH("/api/v1/setting/about/:id", aboutHandler.UpdateAbout)
	r.DELETE("/api/v1/setting/about/:id", aboutHandler.DeleteAbout)

	photoHandler := handlers.NewDefaultPhotoPreviewHandler(store.GetPhotoPreviewStore())
	r.POST("/api/v1/setting/photo/:settingid", photoHandler.CreatePhotoPreview)
	r.GET("/api/v1/setting/photo/", photoHandler.GetAllPhotoPreviews)
	r.GET("/api/v1/setting/photo/:id", photoHandler.GetPhotoPreview)
	r.PATCH("/api/v1/setting/photo/:id", photoHandler.UpdatePhotoPreview)
	r.DELETE("/api/v1/setting/photo/:id", photoHandler.DeletePhotoPreview)

	qualificationHandler := handlers.NewDefaultQualificationHandler(store.GetQualificationStore())
	r.POST("/api/v1/setting/qualification/:settingid", qualificationHandler.CreateQualification)
	r.GET("/api/v1/setting/qualification/", qualificationHandler.GetAllQualifications)
	r.GET("/api/v1/setting/qualification/:id", qualificationHandler.GetQualification)
	r.PATCH("/api/v1/setting/qualification/:id", qualificationHandler.UpdateQualification)
	r.DELETE("/api/v1/setting/qualification/:id", qualificationHandler.DeleteQualification)

	sponcerHandler := handlers.NewDefaultSponcerHandler(store.GetSponcerStore())
	r.POST("/api/v1/setting/sponcer/:settingid", sponcerHandler.CreateSponcer)
	r.GET("/api/v1/setting/sponcer/", sponcerHandler.GetAllSponcers)
	r.GET("/api/v1/setting/sponcer/:id", sponcerHandler.GetSponcer)
	r.PATCH("/api/v1/setting/sponcer/:id", sponcerHandler.UpdateSponcer)
	r.DELETE("/api/v1/setting/sponcer/:id", sponcerHandler.DeleteSponcer)

	timelineHandler := handlers.NewDefaultTimelineHandler(store.GetTimelineStore())
	r.POST("/api/v1/setting/timeline/:settingid", timelineHandler.CreateTimeline)
	r.GET("/api/v1/setting/timeline/", timelineHandler.GetAllTimelines)
	r.GET("/api/v1/setting/timeline/:id", timelineHandler.GetTimeline)
	r.PATCH("/api/v1/setting/timeline/:id", timelineHandler.UpdateTimeline)
	r.DELETE("/api/v1/setting/timeline/:id", timelineHandler.DeleteTimeline)

	if err := r.Listen(":" + config.App.Port); err != nil {
		panic(err)
	}
}

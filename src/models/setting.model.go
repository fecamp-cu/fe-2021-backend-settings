package models

import (
	"time"

	"gorm.io/gorm"
)

type Setting struct {
	gorm.Model
	ID         uint   `gorm:"primary_key" json:"id"`
	Title      string `gorm:"type:text;not null" json:"title"`
	YouTubeURL string `gorm:"type:text;not null" json:"youtube_url"`
	IsActive   bool   `gorm:"type:boolean;not null" json:"is_active"`
}
type About struct {
	gorm.Model
	Order     int     `gorm:"type:integer;not null" json:"order"`
	Text      string  `gorm:"type:text;not null" json:"text"`
	SettingID uint    `gorm:"type:integer;not null" json:"setting_id"`
	Setting   Setting `gorm:"foreignkey:SettingID;references:ID" json:"setting"`
}

type Qualification struct {
	gorm.Model
	Order     int     `gorm:"type:integer;not null" json:"order"`
	Text      string  `gorm:"type:text;not null" json:"text"`
	SettingID uint    `gorm:"type:integer;not null" json:"setting_id"`
	Setting   Setting `gorm:"foreignkey:SettingID;references:ID" json:"setting"`
}

type Sponcer struct {
	gorm.Model
	Order     int     `gorm:"type:integer;not null" json:"order"`
	Text      string  `gorm:"type:text;not null" json:"text"`
	SettingID uint    `gorm:"type:integer;not null" json:"setting_id"`
	Setting   Setting `gorm:"foreignkey:SettingID;references:ID" json:"setting"`
}

type PhotoPreview struct {
	gorm.Model
	Order     int     `gorm:"type:integer;not null" json:"order"`
	ImageUrl  string  `gorm:"type:text;not null" json:"image_url"`
	SettingID uint    `gorm:"type:integer;not null" json:"setting_id"`
	Setting   Setting `gorm:"foreignkey:SettingID;references:ID" json:"setting"`
}

type Timeline struct {
	gorm.Model
	Text           string    `gorm:"type:text;not null" json:"text"`
	EventStartDate time.Time `gorm:"type:date;not null" json:"event_start_date"`
	EventEndDate   time.Time `gorm:"type:date;not null" json:"event_end_date"`
	SettingID      uint      `gorm:"type:integer;not null" json:"setting_id"`
	Setting        Setting   `gorm:"foreignkey:SettingID;references:ID" json:"setting"`
}

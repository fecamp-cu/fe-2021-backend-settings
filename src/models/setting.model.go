package models

import (
	"time"

	"gorm.io/gorm"
)

type Setting struct {
	gorm.Model
	Title          string          `gorm:"type:text;not null" json:"title"`
	YouTubeURL     string          `gorm:"type:text;not null" json:"youtube_url"`
	IsActive       bool            `gorm:"type:boolean;not null" json:"is_active"`
	Abouts         []About         `gorm:"foreignkey:SettingID;references:ID" json:"about"`
	Qualifications []Qualification `gorm:"foreignkey:SettingID;references:ID" json:"qualification"`
	Sponcers       []Sponcer       `gorm:"foreignkey:SettingID;references:ID" json:"sponcer"`
	PhotoPreviews  []PhotoPreview  `gorm:"foreignkey:SettingID;references:ID" json:"photo_preview"`
	Timelines      []Timeline      `gorm:"foreignkey:SettingID;references:ID" json:"timeline"`
}
type About struct {
	gorm.Model
	Order     int    `gorm:"type:integer;not null" json:"order"`
	Text      string `gorm:"type:text;not null" json:"text"`
	SettingID uint   `gorm:"type:integer;not null" json:"setting_id"`
}

type Qualification struct {
	gorm.Model
	Order     int    `gorm:"type:integer;not null" json:"order"`
	Text      string `gorm:"type:text;not null" json:"text"`
	SettingID uint   `gorm:"type:integer;not null" json:"setting_id"`
}

type Sponcer struct {
	gorm.Model
	Order     int    `gorm:"type:integer;not null" json:"order"`
	Text      string `gorm:"type:text;not null" json:"text"`
	SettingID uint   `gorm:"type:integer;not null" json:"setting_id"`
}

type PhotoPreview struct {
	gorm.Model
	Order     int    `gorm:"type:integer;not null" json:"order"`
	ImageUrl  string `gorm:"type:text;not null" json:"image_url"`
	SettingID uint   `gorm:"type:integer;not null" json:"setting_id"`
}

type Timeline struct {
	gorm.Model
	Text           string    `gorm:"type:text;not null" json:"text"`
	EventStartDate time.Time `gorm:"type:date;not null" json:"event_start_date"`
	EventEndDate   time.Time `gorm:"type:date;not null" json:"event_end_date"`
	SettingID      uint      `gorm:"type:integer;not null" json:"setting_id"`
}

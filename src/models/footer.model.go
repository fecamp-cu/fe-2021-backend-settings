package models

type Footer struct {
	ID        uint   `gorm:"primary_key" json:"-"`
	Name      string `gorm:"type:text;not null" json:"name"`
	Place     string `gorm:"type:text;not null" json:"place"`
	Facebook  string `gorm:"type:text;not null" json:"facebook"`
	Instagram string `gorm:"type:text;not null" json:"instagram"`
	Youtube   string `gorm:"type:text;not null" json:"youtube"`
	Copyright string `gorm:"type:text;not null" json:"copy_right"`
}

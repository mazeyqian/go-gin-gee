package models

type WebSite struct {
	Name string `gorm:"column:name;not null;" json:"name"`
	Link string `gorm:"column:link;not null;" json:"link"`
	Code int    `gorm:"column:code;not null;" json:"code"`
}

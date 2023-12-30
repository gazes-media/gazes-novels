package models

import (
	"gorm.io/gorm"
)

type Novel struct {
	gorm.Model
	Title    string `json:"title"`
	Synopsis string `json:"author"`

	Chapters []Chapter `json:"chapters" gorm:"foreignKey:NovelID"`
}

type Chapter struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content" gorm:"type:text"`

	Novel   Novel `json:"novel" gorm:"foreignKey:NovelID"`
	NovelID uint  `json:"-"`
}

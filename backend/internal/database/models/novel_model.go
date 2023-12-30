package models

import (
	"github.com/gazes-media/gazes-novels/internal/database"
	"gorm.io/gorm"
)

type Novel struct {
	gorm.Model
	Title    string    `json:"title"`
	Synopsis string    `json:"synopsis"`
	Author   User      `json:"author" gorm:"foreignKey:AuthorID"`
	AuthorID uint      `json:"-"`
	Chapters []Chapter `json:"chapters" gorm:"foreignKey:NovelID"`
}

func CreateNovel(title, synopsis string) (*Novel, error) {
	novel := &Novel{
		Title:    title,
		Synopsis: synopsis,
	}

	if err := database.DB().Create(novel).Error; err != nil {
		return nil, err
	}

	return novel, nil
}

func (n *Novel) AddChapter(title, content string) (*Chapter, error) {
	chapter := &Chapter{
		Title:   title,
		Content: content,
		NovelID: n.ID,
	}

	tx := database.DB().Begin()
	if err := tx.Create(chapter).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	n.Chapters = append(n.Chapters, *chapter)

	if err := tx.Model(n).Association("Chapters").Append(chapter); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return chapter, nil
}

func GetNovelByID(id uint) (*Novel, error) {
	var novel Novel
	if err := database.DB().Preload("Chapters").First(&novel, id).Error; err != nil {
		return nil, err
	}

	return &novel, nil
}

func GetAllNovels() ([]Novel, error) {
	var novels []Novel
	if err := database.DB().Preload("Chapters").Find(&novels).Error; err != nil {
		return nil, err
	}

	return novels, nil
}

type Chapter struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content" gorm:"type:text"`

	Novel   Novel `json:"novel" gorm:"foreignKey:NovelID"`
	NovelID uint  `json:"-"`
}

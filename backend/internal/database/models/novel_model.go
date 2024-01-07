// Package models manages entities related to novels and chapters.
// It defines the Novel and Chapter structs along with methods to interact with novel-related data in the database.
package models

import (
	"fmt"

	"github.com/gazes-media/gazes-novels/internal/database"
	"gorm.io/gorm"
)

// Novel represents a literary work with its associated fields.
type Novel struct {
	gorm.Model           // gorm.Model provides ID, CreatedAt, UpdatedAt, DeletedAt fields
	Title      string    `json:"title"`                              // Title of the novel
	Synopsis   string    `json:"synopsis"`                           // Synopsis of the novel
	Author     User      `json:"author" gorm:"foreignKey:AuthorID"`  // Author of the novel
	AuthorID   uint      `json:"-"`                                  // ID of the author
	Chapters   []Chapter `json:"chapters" gorm:"foreignKey:NovelID"` // Chapters in the novel
}

// CreateNovel creates a new novel with the provided title, synopsis, and author ID.
// It returns a pointer to the created novel and an error if any.
func CreateNovel(title, synopsis string, authorID uint) (*Novel, error) {
	// Transaction begins for creating the novel
	tx := database.GetDB().Begin()

	novel := &Novel{
		Title:    title,
		Synopsis: synopsis,
		AuthorID: authorID,
	}

	if err := tx.Create(novel).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create novel: %w", err)
	}

	// Preloading author information
	if err := tx.Preload("Author").First(novel, novel.ID).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to preload author information: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return novel, nil
}

// AddChapter adds a new chapter to the novel with the provided title and content.
// It returns a pointer to the created chapter and an error if any.
func (n *Novel) AddChapter(title, content string) (*Chapter, error) {
	chapter := &Chapter{
		Title:   title,
		Content: content,
		NovelID: n.ID,
	}

	tx := database.GetDB().Begin()
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

// GetNovelByID retrieves a novel by its ID including its chapters.
// It returns a pointer to the novel and an error if any.
func GetNovelByID(id uint) (*Novel, error) {
	var novel Novel
	if err := database.GetDB().Preload("Chapters").First(&novel, id).Error; err != nil {
		return nil, err
	}

	return &novel, nil
}

// GetAllNovels retrieves all novels including their chapters.
// It returns a slice of novels and an error if any.
func GetAllNovels() ([]Novel, error) {
	var novels []Novel
	if err := database.GetDB().Preload("Chapters").Find(&novels).Error; err != nil {
		return nil, err
	}

	return novels, nil
}

// Chapter represents a section of a novel with its associated fields.
type Chapter struct {
	gorm.Model        // gorm.Model provides ID, CreatedAt, UpdatedAt, DeletedAt fields
	Title      string `json:"title"`                    // Title of the chapter
	Content    string `json:"content" gorm:"type:text"` // Content of the chapter (text)

	Novel   Novel `json:"novel" gorm:"foreignKey:NovelID"` // Novel to which the chapter belongs
	NovelID uint  `json:"-"`                               // ID of the novel
}

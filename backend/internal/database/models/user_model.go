// Package models provides structures and methods for managing user-related data.
// It defines the User struct and functions to interact with user data in the database.
package models

import (
	"time"

	"github.com/gazes-media/gazes-novels/internal/database"
	"gorm.io/gorm"
)

// User represents a user entity with its associated fields.
type User struct {
	gorm.Model           // gorm.Model provides ID, CreatedAt, UpdatedAt, DeletedAt fields
	Username   string    `json:"username"`                          // Username of the user
	Birthdate  time.Time `json:"birthdate"`                         // Birthdate of the user
	Novels     []Novel   `json:"novels" gorm:"foreignKey:AuthorID"` // Novels written by the user
}

// CreateUser creates a new user with the provided username and birthdate.
// It returns a pointer to the created user and an error if any.
func CreateUser(username string, birthdate time.Time) (*User, error) {
	user := &User{
		Username:  username,
		Birthdate: birthdate,
	}

	if err := database.DB().Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

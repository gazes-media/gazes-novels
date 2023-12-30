package models

import (
	"time"

	"github.com/gazes-media/gazes-novels/internal/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Birthdate time.Time `json:"birthdate"`
	Novels    []Novel   `json:"novels" gorm:"foreignKey:AuthorID"`
}

func CreateUser(username, email string, birthdate time.Time) (*User, error) {
	user := &User{
		Username:  username,
		Email:     email,
		Birthdate: birthdate,
	}

	if err := database.DB().Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

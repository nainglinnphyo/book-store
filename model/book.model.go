package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

// Users struct
type Books struct {
	Books []Book `json:"books"`
}

func (book *Book) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	book.ID = uuid.New()
	return
}

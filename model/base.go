package model

import (
	"gorm.io/gorm"
	"time"
)

// describes the structure.
type BaseModel struct {
	CreatedAt time.Time       `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" swaggertype:"string"`
}

type UriParse struct {
	ID []string `json:"id" uri:"id"`
}

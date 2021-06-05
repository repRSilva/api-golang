package models

import (
	"time"
)

type Users struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Age       uint      `json:"age"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
	DeletedAt time.Time `gorm:"index" json:"deleted"`
}

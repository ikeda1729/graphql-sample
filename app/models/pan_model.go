package models

import "time"

type Pan struct {
	ID        string    `gorm:"primary" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

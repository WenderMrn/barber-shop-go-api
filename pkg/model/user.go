package model

import (
	"time"
)

//User model
type User struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Document  string    `json:"document"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique_index"`
	Password  string    `json:"password"`
	Birthday  time.Time `json:"birthday"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

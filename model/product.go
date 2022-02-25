package model

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Id       uint `json:"id" gorm:"primaryKey"`
	CreateAt time.Time
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

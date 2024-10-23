package models

import "gorm.io/gorm"

type Item struct {
    gorm.Model
    Name     string  `json:"name" gorm:"not null"`
    Price    float64 `json:"price" gorm:"not null"`
    StoreID  uint    `json:"store_id" gorm:"not null"` // Foreign key to Store
}
package models

import "gorm.io/gorm"

type Store struct {
    gorm.Model
    Name    string  `json:"name" gorm:"not null"`
    UserID  uint    `json:"user_id" gorm:"not null"` // Foreign key to User
    Items   []Item  `json:"items" gorm:"foreignKey:StoreID"` // One-to-many relationship with Item
}
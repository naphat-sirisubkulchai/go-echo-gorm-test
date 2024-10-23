package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name     string  `json:"name"`
    Email    string  `json:"email" gorm:"unique"`
    Password string  `json:"-"`
    Stores   []Store `json:"stores" gorm:"foreignKey:UserID"` // One-to-many relationship with Store
}
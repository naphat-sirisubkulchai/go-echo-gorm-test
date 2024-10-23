package repositories

import (
    "get-echo-project/config"
    "get-echo-project/models"
	"fmt"
	"gorm.io/gorm"
)

func CreateUser(user *models.User) (*models.User, error) {
    // Create the user in the database
    if err := config.DB.Create(user).Error; err != nil {
        // Check if the error is due to duplicate email
        if err == gorm.ErrDuplicatedKey {
            return nil, fmt.Errorf("email already exists")
        }
        return nil, err // Return other errors
    }
    return user, nil // Return the created user with ID
}

func GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    err := config.DB.Where("email = ?", email).First(&user).Error
    return &user, err
}

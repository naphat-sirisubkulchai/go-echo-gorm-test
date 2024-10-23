package controllers

import (
    "get-echo-project/models"
    "get-echo-project/repositories"
    "get-echo-project/utils"
    "github.com/labstack/echo/v4"
    "net/http"
	"github.com/dgrijalva/jwt-go"
)

func RegisterUser(c echo.Context) error {
    user := new(models.User)
    if err := c.Bind(user); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid request")
    }

    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to hash password")
    }
    user.Password = hashedPassword

    createdUser, err := repositories.CreateUser(user)
    if err != nil {
        if err.Error() == "email already exists" {
            return c.JSON(http.StatusConflict, "Email already registered")
        }
        return c.JSON(http.StatusInternalServerError, "Failed to create user")
    }

    return c.JSON(http.StatusCreated, createdUser) // Return the created user
}
func GetProfile(c echo.Context) error {
    // Retrieve user from JWT token
    userToken := c.Get("user").(*jwt.Token)
    claims := userToken.Claims.(jwt.MapClaims)
    userEmail := claims["email"].(string)

    // Fetch user from the database by email
    user, err := repositories.GetUserByEmail(userEmail)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "User not found")
    }

    // Respond with the user's profile
    return c.JSON(http.StatusOK, user)
}
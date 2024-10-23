package controllers

import (
    "get-echo-project/models"
    "get-echo-project/repositories"
    "get-echo-project/utils"
    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo/v4"
    "net/http"
    "os"
    "time"
)

func LoginUser(c echo.Context) error {
    user := new(models.User)
    if err := c.Bind(user); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    // Fetch the user by email
    existingUser, err := repositories.GetUserByEmail(user.Email)
    if err != nil || !utils.CheckPasswordHash(user.Password, existingUser.Password) {
        return c.JSON(http.StatusUnauthorized, "Invalid credentials")
    }

    // Generate JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": existingUser.Email,
        "user_id": existingUser.ID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expiration
    })

    // Sign the token with the secret key
    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Could not generate token")
    }

    return c.JSON(http.StatusOK, echo.Map{
        "token": tokenString,
    })
}

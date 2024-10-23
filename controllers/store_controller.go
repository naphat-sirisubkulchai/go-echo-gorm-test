package controllers

import (
	"get-echo-project/models"
	"get-echo-project/repositories"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

// CreateStore creates a new store
func CreateStore(c echo.Context) error {
    // Parse user_id from the request body or query parameters
    userID, err := strconv.Atoi(c.FormValue("user_id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid user ID")
    }

    store := new(models.Store)
    if err := c.Bind(store); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    // Set the UserID for the store
    store.UserID = uint(userID)

    // Handle store creation with associated items
    if err := repositories.CreateStore(store); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to create store")
    }

    return c.JSON(http.StatusCreated, store)
}

// GetStores retrieves all stores
func GetStores(c echo.Context) error {
    stores, err := repositories.GetAllStores()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to retrieve stores")
    }
    return c.JSON(http.StatusOK, stores)
}
// GetStoresByUserID fetches stores owned by a user
func GetStoresByUserID(c echo.Context) error {
    userID, err := strconv.Atoi(c.Param("user_id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid user ID")
    }

    stores, err := repositories.GetStoresByUserID(uint(userID))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to retrieve stores")
    }
    return c.JSON(http.StatusOK, stores)
}
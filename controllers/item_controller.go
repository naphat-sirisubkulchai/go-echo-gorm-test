package controllers

import (
    "get-echo-project/models"
    "get-echo-project/repositories"
    "net/http"
	"strconv"
    "github.com/labstack/echo/v4"
)

// CreateItem creates a new item in a store
func CreateItem(c echo.Context) error {
    storeID, err := strconv.Atoi(c.Param("store_id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid store ID")
    }

    item := new(models.Item)
    if err := c.Bind(item); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    if err := repositories.CreateItem(uint(storeID), item); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to create item")
    }

    return c.JSON(http.StatusCreated, item)
}

// GetItemsByStore retrieves all items for a specific store
func GetItemsByStore(c echo.Context) error {
    storeID := c.Param("store_id")
    items, err := repositories.GetItemsByStore(storeID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to retrieve items")
    }
    return c.JSON(http.StatusOK, items)
}
func GetItemsByName(c echo.Context) error {
    itemName := c.Param("name")

    items, err := repositories.GetItemsByName(itemName)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to retrieve items")
    }
    return c.JSON(http.StatusOK, items)
}
func GetItemsByPriceRange(c echo.Context) error {
    lowPrice, err := strconv.ParseFloat(c.QueryParam("low"), 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid low price")
    }
    
    highPrice, err := strconv.ParseFloat(c.QueryParam("high"), 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid high price")
    }

    items, err := repositories.GetItemsByPriceRange(lowPrice, highPrice)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to retrieve items")
    }
    return c.JSON(http.StatusOK, items)
}
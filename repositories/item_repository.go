package repositories

import (
    "get-echo-project/config"
    "get-echo-project/models"
)

// CreateItem inserts a new item into the database
func CreateItem(storeID uint, item *models.Item) error {
    item.StoreID = storeID // Associate item with the store
    return config.DB.Create(item).Error
}

// GetItemsByStore retrieves all items for a specific store
func GetItemsByStore(storeID string) ([]models.Item, error) {
    var items []models.Item
    err := config.DB.Where("store_id = ?", storeID).Find(&items).Error
    return items, err
}
func GetItemsByName(name string) ([]models.Item, error) {
    var items []models.Item
    err := config.DB.Where("name = ?", name).Find(&items).Error
    return items, err
}
func GetItemsByPriceRange(low, high float64) ([]models.Item, error) {
    var items []models.Item
    err := config.DB.Where("price BETWEEN ? AND ?", low, high).Find(&items).Error
    return items, err
}
package repositories

import (
	"get-echo-project/config"
	"get-echo-project/models"
	"gorm.io/gorm"
)

// CreateStore inserts a new store into the database
func CreateStore(store *models.Store) error {
    // Create the store and associated items in one transaction
    return config.DB.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(store).Error; err != nil {
            return err
        }

        // If there are items in the store, save them
        for _, item := range store.Items {
            item.StoreID = store.ID // Set StoreID for each item
            if err := tx.Create(&item).Error; err != nil {
                return err
            }
        }

        return nil
    })
}

// GetAllStores retrieves all stores from the database
func GetAllStores() ([]models.Store, error) {
    var stores []models.Store
    err := config.DB.Find(&stores).Error
    return stores, err
}
func GetStoresByUserID(userID uint) ([]models.Store, error) {
    var stores []models.Store
    err := config.DB.Where("owner_id = ?", userID).Find(&stores).Error
    return stores, err
}
package routes

import (
    "get-echo-project/controllers"
    "get-echo-project/middlewares"
    "github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
    //User
    e.POST("/users/register", controllers.RegisterUser)
	e.GET("/users/profile", controllers.GetProfile, middlewares.JWTMiddleware())

	//Login
    e.POST("/users/login", controllers.LoginUser)


    // Store routes
    e.POST("/stores", controllers.CreateStore, middlewares.JWTMiddleware())
    e.GET("/stores", controllers.GetStores)
    e.GET("/stores/user/:user_id", controllers.GetStoresByUserID) // Get stores by user ID

    // Item routes
    e.POST("/stores/:store_id/items", controllers.CreateItem)
    e.GET("/stores/:store_id/items", controllers.GetItemsByStore)
    e.GET("/items/name/:name", controllers.GetItemsByName)           // Get items by name
    e.GET("/items/price-range", controllers.GetItemsByPriceRange)    // Get items by price range


}
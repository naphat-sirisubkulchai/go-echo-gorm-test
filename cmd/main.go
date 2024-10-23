package main

import (
    "get-echo-project/config"
    "get-echo-project/routes"
    "github.com/labstack/echo/v4"
)

func main() {
    // Initialize Echo
    e := echo.New()

    // Initialize DB connection
    config.ConnectDB()

    // Load API routes
    routes.InitRoutes(e)

    // Start server
    e.Start(":8080")
}

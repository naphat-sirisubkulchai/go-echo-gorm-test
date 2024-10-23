package config

import (
	"fmt"
	"get-echo-project/models"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
  )
  

var DB *gorm.DB

func ConnectDB() {
    dsn := fmt.Sprintf("host=%s port=%d user=%s "+
  "password=%s dbname=%s sslmode=disable",
  host, port, user, password, dbname)
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    } else {
        log.Println("Database connected")
    }
	err = DB.AutoMigrate(
        &models.User{},
        &models.Store{},
        &models.Item{},
    )
    if err != nil {
        log.Fatal("Failed to migrate database: ", err)
    }
}

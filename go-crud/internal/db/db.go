package db

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "go-crud/internal/models"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "host=db user=postgres password=postgres dbname=crudapp port=5432 sslmode=disable TimeZone=Africa/Lagos"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database: ", err)
    }

    // Auto-migrate the User model
    DB.AutoMigrate(&models.User{})
}

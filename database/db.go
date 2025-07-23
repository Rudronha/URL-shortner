package database

import (
    "log"
    "url-shortener/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "host=localhost user=test_user password=test@123 dbname=test_db port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }
    
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatal("Failed to get database instance: ", err)
    }
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetMaxIdleConns(10)

    // Auto-migrate table structure
    if err := DB.AutoMigrate(&models.URL{}); err != nil {
        log.Fatal("Failed to auto-migrate: ", err)
    }

    log.Println("Database connected and migrated!")
}
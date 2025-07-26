package database

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "url-shortener/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Failed to load .env file: ", err)
    }

    // Retrieve environment variables
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")
    sslmode := os.Getenv("DB_SSLMODE")

    // Construct DSN
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        host, user, password, dbname, port, sslmode)

    // Connect to database
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // Configure connection pool
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
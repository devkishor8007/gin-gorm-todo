package models

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
    "github.com/joho/godotenv"
)

var DB *gorm.DB

func Init() *gorm.DB {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading from .env file")
    }

    dbURL := os.Getenv("DB_URL")

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&Todo{})

    DB = db

	return DB;
}
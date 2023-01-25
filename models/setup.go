package models

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
    dbURL := "postgres://postgres:1234@localhost:5432/crud"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&Todo{})

    DB = db

	return DB;
}
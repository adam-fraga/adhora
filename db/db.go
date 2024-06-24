package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Client *gorm.DB
}

func (db *DB) NewConnection() *DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGPORT"),
	)
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Client = dbConn
	return db
}

func (db *DB) Close() {
	sqlDB, err := db.Client.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
	log.Println("Database connection closed")
}

func (db *DB) Ping() {
	sqlDB, err := db.Client.DB()
	if err != nil {
		log.Fatal(err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection successful")
}

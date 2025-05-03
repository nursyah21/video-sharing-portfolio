package database

import (
	"log"
	"os"
	"video-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=myuser password=mypassword dbname=video port=5432 sslmode=disable"
	}

	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to database")
}

func Drop() {
	err := Db.Migrator().DropTable(&models.Video{})
	if err != nil {
		log.Fatal("Drop Table failed:", err)
	}
	log.Println("Database drop table completed successfully")
}

func Migrate() {
	err := Db.AutoMigrate(&models.Video{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Database migration completed successfully")
}

func CloseDB() {
	sqlDB, err := Db.DB()
	if err != nil {
		log.Println("Failed to get SQL DB:", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Println("Failed to close database:", err)
	} else {
		log.Println("Database connection closed")
	}
}

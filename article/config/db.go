// package config

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func InitDB() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Println("No .env file found, using system environment")
// 	}

// 	host := os.Getenv("DB_HOST")
// 	port := os.Getenv("DB_PORT")
// 	user := os.Getenv("DB_USER")
// 	password := os.Getenv("DB_PASSWORD")
// 	dbname := os.Getenv("DB_NAME")

// 	dsn := fmt.Sprintf(
// 		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
// 		host, user, password, dbname, port,
// 	)

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("❌ Failed to connect to database: %v", err)
// 	}

// 	log.Println("✅ Connected to PostgreSQL using GORM")
// 	DB = db
// }

package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment")
	}

	// Get environment variables for DB configuration
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER") // DB_USER (root)
	dbname := os.Getenv("DB_NAME")

	// Construct DSN for MySQL (no password for root user)
	dsn := fmt.Sprintf(
		"%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, host, port, dbname,
	)

	// Open the connection to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	log.Println("✅ Connected to MySQL using GORM")
	DB = db
}

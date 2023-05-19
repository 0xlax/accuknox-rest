package database

import (
	"fmt"
	"log"
	"os"

	"github.com/0xlax/accuknox-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Dbinstance holds the database instance
type Dbinstance struct {
	Db *gorm.DB
}

// DB variable to store the Dbinstance
var DB Dbinstance

// ConnectDb establishes a connection to the PostgreSQL database using the provided environment variables.
// It initializes the database instance and performs necessary migrations.
func ConnectDb() {

	// Create the connection string using environment variables
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable Timezone=Asia/Kolkata",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// Open a connection to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to the database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected")

	// Set the logger mode for the database
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	// Perform necessary database migrations
	db.AutoMigrate(&models.Word{})

	// Store the database instance in the DB variable
	DB = Dbinstance{
		Db: db,
	}
}

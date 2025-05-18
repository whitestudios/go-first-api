package config

import (
	"os"

	"github.com/whitestudios/go-first-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlite() (*gorm.DB, error) {
	// Create logger
	logger := GetLogger("sqlite")
	dbDir := "./db"
	dbPath := "./db/openings.db"

	// Check if the database file exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")

		// Create db file and directory
		err = os.MkdirAll(dbDir, os.ModePerm)

		if err != nil {
			logger.Err("File to create db directory")
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			logger.Err("File to create db file")
			return nil, err
		}

		file.Close()
	}

	// Create DB and Connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errf("Sqlite opening error: %v", err)
		return nil, err
	}

	// Migrate Schemas
	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errf("Sqlite opening error: %v", err)
		return nil, err
	}

	// return the DB
	return db, nil
}

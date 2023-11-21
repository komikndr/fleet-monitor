package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// OpenDB opens the SQLite database and performs migrations.
// Example
// db, err := db.OpenDB("tasks.db")
func OpenDB(databasePath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&User{}, &Drone{}, &Task{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

package app

import (
	"github.com/jinzhu/gorm"
)

// DBMigrate check schema on db
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Name{})
	hasName := db.HasTable(&Name{})
	if !hasName {
		db.CreateTable(&Name{})
	}

	return db
}

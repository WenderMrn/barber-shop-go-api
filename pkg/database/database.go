package database

import (
	log "github.com/sirupsen/logrus"
	"github.com/wendermrn/barber-shop-go-api/pkg/model"

	"gorm.io/driver/sqlite" //import sqlite
	"gorm.io/gorm"
)

var db *gorm.DB

// GetDBG return alias Gorm connection
func GetDBG() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("/tmp/barbershopdb.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
		return nil, err
	}
	db.AutoMigrate(&model.Customer{})
	return db, nil
}

package database

import (
	log "github.com/sirupsen/logrus"
	"github.com/wendermrn/barber-shop-go-api/pkg/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //import sqlite
)

var db *gorm.DB

// GetDBG return alias Gorm connection
func GetDBG() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "/tmp/barbershopdb.db")
	if err != nil {
		log.Fatal("failed to connect database")
		return nil, err
	}
	db.AutoMigrate(&model.Customer{})
	return db, nil
}

package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Employee struct {
	gorm.Model
	Name   string `gorm:"unique" json:"name"`
	City   string `json:"city"`
	Age    int    `json:"age"`
	Status bool   `json:"status"`
}

func (e *Employee) Disable() {
	e.Status = false
}

func (e *Employee) Enable() {
	e.Status = true
}

//DB migrate will create and migrate the tables and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Employee{})
	return db
}

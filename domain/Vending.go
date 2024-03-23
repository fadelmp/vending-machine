// domain/Vending.go

package domain

import (
	"github.com/jinzhu/gorm"
)

// Database Design

// @Summary Vending object
// @Description Represents a vending machine menu
type Vending struct {
	gorm.Model
	Base  Base
	Id    uint   `gorm:"primaryKey;autoIncrement:true;Index"`
	Name  string `gorm:"type:VARCHAR(255);NOT NULL"`
	Price int    `gorm:"type:int;NOT NULL"`
}

// Set the custom table name for Title
func (v Vending) TableName() string {
	return "vending"
}

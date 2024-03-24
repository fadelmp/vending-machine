// domain/Vending.go

package domain

// Database Design

// @Summary Vending object
// @Description Represents a vending machine menu
type Vending struct {
	Base  Base
	Id    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"type:VARCHAR(255);notNull" json:"name"`
	Price int    `gorm:"type:int;notNull" json:"price"`
}

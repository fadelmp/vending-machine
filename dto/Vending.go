// dto/Vending.go

package dto

// Payload Design

// @Summary VendingDto object
// @Description Represents a VendingDto Machine VendingDto
type Vending struct {
	Base  Base
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

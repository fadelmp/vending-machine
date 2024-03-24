// dto/Vending.go

package dto

// Payload Design

// @Summary VendingDto object
// @Description Represents a VendingDto Machine VendingDto
type Vending struct {
	Id    uint   `json:"id"`
	Name  string `json:"name" validate:"required,min=1,max=30"`
	Price int    `json:"price" validate:"required,gte=0"`
	Base
}

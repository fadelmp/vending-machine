package dto

// Payload Design
type Vending struct {
	Base  Base
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

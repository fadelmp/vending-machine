// dto/Vending.go

package dto

import "time"

// Payload Design

// @Summary VendingDto object
// @Description Represents a VendingDto Machine VendingDto
type Vending struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name" validate:"required,min=1,max=50"`
	Price     int       `json:"price" validate:"required,gte=0"`
	IsActived bool      `json:"is_actived"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
	DeletedAt time.Time `json:"deleted_at"`
	DeletedBy string    `json:"deleted_by"`
}

// dto/Base.go

package dto

import "time"

// Payload Design

// @Summary BaseDto object
// @Description Represents a BaseDto Machine BaseDto
type Base struct {
	IsActived bool      `json:"is_actived"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
	DeletedAt time.Time `json:"deleted_at"`
	DeletedBy string    `json:"deleted_by"`
}

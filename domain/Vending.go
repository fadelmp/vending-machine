// domain/Vending.go

package domain

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Database Design

// @Summary Vending object
// @Description Represents a vending machine menu
type Vending struct {
	Id        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:VARCHAR(255);notNull" json:"name"`
	Price     int       `gorm:"type:int;notNull" json:"price"`
	IsActived bool      `gorm:"type:boolean;notNull" json:"is_actived"`
	IsDeleted bool      `gorm:"type:boolean;notNull" json:"is_deleted"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	CreatedBy string    `gorm:"type:VARCHAR(255)" json:"created_by"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
	UpdatedBy string    `gorm:"type:VARCHAR(255)" json:"updated_by"`
	DeletedAt time.Time `gorm:"type:timestamp" json:"deleted_at"`
	DeletedBy string    `gorm:"type:VARCHAR(255)" json:"deleted_by"`
}

func (v *Vending) BeforeSave(tx *gorm.DB) (err error) {

	v.IsActived = true
	v.IsDeleted = false
	v.UpdatedAt = time.Time{}
	v.DeletedAt = time.Time{}

	return
}

func (v *Vending) BeforeUpdate(tx *gorm.DB) (err error) {

	v.IsActived = true
	v.IsDeleted = false

	return
}

func (v *Vending) BeforeDelete(tx *gorm.DB) (err error) {

	v.IsActived = false
	v.IsDeleted = true

	return
}

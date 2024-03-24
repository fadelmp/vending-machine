// domain/Base.go

package domain

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Base Design

type Base struct {
	IsActived bool      `gorm:"type:boolean;notNull" json:"is_actived"`
	IsDeleted bool      `gorm:"type:boolean;notNull" json:"is_deleted"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	CreatedBy string    `gorm:"type:VARCHAR(255)" json:"created_by"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
	UpdatedBy string    `gorm:"type:VARCHAR(255)" json:"updated_by"`
	DeletedAt time.Time `gorm:"type:timestamp" json:"deleted_at"`
	DeletedBy string    `gorm:"type:VARCHAR(255)" json:"deleted_by"`
}

func (b *Base) BeforeSave(tx *gorm.DB) (err error) {

	b.IsActived = true
	b.IsDeleted = false
	b.UpdatedAt = time.Time{}
	b.DeletedAt = time.Time{}

	return
}

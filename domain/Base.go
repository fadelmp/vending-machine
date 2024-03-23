package domain

import "time"

type Base struct {
	IsActived bool      `gorm:"type:boolean;NOT NULL"`
	IsDeleted bool      `gorm:"type:boolean;NOT NULL"`
	CreatedAt time.Time `gorm:"type:timestamp;NOT NULL"`
	CreatedBy string    `gorm:"type:VARCHAR(255);NOT NULL"`
	UpdatedAt time.Time `gorm:"type:timestamp;NOT NULL"`
	UpdatedBy string    `gorm:"type:VARCHAR(255);NOT NULL"`
}

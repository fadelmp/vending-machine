package mapper

import (
	"vending-machine/domain"
	"vending-machine/dto"
)

// Interface
type BaseMapperContract interface {
	Create(*domain.Base, string)
	Update(*domain.Base, string)
	Delete(*domain.Base, string)

	ToBaseDto(base domain.Base) dto.Base
}

// Class
type BaseMapper struct {
}

// Constructor
func NewBaseMapper() *BaseMapper {
	return &BaseMapper{}
}

// Implementation

func (m *BaseMapper) Create(base *domain.Base, name string) {

	base.IsActived = true
	base.IsDeleted = false
	base.CreatedBy = name
	base.UpdatedBy = name
}

func (m *BaseMapper) Update(Base *domain.Base, name string) {

	Base.IsActived = true
	Base.IsDeleted = false
	Base.UpdatedBy = name
}

func (m *BaseMapper) Delete(Base *domain.Base, name string) {

	Base.IsActived = false
	Base.IsDeleted = true
	Base.UpdatedBy = name
	Base.DeletedBy = name
}

func (m *BaseMapper) ToBaseDto(base domain.Base) dto.Base {

	return dto.Base{
		IsActived: base.IsActived,
		IsDeleted: base.IsDeleted,
		CreatedAt: base.CreatedAt,
		CreatedBy: base.CreatedBy,
		UpdatedAt: base.UpdatedAt,
		UpdatedBy: base.UpdatedBy,
		DeletedAt: base.DeletedAt,
		DeletedBy: base.DeletedBy,
	}
}

package mapper

import (
	"vending-machine/domain"
	"vending-machine/dto"
)

// Interface
type VendingMapperContract interface {
	ToVendingDto(domain.Vending) dto.Vending
	ToVendingDtoList([]domain.Vending) []dto.Vending
	ToVending(dto.Vending, domain.Base) domain.Vending
}

// Class
type VendingMapper struct {
}

// Constructor
func NewVendingMapper() *VendingMapper {
	return &VendingMapper{}
}

// Implementation

func (m *VendingMapper) ToVendingDto(vending domain.Vending) dto.Vending {

	return dto.Vending{
		Id:    vending.Id,
		Name:  vending.Name,
		Price: vending.Price,
		Base:  ToBaseDto(vending.Base),
	}
}

func (m *VendingMapper) ToVendingDtoList(vendings []domain.Vending) []dto.Vending {

	vendingDtos := make([]dto.Vending, len(vendings))

	for i, value := range vendings {
		vendingDtos[i] = m.ToVendingDto(value)
	}

	return vendingDtos
}

func (m *VendingMapper) ToVending(dto dto.Vending, base domain.Base) domain.Vending {

	return domain.Vending{
		Id:    dto.Id,
		Name:  dto.Name,
		Price: dto.Price,
		Base:  base,
	}
}

package mapper

import (
	"time"
	"vending-machine/domain"
	"vending-machine/dto"
)

func ToBaseDto(domain domain.Base) dto.Base {

	return dto.Base{
		IsActived: domain.IsActived,
		IsDeleted: domain.IsDeleted,
		CreatedBy: domain.CreatedBy,
		UpdatedBy: domain.UpdatedBy,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func Create(name string) domain.Base {

	return domain.Base{
		IsActived: true,
		IsDeleted: false,
		CreatedBy: name,
		UpdatedBy: name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func Update(name string) domain.Base {

	return domain.Base{
		UpdatedBy: name,
		UpdatedAt: time.Now(),
	}
}

func Delete(name string) domain.Base {

	return domain.Base{
		IsActived: false,
		IsDeleted: true,
		UpdatedBy: name,
		UpdatedAt: time.Now(),
	}
}

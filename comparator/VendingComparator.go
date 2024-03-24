package comparator

import (
	"errors"
	"vending-machine/dto"
	"vending-machine/message"
	"vending-machine/repository"
)

// Interface
type VendingComparatorContract interface {
	CheckId(uint) error
	CheckName(dto.Vending) error
}

// Class
type VendingComparator struct {
	repo repository.VendingRepositoryContract
}

// Constructor
func NewVendingComparator(repo repository.VendingRepositoryContract) *VendingComparator {
	return &VendingComparator{
		repo: repo,
	}
}

// Implementation

func (c *VendingComparator) CheckId(id uint) error {

	// Get Data By Id
	vending := c.repo.GetById(id)

	// Return Error If Data Not Found
	if vending.Id == 0 {

		return errors.New(message.NotFound)
	}

	return nil
}

func (c *VendingComparator) CheckName(dto dto.Vending) error {

	// Get Data By Name
	vending := c.repo.GetByName(dto.Name)

	// Return Error If Data Exists
	if vending.Id != 0 && vending.Id != dto.Id {

		return errors.New(message.NameExists)
	}

	return nil
}

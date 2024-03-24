package comparator

import (
	"errors"
	"fmt"
	"vending-machine/dto"
	"vending-machine/message"
	"vending-machine/repository"
)

// Interface
type VendingComparatorContract interface {
	CheckName(dto.Vending) error
	CheckId(uint) error
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

func (c *VendingComparator) CheckName(dto dto.Vending) error {

	// Get data by name
	vending := c.repo.GetByName(dto.Name)

	fmt.Println("result")
	fmt.Println(vending)

	// Return error if data exists
	if vending.Id != 0 && vending.Id != dto.Id {
		return errors.New(message.NameExists)
	}

	return nil
}

func (c *VendingComparator) CheckId(id uint) error {

	// Get data by Id
	vending := c.repo.GetById(id)

	fmt.Println(id)
	fmt.Println("vending data")
	fmt.Println(vending)

	// Return error if data not found
	if vending.Id == 0 {
		fmt.Println("goto if else")
		return errors.New(message.NotFound)
	}

	return nil
}

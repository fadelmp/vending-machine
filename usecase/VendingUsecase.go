package usecase

import (
	"errors"
	"vending-machine/comparator"
	"vending-machine/dto"
	"vending-machine/mapper"
	"vending-machine/message"
	"vending-machine/repository"
)

// Interface
type VendingUsecaseContract interface {
	GetAll() []dto.Vending
	GetById(uint) dto.Vending

	Create(dto.Vending) error
	Update(dto.Vending) error
	Delete(dto.Vending) error
}

// Class
type VendingUsecase struct {
	mapper     mapper.VendingMapperContract
	comparator comparator.VendingComparatorContract
	repo       repository.VendingRepositoryContract
}

// Constructor
func NewVendingUsecase(
	repo repository.VendingRepositoryContract,
	mapper mapper.VendingMapperContract,
	comparator comparator.VendingComparatorContract) *VendingUsecase {
	return &VendingUsecase{
		repo:       repo,
		mapper:     mapper,
		comparator: comparator,
	}
}

// Implementation

func (u *VendingUsecase) GetAll() []dto.Vending {

	// Get all data
	vendings := u.repo.GetAll()

	// Map and Return Vending to Vending Dto
	return u.mapper.ToVendingDtoList(vendings)
}

func (u *VendingUsecase) GetById(id uint) dto.Vending {

	// Get By Id
	vending := u.repo.GetById(id)

	// Map and Return Vending to VendingDto
	return u.mapper.ToVendingDto(vending)
}

func (u *VendingUsecase) Create(dto dto.Vending) error {

	// Check Name and Return Error if Name Exists
	if err := u.comparator.CheckName(dto); err != nil {
		return err
	}

	// Create base data
	base := mapper.Create(dto.Base.CreatedBy)

	// map Vending dto to Vending domain
	vending := u.mapper.ToVending(dto, base)

	// create Vending
	_, err := u.repo.Create(vending)

	// Return Error if err not nil
	if err != nil {
		return errors.New(message.CreateFailed)
	}

	return err
}

func (u *VendingUsecase) Update(dto dto.Vending) error {

	// Check Id whether not found
	if err := u.comparator.CheckId(dto.Id); err != nil {
		return err
	}

	// Check Name whether name exists
	if err := u.comparator.CheckName(dto); err != nil {
		return err
	}

	// Update base data
	base := mapper.Update(dto.Base.CreatedBy)

	// Map Vending dto to Vending domain
	vending := u.mapper.ToVending(dto, base)

	// Update Vending and return
	_, err := u.repo.Update(vending)

	// Return Error if err not nil
	if err != nil {
		return errors.New(message.UpdateFailed)
	}

	return err
}

func (u *VendingUsecase) Delete(dto dto.Vending) error {

	// Check Id whether not found
	err := u.comparator.CheckId(dto.Id)
	if err != nil {
		return err
	}

	// Delete Base data
	base := mapper.Delete(dto.Base.CreatedBy)

	// Delete Vending and return
	err = u.repo.Delete(dto.Id, base)

	// Return Error if err not nil
	if err != nil {
		return errors.New(message.DeleteFailed)
	}

	return err
}

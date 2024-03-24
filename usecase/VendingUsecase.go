package usecase

import (
	"errors"
	"vending-machine/comparator"
	"vending-machine/domain"
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
	baseMapper mapper.BaseMapperContract
	mapper     mapper.VendingMapperContract
	comparator comparator.VendingComparatorContract
	repo       repository.VendingRepositoryContract
}

// Constructor
func NewVendingUsecase(
	baseMapper mapper.BaseMapperContract,
	mapper mapper.VendingMapperContract,
	repo repository.VendingRepositoryContract,
	comparator comparator.VendingComparatorContract) *VendingUsecase {
	return &VendingUsecase{
		baseMapper: baseMapper,
		mapper:     mapper,
		repo:       repo,
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

	// Map Vending Dto to Vending Domain
	vending := u.mapper.ToVending(dto)

	// Set Created Value
	u.baseMapper.Create(&vending.Base, dto.Base.CreatedBy)

	// Create Vending
	err := u.repo.Create(&vending)

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

	// Map Vending dto to Vending domain
	vending := u.mapper.ToVending(dto)

	u.baseMapper.Update(&vending.Base, dto.Base.UpdatedBy)

	// Update Vending and return
	err := u.repo.Update(&vending)

	// Return Error if err not nil
	if err != nil {
		return errors.New(message.UpdateFailed)
	}

	return err
}

func (u *VendingUsecase) Delete(dto dto.Vending) error {

	var vending domain.Vending
	// Check Id whether not found
	err := u.comparator.CheckId(dto.Id)
	if err != nil {
		return err
	}

	u.baseMapper.Delete(&vending.Base, dto.Base.UpdatedBy)

	// Delete Vending and return
	err = u.repo.Delete(&vending)

	// Return Error if err not nil
	if err != nil {
		return errors.New(message.DeleteFailed)
	}

	return err
}

package usecase_test

import (
	"testing"
	usecase "vending-machine/Usecase"
	"vending-machine/comparator"
	"vending-machine/dto"
	"vending-machine/mapper"
	"vending-machine/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestSuite represents the test suite for the VendingUsecase
type TestSuite struct {
	suite.Suite
	baseMapper *mapper.BaseMapper
	mapper     *mapper.VendingMapper
	repo       *repository.VendingRepository
	comparator *comparator.VendingComparator
}

func (suite *TestSuite) TestVendingUsecase_GetAll(t *testing.T) {

	// Setup
	vendingUsecase := usecase.NewVendingUsecase(suite.baseMapper, suite.mapper, suite.repo, suite.comparator)

	// Test
	vendings := vendingUsecase.GetAll()

	// Assert
	assert.NotNil(t, vendings)
	assert.NotEmpty(t, vendings)
}

func (suite *TestSuite) TestVendingUsecase_GetById(t *testing.T) {

	// Setup
	vendingUsecase := usecase.NewVendingUsecase(suite.baseMapper, suite.mapper, suite.repo, suite.comparator)

	// Test
	vendings := vendingUsecase.GetById(1)

	// Assert
	assert.NotNil(t, vendings)
	assert.NotEmpty(t, vendings)
}

func (suite *TestSuite) TestVendingUsecase_Create(t *testing.T) {

	// Setup
	vendingUsecase := usecase.NewVendingUsecase(suite.baseMapper, suite.mapper, suite.repo, suite.comparator)

	// VendingDto
	vendingDto := dto.Vending{
		Name:  "Lychee",
		Price: 1000,
		Base: dto.Base{
			CreatedBy: "System",
		},
	}

	// Test
	err := vendingUsecase.Create(vendingDto)

	// Assert
	assert.NoError(t, err)
}

func (suite *TestSuite) TestVendingUsecase_Update(t *testing.T) {

	// Setup
	vendingUsecase := usecase.NewVendingUsecase(suite.baseMapper, suite.mapper, suite.repo, suite.comparator)

	// VendingDto
	vendingDto := dto.Vending{
		Id:    1,
		Name:  "Lychee",
		Price: 1000,
		Base: dto.Base{
			UpdatedBy: "System",
		},
	}

	// Test
	err := vendingUsecase.Update(vendingDto)

	// Assert
	assert.NoError(t, err)
}

func (suite *TestSuite) TestVendingUsecase_Delete(t *testing.T) {

	// Setup
	vendingUsecase := usecase.NewVendingUsecase(suite.baseMapper, suite.mapper, suite.repo, suite.comparator)

	// VendingDto
	vendingDto := dto.Vending{
		Id: 1,
		Base: dto.Base{
			DeletedBy: "System",
		},
	}

	// Test
	err := vendingUsecase.Delete(vendingDto)

	// Assert
	assert.NoError(t, err)
}

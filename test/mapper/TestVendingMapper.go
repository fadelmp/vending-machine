package mapper_test

import (
	"testing"
	"time"
	"vending-machine/domain"
	"vending-machine/dto"
	"vending-machine/mapper"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestSuite represents the test suite for the VendingMapper
type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) TestVendingMapper_ToVending(t *testing.T) {

	// Setup
	vendingMap := mapper.NewVendingMapper()

	// VendingDto
	vendingDto := dto.Vending{
		Id:    1,
		Name:  "Lychee",
		Price: 1000,
	}

	// Test
	vending := vendingMap.ToVending(vendingDto)

	// Assert
	assert.NotNil(t, vending)
	assert.NotEmpty(t, vending)
}

func (suite *TestSuite) TestVendingMapper_ToVendingDto(t *testing.T) {

	// Setup
	vendingMap := mapper.NewVendingMapper()

	// VendingDto
	vending := domain.Vending{
		Id:    1,
		Name:  "Lychee",
		Price: 1000,
		Base: domain.Base{
			IsActived: true,
			IsDeleted: false,
			CreatedAt: time.Now(),
			CreatedBy: "System",
			UpdatedAt: time.Now(),
			UpdatedBy: "System",
			DeletedAt: time.Now(),
			DeletedBy: "System",
		},
	}

	// Test
	vendingDto := vendingMap.ToVendingDto(vending)

	// Assert
	assert.NotNil(t, vendingDto)
	assert.NotEmpty(t, vendingDto)
}

func (suite *TestSuite) TestVendingMapper_ToVendingDtoList(t *testing.T) {

	// Setup
	vendingMap := mapper.NewVendingMapper()

	var vendings []domain.Vending

	// VendingDto
	vending := domain.Vending{
		Id:    1,
		Name:  "Lychee",
		Price: 1000,
		Base: domain.Base{
			IsActived: true,
			IsDeleted: false,
			CreatedAt: time.Now(),
			CreatedBy: "System",
			UpdatedAt: time.Now(),
			UpdatedBy: "System",
			DeletedAt: time.Now(),
			DeletedBy: "System",
		},
	}

	vendings = append(vendings, vending)
	vendings = append(vendings, vending)
	vendings = append(vendings, vending)
	vendings = append(vendings, vending)
	vendings = append(vendings, vending)

	// Test
	vendingDtos := vendingMap.ToVendingDtoList(vendings)

	// Assert
	assert.NotNil(t, vendingDtos)
	assert.NotEmpty(t, vendingDtos)
}

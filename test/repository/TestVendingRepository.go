package repository_test

import (
	"testing"
	"time"
	"vending-machine/domain"
	"vending-machine/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestSuite represents the test suite for the VendingRepository
type TestSuite struct {
	suite.Suite
	repo *repository.VendingRepository
}

func (suite *TestSuite) TestVendingRepository_GetAll(t *testing.T) {

	// Setup
	repo := repository.NewVendingRepository(suite.repo.DB, suite.repo.Redis)

	// Test
	vendings := repo.GetAll()

	// Assert
	assert.NotNil(t, vendings)
	assert.NotEmpty(t, vendings)
}

func (suite *TestSuite) TestVendingRepository_GetById(t *testing.T) {

	// Setup
	repo := repository.NewVendingRepository(suite.repo.DB, suite.repo.Redis)

	// Test
	vending := repo.GetById(1)

	// Assert
	assert.NotNil(t, vending)
	assert.NotEmpty(t, vending)
}

func (suite *TestSuite) TestVendingRepository_GetByName(t *testing.T) {

	// Setup
	repo := repository.NewVendingRepository(suite.repo.DB, suite.repo.Redis)

	// Test
	vending := repo.GetByName("Milo")

	// Assert
	assert.NotNil(t, vending)
	assert.NotEmpty(t, vending)
}

func (suite *TestSuite) TestVendingRepository_Create(t *testing.T) {

	// Setup
	repo := repository.NewVendingRepository(suite.repo.DB, suite.repo.Redis)

	// Vending
	vending := &domain.Vending{
		Name:  "Lychee",
		Price: 1000,
		Base: domain.Base{
			IsActived: true,
			IsDeleted: false,
			CreatedAt: time.Now(),
			CreatedBy: "System",
			UpdatedAt: time.Now(),
			UpdatedBy: "System",
		},
	}

	// Test
	err := repo.Create(vending)

	// Assert
	assert.NoError(t, err)
}

func (suite *TestSuite) TestVendingRepository_Update(t *testing.T) {

	// Setup
	repo := repository.NewVendingRepository(suite.repo.DB, suite.repo.Redis)

	// Vending
	vending := &domain.Vending{
		Id:    1,
		Name:  "Lychee",
		Price: 1000,
		Base: domain.Base{
			IsActived: true,
			IsDeleted: false,
			UpdatedAt: time.Now(),
			UpdatedBy: "System",
		},
	}

	// Test
	err := repo.Update(vending)

	// Assert
	assert.NoError(t, err)
}

func (suite *TestSuite) TestVendingRepository_Delete(t *testing.T) {

	// Setup
	repo := repository.NewVendingRepository(suite.repo.DB, suite.repo.Redis)

	// Vending
	vending := &domain.Vending{
		Id: 1,
		Base: domain.Base{
			IsActived: false,
			IsDeleted: true,
			UpdatedAt: time.Now(),
			UpdatedBy: "System",
			DeletedAt: time.Now(),
			DeletedBy: "System",
		},
	}

	// Test
	err := repo.Delete(vending)

	// Assert
	assert.NoError(t, err)
}

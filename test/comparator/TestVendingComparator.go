package comparator_test

import (
	"testing"
	comparator "vending-machine/Comparator"
	"vending-machine/dto"
	"vending-machine/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestSuite represents the test suite for the VendingComparator
type TestSuite struct {
	suite.Suite
	repository repository.VendingRepository
}

func (suite *TestSuite) TestVendingComparator_CheckId(t *testing.T) {

	// Setup
	vendingComparator := comparator.NewVendingComparator(&suite.repository)

	// Test
	err := vendingComparator.CheckId(1)

	// Assert
	assert.NoError(t, err)
}

func (suite *TestSuite) TestVendingComparator_CheckName(t *testing.T) {

	// Setup
	vendingComparator := comparator.NewVendingComparator(&suite.repository)

	// VendingDto
	vendingDto := dto.Vending{
		Id:   1,
		Name: "Milo",
	}

	// Test
	err := vendingComparator.CheckName(vendingDto)

	// Assert
	assert.NoError(t, err)
}

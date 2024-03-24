package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	Handler "vending-machine/Handler"
	"vending-machine/usecase"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestSuite represents the test suite for the VendingHandler
type TestSuite struct {
	suite.Suite
	usecase *usecase.VendingUsecase
}

func (suite *TestSuite) TestVendingHandler_GetAll(t *testing.T) {

	// Setup
	vendingHandler := Handler.NewVendingHandler(suite.usecase)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/vendings", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	vendings := vendingHandler.GetAll(c)

	// Assert
	assert.NotNil(t, vendings)
	assert.NotEmpty(t, vendings)
}

func (suite *TestSuite) TestVendingHandler_GetById(t *testing.T) {

	// Setup
	vendingHandler := Handler.NewVendingHandler(suite.usecase)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/vendings/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	vendings := vendingHandler.GetById(c)

	// Assert
	assert.NotNil(t, vendings)
	assert.NotEmpty(t, vendings)
}

func (suite *TestSuite) TestVendingHandler_Create(t *testing.T) {

	// Setup
	vendingHandler := Handler.NewVendingHandler(suite.usecase)

	e := echo.New()
	requestBody := []byte(`{"name": "Milo", "price": 10000}`)
	req := httptest.NewRequest(http.MethodPost, "/vendings", bytes.NewBuffer(requestBody))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	vendings := vendingHandler.Create(c)

	// Assert
	assert.NotNil(t, vendings)
	assert.NotEmpty(t, vendings)
}

func (suite *TestSuite) TestVendingHandler_Update(t *testing.T) {

	// Setup
	vendingHandler := Handler.NewVendingHandler(suite.usecase)

	e := echo.New()
	requestBody := []byte(`{"name": "Milo", "price": 10000}`)
	req := httptest.NewRequest(http.MethodPost, "/vendings/1", bytes.NewBuffer(requestBody))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	vendings := vendingHandler.Update(c)

	// Assert
	assert.NotNil(t, vendings)
	assert.NotEmpty(t, vendings)
}

func (suite *TestSuite) TestVendingHandler_Delete(t *testing.T) {

	// Setup
	vendingHandler := Handler.NewVendingHandler(suite.usecase)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/vendings/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	vendings := vendingHandler.Delete(c)

	// Assert
	assert.NotNil(t, vendings)
	assert.NotEmpty(t, vendings)
}

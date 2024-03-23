package handler

import (
	"vending-machine/dto"
	"vending-machine/message"
	"vending-machine/usecase"

	"github.com/labstack/echo/v4"
)

// Interface
type VendingHandlerContract interface {
	GetAll(e echo.Context) error
	GetById(e echo.Context) error

	Create(e echo.Context) error
	Update(e echo.Context) error
	Delete(e echo.Context) error
}

// Class
type VendingHandler struct {
	usecase usecase.VendingUsecaseContract
}

// Constructor
func NewVendingHandler(usecase usecase.VendingUsecaseContract) *VendingHandler {
	return &VendingHandler{
		usecase: usecase,
	}
}

// @Summary Get All Vending
// @Description Retrieve a list of all Vending
// @Success 200 {object} dto.Response
func (h *VendingHandler) GetAll(e echo.Context) error {

	vendingDtos := h.usecase.GetAll()

	if len(vendingDtos) == 0 {
		return NotFound(e, message.NotFound)
	}

	return Success(e, message.GetSuccess, vendingDtos)
}

// @Summary Get Vending by ID
// @Description Retrieve an Vending from the vending machine by its ID
// @Success 200 {object} dto.Response
func (h *VendingHandler) GetById(e echo.Context) error {

	var vending dto.Vending

	if e.Bind(&vending) != nil {
		return BadRequest(e)
	}

	vendingDto := h.usecase.GetById(vending.Id)

	if vendingDto.Id == 0 {
		return NotFound(e, message.NotFound)
	}

	return Success(e, message.GetSuccess, vendingDto)
}

// @Summary Create a new Vending
// @Description Add a new Vending to the vending machine
// @Success 200 {object} dto.Response
func (h *VendingHandler) Create(e echo.Context) error {

	var Vending dto.Vending

	if e.Bind(&Vending) != nil {
		return BadRequest(e)
	}

	if err := h.usecase.Create(Vending); err != nil {
		return Error(e, err)
	}

	return Success(e, message.CreateSuccess, "")
}

// @Summary Update an existing Vending
// @Description Update an Vending in the vending machine by its ID
// @Success 200 {object} dto.Response
func (h *VendingHandler) Update(e echo.Context) error {

	var vending dto.Vending

	if e.Bind(&vending) != nil {
		return BadRequest(e)
	}

	if err := h.usecase.Update(vending); err != nil {
		return Error(e, err)
	}

	return Success(e, message.UpdateSuccess, "")
}

// @Summary Delete an dto.Vending
// @Description Remove an dto.Vending from the vending machine by its ID
// @Success 200 {object} dto.Response
func (h *VendingHandler) Delete(e echo.Context) error {

	var vending dto.Vending

	if e.Bind(&vending) != nil {
		return BadRequest(e)
	}

	if err := h.usecase.Delete(vending); err != nil {
		return Error(e, message.DeleteFailed)
	}

	return Success(e, message.DeleteSuccess, "")
}

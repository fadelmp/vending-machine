package handler

import (
	"errors"
	"vending-machine/dto"

	"github.com/go-playground/validator"
	_ "github.com/labstack/echo/v4"
)

func Validate(dto *dto.Vending) error {

	validate := validator.New()

	if err := validate.Struct(*dto); err != nil {
		return errors.New("bad request")
	}

	return nil
}

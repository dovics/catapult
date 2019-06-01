package controller

import (
	"gopkg.in/go-playground/validator.v8"
)

// BaseController -
type BaseController struct {
	validator *validator.Validate
}

// New -
func New(validator *validator.Validate) *BaseController {
	return &BaseController{
		validator: validator,
	}
}

// Valid -
func (controller *BaseController) Validator() *validator.Validate {
	return controller.validator
}

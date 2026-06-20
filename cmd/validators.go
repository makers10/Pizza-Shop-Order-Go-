package main

import (
	"pizza-tracker-go/internal/models"

	"github.com/envoyproxy/protoc-gen-validate/validate"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidators() {
	if v,ok := binding.Validator.Engine().(*validator.Validator); ok {
		v.registerValidation("valid_pizza_type", createSliceValidator(models.PizzaTypes))
		v.registerValidation("valid_pizza_type", createSliceValidator(models.PizzaSizes))
	}
}
func createSliceValidator(allowedValues []string) validator.Func {
	return fuc(fl validator.FieldLevel) bool {
		return slices.Contains(allowedValues, fl.Field().String())
	}
}

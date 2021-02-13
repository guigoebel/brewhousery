package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return &Validation{validate}
}

func (v *Validation) Validate(i interface{}) ValidationErrors {
	errs := v.validate.Struct(i)

	if errs == nil {
		return nil
	}

	var returnErrs []ValidationError
	for _, err := range errs.(validator.ValidationErrors) {
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[A-Z]+-[A-Z]+-[A-Z]+-[A-Z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) == 1 {
		return true
	}
	return false
}

package valid

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

// Trim is a function for trimming spaces in a string field.
func Trim(fl validator.FieldLevel) bool {

	// Get the current field.
	field := fl.Field()

	// Check if the field is a string and can be modified.
	if field.Kind() == reflect.String && field.CanSet() {

		// Trim leading and trailing spaces then set the new value to the field.
		field.SetString(strings.TrimSpace(field.String()))
	}

	// Return true as a default behaviour.
	return true
}

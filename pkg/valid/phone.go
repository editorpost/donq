package valid

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/nyaruka/phonenumbers"
	"reflect"
	"strings"
)

const (
	// PhoneLength limit following E.164
	PhoneLength = 20
)

var (
	ErrPhoneInvalid = errors.New("contact.phone.invalid")
)

// Phone normalizes and validates an phone number.
// Length must be upto 20 characters (E.164).
// Original value replaced with empty string in case of invalid.
func Phone(fl validator.FieldLevel) bool {

	// Get the current field.
	field := fl.Field()

	// Check if the field is a string and can be modified.
	if field.Kind() == reflect.String && field.CanSet() {

		// Trim redundant spaces, validate, and sanitize the phone address.
		refinedPhone, err := SanitizePhone(field.String())

		// Update the phone field with the sanitized value.
		field.SetString(refinedPhone)

		// invalid phone
		if err != nil {
			return false
		}

		if len(refinedPhone) > PhoneLength {
			return false
		}

	}

	// Default response.
	return true
}

// SanitizePhone trim, validate and sanitize phone number
// The method expects international format phone string
func SanitizePhone(input string) (string, error) {

	number := strings.TrimSpace(input)

	// add leading plus
	if number[0] != '+' {
		number = "+" + number
	}

	phone, err := phonenumbers.Parse(number, "US")

	if err != nil {
		return number, err
	}

	if !phonenumbers.IsValidNumber(phone) {
		return number, err
	}

	// replace phone with normalized string
	return phonenumbers.Format(phone, phonenumbers.E164), nil
}

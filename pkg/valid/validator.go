package valid

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
)

var (
	Validate *validator.Validate
)

func init() {

	Validate = validator.New(validator.WithRequiredStructEnabled())

	// aliases it is a set of predefined validators
	RegisterExtraAliases()

	// Register custom validators.
	validators := map[string]func(fl validator.FieldLevel) bool{
		"phone": Phone,
		"trim":  Trim,
	}

	for name, fn := range validators {
		if err := Validate.RegisterValidation(name, fn); err != nil {
			log.Fatal(err)
			return
		}
	}
}

// Struct validates a structs exposed fields, and automatically validates nested structs, unless otherwise specified.
func Struct(s any) (err error) {
	return Validate.Struct(s)
}

// StructPartial validates a structs exposed fields, and automatically validates nested structs, unless otherwise specified.
func StructPartial(s any, fields ...string) (err error) {
	return Validate.StructPartial(s, fields...)
}

// RegisterExtraAliases registers extra aliases for existing validators.
func RegisterExtraAliases() {
	Validate.RegisterAlias("pass", fmt.Sprintf("trim,min=6,max=128"))
	Validate.RegisterAlias("bcrypt", fmt.Sprintf("trim,max=128"))
	Validate.RegisterAlias("emailx", fmt.Sprintf("trim,email,max=64"))
	Validate.RegisterAlias("lang", fmt.Sprintf("bcp47_language_tag"))
	Validate.RegisterAlias("name", fmt.Sprintf("trim,max=64"))
	Validate.RegisterAlias("optional", fmt.Sprintf("trim,omitempty"))
}

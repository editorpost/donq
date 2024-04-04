package valid

import (
	"fmt"
)

// PasswordAlias creates "pass" alias combining trim and length check.
// In case of invalid returned trimmed original value.
func PasswordAlias(min, max int) string {
	return fmt.Sprintf("trim,min=%d,max=%d", min, max)
}

// BCryptAlias creates "bcrypt" alias combining trim and length check.
// In case of invalid returned trimmed original value.
func BCryptAlias() {
	Validate.RegisterAlias("bcrypt", "trim,max=128")
}

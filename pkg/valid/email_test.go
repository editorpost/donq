package valid_test

import (
	"github.com/editorpost/donq/pkg/valid"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"testing"
)

type (
	EmailCase struct {
		name     string
		input    string
		expected string
		isValid  bool
	}
	EmailSuspected struct {
		Field string `validate:"email"`
	}
	EmailOptional struct {
		Field string `validate:"trim,omitempty,email"`
	}
	EmailTrimmed struct {
		Field string `validate:"trim,email"`
	}
)

var (
	EmailDefaultCases = []EmailCase{
		{"empty", "", "", false},
		{"valid email", "test@example.com", "test@example.com", true},
		{"local", "test@example", "test@example", false},
		{"no@", "test$example.com", "test$example.com", false},
		{"double", "test@example.comtest@example.com", "test@example.comtest@example.com", false},
	}

	EmailTrimmedCases = []EmailCase{
		{"trim space", " test@example.com ", "test@example.com", true},
		{"trim tab", "	test@example.com	", "test@example.com", true},
		{"trim tab regex", "\ttest@example.com", "test@example.com", true},
		{"double", "\ttest@example.comtest@example.com", "test@example.comtest@example.com", false},
	}

	EmailOptionalCases = []EmailCase{
		{"empty", "", "", true},
		{"space", " ", "", true},
	}
)

// TestEmail table-tests EmailX function with valid.Struct
func TestEmail(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("trim", valid.Trim)
	require.NoError(t, err)

	for _, tc := range EmailDefaultCases {
		t.Run(tc.name, func(t *testing.T) {
			emailTest := &EmailSuspected{Field: tc.input}
			err = validate.Struct(emailTest)

			require.Equal(t, tc.isValid, err == nil, "Expected validity: %v, got error: %v", tc.isValid, err)
			require.Equal(t, tc.expected, emailTest.Field, "Expected: %s, got: %s", tc.expected, emailTest.Field)
		})
	}
}

// TestEmail table-tests EmailX function with valid.Struct
func TestEmailTrimmed(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("trim", valid.Trim)
	require.NoError(t, err)

	for _, tc := range EmailTrimmedCases {
		t.Run(tc.name, func(t *testing.T) {
			emailTest := &EmailTrimmed{Field: tc.input}
			err = validate.Struct(emailTest)

			require.Equal(t, tc.isValid, err == nil, "Expected validity: %v, got error: %v", tc.isValid, err)
			require.Equal(t, tc.expected, emailTest.Field, "Expected: %s, got: %s", tc.expected, emailTest.Field)
		})
	}
}

// TestEmail table-tests EmailX function with valid.Struct
func TestEmailOptional(t *testing.T) {

	validate := validator.New()
	err := validate.RegisterValidation("trim", valid.Trim)
	require.NoError(t, err)

	for _, tc := range EmailOptionalCases {
		t.Run(tc.name, func(t *testing.T) {
			emailTest := &EmailOptional{Field: tc.input}
			err = validate.Struct(emailTest)

			require.Equal(t, tc.isValid, err == nil, "Expected validity: %v, got error: %v", tc.isValid, err)
			require.Equal(t, tc.expected, emailTest.Field, "Expected: %s, got: %s", tc.expected, emailTest.Field)
		})
	}
}

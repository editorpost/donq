package valid_test

import (
	"github.com/editorpost/donq/pkg/valid"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type phoneCase struct {
	name     string
	input    string
	expected string
	isValid  bool
}

func PhoneCases(t *testing.T) []phoneCase {
	t.Helper()
	return []phoneCase{
		{"trim space", " +66909808909 ", "+66909808909", true},
		{"trim tab", "	+66909808909	", "+66909808909", true},
		{"trim tab regex", "\t+66909808909", "+66909808909", true},
		{"valid phone", "+66909808909", "+66909808909", true},
		{"invalid phone", "\t+669098 ", "+669098", false},
	}
}

// TestPhoneX shows difference between phone and phonex
func TestPhoneXLeadingPlus(t *testing.T) {

	validate := validator.New()
	err := validate.RegisterValidation("trim", valid.Trim)
	require.NoError(t, err)
	err = validate.RegisterValidation("phone", valid.Phone)
	require.NoError(t, err)

	type PlusX struct {
		Phone string `validate:"phone"`
	}

	type Plus struct {
		Phone string `validate:"trim,e164"`
	}

	input := "66909808909"
	phone := validate.Struct(Plus{input})
	phonex := validate.Struct(PlusX{input})

	assert.Error(t, phone)
	assert.NoError(t, phonex)
}

// TestPhoneX table-tests TestPhoneX function with valid.Struct
// Cases:
// trim space " +66909808909 " => "+66909808909", nil
// valid email "+66909808909" => "+66909808909", nil
// invalid email " +669098 " => "+669098", nil
func TestPhoneX(t *testing.T) {

	validate := validator.New()
	err := validate.RegisterValidation("phonex", valid.Phone)
	require.NoError(t, err)

	type TestPhoneXPayload struct {
		Field string `validate:"phonex"`
	}

	cases := append(PhoneCases(t), phoneCase{"leading plus", "66909808909", "+66909808909", true})

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			phoneTest := &TestPhoneXPayload{Field: tc.input}
			err = validate.Struct(phoneTest)

			require.Equal(t, tc.isValid, err == nil, "Expected validity: %v, got error: %v", tc.isValid, err)
			require.Equal(t, tc.expected, phoneTest.Field, "Expected: %s, got: %s", tc.expected, phoneTest.Field)
		})
	}
}

// TestPhoneX table-tests TestPhoneX function with valid.Struct
// Cases:
// trim space " +66909808909 " => "+66909808909", nil
// valid email "+66909808909" => "+66909808909", nil
// invalid email " +669098 " => "+669098", nil
func TestPhone(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("trim", valid.Trim)
	require.NoError(t, err)

	type TestPhonePayload struct {
		Field string `validate:"trim,e164"`
	}

	for _, tc := range PhoneCases(t) {
		t.Run(tc.name, func(t *testing.T) {
			phoneTest := &TestPhonePayload{Field: tc.input}
			err = validate.Struct(phoneTest)

			require.Equal(t, tc.isValid, err == nil, "Expected validity: %v, got error: %v", tc.isValid, err)
			require.Equal(t, tc.expected, phoneTest.Field, "Expected: %s, got: %s", tc.expected, phoneTest.Field)
		})
	}
}

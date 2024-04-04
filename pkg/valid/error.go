package valid

import "github.com/go-playground/validator/v10"

func ErrorIs(err error) (validator.ValidationErrors, bool) {
	invalid, ok := err.(validator.ValidationErrors)
	return invalid, ok
}

// AssertFieldError checks if a field error with the given tag exists.
func AssertFieldError(err error, field string, tags ...string) bool {

	errs, ok := ErrorIs(err)
	if !ok {
		return false
	}

	hasTags := len(tags) > 0

	for _, e := range errs {

		if e.Field() != field {
			continue
		}

		if !hasTags {
			return true
		}

		for _, tag := range tags {
			if e.Tag() == tag {
				return true
			}
		}

		return false
	}
	return ok
}

// AssertFieldErrorFn returns a function that checks if a field error with the given tag exists.
func AssertFieldErrorFn(field string, tags ...string) func(error) bool {
	return func(err error) bool {
		return AssertFieldError(err, field, tags...)
	}
}

package valid_test

import (
	"github.com/editorpost/donq/pkg/valid"
	"testing"
)

type (
	subject struct {
		Name  string `validate:"trim"`
		Count int    `validate:"trim"`
	}
)

// TestTrim tests Trim function with validate.Struct
// Expected result: trimmed string " test " to "test"
func TestTrim(t *testing.T) {
	s := subject{
		Name: " test ",
	}
	err := valid.Struct(&s)
	if err != nil {
		t.Error(err)
		return
	}
	if s.Name != "test" {
		t.Errorf("Expected: %s, got: %s", "test", s.Name)
		return
	}
}

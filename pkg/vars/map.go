package vars

import (
	"encoding/json"
	"github.com/editorpost/donq/pkg/valid"
)

// FromJSON any struct from JSON data
// Usage:
//
//	var res YourType
//	data, err := wmill.GetResource(name)
//	return FromJSON[T](data)
func FromJSON[T any](data any, typ *T) error {

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, typ)
	if err != nil {
		return err
	}

	err = valid.Validate.Struct(typ)
	if err != nil {
		return err
	}

	return nil
}

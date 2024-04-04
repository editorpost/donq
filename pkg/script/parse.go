package script

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/editorpost/donq/pkg/valid"
)

func ParseArgs[T any](from any, to *T) error {

	m, ok := from.(map[string]interface{})
	if !ok {
		return errors.New("invalid input arguments")
	}

	// Convert the map to JSON
	data, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("failed to marshal input arguments: %w", err)
	}

	// Convert the JSON to a struct
	if err = json.Unmarshal(data, to); err != nil {
		return fmt.Errorf("failed to unmarshal input arguments: %w", err)
	}

	// validate struct
	return valid.Validate.Struct(to)
}

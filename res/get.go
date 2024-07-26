package res

import (
	"encoding/json"
	wmill "github.com/windmill-labs/windmill-go-client"
)

// Unmarshal resource from Windmill by path.
func Unmarshal[T any](path string) (*T, error) {

	b, err := Load(path)
	if err != nil {
		return nil, err
	}

	res := new(T)

	if err = json.Unmarshal(b, res); err != nil {
		return nil, err
	}

	return res, nil
}

func Load(path string) ([]byte, error) {
	data, err := wmill.GetResource(path)
	if err != nil {
		return nil, err
	}
	return json.Marshal(data)
}

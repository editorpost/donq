package vars

import (
	wmill "github.com/windmill-labs/windmill-go-client"
)

func FromResource[T any](name string, to *T) error {

	data, err := wmill.GetResource(name)
	if err != nil {
		return err
	}

	return FromJSON[T](data, to)
}

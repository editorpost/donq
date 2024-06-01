package vars

import (
	"github.com/caarlos0/env/v10"
	"sync"
)

var (
	_env     = &Windmill{}
	_envOnce = &sync.Once{}
)

// FromEnv loads the Windmill config instance from the environment.
// Variables loaded with sync.Once.
func FromEnv() *Windmill {

	_envOnce.Do(func() {
		if err := env.Parse(_env); err != nil {
			panic(err)
		}
	})

	return _env
}

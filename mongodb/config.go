package mongodb

import (
	"fmt"
	wmill "github.com/windmill-labs/windmill-go-client"
)

// Windmill Mongo Config example:
//	{
//		"host": "mongo",
//		"port": 27017,
//		"credential": {
//		"password": "nopass",
//			"username": "root"
//		}
//	}

type Config struct {
	Db         string `json:"db"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Credential struct {
		Password string `json:"password"`
		Username string `json:"username"`
	} `json:"credential"`
	DSN string `json:"dsn"`
}

func MustResource(name string) *Config {
	res, err := GetResource(name)
	if err != nil {
		panic(err)
	}
	return res
}

func GetResource(name string) (*Config, error) {
	res, _ := wmill.GetResource(name)
	secret, ok := res.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("invalid spider mongo resource type: %T", res)
	}

	auth, ok := secret["credential"].(map[string]string)
	if !ok {
		return nil, fmt.Errorf("invalid spider mongo credential type: %T", secret["credential"])
	}

	return &Config{
		Db:   secret["db"].(string),
		Host: secret["host"].(string),
		Port: secret["port"].(int),
		Credential: struct {
			Password string `json:"password"`
			Username string `json:"username"`
		}{
			Password: auth["password"],
			Username: auth["username"],
		},
		DSN: fmt.Sprintf("mongodb://%s:%s@%s:%d",
			auth["username"],
			auth["password"],
			secret["host"],
			secret["port"],
		),
	}, nil
}

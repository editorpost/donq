package mongodb

import (
	"fmt"
	wmill "github.com/windmill-labs/windmill-go-client"
)

// Windmill Mongo Config example:
//	{
//  "db": "spider",
//  "tls": false,
//  "credential": {
//    "password": "nopass",
//    "username": "root"
//  },
//  "servers": [
//    {
//      "host": "mongo",
//      "port": 27017
//    }
//  ]
//}
//}

type (
	Config struct {
		Db   string `json:"db"`
		Host string `json:"host"`
		Port int    `json:"port"`
		DSN  string `json:"dsn"`
		User string `json:"password"`
		Pass string `json:"username"`
	}
)

func MustResource(name string) *Config {
	res, err := GetResource(name)
	if err != nil {
		panic(err)
	}
	return res
}

func GetResource(name string) (*Config, error) {
	res, _ := wmill.GetResource(name)
	return ConfigFromResource(res)
}

func ConfigFromResource(res any) (*Config, error) {

	secret, ok := res.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("invalid spider mongo resource type: %T", res)
	}

	servers, ok := secret["servers"].([]any)
	if !ok {
		return nil, fmt.Errorf("invalid spider mongo servers type: %T", secret["servers"])
	}

	server := map[string]any{}
	for _, srv := range servers {
		server, ok = srv.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("invalid spider mongo server type: %T", server)
		}
		break
	}

	credentials, ok := secret["credential"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("invalid spider mongo credential type: %T", secret["credential"])
	}

	return &Config{
		Db:   secret["db"].(string),
		Host: server["host"].(string),
		Port: int(server["port"].(float64)),
		User: credentials["username"].(string),
		Pass: credentials["password"].(string),
		DSN: fmt.Sprintf("mongodb://%s:%s@%s:%d",
			credentials["username"],
			credentials["password"],
			server["host"],
			int(server["port"].(float64)),
		),
	}, nil
}

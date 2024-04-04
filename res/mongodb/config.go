package mongodb

import (
	"encoding/json"
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
	// Config for MongoDB connection
	Config struct {
		Db   string `json:"db"`
		Host string `json:"host"`
		Port int    `json:"port"`
		DSN  string `json:"dsn"`
		User string `json:"password"`
		Pass string `json:"username"`
		TLS  bool   `json:"tls"`
	}
)

// MustResource for MongoDB from Windmill Resources or panic on error
func MustResource(name string) *Config {
	res, err := GetResource(name)
	if err != nil {
		panic(err)
	}
	return res
}

// GetResource for MongoDB from Windmill Resources
func GetResource(name string) (*Config, error) {
	res, _ := wmill.GetResource(name)
	return ConfigFromResource(res)
}

// ConfigFromResource parse Windmill resource to MongoDB Config
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

	cred, ok := secret["credential"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("invalid spider mongo credential type: %T", secret["credential"])
	}

	port := ParseInt(server["port"])
	if port == 0 {
		return nil, fmt.Errorf("invalid spider mongo port type: %T", server["port"])
	}

	return &Config{
		Db:   secret["db"].(string),
		Host: server["host"].(string),
		Port: port,
		User: cred["username"].(string),
		Pass: cred["password"].(string),
		DSN: fmt.Sprintf("mongodb://%s:%s@%s:%d",
			cred["username"],
			cred["password"],
			server["host"],
			port,
		),
	}, nil
}

// ParseInt from any type
func ParseInt(i any) int {
	switch v := i.(type) {
	case int:
		return v
	case int64:
		return int(v)
	case float64:
		return int(v)
	case json.Number:
		n, _ := v.Int64()
		return int(n)
	default:
		return 0
	}
}

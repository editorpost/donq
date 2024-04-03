package mongodb_test

import (
	"fmt"
	"testing"

	"github.com/editorpost/donq/mongodb"
	"github.com/stretchr/testify/assert"
)

func TestConfigFromResource(t *testing.T) {
	validResource := map[string]interface{}{
		"db": "exampleDb",
		"servers": []interface{}{
			map[string]interface{}{
				"host": "exampleHost",
				"port": 1234,
				"credential": map[string]interface{}{
					"password": "examplePassword",
					"username": "exampleUsername",
				},
			},
		},
	}

	config, err := mongodb.ConfigFromResource(validResource)

	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, config, "Expected non-nil config")

	assert.Equal(t, "exampleDb", config.Db, "Expected db value to match")
	assert.Equal(t, "exampleHost", config.Host, "Expected host value to match")
	assert.Equal(t, 1234, config.Port, "Expected port value to match")
	assert.Equal(t, "examplePassword", config.Credential.Password, "Expected password value to match")
	assert.Equal(t, "exampleUsername", config.Credential.Username, "Expected username value to match")
	assert.Equal(t, "mongodb://exampleUsername:examplePassword@exampleHost:1234", config.DSN, "Expected DSN value to match")
}

func TestConfigFromResource_InvalidResourceType(t *testing.T) {
	invalidResource := "invalid resource"

	config, err := mongodb.ConfigFromResource(invalidResource)

	assert.Error(t, err, "Expected an error")
	assert.Nil(t, config, "Expected nil config")

	expectedError := fmt.Sprintf("invalid spider mongo resource type: %T", invalidResource)
	assert.EqualError(t, err, expectedError, "Expected error message to match")
}

func TestConfigFromResource_InvalidServersType(t *testing.T) {
	invalidResource := map[string]interface{}{
		"servers": "invalid servers type",
	}

	config, err := mongodb.ConfigFromResource(invalidResource)

	assert.Error(t, err, "Expected an error")
	assert.Nil(t, config, "Expected nil config")
}

func TestConfigFromResource_InvalidServerType(t *testing.T) {
	invalidResource := map[string]interface{}{
		"servers": []interface{}{"invalid server type"},
	}

	config, err := mongodb.ConfigFromResource(invalidResource)

	assert.Error(t, err, "Expected an error")
	assert.Nil(t, config, "Expected nil config")
}

func TestConfigFromResource_InvalidCredentialType(t *testing.T) {
	invalidResource := map[string]interface{}{
		"servers": []interface{}{
			map[string]interface{}{
				"credential": "invalid credential type",
			},
		},
	}

	config, err := mongodb.ConfigFromResource(invalidResource)

	assert.Error(t, err, "Expected an error")
	assert.Nil(t, config, "Expected nil config")
}

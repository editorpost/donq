package vars_test

import (
	"github.com/editorpost/donq/pkg/vars"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Args struct {
	// StartURL is the URL to start crawling, e.g. http://example.com
	StartURL string `json:"StartURL"`
	// AllowedURL is the regex to match the URLs, e.g. ".*"
	MatchURL string `json:"AllowedURL"`
	// Depth is the number of levels to follow the links
	Depth int `json:"Depth"`
	// ExtractSelector CSS to match the entities to extract, e.g. ".article--ssr"
	Selector string `json:"ExtractSelector"`
}

func TestFromJSON(t *testing.T) {

	args := Args{}
	var input any = map[string]interface{}{
		"StartURL":        "http://example.com",
		"AllowedURL":      ".*",
		"Depth":           1,
		"ExtractSelector": ".article--ssr",
	}

	// Test the Parse function
	err := vars.FromJSON(input, &args)

	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, "http://example.com", args.StartURL)
	assert.Equal(t, ".*", args.MatchURL)
	assert.Equal(t, 1, args.Depth)
	assert.Equal(t, ".article--ssr", args.Selector)
}

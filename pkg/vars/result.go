package vars

import (
	"encoding/json"
	"os"
)

func WriteScriptResult(jsonData any, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	return enc.Encode(jsonData)
}

func WriteScriptTextResult(textData string, path string) error {

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(textData)
	return err
}

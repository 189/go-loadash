package file

import (
	"encoding/json"
	"fmt"
	"os"
)

// ReadJSONFile reads a JSON file from the specified filepath and decodes its content into a value of type E.
// The function returns a pointer to the decoded value and an error if any occurred during the process.
//
// Type Parameters:
//
//	E - The type into which the JSON content will be decoded.
//
// Parameters:
//
//	filepath - The path to the JSON file to be read.
//
// Returns:
//
//	*E - A pointer to the decoded value of type E.
//	error - An error if any occurred during the file reading or decoding process.
func ReadJSONFile[E any](filepath string) (*E, error) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist: %w", err)
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("fail to open file: %w", err)
	}
	defer file.Close()

	var result E
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("fail to decode file content %w", err)
	}
	return &result, nil
}

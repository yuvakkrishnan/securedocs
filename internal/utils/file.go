package utils

import (
	"io/ioutil"
)

// ReadFile reads a file and returns its content as a string.
func ReadFile(filepath string) (string, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteFile writes a string to a file.
func WriteFile(filepath, content string) error {
	return ioutil.WriteFile(filepath, []byte(content), 0644)
}

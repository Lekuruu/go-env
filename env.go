package env

import (
	"fmt"
	"os"
)

// FromFile reads an env file and unmarshals it into an interface
func FromFile(filename string, config interface{}) error {
	envFile, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("Error reading env file: %v", err)
	}

	envMap, err := EnvToMap(envFile)
	if err != nil {
		return fmt.Errorf("Error parsing env file: %v", err)
	}

	return UnmarshalMap(envMap, config)
}

// ToFile writes an interface to an env file
func ToFile(filename string, config interface{}) error {
	envMap, err := MarshalMap(config)
	if err != nil {
		return fmt.Errorf("Error marshalling config to env: %v", err)
	}

	envFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("Error opening env file: %v", err)
	}

	return MapToEnv(envMap, envFile)
}

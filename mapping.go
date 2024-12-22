package env

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// EnvToMap reads an environment file and returns a map of key-value pairs
func EnvToMap(file io.Reader) (map[string]string, error) {
	envMap := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Ignore comments and empty lines
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		// Split key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.Trim(strings.TrimSpace(parts[1]), `"'`) // Remove quotes

		envMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return envMap, nil
}

// MapToEnv writes a map of key-value pairs to an environment file
func MapToEnv(envMap map[string]string, file io.Writer) error {
	writer := bufio.NewWriter(file)

	for key, value := range envMap {
		_, err := writer.WriteString(fmt.Sprintf("%s=%s\n", key, value))
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

package env

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

// UnmarshalMap maps environment variables to struct fields based on `env` tags
func UnmarshalMap(envMap map[string]string, config interface{}) error {
	val := reflect.ValueOf(config)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("config must be a pointer to a struct")
	}

	val = val.Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// Get the `env` tag
		tag := fieldType.Tag.Get("env")
		if tag == "" {
			continue
		}

		// Get the value from the environment map
		envValue, exists := envMap[tag]
		if !exists {
			continue
		}

		// Set the field value
		if !field.CanSet() {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(envValue)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intValue, err := strconv.Atoi(envValue)
			if err != nil {
				return fmt.Errorf("invalid value for %s: %v", tag, err)
			}
			field.SetInt(int64(intValue))
		case reflect.Bool:
			boolValue, err := strconv.ParseBool(envValue)
			if err != nil {
				return fmt.Errorf("invalid value for %s: %v", tag, err)
			}
			field.SetBool(boolValue)
		default:
			return fmt.Errorf("unsupported field type: %s", field.Kind())
		}
	}

	return nil
}

func MarshalMap(config interface{}) (map[string]string, error) {
	val := reflect.ValueOf(config)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("config must be a pointer to a struct")
	}

	val = val.Elem()
	typ := val.Type()

	envMap := make(map[string]string)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// Get the `env` tag
		tag := fieldType.Tag.Get("env")
		if tag == "" {
			continue
		}

		// Get the field value
		var fieldValue string
		switch field.Kind() {
		case reflect.String:
			fieldValue = field.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fieldValue = strconv.FormatInt(field.Int(), 10)
		case reflect.Bool:
			fieldValue = strconv.FormatBool(field.Bool())
		default:
			return nil, fmt.Errorf("unsupported field type: %s", field.Kind())
		}

		envMap[tag] = fieldValue
	}

	// Sort the envMap by key
	sortedEnvMap := make(map[string]string)
	keys := make([]string, 0, len(envMap))
	for k := range envMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		sortedEnvMap[k] = envMap[k]
	}

	return sortedEnvMap, nil
}

package env

import (
	"testing"
)

type TestConfig struct {
	AppName     string `env:"APP_NAME"`
	Port        int    `env:"PORT"`
	Debug       bool   `env:"DEBUG"`
	DatabaseUrl string `env:"DATABASE_URL"`
}

func TestEnv(t *testing.T) {
	inputEnvFile := ".example.env"
	outputEnvFile := ".env"

	config := TestConfig{}
	err := FromFile(inputEnvFile, &config)
	if err != nil {
		t.Errorf("Error reading env file: %v", err)
	}

	if config.AppName != "MyApp" {
		t.Errorf("AppName is not correct")
	}

	if config.Port != 8080 {
		t.Errorf("Port is not correct")
	}

	if config.Debug != true {
		t.Errorf("Debug is not correct")
	}

	if config.DatabaseUrl != "postgres://user:pass@localhost:5432/mydb" {
		t.Errorf("DatabaseUrl is not correct")
	}

	err = ToFile(outputEnvFile, &config)
	if err != nil {
		t.Errorf("Error writing env file: %v", err)
	}
}

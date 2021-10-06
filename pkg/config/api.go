package config

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

const defaultPort = 3000

// Generate retrieves necessary information from the environment to configure the service
func Generate() (Config, error) {
	port, err := getInt("PORT", defaultPort)
	if err != nil {
		return Config{}, fmt.Errorf("parsing port: %w", err)
	}

	return Config{
		BaseURL: getString("BASE_URL", ""),
		Port:    port,
	}, nil
}

// Validate ensures necessary configuration is valid and available
func (receiver Config) Validate() error {
	return validation.ValidateStruct(&receiver,
		validation.Field(&receiver.BaseURL, validation.Required),
		validation.Field(&receiver.Port, validation.Min(minimumPort), validation.Max(maximumPort)),
	)
}

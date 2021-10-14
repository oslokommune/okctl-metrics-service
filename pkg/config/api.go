package config

import (
	"fmt"
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	defaultPort         = 3000
	defaultUserAgent    = "okctl"
	defaultUserAgentDev = "okctldev"
)

// Generate retrieves necessary information from the environment to configure the service
func Generate() (cfg Config, err error) {
	cfg.BaseURL = getString(os.Getenv, "BASE_URL", "")

	cfg.Port, err = getInt(os.Getenv, "PORT", defaultPort)
	if err != nil {
		return Config{}, fmt.Errorf("parsing port: %w", err)
	}

	cfg.LegalAgents = getStringList(os.Getenv, "LEGAL_USER_AGENTS", []string{
		defaultUserAgent,
		defaultUserAgentDev,
	})

	return cfg, nil
}

// Validate ensures necessary configuration is valid and available
func (receiver Config) Validate() error {
	return validation.ValidateStruct(&receiver,
		validation.Field(&receiver.BaseURL, validation.Required),
		validation.Field(&receiver.Port, validation.Min(minimumPort), validation.Max(maximumPort)),
		validation.Field(&receiver.LegalAgents, validation.Length(1, 0)),
	)
}

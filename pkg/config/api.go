package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	defaultPort         = 3000
	defaultUserAgent    = "okctl"
	defaultUserAgentDev = "okctldev"
)

// Generate retrieves necessary information from the environment to configure the service
func Generate() (cfg Config, err error) {
	// Valid levels can be found here: https://pkg.go.dev/github.com/sirupsen/logrus@v1.6.0#Level
	cfg.LogLevel, err = getLogLevel(os.Getenv, "LOG_LEVEL", logrus.InfoLevel)
	if err != nil {
		return Config{}, fmt.Errorf("acquiring log level: %w", err)
	}

	cfg.Port, err = getInt(os.Getenv, "PORT", defaultPort)
	if err != nil {
		return Config{}, fmt.Errorf("parsing port: %w", err)
	}

	cfg.BaseURL = getString(os.Getenv, "BASE_URL", "")

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

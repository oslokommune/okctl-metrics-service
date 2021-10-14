package config

import "github.com/sirupsen/logrus"

const (
	minimumPort = 0
	maximumPort = 65535
)

// Config contains necessary configuration for the service
type Config struct {
	BaseURL     string
	LogLevel    logrus.Level
	LegalAgents []string
	Port        int
}

type stringValueGetter func(key string) string

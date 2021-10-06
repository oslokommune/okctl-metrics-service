package config

const (
	minimumPort = 0
	maximumPort = 65535
)

// Config contains necessary configuration for the service
type Config struct {
	BaseURL string
	Port    int
}

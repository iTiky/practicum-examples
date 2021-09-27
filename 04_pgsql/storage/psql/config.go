package psql

import "fmt"

const (
	defaultConfigEndpoint = "http://alpha.local:5432"
)

// Config keeps Storage configuration.
type Config struct {
	DSN string `mapstructure:"dsn"`
}

// Validate performs a basic validation.
func (c Config) Validate() error {
	if c.DSN == "" {
		return fmt.Errorf("%s field: empty", "DSN")
	}

	return nil
}

// NewDefaultConfig builds a Config with default values.
func NewDefaultConfig() Config {
	return Config{
		DSN: defaultConfigEndpoint,
	}
}

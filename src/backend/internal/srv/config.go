package srv

import "errors"

// Config holds configuration data for the app.
type Config struct {
	// Name of the app.
	Name string `json:"name"`
	// Port of the application server
	Port string `json:"port"`
}

// Validate checks the whether the Config is valid. Returns a non-nil error if
// Config is invalid.
func (c Config) Validate() error {
	if c.Name == "" {
		return errors.New("app name cannot be empty")
	}
	return nil
}

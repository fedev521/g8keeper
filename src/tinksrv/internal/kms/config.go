package kms

import "errors"

// Config holds the configuration data for the KMS portion of the app.
type Config struct {
	KekFile string
}

// Validate checks the whether the Config is valid. Returns a non-nil error if
// Config is invalid.
func (c Config) Validate() error {
	if c.KekFile == "" {
		return errors.New("the path to the file containing the KEK cannot be empty")
	}
	return nil
}

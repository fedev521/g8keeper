package svc

import "errors"

// TinkSvcConfig holds configuration for the tinksrv service.
type TinkSvcConfig struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	KekUri string `json:"kekUri"`
}

func (c TinkSvcConfig) Validate() error {
	if c.Host == "" {
		return errors.New("tinksvc host cannot be empty")
	}
	if c.Port == "" {
		return errors.New("tinksvc port cannot be empty")
	}
	if c.KekUri == "" {
		return errors.New("tinksvc KEK URI cannot be empty")
	}
	return nil
}

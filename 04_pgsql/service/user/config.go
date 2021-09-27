package user

import (
	"fmt"

	"github.com/ttacon/libphonenumber"
)

// Config keeps Processor params.
type Config struct {
	AllowedRegionCodes []string `mapstructure:"allowed_region_codes"`
}

// Validate performs a basic validation.
func (c Config) Validate() error {
	allRegions := libphonenumber.GetSupportedRegions()
	for _, regionCode := range c.AllowedRegionCodes {
		if _, found := allRegions[regionCode]; !found {
			return fmt.Errorf("%s field: %s: invalid", "AllowedRegionCodes", regionCode)
		}
	}

	return nil
}

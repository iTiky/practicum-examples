package validator

import (
	"fmt"

	"github.com/ttacon/libphonenumber"
)

// ValidatePhoneNumberWithCountryCode validates region code and phone number (if provided).
func ValidatePhoneNumberWithCountryCode(phoneNumber, regionCode string) (*libphonenumber.PhoneNumber, error) {
	if regionCode == "" {
		return nil, fmt.Errorf("regionCode: empty")
	}

	if code := libphonenumber.GetCountryCodeForRegion(regionCode); code == 0 {
		return nil, fmt.Errorf("regionCode: invalid")
	}

	if phoneNumber == "" {
		return nil, nil
	}

	number, err := libphonenumber.Parse(phoneNumber, regionCode)
	if err != nil {
		return nil, fmt.Errorf("phoneNumber: %w", err)
	}

	if !libphonenumber.IsValidNumberForRegion(number, regionCode) {
		return nil, fmt.Errorf("phoneNumber: invalid for region")
	}

	return number, nil
}

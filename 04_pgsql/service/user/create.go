package user

import (
	"context"
	"fmt"

	"github.com/itiky/practicum-examples/04_pgsql/model"
	"github.com/itiky/practicum-examples/04_pgsql/pkg"
	"github.com/itiky/practicum-examples/04_pgsql/service/user/validator"
	"github.com/ttacon/libphonenumber"
)

func (svc *Processor) CreateUser(ctx context.Context, name, regionCode, email, phoneNumberRaw string) (model.User, error) {
	logger := svc.Logger()

	// Input checks
	if name == "" {
		return model.User{}, fmt.Errorf("%w: name: empty", pkg.ErrInvalidInput)
	}

	if err := validator.ValidateEmail(email); err != nil {
		return model.User{}, fmt.Errorf("%w: email: %v", pkg.ErrInvalidInput, err)
	}

	phoneNumber, err := validator.ValidatePhoneNumberWithCountryCode(phoneNumberRaw, regionCode)
	if err != nil {
		return model.User{}, fmt.Errorf("%w: phoneNumber / regionCode: %v", pkg.ErrInvalidInput, err)
	}

	if _, found := svc.allowedRegionCodes[regionCode]; !found {
		return model.User{}, fmt.Errorf("%w: regionCode is not allowed", pkg.ErrInvalidInput)
	}

	// Build input
	input := model.User{
		Name:       name,
		Email:      email,
		RegionCode: regionCode,
	}
	if phoneNumber != nil {
		input.PhoneNumber = libphonenumber.Format(phoneNumber, libphonenumber.INTERNATIONAL)
	}

	logger.UpdateContext(input.GetLoggerContext)
	logger.Info().Msg("Creating user")

	// Create
	user, err := svc.userStorage.CreateUser(ctx, input)
	if err != nil {
		logger.Warn().Err(err).Msg("Creating user")
		return model.User{}, fmt.Errorf("creating user: %w", err)
	}

	logger.UpdateContext(user.GetLoggerContext)
	logger.Info().Msg("User created")

	return user, nil
}

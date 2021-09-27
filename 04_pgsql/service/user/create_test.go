package user

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/itiky/practicum-examples/04_pgsql/model"
	"github.com/itiky/practicum-examples/04_pgsql/pkg"
	storagemock "github.com/itiky/practicum-examples/04_pgsql/storage/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserService_CreateUser(t *testing.T) {
	type (
		testInput struct {
			name        string
			regionCode  string
			email       string
			phoneNumber string
		}

		testCase struct {
			name         string
			prepareMocks func(userWriterMock *storagemock.MockUserWriter) testInput
			errExpected  bool
			errTarget    error
			errContains  string
		}
	)

	allowedRegionCodes := []string{"RU"}

	testCases := []testCase{
		{
			name: "Fail: invalid input (name empty)",
			prepareMocks: func(userWriterMock *storagemock.MockUserWriter) testInput {
				return testInput{
					name:       "",
					regionCode: "RU",
					email:      "bob@gmail.com",
				}
			},
			errExpected: true,
			errTarget:   pkg.ErrInvalidInput,
			errContains: "name",
		},
		{
			name: "Fail: invalid input (invalid region)",
			prepareMocks: func(userWriterMock *storagemock.MockUserWriter) testInput {
				return testInput{
					name:       "Mock",
					regionCode: "RUUUU",
					email:      "bob@gmail.com",
				}
			},
			errExpected: true,
			errTarget:   pkg.ErrInvalidInput,
			errContains: "regionCode",
		},
		{
			name: "Fail: invalid input (invalid email)",
			prepareMocks: func(userWriterMock *storagemock.MockUserWriter) testInput {
				return testInput{
					name:       "Mock",
					regionCode: "RU",
					email:      "@gmail.com",
				}
			},
			errExpected: true,
			errTarget:   pkg.ErrInvalidInput,
			errContains: "email",
		},
		{
			name: "Fail: invalid input (invalid phoneNumber)",
			prepareMocks: func(userWriterMock *storagemock.MockUserWriter) testInput {
				return testInput{
					name:        "Mock",
					regionCode:  "RU",
					email:       "mock@gmail.com",
					phoneNumber: "+3 123",
				}
			},
			errExpected: true,
			errTarget:   pkg.ErrInvalidInput,
			errContains: "phoneNumber",
		},
		{
			name: "Fail: disallowed region code",
			prepareMocks: func(userWriterMock *storagemock.MockUserWriter) testInput {
				return testInput{
					name:        "Mock",
					regionCode:  "US",
					email:       "mock@gmail.com",
					phoneNumber: "+1 201-555-0123",
				}
			},
			errExpected: true,
			errTarget:   pkg.ErrInvalidInput,
			errContains: "not allowed",
		},
		{
			name: "OK",
			prepareMocks: func(userWriterMock *storagemock.MockUserWriter) testInput {
				input := testInput{
					name:        "Mock",
					regionCode:  "RU",
					email:       "mock@gmail.com",
					phoneNumber: "+7-301-123-45-67",
				}

				user := model.User{
					Name:        input.name,
					Email:       input.email,
					PhoneNumber: "+7 301 123-45-67",
					RegionCode:  input.regionCode,
				}

				userWriterMock.EXPECT().
					CreateUser(gomock.Any(), user).
					Return(user, nil)

				return input
			},
			errExpected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			userWriterMock := storagemock.NewMockUserWriter(mockCtrl)

			input := tc.prepareMocks(userWriterMock)

			cfg := Config{
				AllowedRegionCodes: allowedRegionCodes,
			}

			svc, err := New(
				WithUserWriter(userWriterMock),
				WithConfig(cfg),
			)
			require.NoError(t, err)

			_, err = svc.CreateUser(context.TODO(), input.name, input.regionCode, input.email, input.phoneNumber)
			if tc.errExpected {
				assert.Error(t, err)
				if tc.errTarget != nil {
					assert.True(t, errors.Is(err, tc.errTarget))
				}
				if tc.errContains != "" {
					assert.Contains(t, err.Error(), tc.errContains)
				}
				return
			}

			assert.NoError(t, err)
		})
	}
}

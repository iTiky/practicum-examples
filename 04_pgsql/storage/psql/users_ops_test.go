package psql

import (
	"errors"

	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/itiky/practicum-examples/04_pgsql/model"
	"github.com/itiky/practicum-examples/04_pgsql/pkg"
)

func (s *TestSuite) TestUsers_CreateUser() {
	user := model.User{
		Name:       fake.CharactersN(8),
		Email:      fake.CharactersN(8),
		RegionCode: "GB",
	}

	s.Run("Create non-existing user", func() {
		res, err := s.storage.CreateUser(s.ctx, user)
		s.Require().NoError(err)

		s.Assert().EqualValues(user.Name, res.Name)
		s.Assert().EqualValues(user.Email, res.Email)
		s.Assert().NotEqual(uuid.Nil, res.ID)
		s.Assert().NotEmpty(res.CreatedAt)
		s.Assert().NotEmpty(res.UpdatedAt)
		s.Assert().Empty(res.DeletedAt)
	})

	s.Run("Try to create existing user", func() {
		_, err := s.storage.CreateUser(s.ctx, user)
		s.Require().Error(err)
		s.Require().True(errors.Is(err, pkg.ErrAlreadyExists))
	})
}

func (s *TestSuite) TestUsers_GetUserByEmail() {
	s.Run("Get non-existing user", func() {
		res, err := s.storage.GetUserByEmail(s.ctx, fake.CharactersN(8))
		s.Require().NoError(err)
		s.Require().Nil(res)
	})

	s.Run("Get Bob", func() {
		expectedUser, err := s.fixtures.Users[0].ToCanonical()
		s.Require().NoError(err)

		res, err := s.storage.GetUserByEmail(s.ctx, expectedUser.Email)
		s.Require().NoError(err)
		s.Assert().NotNil(res)
		s.Assert().EqualValues(expectedUser, *res)
	})

	s.Run("Get Alice", func() {
		expectedUser, err := s.fixtures.Users[1].ToCanonical()
		s.Require().NoError(err)

		res, err := s.storage.GetUserByEmail(s.ctx, expectedUser.Email)
		s.Require().NoError(err)
		s.Assert().NotNil(res)
		s.Assert().EqualValues(expectedUser, *res)
	})
}

package psql

import (
	"errors"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/itiky/practicum-examples/04_pgsql/model"
	"github.com/itiky/practicum-examples/04_pgsql/pkg"
	"github.com/itiky/practicum-examples/04_pgsql/pkg/input"
	"github.com/itiky/practicum-examples/04_pgsql/storage/psql/schema"
)

func (s *TestSuite) TestOrders_CreateOrder() {
	s.Run("Create order for Alice", func() {
		order := model.Order{
			UserID: s.fixtures.Users[1].ID,
			Status: model.OrderStatusDelivered,
		}

		res, err := s.storage.CreateOrder(s.ctx, order)
		s.Require().NoError(err)

		s.Assert().EqualValues(order.UserID, res.UserID)
		s.Assert().EqualValues(order.Status, res.Status)
		s.Assert().NotEqual(uuid.Nil, res.ID)
		s.Assert().NotEmpty(res.CreatedAt)
		s.Assert().Empty(res.DeletedAt)
	})

	s.Run("Create order for non-existing user", func() {
		order := model.Order{
			UserID: uuid.New(),
			Status: model.OrderStatusDelivered,
		}

		_, err := s.storage.CreateOrder(s.ctx, order)
		s.Require().Error(err)
		s.Assert().True(errors.Is(err, pkg.ErrNotExists))
	})
}

func (s *TestSuite) TestOrders_GetOrderByID() {
	s.Run("Get non-existing order", func() {
		res, err := s.storage.GetOrderByID(s.ctx, uuid.New())
		s.Require().NoError(err)
		s.Require().Nil(res)
	})

	s.Run("Get Bob's 1st order", func() {
		expectedOrder, err := s.fixtures.Orders[0].ToCanonical()
		s.Require().NoError(err)

		res, err := s.storage.GetOrderByID(s.ctx, expectedOrder.ID)
		s.Require().NoError(err)
		s.Assert().NotNil(res)
		s.Assert().EqualValues(expectedOrder, *res)
	})
}

func (s *TestSuite) TestOrders_GetOrdersForUser() {
	type testCase struct {
		name           string
		userID         uuid.UUID
		rangeStart     time.Time
		rangeEnd       time.Time
		pagination     input.PageParams
		expectedDbObjs schema.Orders
	}

	bobID := s.fixtures.Users[0].ID
	bobOrders := s.fixtures.Orders[0:3]

	testCases := []testCase{
		{
			name:       "Get orders for non-existing user",
			userID:     uuid.New(),
			pagination: input.NewDefaultPageParams(),
		},
		{
			name:           "Get Bob's orders: all",
			userID:         bobID,
			pagination:     input.NewDefaultPageParams(),
			expectedDbObjs: bobOrders,
		},
		{
			name:   "Get Bob's orders: 1st page",
			userID: bobID,
			pagination: input.PageParams{
				Offset: 0,
				Limit:  1,
			},
			expectedDbObjs: schema.Orders{bobOrders[0]},
		},
		{
			name:   "Get Bob's orders: 2nd page",
			userID: bobID,
			pagination: input.PageParams{
				Offset: 1,
				Limit:  2,
			},
			expectedDbObjs: schema.Orders{bobOrders[1], bobOrders[2]},
		},
		{
			name:           "Get Bob's orders: with lower createdAt bound",
			userID:         bobID,
			pagination:     input.NewDefaultPageParams(),
			rangeStart:     bobOrders[1].CreatedAt,
			expectedDbObjs: schema.Orders{bobOrders[1], bobOrders[2]},
		},
		{
			name:           "Get Bob's orders: with upper createdAt bound",
			userID:         bobID,
			pagination:     input.NewDefaultPageParams(),
			rangeEnd:       bobOrders[1].CreatedAt,
			expectedDbObjs: schema.Orders{bobOrders[0], bobOrders[1]},
		},
		{
			name:           "Get Bob's orders: with createdAt double bounded",
			userID:         bobID,
			pagination:     input.NewDefaultPageParams(),
			rangeStart:     bobOrders[1].CreatedAt.Add(-1 * time.Nanosecond),
			rangeEnd:       bobOrders[1].CreatedAt.Add(1 * time.Nanosecond),
			expectedDbObjs: schema.Orders{bobOrders[1]},
		},
		{
			name:       "Get Bob's orders: with createdAt out of range",
			userID:     bobID,
			pagination: input.NewDefaultPageParams(),
			rangeStart: bobOrders[2].CreatedAt.Add(1 * time.Minute),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			var rangeStart, rangeEnd *time.Time
			if !tc.rangeStart.IsZero() {
				rangeStart = &tc.rangeStart
			}
			if !tc.rangeEnd.IsZero() {
				rangeEnd = &tc.rangeEnd
			}

			resExpected, err := tc.expectedDbObjs.ToCanonical()
			s.Require().NoError(err)

			resReceived, err := s.storage.GetOrdersForUser(s.ctx, tc.userID, rangeStart, rangeEnd, tc.pagination)
			s.Require().NoError(err)
			s.Assert().ElementsMatch(resExpected, resReceived)

			resReceivedIsSorted := sort.SliceIsSorted(resReceived, func(i, j int) bool {
				return resReceived[i].CreatedAt.Before(resReceived[j].CreatedAt)
			})
			s.Assert().True(resReceivedIsSorted)
		})
	}
}

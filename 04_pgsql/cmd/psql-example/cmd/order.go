package cmd

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/itiky/practicum-examples/04_pgsql/cmd/psql-example/cmd/common"
	"github.com/itiky/practicum-examples/04_pgsql/model"
	"github.com/itiky/practicum-examples/04_pgsql/pkg"
	"github.com/itiky/practicum-examples/04_pgsql/pkg/input"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	flagTimeRangeStart = "range-start"
	flagTimeRangeEnd   = "range-end"
)

// newUserCmd creates a new user cmd.
func newOrderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "order",
		Short: "Order sub-commands",
	}

	cmd.AddCommand(newOrderCreateCmd())
	cmd.AddCommand(newOrderListCmd())

	return cmd
}

// newOrderCreateCmd creates a new order.create cmd.
func newOrderCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create <user ID> <status>",
		Short:   "Creates a new order",
		Example: `create 123e4567-e89b-12d3-a456-426614174000 created`,
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Init
			config := common.GetConfigFromCmdCtx(cmd)
			timeout := common.GetTimeoutToCmdCtx(cmd)

			svc, err := config.BuildOrderService()
			if err != nil {
				return err
			}
			defer func() {
				if err := svc.Close(); err != nil {
					log.Error().Err(err).Msg("Shutting down the app")
				}
			}()

			// Prepare input
			userID, err := uuid.Parse(args[0])
			if err != nil {
				return fmt.Errorf("parsing %q argument: %w", "user ID", err)
			}

			status := model.OrderStatus(strings.ToTitle(args[1]))

			// Exec
			ctx, ctxCancel := context.WithTimeout(context.Background(), timeout)
			defer ctxCancel()

			user, err := svc.CreateOrder(ctx, userID, status)
			if err != nil {
				if errors.Is(err, pkg.ErrInvalidInput) || errors.Is(err, pkg.ErrNotExists) {
					log.Error().Err(err).Msg("Invalid input")
					return nil
				}

				return err
			}
			common.PrintAsYAML("Order data:", user)

			return nil
		},
	}

	return cmd
}

// newOrderListCmd creates a new order.list cmd.
func newOrderListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list <user ID>",
		Short:   "List orders for a specific user",
		Example: `list 123e4567-e89b-12d3-a456-426614174000`,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Init
			config := common.GetConfigFromCmdCtx(cmd)
			timeout := common.GetTimeoutToCmdCtx(cmd)

			svc, err := config.BuildOrderService()
			if err != nil {
				return err
			}
			defer func() {
				if err := svc.Close(); err != nil {
					log.Error().Err(err).Msg("Shutting down the app")
				}
			}()

			// Prepare input
			userID, err := uuid.Parse(args[0])
			if err != nil {
				return fmt.Errorf("parsing %q argument: %w", "user ID", err)
			}

			rangeStart, err := getTimeParam(flagTimeRangeStart, cmd)
			if err != nil {
				return err
			}
			rangeEnd, err := getTimeParam(flagTimeRangeEnd, cmd)
			if err != nil {
				return err
			}

			pageParams, err := input.ParsePageParams(cmd)
			if err != nil {
				return err
			}

			// Exec
			ctx, ctxCancel := context.WithTimeout(context.Background(), timeout)
			defer ctxCancel()

			user, err := svc.GetOrdersForUser(ctx, userID, rangeStart, rangeEnd, pageParams)
			if err != nil {
				if errors.Is(err, pkg.ErrInvalidInput) {
					log.Error().Err(err).Msg("Invalid input")
					return nil
				}

				return err
			}
			common.PrintAsYAML("Orders data:", user)

			return nil
		},
	}

	cmd.Flags().String(flagTimeRangeStart, "", "Time range filter start value (optional, RFC3339 format)")
	cmd.Flags().String(flagTimeRangeEnd, "", "Time range filter end value (optional, RFC3339 format)")
	input.AddPageParamsToCmd(cmd)

	return cmd
}

// getTimeParam reads and parses time parameter.
func getTimeParam(pName string, cmd *cobra.Command) (*time.Time, error) {
	vStr, err := cmd.Flags().GetString(pName)
	if err != nil {
		return nil, fmt.Errorf("reading %q parameter: %w", pName, err)
	}

	if vStr == "" {
		return nil, nil
	}

	vTime, err := time.Parse(time.RFC3339, vStr)
	if err != nil {
		return nil, fmt.Errorf("reading %q parameter: parsing time (RFC3339): %w", pName, err)
	}

	return &vTime, nil
}

package cmd

import (
	"context"
	"errors"

	"github.com/itiky/practicum-examples/04_pgsql/cmd/psql-example/cmd/common"
	"github.com/itiky/practicum-examples/04_pgsql/pkg"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// newUserCmd creates a new user cmd.
func newUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "User sub-commands",
	}

	cmd.AddCommand(newUserCreateCmd())

	return cmd
}

// newUserCreateCmd creates a new user.create cmd.
func newUserCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create <name> <email> <region code> [phone number]",
		Short:   "Creates a new user",
		Example: `create "Bob Mock" bob@google.com RU`,
		Args:    cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Init
			config := common.GetConfigFromCmdCtx(cmd)
			timeout := common.GetTimeoutToCmdCtx(cmd)

			svc, err := config.BuildUserService()
			if err != nil {
				return err
			}
			defer func() {
				if err := svc.Close(); err != nil {
					log.Error().Err(err).Msg("Shutting down the app")
				}
			}()

			// Prepare input
			name, email, regionCode, phoneNumber := args[0], args[1], args[2], ""
			if len(args) > 3 {
				phoneNumber = args[3]
			}

			// Exec
			ctx, ctxCancel := context.WithTimeout(context.Background(), timeout)
			defer ctxCancel()

			user, err := svc.CreateUser(ctx, name, regionCode, email, phoneNumber)
			if err != nil {
				if errors.Is(err, pkg.ErrInvalidInput) {
					log.Error().Err(err).Msg("Invalid input")
					return nil
				}

				if errors.Is(err, pkg.ErrAlreadyExists) {
					log.Error().Msg("Specified user exists")
					return nil
				}

				return err
			}
			common.PrintAsYAML("User data:", user)

			return nil
		},
	}

	return cmd
}

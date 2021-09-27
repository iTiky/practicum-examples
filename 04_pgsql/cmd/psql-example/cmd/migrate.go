package cmd

import (
	"context"

	"github.com/itiky/practicum-examples/04_pgsql/cmd/psql-example/cmd/common"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// newMigrateCmd creates a new migrate cmd.
func newMigrateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate DB to the latest version",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Init
			config := common.GetConfigFromCmdCtx(cmd)
			timeout := common.GetTimeoutToCmdCtx(cmd)

			st, err := config.BuildPsqlStorage()
			if err != nil {
				return err
			}
			defer func() {
				if err := st.Close(); err != nil {
					log.Error().Err(err).Msg("Shutting down the app")
				}
			}()

			// Exec
			ctx, ctxCancel := context.WithTimeout(context.Background(), timeout)
			defer ctxCancel()

			if err := st.Migrate(ctx); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}

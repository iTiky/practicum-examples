package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/itiky/practicum-examples/04_pgsql/cmd/psql-example/cmd/common"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagLogLevel         = "log-level"
	flagConfigPath       = "config"
	flagOperationTimeout = "timeout"
)

// Execute prepares cobra.Command context and executes root cmd.
func Execute() error {
	return newRootCmd().ExecuteContext(common.NewBaseCmdCtx())
}

// newRootCmd creates a new root cmd.
func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "psql-example",
		Short: "Basic PostgreSQL example app",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := setupLogger(cmd); err != nil {
				return fmt.Errorf("app initialization: %w", err)
			}

			if err := setupConfig(cmd); err != nil {
				return fmt.Errorf("app initialization: %w", err)
			}

			opTimeout, err := cmd.Flags().GetDuration(flagOperationTimeout)
			if err != nil {
				return fmt.Errorf("app initialization: reading flag %s: %w", flagOperationTimeout, err)
			}
			common.SetTimeoutToCmdCtx(cmd, opTimeout)

			return nil
		},
	}

	cmd.PersistentFlags().String(flagLogLevel, "info", "Logger level [debug,info,warn,error,fatal]")
	cmd.PersistentFlags().String(flagConfigPath, "./config.toml", "Config file path")
	cmd.PersistentFlags().Duration(flagOperationTimeout, 30*time.Second, "Operation timeout")

	cmd.AddCommand(newMigrateCmd())
	cmd.AddCommand(newUserCmd())
	cmd.AddCommand(newOrderCmd())

	return cmd
}

// setupLogger configures global logger.
func setupLogger(cmd *cobra.Command) error {
	// Parse flag
	logLevelBz, err := cmd.Flags().GetString(flagLogLevel)
	if err != nil {
		return fmt.Errorf("%s flag reading: %w", flagLogLevel, err)
	}
	logLevel, err := zerolog.ParseLevel(logLevelBz)
	if err != nil {
		return fmt.Errorf("%s flag parsing: %w", flagLogLevel, err)
	}

	// Setup
	logWriter := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	}
	log.Logger = log.Output(logWriter).Level(logLevel)

	return nil
}

// setupConfig reads app config and stores it to cobra.Command context.
func setupConfig(cmd *cobra.Command) error {
	// Parse flag
	configPath, err := cmd.Flags().GetString(flagConfigPath)
	if err != nil {
		return fmt.Errorf("%s flag reading: %w", flagConfigPath, err)
	}

	// Viper common setup
	viper.AutomaticEnv()
	viper.SetEnvPrefix("PSQLEXAMPLE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigFile(configPath)
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("reading config: %w", err)
	}

	// Override config parameters with ENVs
	for _, key := range viper.AllKeys() {
		val := viper.Get(key)
		viper.Set(key, val)
	}

	// Get config and persist it to cobra.Command context
	config := common.Config{}
	if err := viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("config unmarshal: %w", err)
	}
	common.SetConfigToCmdCtx(cmd, config)

	return nil
}

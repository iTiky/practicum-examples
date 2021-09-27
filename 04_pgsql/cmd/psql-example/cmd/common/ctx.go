package common

import (
	"context"
	"fmt"
	"time"

	"github.com/itiky/practicum-examples/04_pgsql/pkg"
	"github.com/spf13/cobra"
)

const (
	ctxKeyConfig  = pkg.ContextKey("cli.config")
	ctxKeyTimeout = pkg.ContextKey("cli.timeout")
)

// SetConfigToCmdCtx adds Config to cobra.Command context.
func SetConfigToCmdCtx(cmd *cobra.Command, config Config) {
	v := cmd.Context().Value(ctxKeyConfig)
	if v == nil {
		panic(fmt.Errorf("%s context: not set", ctxKeyConfig))
	}

	ctxPtr := v.(*Config)
	*ctxPtr = config
}

// GetConfigFromCmdCtx gets stored Config from cobra.Command context.
func GetConfigFromCmdCtx(cmd *cobra.Command) Config {
	v := cmd.Context().Value(ctxKeyConfig)
	if v == nil {
		panic(fmt.Errorf("%s context: not set", ctxKeyConfig))
	}
	config := v.(*Config)

	return *config
}

// SetTimeoutToCmdCtx adds timeout duration to cobra.Command context.
func SetTimeoutToCmdCtx(cmd *cobra.Command, dur time.Duration) {
	v := cmd.Context().Value(ctxKeyTimeout)
	if v == nil {
		panic(fmt.Errorf("%s context: not set", ctxKeyTimeout))
	}

	ctxPtr := v.(*time.Duration)
	*ctxPtr = dur
}

// GetTimeoutToCmdCtx gets stored timeout duration from cobra.Command context.
func GetTimeoutToCmdCtx(cmd *cobra.Command) time.Duration {
	v := cmd.Context().Value(ctxKeyTimeout)
	if v == nil {
		panic(fmt.Errorf("%s context: not set", ctxKeyTimeout))
	}
	dur := v.(*time.Duration)

	return *dur
}

// NewBaseCmdCtx creates an empty base context used for storing values for cobra.Command.
func NewBaseCmdCtx() context.Context {
	var dur time.Duration

	ctx := context.Background()
	ctx = context.WithValue(ctx, ctxKeyConfig, &Config{})
	ctx = context.WithValue(ctx, ctxKeyTimeout, &dur)

	return ctx
}

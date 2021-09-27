package psql

import (
	"github.com/itiky/practicum-examples/04_pgsql/pkg/logging"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// loggerOption customises logger context fields.
type loggerOption func(logCtx zerolog.Context) zerolog.Context

// withTable sets table name logging context.
// nolint: unused
func withTable(tableName string) loggerOption {
	return func(logCtx zerolog.Context) zerolog.Context {
		return logCtx.Str(dbTableLoggingKey, tableName)
	}
}

// withOperation sets DB operation logging context.
func withOperation(opID string) loggerOption {
	return func(logCtx zerolog.Context) zerolog.Context {
		return logCtx.Str(dbOperationLoggingKey, opID)
	}
}

// Logger returns logger with service context.
func (st Storage) Logger(opts ...loggerOption) zerolog.Logger {
	logCtx := log.With().Str(logging.ServiceKey, serviceName)
	for _, opt := range opts {
		logCtx = opt(logCtx)
	}

	return logCtx.Logger()
}

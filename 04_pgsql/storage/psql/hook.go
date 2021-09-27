package psql

import (
	"context"
	"time"

	"github.com/itiky/practicum-examples/04_pgsql/pkg/logging"
	"github.com/uptrace/bun"
)

type queryHook struct {
	st *Storage
}

func newQueryHook(storage *Storage) queryHook {
	return queryHook{
		st: storage,
	}
}

func (h queryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (h queryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	logger := h.st.Logger()
	logger.Debug().
		Dur(logging.RequestDurKey, time.Since(event.StartTime)).
		Msg(event.Query)
}

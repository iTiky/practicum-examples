package psql

import (
	"github.com/itiky/practicum-examples/04_pgsql/pkg/input"
	"github.com/uptrace/bun"
)

// paginateQuery adds pagination parameters to select query.
func paginateQuery(q *bun.SelectQuery, params input.PageParams) {
	q.Offset(params.Offset)
	q.Limit(params.Limit)
}

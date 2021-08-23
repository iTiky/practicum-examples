package model

import (
	"github.com/itiky/practicum-examples/01_log_and_trace_context/pkg/logging"
	"github.com/rs/zerolog"
)

type (
	Order struct {
		ID    string    `json:"id"`
		Items []Product `json:"items"`
	}

	Product struct {
		Barcode  string        `json:"barcode"`
		Name     string        `json:"name,omitempty"`
		Status   ProductStatus `json:"status"`
		Quantity int           `json:"quantity"`
	}

	ProductStatus string
)

const (
	ProductStatusDispatched ProductStatus = "dispatched"
	ProductStatusReturned   ProductStatus = "returned"
)

// String implements fmt.Stringer interface.
func (s ProductStatus) String() string {
	return string(s)
}

// GetLoggerContext enriches logger context with essential Order fields.
func (o Order) GetLoggerContext(logCtx zerolog.Context) zerolog.Context {
	return logCtx.
		Str(logging.OrderIDKey, o.ID)
}

// GetLoggerContext enriches logger context with essential Order fields.
func (p Product) GetLoggerContext(logCtx zerolog.Context) zerolog.Context {
	return logCtx.
		Str(logging.BarcodeKey, p.Barcode)
}

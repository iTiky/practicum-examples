package main

import (
	"encoding/json"
	"fmt"

	"github.com/itiky/practicum-examples/01_log_and_trace_context/api/handler"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/model"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/pkg/tracing"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/provider/product_library"
	orderService "github.com/itiky/practicum-examples/01_log_and_trace_context/service/order"
	orderStorage "github.com/itiky/practicum-examples/01_log_and_trace_context/storage/order"
)

func buildDependencies() (*handler.OrderHandler, error) {
	orderProcessorSvc, err := orderService.NewProcessor(
		orderService.WithProductNameProvider(product_library.NewProvider()),
		orderService.WithProductStorage(orderStorage.NewStorage()),
	)
	if err != nil {
		return nil, fmt.Errorf("orderProcessorSvc init: %w", err)
	}

	return handler.NewOrderHandler(orderProcessorSvc), nil
}

func buildRecords(orders ...*model.Order) [][]byte {
	records := make([][]byte, len(orders))
	for recordIdx, order := range orders {
		if order == nil {
			continue
		}

		orderBz, _ := json.Marshal(order)
		records[recordIdx] = orderBz
	}

	return records
}

func main() {
	tracerCloser, err := tracing.SetupGlobalJaegerTracer()
	if err != nil {
		panic(err)
	}
	defer tracerCloser.Close()

	h, err := buildDependencies()
	if err != nil {
		panic(err)
	}

	h.HandleOrderRecords(
		buildRecords(
			&model.Order{
				ID: "01",
				Items: []model.Product{
					{
						Barcode:  "001",
						Status:   model.ProductStatusDispatched,
						Quantity: 1,
					},
					{
						Barcode:  "002",
						Status:   model.ProductStatusReturned,
						Quantity: 2,
					},
				},
			},
			nil,
		)...,
	)

	h.HandleOrderRecords(
		buildRecords(
			&model.Order{
				ID: "02",
				Items: []model.Product{
					{
						Barcode:  "",
						Status:   model.ProductStatusDispatched,
						Quantity: 1,
					},
				},
			},
			&model.Order{
				ID: "03",
				Items: []model.Product{
					{
						Barcode:  "003",
						Status:   model.ProductStatusDispatched,
						Quantity: -1,
					},
				},
			},
		)...,
	)
}

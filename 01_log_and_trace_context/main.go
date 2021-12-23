package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/itiky/practicum-examples/01_log_and_trace_context/api/stream"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/model"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/pkg/tracing"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/provider/prodlibrary/http"
	orderService "github.com/itiky/practicum-examples/01_log_and_trace_context/service/order/v1"
	orderStorage "github.com/itiky/practicum-examples/01_log_and_trace_context/storage/test"
)

func buildDependencies() (*stream.OrderHandler, error) {
	orderProcessorSvc, err := orderService.NewProcessor(
		orderService.WithProductNameProvider(http.NewProvider()),
		orderService.WithProductStorage(orderStorage.NewStorage()),
	)
	if err != nil {
		return nil, fmt.Errorf("orderProcessorSvc init: %w", err)
	}

	h, err := stream.NewOrderHandler(orderProcessorSvc)
	if err != nil {
		return nil, fmt.Errorf("records stream init: %w", err)
	}

	return h, nil
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

	fmt.Println()
	h.HandleRecords(
		context.Background(),
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

	fmt.Println()
	h.HandleRecords(
		context.Background(),
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

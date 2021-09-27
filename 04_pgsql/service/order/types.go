package order

import "github.com/itiky/practicum-examples/04_pgsql/storage"

type StorageExpected interface {
	storage.OrderWriter
	storage.OrderReader
}

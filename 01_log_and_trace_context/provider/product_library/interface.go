//go:generate mockgen -source=interface.go -destination=./mock/provider.go -package=prodnameprovidermock
package product_library

import "context"

type ProductNameProvider interface {
	// GetProductName returns product name by the barcode.
	GetProductName(ctx context.Context, barcode string) (string, error)
}

//go:generate mockgen -source=interface.go -destination=./mock/provider.go -package=providermock
package prodlibrary

import "context"

type ProductNameProvider interface {
	// GetProductName returns product name by the barcode.
	GetProductName(ctx context.Context, barcode string) (string, error)
}

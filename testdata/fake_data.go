package testdata

import (
	"github.com/shopspring/decimal"

	"github.com/aadiths/unifize-discount-engine/internal/models"
)

// PumaTshirtCart models the assignment scenario:
// PUMA T-shirt with brand, category, and ICICI bank discounts.
func PumaTshirtCart() []models.CartItem {
	product := models.Product{
		ID:        "p1",
		Brand:     models.BrandPuma,
		Category:  models.CategoryTShirts,
		BasePrice: decimal.NewFromInt(1000),
	}

	return []models.CartItem{
		{
			Product:  product,
			Quantity: 1,
			Size:     "M",
		},
	}
}

// NikeTshirtCart is used to exercise SUPER69 (PUMA is excluded from that voucher).
func NikeTshirtCart() []models.CartItem {
	product := models.Product{
		ID:        "n1",
		Brand:     models.BrandNike,
		Category:  models.CategoryTShirts,
		BasePrice: decimal.NewFromInt(1000),
	}

	return []models.CartItem{
		{
			Product:  product,
			Quantity: 1,
			Size:     "L",
		},
	}
}

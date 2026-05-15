package testdata

import (
	"github.com/shopspring/decimal"

	"github.com/aadiths/unifize-discount-engine/internal/models"
)

func PumaTshirtCart() []models.CartItem {

	product := models.Product{
		ID:        "p1",
		Brand:     "PUMA",
		Category:  "T-shirts",
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
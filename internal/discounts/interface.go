package discounts

import (
	"github.com/shopspring/decimal"
	"github.com/aadiths/unifize-discount-engine/internal/models"
)

type DiscountRule interface {
	Name() string

	Apply(
		items []models.CartItem,
		current decimal.Decimal,
	) (decimal.Decimal, decimal.Decimal)
}
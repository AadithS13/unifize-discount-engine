package discounts

import (
	"github.com/aadiths/unifize-discount-engine/internal/models"
	"github.com/shopspring/decimal"
)

type DiscountRule interface {
	Name() string

	Apply(
		items []models.CartItem,
		current decimal.Decimal,
	) (decimal.Decimal, decimal.Decimal)
}

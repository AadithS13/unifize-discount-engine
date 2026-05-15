package discounts

import (

	"github.com/shopspring/decimal"
	"github.com/aadiths/unifize-discount-engine/internal/models"
)

type CategoryDiscount struct{}

func (c CategoryDiscount) Name() string {
	return "Category Discount"
}

func (c CategoryDiscount) Apply(
	items []models.CartItem,
	current decimal.Decimal,
) (decimal.Decimal, decimal.Decimal) {

	discount := current.Mul(decimal.NewFromFloat(0.10))

	final := current.Sub(discount)

	return final, discount
}
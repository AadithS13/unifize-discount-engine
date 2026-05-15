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

	discount := decimal.Zero

	for _, item := range items {

		if item.Product.Category == models.CategoryTShirts {

			itemTotal := item.Product.BasePrice.
				Mul(decimal.NewFromInt(int64(item.Quantity)))

			itemDiscount := itemTotal.Mul(decimal.NewFromFloat(0.10))

			discount = discount.Add(itemDiscount)
		}
	}

	final := current.Sub(discount)

	return final, discount
}

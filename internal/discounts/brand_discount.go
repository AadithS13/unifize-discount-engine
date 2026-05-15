package discounts

import (
	"strings"

	"github.com/shopspring/decimal"
	"github.com/aadiths/unifize-discount-engine/internal/models"
)

type BrandDiscount struct{}

func (b BrandDiscount) Name() string {
	return "Brand Discount"
}

func (b BrandDiscount) Apply(
	items []models.CartItem,
	current decimal.Decimal,
) (decimal.Decimal, decimal.Decimal) {

	discount := decimal.Zero

	for _, item := range items {

		if strings.EqualFold(item.Product.Brand, "PUMA") {

			itemTotal := item.Product.BasePrice.
				Mul(decimal.NewFromInt(int64(item.Quantity)))

			itemDiscount := itemTotal.Mul(decimal.NewFromFloat(0.40))

			discount = discount.Add(itemDiscount)
		}
	}

	final := current.Sub(discount)

	return final, discount
}
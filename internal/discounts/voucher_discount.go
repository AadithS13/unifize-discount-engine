package discounts

import (
	"strings"

	"github.com/shopspring/decimal"

	"github.com/aadiths/unifize-discount-engine/internal/models"
)

type VoucherDiscount struct {
	Code string
}

func (v VoucherDiscount) Name() string {
	return "Voucher Discount"
}

func (v VoucherDiscount) Apply(
	items []models.CartItem,
	current decimal.Decimal,
) (decimal.Decimal, decimal.Decimal) {

	if !strings.EqualFold(v.Code, "SUPER69") {
		return current, decimal.Zero
	}

	discount := current.Mul(decimal.NewFromFloat(0.69))

	final := current.Sub(discount)

	return final, discount
}

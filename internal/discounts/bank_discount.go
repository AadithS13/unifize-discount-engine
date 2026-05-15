package discounts

import (
	"strings"

	"github.com/shopspring/decimal"
	"github.com/aadiths/unifize-discount-engine/internal/models"
)

type BankDiscount struct {
	Payment *models.PaymentInfo
}

func (b BankDiscount) Name() string {
	return "Bank Discount"
}

func (b BankDiscount) Apply(
	items []models.CartItem,
	current decimal.Decimal,
) (decimal.Decimal, decimal.Decimal) {

	if b.Payment == nil || b.Payment.BankName == nil {
		return current, decimal.Zero
	}

	if !strings.EqualFold(*b.Payment.BankName, "ICICI") {
		return current, decimal.Zero
	}

	discount := current.Mul(decimal.NewFromFloat(0.10))

	final := current.Sub(discount)

	return final, discount
}
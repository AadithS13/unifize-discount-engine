package models

import (
	"github.com/shopspring/decimal"
)

type BrandTier string

const (
	BrandTierPremium BrandTier = "premium"
	BrandTierRegular BrandTier = "regular"
	BrandTierBudget  BrandTier = "budget"
)

type Product struct {
	ID           string
	Brand        string
	BrandTier    BrandTier
	Category     string
	BasePrice    decimal.Decimal
	CurrentPrice decimal.Decimal
}

type CartItem struct {
	Product  Product
	Quantity int
	Size     string
}

type PaymentInfo struct {
	Method   string
	BankName *string
	CardType *string
}

type DiscountedPrice struct {
	OriginalPrice    decimal.Decimal
	FinalPrice       decimal.Decimal
	AppliedDiscounts map[string]decimal.Decimal
	Message          string
}

type CustomerProfile struct {
	ID   string
	Tier string
}
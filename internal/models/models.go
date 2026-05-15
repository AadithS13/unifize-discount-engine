package models

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type BrandTier string

const (
	BrandTierPremium BrandTier = "premium"
	BrandTierRegular BrandTier = "regular"
	BrandTierBudget  BrandTier = "budget"
)

type Brand string

const (
	BrandPuma Brand = "PUMA"
)

type Category string

const (
	CategoryTShirts Category = "T-shirts"
)

type CustomerTier string

const (
	CustomerTierPremium CustomerTier = "PREMIUM"
	CustomerTierRegular CustomerTier = "REGULAR"
)

type PaymentMethod string

const (
	PaymentMethodCard PaymentMethod = "CARD"
	PaymentMethodUPI  PaymentMethod = "UPI"
)

type VoucherCode string

const (
	VoucherSuper69 VoucherCode = "SUPER69"
)

type Product struct {
	ID           string
	Brand        Brand
	BrandTier    BrandTier
	Category     Category
	BasePrice    decimal.Decimal
	CurrentPrice decimal.Decimal
}

type CartItem struct {
	Product  Product
	Quantity int
	Size     string
}

type PaymentInfo struct {
	Method   PaymentMethod
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
	Tier CustomerTier
}

func ValidateProduct(product Product) error {

	switch product.Brand {
	case BrandPuma:
	default:
		return fmt.Errorf("invalid brand: %s", product.Brand)
	}

	switch product.Category {
	case CategoryTShirts:
	default:
		return fmt.Errorf("invalid category: %s", product.Category)
	}

	return nil
}

func ValidateCustomer(customer CustomerProfile) error {

	switch customer.Tier {
	case CustomerTierPremium, CustomerTierRegular:
	default:
		return fmt.Errorf("invalid customer tier: %s", customer.Tier)
	}

	return nil
}

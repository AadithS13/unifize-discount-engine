package service

import (
	"context"

	"github.com/shopspring/decimal"

	"github.com/aadiths/unifize-discount-engine/internal/discounts"
	"github.com/aadiths/unifize-discount-engine/internal/models"
)

type DiscountService struct{}

func NewDiscountService() *DiscountService {
	return &DiscountService{}
}

func (s *DiscountService) CalculateCartDiscounts(
	ctx context.Context,
	cartItems []models.CartItem,
	customer models.CustomerProfile,
	paymentInfo *models.PaymentInfo,
	voucherCode *string,
) (*models.DiscountedPrice, error) {

	// VALIDATION

	for _, item := range cartItems {

		if err := models.ValidateProduct(item.Product); err != nil {
			return nil, err
		}
	}

	if err := models.ValidateCustomer(customer); err != nil {
		return nil, err
	}

	// ORIGINAL TOTAL

	total := decimal.Zero

	for _, item := range cartItems {

		itemTotal := item.Product.BasePrice.
			Mul(decimal.NewFromInt(int64(item.Quantity)))

		total = total.Add(itemTotal)
	}

	original := total

	applied := map[string]decimal.Decimal{}

	// BASE RULES

	rules := []discounts.DiscountRule{
		discounts.BrandDiscount{},
		discounts.CategoryDiscount{},
	}

	// VOUCHER

	if voucherCode != nil {

		valid, err := s.ValidateDiscountCode(
			ctx,
			*voucherCode,
			cartItems,
			customer,
		)

		if err == nil && valid {

			rules = append(
				rules,
				discounts.VoucherDiscount{
					Code: *voucherCode,
				},
			)
		}
	}

	// BANK OFFER

	rules = append(
		rules,
		discounts.BankDiscount{
			Payment: paymentInfo,
		},
	)

	// APPLY RULES

	for _, rule := range rules {

		newTotal, discount := rule.Apply(cartItems, total)

		applied[rule.Name()] = discount

		total = newTotal
	}

	// RESPONSE

	return &models.DiscountedPrice{
		OriginalPrice:    original,
		FinalPrice:       total,
		AppliedDiscounts: applied,
		Message:          "Discounts applied successfully",
	}, nil
}

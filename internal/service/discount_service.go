package service

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/aadiths/unifize-discount-engine/internal/discounts"
	"github.com/aadiths/unifize-discount-engine/internal/models"
)

type discountService struct{}

func NewDiscountService() FullDiscountService {
	return &discountService{}
}

func (s *discountService) CalculateCartDiscounts(
	ctx context.Context,
	cartItems []models.CartItem,
	customer models.CustomerProfile,
	paymentInfo *models.PaymentInfo,
) (*models.DiscountedPrice, error) {
	return s.CalculateCartDiscountsWithVoucher(ctx, cartItems, customer, paymentInfo, nil)
}

func (s *discountService) CalculateCartDiscountsWithVoucher(
	ctx context.Context,
	cartItems []models.CartItem,
	customer models.CustomerProfile,
	paymentInfo *models.PaymentInfo,
	voucherCode *string,
) (*models.DiscountedPrice, error) {

	for _, item := range cartItems {
		if err := models.ValidateProduct(item.Product); err != nil {
			return nil, err
		}
	}

	if err := models.ValidateCustomer(customer); err != nil {
		return nil, err
	}

	total := decimal.Zero
	for _, item := range cartItems {
		itemTotal := item.Product.BasePrice.
			Mul(decimal.NewFromInt(int64(item.Quantity)))
		total = total.Add(itemTotal)
	}

	original := total
	applied := map[string]decimal.Decimal{}

	rules := []discounts.DiscountRule{
		discounts.BrandDiscount{},
		discounts.CategoryDiscount{},
	}

	if voucherCode != nil {
		valid, err := s.ValidateDiscountCode(ctx, *voucherCode, cartItems, customer)
		if err != nil {
			return nil, err
		}
		if !valid {
			return nil, fmt.Errorf("voucher %s cannot be applied", *voucherCode)
		}

		rules = append(rules, discounts.VoucherDiscount{Code: *voucherCode})
	}

	rules = append(rules, discounts.BankDiscount{Payment: paymentInfo})

	for _, rule := range rules {
		newTotal, discount := rule.Apply(cartItems, total)
		applied[rule.Name()] = discount
		total = newTotal
	}

	return &models.DiscountedPrice{
		OriginalPrice:    original,
		FinalPrice:       total,
		AppliedDiscounts: applied,
		Message:          "Discounts applied successfully",
	}, nil
}

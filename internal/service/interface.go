package service

import (
	"context"

	"github.com/aadiths/unifize-discount-engine/internal/models"
)

// DiscountService is the core discount engine contract from the assignment spec.
type DiscountService interface {
	CalculateCartDiscounts(
		ctx context.Context,
		cartItems []models.CartItem,
		customer models.CustomerProfile,
		paymentInfo *models.PaymentInfo,
	) (*models.DiscountedPrice, error)

	ValidateDiscountCode(
		ctx context.Context,
		code string,
		cartItems []models.CartItem,
		customer models.CustomerProfile,
	) (bool, error)
}

// FullDiscountService extends DiscountService with optional voucher application at checkout.
type FullDiscountService interface {
	DiscountService

	CalculateCartDiscountsWithVoucher(
		ctx context.Context,
		cartItems []models.CartItem,
		customer models.CustomerProfile,
		paymentInfo *models.PaymentInfo,
		voucherCode *string,
	) (*models.DiscountedPrice, error)
}

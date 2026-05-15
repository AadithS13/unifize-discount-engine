package service

import (
	"context"
	"errors"
	"strings"

	"github.com/aadiths/unifize-discount-engine/internal/models"
)

func (s *DiscountService) ValidateDiscountCode(
	ctx context.Context,
	code string,
	cartItems []models.CartItem,
	customer models.CustomerProfile,
) (bool, error) {

	if !strings.EqualFold(code, "SUPER69") {
		return false, errors.New("invalid voucher code")
	}

	// premium-only voucher
	if customer.Tier != "PREMIUM" {
		return false, errors.New("voucher allowed only for premium customers")
	}

	// exclude PUMA
	for _, item := range cartItems {

		if strings.EqualFold(item.Product.Brand, "PUMA") {
			return false, errors.New("voucher not applicable on PUMA products")
		}
	}

	return true, nil
}
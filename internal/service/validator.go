package service

import (
	"context"
	"fmt"

	"github.com/aadiths/unifize-discount-engine/internal/models"
)

func (s *DiscountService) ValidateDiscountCode(
	ctx context.Context,
	code string,
	cartItems []models.CartItem,
	customer models.CustomerProfile,
) (bool, error) {

	// validate voucher
	if code != string(models.VoucherSuper69) {
		return false, fmt.Errorf("voucher %s is invalid", code)
	}

	// premium-only voucher
	if customer.Tier != models.CustomerTierPremium {
		return false, fmt.Errorf(
			"voucher allowed only for premium customers",
		)
	}

	// example brand exclusion rule
	for _, item := range cartItems {

		if item.Product.Brand == models.BrandPuma {
			return false, fmt.Errorf(
				"voucher not applicable on PUMA products",
			)
		}
	}

	return true, nil
}

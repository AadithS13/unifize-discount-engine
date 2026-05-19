package tests

import (
	"context"
	"testing"

	"github.com/aadiths/unifize-discount-engine/internal/models"
	"github.com/aadiths/unifize-discount-engine/internal/service"
	"github.com/aadiths/unifize-discount-engine/testdata"
)

func TestCalculateCartDiscounts(t *testing.T) {
	bank := "ICICI"
	premium := models.CustomerProfile{ID: "c1", Tier: models.CustomerTierPremium}
	iciciCard := &models.PaymentInfo{
		Method:   models.PaymentMethodCard,
		BankName: &bank,
	}

	tests := []struct {
		name            string
		cart            []models.CartItem
		customer        models.CustomerProfile
		payment         *models.PaymentInfo
		voucher         *string
		wantFinal       string
		wantDiscounts   map[string]string
		wantErr         bool
		useVoucherParam bool
	}{
		{
			name:      "PUMA T-shirt with brand, category, and ICICI bank offer",
			cart:      testdata.PumaTshirtCart(),
			customer:  premium,
			payment:   iciciCard,
			wantFinal: "450",
			wantDiscounts: map[string]string{
				"Brand Discount":    "400",
				"Category Discount": "100",
				"Bank Discount":     "50",
			},
		},
		{
			name:      "PUMA T-shirt without bank offer",
			cart:      testdata.PumaTshirtCart(),
			customer:  premium,
			payment:   nil,
			wantFinal: "500",
			wantDiscounts: map[string]string{
				"Brand Discount":    "400",
				"Category Discount": "100",
				"Bank Discount":     "0",
			},
		},
		{
			name:            "NIKE T-shirt with SUPER69 and ICICI",
			cart:            testdata.NikeTshirtCart(),
			customer:        premium,
			payment:         iciciCard,
			voucher:         strPtr("SUPER69"),
			useVoucherParam: true,
			wantFinal:       "251",
			wantDiscounts: map[string]string{
				"Brand Discount":    "0",
				"Category Discount": "100",
				"Voucher Discount":  "621",
				"Bank Discount":     "28",
			},
		},
		{
			name:            "SUPER69 rejected on PUMA cart",
			cart:            testdata.PumaTshirtCart(),
			customer:        premium,
			payment:         iciciCard,
			voucher:         strPtr("SUPER69"),
			useVoucherParam: true,
			wantErr:         true,
		},
	}

	svc := service.NewDiscountService()
	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				res *models.DiscountedPrice
				err error
			)

			if tt.useVoucherParam {
				res, err = svc.CalculateCartDiscountsWithVoucher(
					ctx, tt.cart, tt.customer, tt.payment, tt.voucher,
				)
			} else {
				res, err = svc.CalculateCartDiscounts(
					ctx, tt.cart, tt.customer, tt.payment,
				)
			}

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if res.FinalPrice.StringFixed(0) != tt.wantFinal {
				t.Fatalf("final price: want %s, got %s", tt.wantFinal, res.FinalPrice.StringFixed(0))
			}

			for name, want := range tt.wantDiscounts {
				got := res.AppliedDiscounts[name].StringFixed(0)
				if got != want {
					t.Fatalf("discount %q: want %s, got %s", name, want, got)
				}
			}
		})
	}
}

func TestDiscountServiceInterface(t *testing.T) {
	var _ service.DiscountService = service.NewDiscountService()
}

func TestVoucherValidation(t *testing.T) {
	svc := service.NewDiscountService()

	valid, err := svc.ValidateDiscountCode(
		context.Background(),
		"SUPER69",
		testdata.PumaTshirtCart(),
		models.CustomerProfile{ID: "c1", Tier: models.CustomerTierPremium},
	)

	if valid {
		t.Fatal("expected voucher to fail for PUMA")
	}
	if err == nil {
		t.Fatal("expected validation error")
	}
}

func TestValidateDiscountCode(t *testing.T) {
	tests := []struct {
		name        string
		cart        []models.CartItem
		customer    models.CustomerProfile
		expectValid bool
	}{
		{
			name: "SUPER69 invalid on PUMA for premium customer",
			cart: testdata.PumaTshirtCart(),
			customer: models.CustomerProfile{
				ID:   "c1",
				Tier: models.CustomerTierPremium,
			},
			expectValid: false,
		},
		{
			name: "SUPER69 valid on NIKE for premium customer",
			cart: testdata.NikeTshirtCart(),
			customer: models.CustomerProfile{
				ID:   "c1",
				Tier: models.CustomerTierPremium,
			},
			expectValid: true,
		},
		{
			name: "SUPER69 invalid for regular customer",
			cart: testdata.NikeTshirtCart(),
			customer: models.CustomerProfile{
				ID:   "c2",
				Tier: models.CustomerTierRegular,
			},
			expectValid: false,
		},
	}

	svc := service.NewDiscountService()
	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := svc.ValidateDiscountCode(ctx, "SUPER69", tt.cart, tt.customer)

			if valid != tt.expectValid {
				t.Fatalf("valid: want %v, got %v (err=%v)", tt.expectValid, valid, err)
			}

			if tt.expectValid && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if !tt.expectValid && err == nil {
				t.Fatal("expected validation error")
			}
		})
	}
}

func strPtr(s string) *string {
	return &s
}

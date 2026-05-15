package tests

import (
	"context"
	"testing"

	"github.com/aadiths/unifize-discount-engine/internal/models"
	"github.com/aadiths/unifize-discount-engine/internal/service"
	"github.com/aadiths/unifize-discount-engine/internal/testdata"
)

func TestDiscountCalculation(t *testing.T) {

	svc := service.NewDiscountService()

	bank := "ICICI"

	res, err := svc.CalculateCartDiscounts(
		context.Background(),
		testdata.PumaTshirtCart(),
		models.CustomerProfile{
			ID:   "c1",
			Tier: "PREMIUM",
		},
		&models.PaymentInfo{
			Method:   "CARD",
			BankName: &bank,
		},
		nil, 
	)

	if err != nil {
		t.Fatal(err)
	}

	expected := "486"

	if res.FinalPrice.StringFixed(0) != expected {
		t.Fatalf(
			"expected %s got %s",
			expected,
			res.FinalPrice.StringFixed(0),
		)
	}
}

func TestVoucherValidation(t *testing.T) {

	svc := service.NewDiscountService()

	valid, err := svc.ValidateDiscountCode(
		context.Background(),
		"SUPER69",
		testdata.PumaTshirtCart(),
		models.CustomerProfile{
			ID:   "c1",
			Tier: "PREMIUM",
		},
	)

	if valid {
		t.Fatal("expected voucher to fail for PUMA")
	}

	if err == nil {
		t.Fatal("expected validation error")
	}
}
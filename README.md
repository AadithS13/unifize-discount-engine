# Unifize Discount Engine

A modular and extensible e-commerce discount engine built in Go.

This project models common fashion e-commerce discount scenarios such as:
- Brand discounts
- Category discounts
- Voucher/coupon codes
- Bank card offers

The implementation focuses on clean architecture, extensibility, validation, and maintainable business logic.

---

# Features

- Brand-specific discounts
- Category-based discounts
- Voucher validation and application
- Bank offer discounts
- Sequential discount stacking
- Typed domain models and validation
- Precision-safe money calculations using decimal
- Table-driven tests
- Lint-clean and idiomatic Go code

---

# Tech Stack

- Go
- shopspring/decimal
- golangci-lint
- Go testing package

---

# Project Structure

```text
unifize-discount-engine/
│
├── internal/
│   ├── discounts/
│   ├── models/
│   └── service/
│
├── testdata/
├── tests/
│
├── go.mod
└── README.md
```

---

# Discount Flow

```text
Brand Discount
→ Category Discount
→ Voucher Discount
→ Bank Offer
```

---

# Example Scenario

Input:
- PUMA T-shirt
- Base price: ₹1000
- PUMA discount: 40%
- T-shirt category discount: 10%
- ICICI bank offer: 10%

Flow (brand and category discounts are each computed on base price; bank applies on the running total):

```text
1000
→ 40% brand discount on base (−400) = 600
→ 10% category discount on base (−100) = 500
→ 10% ICICI bank discount on total (−50) = 450
```

Final Price:

```text
₹450
```

---

# Design Decisions

## Rule-Based Discount Engine

Discounts are modeled using a common interface:

```go
type DiscountRule interface {
    Name() string
    Apply(...)
}
```

This makes the system extensible and allows new discount types to be added easily.

---

## Sequential Discount Application

Rules run in order: brand → category → voucher → bank. Brand and category discounts are each calculated as a percentage of the item’s **base price**, then subtracted from the running cart total. Voucher and bank discounts apply as a percentage of the **current** total after prior steps.

---

## Typed Domain Models

Typed domain constants were introduced for:
- Brands
- Categories
- Customer tiers
- Payment methods
- Voucher codes

This prevents invalid business states from entering the pricing engine.

---

## Validation Layer Separation

Voucher validation is separated from discount calculation.

This keeps:
- eligibility rules
- pricing logic

independent and easier to maintain.

---

## Decimal-Based Money Handling

Money calculations use:

```go
github.com/shopspring/decimal
```

instead of `float64` to avoid floating point precision issues in financial calculations.

---

# Validation Rules Implemented

## Voucher: SUPER69

Rules:
- Only valid for PREMIUM customers
- Not applicable on PUMA products

---

# Assumptions

- One happy path implementation is sufficient
- Brand and category discounts use base price; voucher and bank use the running total
- Invalid voucher codes passed at checkout return an error (not silently ignored)
- `CalculateCartDiscounts` matches the assignment interface; `CalculateCartDiscountsWithVoucher` is an extension on `FullDiscountService`
- No database or persistence layer required
- Only core pricing/business logic is modeled

---

# Testing

Run tests:

```bash
go test ./...
```

---

# Formatting and Linting

Format code:

```bash
go fmt ./...
```

Run lint checks:

```bash
golangci-lint run
```

---

# How to Run

## Install dependencies

```bash
go mod tidy
```

## Run tests

```bash
go test ./...
```

---

# AI-Assisted Development Workflow

Initial implementation was generated using AI assistance.

Subsequent commits focused on:
- improving architecture
- correcting business logic
- strengthening validation
- improving test coverage
- enforcing lint compliance
- introducing typed domain modeling

---

# Future Improvements

Potential future enhancements:
- Config-driven discount rules
- Time-bound offers
- Multiple voucher stacking
- Cart-level promotions
- Customer-specific campaigns
- Database-backed rule management
- REST/gRPC APIs

---

# Author

Aadith S
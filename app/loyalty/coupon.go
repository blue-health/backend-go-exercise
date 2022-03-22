package loyalty

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	// Campaign can apply coupons that belong to it
	// to the shopping card, based on a set of conditions and effects.
	Campaign struct {
		ID         uuid.UUID
		Name       string
		Conditions []Condition
		Effects    []Effect
		Coupons    []Coupon
	}

	// Coupon is represented by a unique ID,
	// a unique code printed on the receipt,
	// and information about its validity.
	Coupon struct {
		ID        uuid.UUID
		Code      string
		Applied   bool
		ValidTill time.Time
	}

	ApplyCouponCommand struct {
		CouponCode string
		Cart       *Cart
	}
)

// ApplyCoupon performs a few actions:
// 1. Checks if the coupon with this code was belongs to this campaign, and is valid.
// 2. Checks if any of the conditions for this campaign are true for the cart.
// 3. If so, it applies all the effects to the cart, and marks the coupon as applied.
func (c *Campaign) ApplyCoupon(_ context.Context, _ ApplyCouponCommand) (*Cart, error) {
	return nil, nil
}

// IsValid checks if coupon is valid.
func (c *Coupon) IsValid() bool { return true }

// WasApplied checks if coupon was applied.
func (c *Coupon) WasApplied() bool { return true }

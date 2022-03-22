package loyalty

type (
	// Effect is applied to the shopping cart
	// to discount it in some way. It mutates the cart.
	Effect interface {
		Name() string
		Apply(*Cart) error
	}

	// DiscountCartEffect discounts the value of the whole cart
	// by a specified percentage (e.g -10%).
	DiscountCartEffect struct{}

	// DiscountCartEffect discounts the value of items belonging to
	// the particular brand by a specified percentage (e.g -10%).
	DiscountBrandEffect struct{}

	// DiscountCartEffect discounts the value of items with
	// particular IDs by a specified percentage (e.g -10%).
	DiscountItemEffect struct{}
)

// Please provide implementations of the effects above.

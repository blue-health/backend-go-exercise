package loyalty

type (
	// Condition can evaluate to a boolean value
	// for a given cart. It is used to check if
	// a coupon can be applied to the cart.
	Condition interface {
		Name() string
		Check(*Cart) bool
	}

	// AlwaysCondition always returns true.
	AlwaysCondition struct{}

	// NeverCondition always returns false.
	NeverCondition struct{}

	// OrCondition returns a logical OR of all of its child conditions.
	OrCondition struct{}

	// AndCondition returns a logical AND of all of its child conditions.
	AndCondition struct{}

	// ContainsBrandCondition checks if the cart contains
	// at least one product of a given brand(s).
	ContainsBrandCondition struct{}

	// ContainsItemCondition checks if the cart contains
	// at least one product with the given ID(s).
	ContainsItemCondition struct{}
)

// AlwaysCondition implements Condition.
var _ Condition = &AlwaysCondition{}

func (c *AlwaysCondition) Name() string { return "ALWAYS" }

func (c *AlwaysCondition) Check(_ *Cart) bool { return true }

// Please provide implementations of the remaining conditions.

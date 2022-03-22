package loyalty

import "github.com/google/uuid"

type (
	// Cart represents the shopping basket
	// of a customer of the pharmacy
	Cart struct {
		Items []Item
	}

	// Item represents the contents of the
	// shopping basket. Each items holds a unique ID,
	// a name and other important data.
	Item struct {
		ID       uuid.UUID
		Category Category
		Brand    Brand
		Quantity uint
		Price    uint
	}

	// Brand of medical products
	Brand string

	// Category of medical products
	Category string
)

const (
	// Brands
	Bayer Brand = "bayer"
	Wick  Brand = "wick"
	Hylo  Brand = "hylo"
	// Categories
	Lotions Category = "lotions"
	Creams  Category = "creams"
	Gels    Category = "gels"
	Pills   Category = "pills"
)

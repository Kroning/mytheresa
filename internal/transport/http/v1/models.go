package v1

type ProductWithDiscountResponse struct {
	Sku      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    Price  `json:"price"`
}

type Price struct {
	Original int     `json:"original"`
	Final    int     `json:"final"`
	Discount *string `json:"discount_percentage"`
	Currency string  `json:"currency"`
}

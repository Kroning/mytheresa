package product

type ProductRow struct {
	Sku      string `db:"sku"`
	Name     string `db:"name"`
	Category string `db:"category"`
	Price    int    `db:"price"`
}

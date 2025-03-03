package domain

type Product struct {
	ID          int64   `json:"id"`
	SupplierID  int64   `json:"supplier_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

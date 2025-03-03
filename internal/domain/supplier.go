package domain

type Supplier struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Site  string `json:"site"`
}

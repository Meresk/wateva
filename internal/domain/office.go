package domain

type Office struct {
	ID      int64  `json:"id"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

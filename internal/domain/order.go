package domain

import "time"

type Order struct {
	ID          int64      `json:"id"`
	OfficeID    int64      `json:"office_id"`
	Description string     `json:"description"`
	Sale        int16      `json:"sale"`
	Cost        float32    `json:"cost"`
	CreatedAt   time.Time  `json:"created_at"`
	Product     []*Product `json:"product,omitempty"`
}

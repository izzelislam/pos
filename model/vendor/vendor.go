package vendor

import "time"

type Vendor struct {
	ID				 int `json:"id"`
	Name			 string `json:"name"`
	Email string `json:"email"`
	NoTlp string `json:"no_tlp"`
	Address string `json:"address"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
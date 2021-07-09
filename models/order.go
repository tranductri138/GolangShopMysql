package models

type Order struct {
	Id           int     `json:"id"`
	Totalproduct int     `json:"total_product"`
	Totalprice   float32 `json:"total_price"`
	Customer     string  `json:"customer"`
	Address      string  `json:"address"`
	Phone        string  `json:"phone"`
	Email        string  `json:"email"`
	Payment      string  `json:"payment"`
}

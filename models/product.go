package models

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	IsSales     bool    `json:"is_sale"`
	Price       float32 `json:"price"`
	PriceSale   float32 `json:"price_sale"`
	Description string  `json:"des"`
	CreatedAt   int64   `json:"created_at"`
	ModifiedAt  int64   `json:"modified_at"`
	CategoryId  int     `json:"category_id"`
}

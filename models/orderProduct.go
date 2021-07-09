package models

type OderProduct struct {
	IdProduct int `json:"id_product"`
	IdOrder   int `json:"id_order"`
	Quantity  int `json:"quantity"`
}

package CartHandler

type RequestBody struct {
	ID     uint `json:"id"`
	Amount int  `json:"amount"`
}

type Product struct {
	ProductName string  `json:"product_name"`
	Amount      int     `json:"amount"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
	ProductId   uint    `json:"product_id"`
}

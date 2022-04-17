package orderHandler

type Product struct {
	ProductName string
	Amount      int
	UnitPrice   float64
	TotalPrice  float64
}

type Order struct {
	TotalPrice float64
	Amount     int
}

package inventory

type Item struct {
	SKU              string
	Name             string
	CurrentStock     int
	ReorderThreshold int
}

// // recommended order quantity calc: threshold * 2 - current stock
// testReport := ReorderReport{
// 	Item: testItem,
// 	Shortfall: 1 //threshold - current stock
// 	QuantityToOrder: 1
// }

type ReorderReport struct {
	Item            Item
	Shortfall       int
	QuantityToOrder int
}

func GenerateReorderReport(items []Item) ([]ReorderReport, error) {
	res := []ReorderReport{
		{Item: items[0], Shortfall: 0, QuantityToOrder: 0},
	}

	return res, nil
}

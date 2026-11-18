package inventory

import (
	"cmp"
	"slices"
)

type Item struct {
	SKU              string
	Name             string
	CurrentStock     int
	ReorderThreshold int
}

type ReorderReport struct {
	Item            Item
	Shortfall       int
	QuantityToOrder int
}

func GenerateReorderReport(items []Item) ([]ReorderReport, error) {
	base := []ReorderReport{}

	if len(items) == 0 {
		return base, nil
	}

	for _, item := range items {
		if item.CurrentStock <= item.ReorderThreshold {
			baseReport := ReorderReport{
				Item: item}
			baseReport.Shortfall = baseReport.Item.ReorderThreshold - baseReport.Item.CurrentStock
			baseReport.QuantityToOrder = (baseReport.Item.ReorderThreshold * 2) - baseReport.Item.CurrentStock
			base = append(base, baseReport)
		}
	}

	slices.SortFunc(base, func(a, b ReorderReport) int {
		return cmp.Compare(b.Shortfall, a.Shortfall)
	})

	return base, nil
}

package inventory

import (
	"cmp"
	"errors"
	"fmt"
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

var badItems []Item
var ErrNegativeStock = errors.New("inventory: negative stock")

func GenerateReorderReport(items []Item) ([]ReorderReport, error) {
	var (
		base     []ReorderReport
		badItems []error
	)

	if len(items) == 0 {
		return base, nil
	}

	for _, item := range items {
		if item.CurrentStock < 0 || item.ReorderThreshold < 0 {
			badItems = append(badItems, fmt.Errorf("%w: sku: %q (stock=%d, threshold=%d)",
				ErrNegativeStock, item.SKU, item.CurrentStock, item.ReorderThreshold))
			continue
		}
		if item.CurrentStock <= item.ReorderThreshold {
			baseReport := ReorderReport{
				Item: item}
			baseReport.Shortfall = baseReport.Item.ReorderThreshold - baseReport.Item.CurrentStock
			baseReport.QuantityToOrder = (baseReport.Item.ReorderThreshold * 2) - baseReport.Item.CurrentStock
			base = append(base, baseReport)
		}
	}

	if len(badItems) > 0 && len(base) == 0 {
		return nil, ErrNegativeStock
	} else if len(badItems) > 0 {
		return nil, errors.Join(badItems...)
	}

	slices.SortFunc(base, func(a, b ReorderReport) int {
		return cmp.Or(cmp.Compare(b.Shortfall, a.Shortfall), cmp.Compare(a.Item.SKU, b.Item.SKU))
	})

	return base, nil
}

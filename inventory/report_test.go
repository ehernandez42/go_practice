package inventory

import (
	"reflect"
	"testing"
)

func TestGenerateReorderReportOverstock(t *testing.T) {
	testItems := []Item{
		{SKU: "SKU-1", Name: "test-item-1", CurrentStock: 2, ReorderThreshold: 1},
	}

	got, err := GenerateReorderReport(testItems)

	if err != nil {
		t.Errorf("runtime error: %q", err)
	}

	if len(got) != 0 {
		t.Error("we're supposed to have an empty list")
	}
}
func TestReorderReportDescending(t *testing.T) {
	testItems := []Item{
		{SKU: "SKU-1", Name: "test-item-1", CurrentStock: 1, ReorderThreshold: 3},
		{SKU: "SKU-2", Name: "test-item-2", CurrentStock: 2, ReorderThreshold: 2},
		{SKU: "SKU-3", Name: "test-item-3", CurrentStock: 3, ReorderThreshold: 7},
	}

	got, err := GenerateReorderReport(testItems)
	want := []ReorderReport{
		{Item: Item{SKU: "SKU-3", Name: "test-item-3", CurrentStock: 3, ReorderThreshold: 7}, Shortfall: 4, QuantityToOrder: 11},
		{Item: Item{SKU: "SKU-1", Name: "test-item-1", CurrentStock: 1, ReorderThreshold: 3}, Shortfall: 2, QuantityToOrder: 5},
		{Item: Item{SKU: "SKU-2", Name: "test-item-2", CurrentStock: 2, ReorderThreshold: 2}, Shortfall: 0, QuantityToOrder: 2},
	}

	if err != nil {
		t.Errorf("runtime error: %q", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

func TestReorderReportNegativeStock(t *testing.T) {
	testItems := []Item{
		{SKU: "SKU-1", Name: "test-item-1", CurrentStock: -1, ReorderThreshold: 0},
	}
	_, err := GenerateReorderReport(testItems)

	if err == nil {
		t.Errorf("error must be present")
	}
}

func TestReorderReportTieBreaker(t *testing.T) {
	testItems := []Item{
		{SKU: "SKU-2", Name: "test-item-2", CurrentStock: 1, ReorderThreshold: 3},
		{SKU: "SKU-1", Name: "test-item-1", CurrentStock: 1, ReorderThreshold: 3},
	}

	got, err := GenerateReorderReport(testItems)
	want := []ReorderReport{
		{Item: Item{SKU: "SKU-1", Name: "test-item-1", CurrentStock: 1, ReorderThreshold: 3}, Shortfall: 2, QuantityToOrder: 5},
		{Item: Item{SKU: "SKU-2", Name: "test-item-2", CurrentStock: 1, ReorderThreshold: 3}, Shortfall: 2, QuantityToOrder: 5},
	}

	if err != nil {
		t.Errorf("runtime error: %q", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

package inventory

import (
	"reflect"
	"testing"
)

func TestGenerateReorderReport(t *testing.T) {
	testItems := []Item{
		{SKU: "SKU-1", Name: "test-item-1", CurrentStock: 2, ReorderThreshold: 1},
	}

	got, err := GenerateReorderReport(testItems)
	wantItem := Item{SKU: "SKU-1", Name: "test-item-1", CurrentStock: 2, ReorderThreshold: 1}

	want := []ReorderReport{
		{Item: wantItem,
			Shortfall:       1, //threshold - current stock
			QuantityToOrder: 1},
	}

	if err != nil {
		t.Errorf("runtime error: %q", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

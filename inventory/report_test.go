package inventory

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestGeneratedReorderReports(t *testing.T) {
	tests := []struct {
		name            string
		items           []Item
		want            []ReorderReport
		wantErr         error
		wantErrContains string
	}{
		{
			name:  "empty input returns an empty ReorderReport",
			items: nil,
		},
		{
			name:  "overstocked items are filtered out",
			items: []Item{{SKU: "SKU-1", Name: "test-item-1", CurrentStock: 2, ReorderThreshold: 1}},
		},
		{
			name: "exactly at threshold is included",
			items: []Item{
				{SKU: "SKU-1", Name: "test-item-1", CurrentStock: 2, ReorderThreshold: 2},
			},
			want: []ReorderReport{
				{Item: Item{SKU: "SKU-1", Name: "test-item-1", CurrentStock: 2, ReorderThreshold: 2}, Shortfall: 0, QuantityToOrder: 2},
			},
		},
		{
			name: "most urgent first",
			items: []Item{
				{SKU: "SKU-1", Name: "test-item-1", CurrentStock: 1, ReorderThreshold: 3},
				{SKU: "SKU-2", Name: "test-item-2", CurrentStock: 2, ReorderThreshold: 2},
				{SKU: "SKU-3", Name: "test-item-3", CurrentStock: 3, ReorderThreshold: 7},
			},
			want: []ReorderReport{
				{Item: Item{SKU: "SKU-3", Name: "test-item-3", CurrentStock: 3, ReorderThreshold: 7}, Shortfall: 4, QuantityToOrder: 11},
				{Item: Item{SKU: "SKU-1", Name: "test-item-1", CurrentStock: 1, ReorderThreshold: 3}, Shortfall: 2, QuantityToOrder: 5},
				{Item: Item{SKU: "SKU-2", Name: "test-item-2", CurrentStock: 2, ReorderThreshold: 2}, Shortfall: 0, QuantityToOrder: 2},
			},
		},
		{
			name: "equal shortfalls tie-break by SKU",
			items: []Item{
				{SKU: "SKU-2", Name: "test-item-2", CurrentStock: 1, ReorderThreshold: 3},
				{SKU: "SKU-1", Name: "test-item-1", CurrentStock: 1, ReorderThreshold: 3},
			},
			want: []ReorderReport{
				{Item: Item{SKU: "SKU-1", Name: "test-item-1", CurrentStock: 1, ReorderThreshold: 3}, Shortfall: 2, QuantityToOrder: 5},
				{Item: Item{SKU: "SKU-2", Name: "test-item-2", CurrentStock: 1, ReorderThreshold: 3}, Shortfall: 2, QuantityToOrder: 5},
			},
		},
		{
			name: "negative stock or threshold is rejected",
			items: []Item{
				{SKU: "SKU-A", Name: "a", CurrentStock: 5, ReorderThreshold: -3},
				{SKU: "SKU-B", Name: "b", CurrentStock: -1, ReorderThreshold: -5},
			},
			wantErr:         ErrNegativeStock,
			wantErrContains: `sku: "SKU-A"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateReorderReport(tt.items)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error = %v, want %v", err, tt.wantErr)
			}
			if tt.wantErrContains != "" && !strings.Contains(err.Error(), tt.wantErrContains) {
				t.Errorf("%q error should mention %q", err, tt.wantErrContains)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("got  %+v\nwant %+v", got, tt.want)
			}
		})
	}
}

# RETAIL-4821 — Design Summary

Quick reference of the design decisions for the Inventory Reorder Report ticket.

## Package & Layout

```text
retail/inventory/          (module)
└── inventory/             (package — no underscores in Go package names)
    ├── report.go          (types + GenerateReorderReport)
    └── report_test.go
```

- Package: `inventory` (matches the ticket's technical notes; underscored names like `inventory_report` are non-idiomatic Go).
- Keep types and logic together in `report.go`; split into `types.go` only when more types arrive.

## Data Types

```go
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

func GenerateReorderReport(items []Item) ([]ReorderReport, error)
```

## Core Rules (from the ACs)

- **Needs reorder:** `CurrentStock <= ReorderThreshold`.
- **Shortfall:** `threshold - current stock`.
- **Quantity to order:** `threshold * 2 - current stock`.
- **Sort:** most urgent first (largest shortfall), SKU ascending as tie-break.
- **Empty input:** empty result, no error. **Malformed input (negatives):** descriptive error, no panic.

## Sorting — `slices.SortFunc` (Go 1.21+)

Chosen over `sort.Slice` for **compile-time type safety and no reflection** (faster), not "memory safety". Requires a 3-way comparator with a consistent strict weak ordering; it is **not stable**, which is fine because the SKU tie-break fully determines order.

```go
slices.SortFunc(reports, func(a, b ReorderReport) int {
	return cmp.Or(
		cmp.Compare(b.Shortfall, a.Shortfall), // most urgent first
		cmp.Compare(a.Item.SKU, b.Item.SKU),  // SKU tie-break
	)
})
```

## Error Handling — collect all bad items, `errors.Join`

Fail on data errors but report **every** invalid item, not just the first.

```go
var ErrInvalidStock = errors.New("inventory: negative stock")

// in the loop:
if it.CurrentStock < 0 || it.ReorderThreshold < 0 {
	badItems = append(badItems, fmt.Errorf("%w: sku %q (stock=%d, threshold=%d)",
		ErrInvalidStock, it.SKU, it.CurrentStock, it.ReorderThreshold))
	continue
}
// ...build reports for valid items...

if len(badItems) > 0 {
	return nil, errors.Join(badItems...)
}
```

- `errors.Join` keeps every wrapped error in the tree, so `errors.Is(err, ErrInvalidStock)` still works.
- Sentinel error gives tests something stable to assert on (no fragile string-matching).
- Return the error **before** sorting; nobody consumes sorted results that are thrown away.
- Open question: one sentinel, or split into `ErrNegativeStock` / `ErrNegativeThreshold` if the data team needs to distinguish the cases.

## Implementation Flow

1. Loop over items once: validate, record bad items, build `ReorderReport` for items at/below threshold.
2. If any bad items, return `nil, errors.Join(badItems...)`.
3. Sort with the `cmp.Or` comparator.
4. Return reports.

## Testing Checklist (from AC #7)

- Normal case
- Empty input → empty result, no error
- Tie-breaking (equal shortfalls → SKU order)
- Exactly at threshold (must be included)
- Malformed input → `errors.Is(err, ErrInvalidStock)`, all bad SKUs reported

`go test ./...` must pass; coverage ≥ 80% on new code (from the DoD).
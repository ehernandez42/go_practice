# RETAIL-4821 — Inventory Reorder Report: flag low-stock items for store managers

| | |
|---|---|
| **Type** | Story |
| **Priority** | Medium |
| **Sprint** | 24 |
| **Reporter** | Sarah Kim (Product) |
| **Assignee** | You |

**Labels:** `backend` · `inventory-api` · `go`

---

## Business Context

Store managers currently check low-stock items manually every morning, which is slow and error-prone. This is the first backend piece of the automated reorder project.

## Description

As a store manager, I want an automated report that flags items running low on stock, so that I can reorder before we run out.

## Acceptance Criteria

1. Implement a function that takes a list of inventory items and returns the items that need reordering.
2. An item needs reordering when its current stock is **at or below** its reorder threshold.
3. Each result must include: SKU, product name, current stock, reorder threshold, and the recommended reorder quantity, which is: `threshold * 2 - current stock`.
4. Results must be sorted by "most urgent first" (largest stock shortfall = threshold minus current stock), and alphabetical by SKU for ties.
5. If input is empty, return an empty result, not an error.
6. Malformed input (negative stock or negative threshold) must return a descriptive error, not panic.
7. Unit tests covering: normal case, empty input, tie-breaking, exactly-at-threshold, and malformed input.

## Technical Notes

- Language: Go (module `retail/inventory`)
- Suggested signatures (adjust if you have a better idea):

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

- No database needed for this ticket, just pure logic and tests. Persistence is **RETAIL-4902**.
- Keep it in the `inventory` package; no changes to the HTTP layer.

## Out of Scope

- Persisting reports
- Notifications / email digests
- Multi-warehouse support (that's **RETAIL-4890**)

## Definition of Done

- Code reviewed and merged
- `go test ./...` passes, coverage on new code ≥ 80%
- PR description includes a short example of the report output
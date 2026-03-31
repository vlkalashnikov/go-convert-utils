# go-convert-utils
Utility helpers for value conversion, slice/map operations, and date/time helpers.

## Installation
```bash
go get github.com/vlkalashnikov/go-convert-utils
```

## Quick Start
```go
package main

import (
	"fmt"
	"time"

	utils "github.com/vlkalashnikov/go-convert-utils"
)

func main() {
	// Chunking
	fmt.Println(utils.ChunkArray([]int{1, 2, 3, 4, 5}, 2)) // [[1 2] [3 4] [5]]
	fmt.Println(utils.ChunkArray([]int{1, 2, 3}, 0))       // nil

	// Deduplication with typed comparable key
	type user struct {
		ID   int
		Name string
	}
	users := []user{{1, "a"}, {2, "b"}, {1, "c"}}
	fmt.Println(utils.UniqueBy(users, func(u user) int { return u.ID })) // [{1 a} {2 b}]

	// Date helper
	d := time.Date(2026, 3, 1, 10, 0, 0, 0, time.UTC)
	fmt.Println(utils.IsFirstDayOfMonth(d)) // true
}
```

## Main Functions
- `ChunkArray`: splits a slice into chunks; returns `nil` for invalid chunk size (`<= 0`).
- `UniqueBy`: generic unique helper using typed comparable key.
- `UniqSlice`: keeps backward compatibility and now safely handles non-comparable keys (for example slices/maps) without panic.
- `JsonB`: implements SQL `Value`/`Scan`; `Scan` supports `nil`, `[]byte`, and `string`.
- `IsFirstDayOfMonth`, `IsLastDayOfMonth`, `DaySlice`, `TruncateToDay`, `EndOfDay`: date helpers.

## Backward Compatibility
Deprecated aliases are still supported:
- `ChunckArray` -> `ChunkArray`
- `IsFitstDayOfMonth` -> `IsFirstDayOfMonth`

## Local Checks
```bash
go test ./...
go vet ./...
```

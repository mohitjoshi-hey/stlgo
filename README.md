# stlgo

High-performance, generic data structures and algorithms for Go 1.23+.

`stlgo` is built for competitive programming, low-latency applications, and high-throughput systems engineering. It provides modern, type-safe data structures and algorithms designed around Go generics, contiguous memory layouts, CPU cache locality, and zero-allocation semantics.

---

## Design Philosophy

* **Contiguous Memory & Cache Locality:** Algorithms operate directly on native Go slices, and containers prioritize slice-backed storage over pointer-heavy node graphs, maximizing CPU L1/L2 cache hits.
* **Zero-Allocation & Inlinable:** Functions avoid dynamic interface boxing (`interface{}` / `any`) by leveraging `cmp.Ordered`, `comparable`, and numeric type constraints, enabling compiler inlining and vectorization without heap escapes.
* **GC-Leak Free:** Container pop/dequeue operations explicitly zero retired slots (`var zero T`), ensuring stale references are immediately eligible for garbage collection.
* **Bring Your Own Concurrency (BYOC):** Containers do not include internal mutexes by default. This eliminates synchronization overhead in single-threaded, competitive programming, and critical-path loops.
* **Predictable Complexity:** Every algorithm and container operation carries strict time and space complexity guarantees matching C++ STL counterparts, verified against the [benchmark results](#benchmarks) below.

---

## Installation

```bash
go get github.com/mohitjoshi-hey/stlgo
```

## Quick Start & Usage

```go
package main

import (
	"fmt"

	"github.com/mohitjoshi-hey/stlgo/algo"
)

func main() {
	// Binary Search & Bounds (O(log N))
	sorted := []int{10, 20, 20, 30, 40}
	lb := algo.LowerBound(sorted, 20) // 1
	ub := algo.UpperBound(sorted, 20) // 3
	idx, ok := algo.BinarySearch(sorted, 30) // (3, true)
	fmt.Printf("LB: %d, UB: %d, Found at: %d (ok: %t)\n", lb, ub, idx, ok)

	// Unsorted Search & Custom Predicates (O(N))
	items := []string{"apple", "banana", "cherry"}
	pos, _ := algo.Find(items, "banana") // (1, true)
	firstC, _ := algo.FindIf(items, func(s string) bool {
		return len(s) > 5 // "banana"
	})
	fmt.Printf("Find: %d, FindIf: %d\n", pos, firstC)

	// Numeric Utilities
	nums := []int{1, 2, 3, 4, 5}
	sum := algo.Sum(nums)              // 15
	prefix := algo.PrefixSum(nums)     // [1, 3, 6, 10, 15]
	maxVal, _ := algo.MaxElement(nums) // (5, true)
	seq := algo.Iota(10, 4)            // [10, 11, 12, 13]

	fmt.Printf("Sum: %d, Prefix: %v, Max: %d, Iota: %v\n", sum, prefix, maxVal, seq)
}
```

## Package Overview

### 1. Algorithms (`algo`)

The `algo` package exposes a facade over specialized subpackages:

| **Subpackage** | **Available Functions** | **Details** |
|---|---|---|
| **`algo/search`** | `LowerBound`, `UpperBound`, `BinarySearch`, `Find`, `FindIf` | $O(\log N)$ binary searching for `cmp.Ordered` slices and linear scans for `comparable` / arbitrary types. |
| **`algo/sort`** | `Sort`, `DescSort`, `SortBy`, `IsSorted`, `NthElement` | Type-safe in-place sorting and $O(N)$ Introselect order statistics. |
| **`algo/numeric`** | `Sum`, `Product`, `Iota`, `MaxElement`, `MinElement`, `PrefixSum` | Generic arithmetic, array generation, and extrema extraction. |
| **`algo/math`** | `GCD`, `LCM`, `IsPrime`, `IsEven`, `IsOdd`, `Max`, `Min`, `Clamp`, `Abs` | Common competitive programming mathematical utilities. |
| **`algo/permutations`** | `NextPermutation`, `PrevPermutation` | $O(N)$ lexicographic permutation stepping via in-place reversal. |

### 2. Containers (`container`)

- **`Stack`**: Contiguous slice-backed LIFO stack with minimal reallocations.
- **`Queue`**: FIFO queue using a sliding window buffer with amortized memory compaction.
- **`Deque`**: Double-ended queue backed by a circular ring buffer, supporting $O(1)$ amortized push/pop from both ends with automatic unwrap-on-grow.

## Benchmarks

Median of 10 runs (`go test -run="^$" -bench="." -benchmem -count=10 ./benchmarks`), measured on Windows.

<img width="1088" height="857" alt="image" src="https://github.com/user-attachments/assets/05404f94-495f-43ca-a308-2f6edcff431a" />

Full raw data (all 10 runs per benchmark) and the theoretical-complexity breakdown are tracked in `stlgo_benchmarks.xlsx`.

## Testing & Benchmarks

Run the complete test suite:

```bash
go test -v ./...
```

Run tests with data race detection:

```bash
go test -v -race ./...
```

Run benchmarks:

```bash
go test -bench=. -benchmem ./...
```

## Roadmap

### Current

- [x] Generic `Stack`
- [x] Generic `Queue` (sliding window with amortized compaction)
- [x] Generic `Deque` (circular ring buffer)
- [x] Complete `algo/search` (`LowerBound`, `UpperBound`, `BinarySearch`, `Find`, `FindIf`)
- [x] Complete `algo/numeric` (`Sum`, `Product`, `Iota`, `MaxElement`, `MinElement`, `PrefixSum`)
- [x] Complete `algo/sort` (`Sort`, `DescSort`, `SortBy`, `IsSorted`, `NthElement`)
- [x] Complete `algo/math` (`GCD`, `LCM`, `IsPrime`, `Clamp`, etc.)
- [x] `NextPermutation` / `PrevPermutation`
- [x] Unit test suites with edge case coverage
- [x] Benchmark suite with tracked complexity guarantees

### In Progress / Planned

- [ ] Generic `PriorityQueue` (Binary Heap with $O(1)$ peek, $O(\log N)$ push/pop)
- [ ] Generic `Set` (Hash Set backed by `map[T]struct{}`)
- [ ] Disjoint Set Union (DSU with path compression and union by rank)
- [ ] Fenwick Tree (Binary Indexed Tree) and Segment Tree
- [ ] Formal benchmark suite against standard library `container/heap`

## Requirements

- **Go 1.23+**

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

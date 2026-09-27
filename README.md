[![Sourcegraph](https://sourcegraph.com/github.com/stlgo/-/badge.svg)](https://sourcegraph.com/github.com/mohitjoshi-hey/stlgo?badge)

<div align="center">
  <img src="/assets/logo.png" width="75%"/>
</div>

# stlgo

[![FFmpeg Health](https://oss-health-monitor.vercel.app/api/badge/mohitjoshi-hey/stlgo?v=2)](https://github.com/mohitjoshi-hey/stlgo)

High-performance, generic data structures and algorithms for Go 1.26+.

`stlgo` is built for competitive programming, low-latency applications, and
high-throughput systems engineering. It provides modern, type-safe data
structures and algorithms designed around Go generics, contiguous memory
layouts, CPU cache locality, and predictable allocation behavior.

---

## Table of Contents

- [Design Philosophy](#design-philosophy)
- [Installation](#installation)
- [Package Overview](#package-overview)
- [Containers](#containers)
  - [Stack](#stack)
  - [Queue](#queue)
  - [Deque](#deque)
- [Algorithms](#algorithms)
  - [Search](#search-algorithms)
  - [Sorting](#sorting-algorithms)
  - [Numeric](#numeric-algorithms)
  - [Math](#mathematical-algorithms)
  - [Permutations](#permutation-algorithms)
- [Testing](#testing)
- [Fuzz Testing](#fuzz-testing)
- [Benchmarks](#benchmarks)
- [Published Benchmark Results](#published-benchmark-results)
- [Troubleshooting](#troubleshooting)
- [Requirements](#requirements)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Credits](#credits)

---

## Design Philosophy

- **Contiguous Memory & Cache Locality** — Algorithms operate directly on
  native Go slices, while containers prioritize slice-backed storage over
  pointer-heavy node graphs.
- **Low Allocation** — Hot-path operations are designed to avoid unnecessary
  allocations. Operations that require container growth may allocate, while
  steady-state operations can achieve zero allocations.
- **GC-Friendly Containers** — Container pop/dequeue operations explicitly
  zero retired slots (`var zero T`) so references held by removed elements do
  not unnecessarily remain live.
- **Bring Your Own Concurrency (BYOC)** — Containers do not include internal
  mutexes by default. This avoids synchronization overhead when concurrency is
  not required.
- **Predictable Complexity** — Operations provide explicit time-complexity
  guarantees inspired by C++ STL-style interfaces.
- **Generic & Type-Safe** — APIs use Go generics instead of `interface{}` /
  `any`-based APIs wherever practical.

---

## Installation

```bash
go get github.com/mohitjoshi-hey/stlgo
```

Then import the packages you need:

```go
import (
    "github.com/mohitjoshi-hey/stlgo/algo"
    "github.com/mohitjoshi-hey/stlgo/container"
)
```

---

## Package Overview

### Algorithms

| Package             | Functions                                                                |
| -------------------- | ------------------------------------------------------------------------ |
| `algo/search`        | `LowerBound`, `UpperBound`, `BinarySearch`, `Find`, `FindIf`             |
| `algo/sort`          | `Sort`, `DescSort`, `SortBy`, `IsSorted`, `NthElement`                   |
| `algo/numeric`       | `Sum`, `Product`, `Iota`, `MaxElement`, `MinElement`, `PrefixSum`        |
| `algo/math`          | `GCD`, `LCM`, `IsPrime`, `IsEven`, `IsOdd`, `Max`, `Min`, `Clamp`, `Abs` |
| `algo/permutations`  | `NextPermutation`, `PrevPermutation`                                     |

### Containers

| Container | Description                                    |
| --------- | ----------------------------------------------- |
| `Stack`   | Slice-backed LIFO container                     |
| `Queue`   | FIFO container with sliding-window buffering    |
| `Deque`   | Double-ended queue backed by a circular buffer  |

---

## Containers

### Stack

`Stack` provides LIFO (Last-In, First-Out) behavior.

**Usage**

```go
package main

import (
	"fmt"

	"github.com/mohitjoshi-hey/stlgo/container"
)

func main() {
	stack := container.NewStack[int]()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println(stack.Len())
	// 3

	value, ok := stack.Pop()
	fmt.Println(value, ok)
	// 30 true

	value, ok = stack.Peek()
	fmt.Println(value, ok)
	// 20 true
}
```

---

### Queue

`Queue` provides FIFO (First-In, First-Out) behavior.

**Usage**

```go
package main

import (
	"fmt"

	"github.com/mohitjoshi-hey/stlgo/container"
)

func main() {
	queue := container.NewQueue[int]()

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	fmt.Println(queue.Len())
	// 3

	value, ok := queue.Dequeue()
	fmt.Println(value, ok)
	// 10 true

	value, ok = queue.Front()
	fmt.Println(value, ok)
	// 20 true
}
```

---

### Deque

`Deque` is a double-ended queue that supports insertion and removal from both
ends.

**Usage**

```go
package main

import (
	"fmt"

	"github.com/mohitjoshi-hey/stlgo/container"
)

func main() {
	deque := container.NewDeque[int]()

	deque.PushBack(20)
	deque.PushFront(10)
	deque.PushBack(30)

	fmt.Println(deque.Len())
	// 3

	value, ok := deque.Front()
	fmt.Println(value, ok)
	// 10 true

	value, ok = deque.Back()
	fmt.Println(value, ok)
	// 30 true

	value, ok = deque.PopFront()
	fmt.Println(value, ok)
	// 10 true

	value, ok = deque.PopBack()
	fmt.Println(value, ok)
	// 30 true
}
```

---

## Algorithms

### Search Algorithms

Package: `github.com/mohitjoshi-hey/stlgo/algo/search`
(also exposed at the top-level `algo` package)

**LowerBound** — returns the first position where an element is greater than
or equal to the target.

```go
values := []int{10, 20, 20, 30, 40}
index := algo.LowerBound(values, 20)
fmt.Println(index)
// 1
```

**UpperBound** — returns the first position where an element is greater than
the target.

```go
values := []int{10, 20, 20, 30, 40}
index := algo.UpperBound(values, 20)
fmt.Println(index)
// 3
```

**BinarySearch** — searches for an element in a sorted slice.

```go
values := []int{10, 20, 30, 40}
index, found := algo.BinarySearch(values, 30)
fmt.Println(index, found)
// 2 true
```

**Find** — performs a linear search.

```go
values := []string{"apple", "banana", "cherry"}
index, found := algo.Find(values, "banana")
fmt.Println(index, found)
// 1 true
```

**FindIf** — searches using a custom predicate.

```go
values := []string{"apple", "banana", "cherry"}
index, found := algo.FindIf(values, func(value string) bool {
	return len(value) > 5
})
fmt.Println(index, found)
// 1 true
```

---

### Sorting Algorithms

Package: `github.com/mohitjoshi-hey/stlgo/algo`

**Sort** — sorts a slice in ascending order.

```go
values := []int{50, 10, 40, 20, 30}
algo.Sort(values)
fmt.Println(values)
// [10 20 30 40 50]
```

**DescSort** — sorts a slice in descending order.

```go
values := []int{50, 10, 40, 20, 30}
algo.DescSort(values)
fmt.Println(values)
// [50 40 30 20 10]
```

**SortBy** — sorts using a custom comparison function.

```go
type Person struct {
	Name string
	Age  int
}

people := []Person{
	{Name: "Alice", Age: 30},
	{Name: "Bob", Age: 20},
	{Name: "Charlie", Age: 25},
}

algo.SortBy(people, func(a, b Person) bool {
	return a.Age < b.Age
})
```

**IsSorted** — checks whether a slice is already sorted.

```go
values := []int{10, 20, 30, 40}
fmt.Println(algo.IsSorted(values))
// true
```

**NthElement** — rearranges a slice so that the element at a specified
position is the same element that would appear there after sorting, without
fully sorting the slice.

```go
values := []int{9, 2, 7, 4, 1, 8, 3}
algo.NthElement(values, 3)
fmt.Println(values)
```

After the operation, `values[3]` contains the element that belongs at index
`3` in the sorted ordering.

> `NthElement` is useful when you only need an order statistic such as a
> median, percentile, or kth-smallest element and do not need the entire
> slice sorted.

---

### Numeric Algorithms

Package: `github.com/mohitjoshi-hey/stlgo/algo`

**Sum**

```go
values := []int{1, 2, 3, 4, 5}
fmt.Println(algo.Sum(values))
// 15
```

**Product**

```go
values := []int{1, 2, 3, 4, 5}
fmt.Println(algo.Product(values))
// 120
```

**PrefixSum**

```go
values := []int{1, 2, 3, 4, 5}
fmt.Println(algo.PrefixSum(values))
// [1 3 6 10 15]
```

**Iota**

```go
values := algo.Iota(10, 5)
fmt.Println(values)
// [10 11 12 13 14]
```

**MaxElement**

```go
values := []int{10, 50, 20, 40, 30}
value, found := algo.MaxElement(values)
fmt.Println(value, found)
// 50 true
```

**MinElement**

```go
values := []int{10, 50, 20, 40, 30}
value, found := algo.MinElement(values)
fmt.Println(value, found)
// 10 true
```

---

### Mathematical Algorithms

Package: `github.com/mohitjoshi-hey/stlgo/algo`

**GCD**

```go
fmt.Println(algo.GCD(48, 18))
// 6
```

**LCM**

```go
fmt.Println(algo.LCM(12, 18))
// 36
```

**IsPrime**

```go
fmt.Println(algo.IsPrime(97))
// true

fmt.Println(algo.IsPrime(100))
// false
```

**IsEven / IsOdd**

```go
fmt.Println(algo.IsEven(10))
// true

fmt.Println(algo.IsOdd(7))
// true
```

**Max / Min**

```go
fmt.Println(algo.Max(10, 20, 5, 30))
// 30

```

**Clamp**

```go
fmt.Println(algo.Clamp(150, 0, 100))
// 100
```

**Abs**

```go
fmt.Println(algo.Abs(-42))
// 42
```

---

### Permutation Algorithms

Package: `github.com/mohitjoshi-hey/stlgo/algo`

**NextPermutation** — transforms a sequence into its next lexicographical
permutation.

```go
values := []int{1, 2, 3}
ok := algo.NextPermutation(values)
fmt.Println(values, ok)
// [1 3 2] true
```

**PrevPermutation** — transforms a sequence into its previous lexicographical
permutation.

```go
values := []int{3, 2, 1}
ok := algo.PrevPermutation(values)
fmt.Println(values, ok)
// [3 1 2] true
```

---

## Testing

`stlgo` includes unit tests, fuzz tests, and benchmarks. All three are
separate workflows.

### Unit Tests

Run the complete unit-test suite:

```powershell
go test ./...
```

Run with verbose output:

```powershell
go test -v ./...
```

Run tests for a specific package:

```powershell
go test ./container/...
```

Run a specific test:

```powershell
go test ./container -run TestStack
```

Run the complete test suite with the race detector:

```powershell
go test -race ./...
```

Verbose race testing:

```powershell
go test -v -race ./...
```

---

## Fuzz Testing

Go's built-in fuzzing framework is used for fuzz tests.

List available fuzz tests:

```powershell
go test ./... -list "Fuzz"
```

Run a specific fuzz test:

```powershell
go test ./path/to/package -run=^$ -fuzz=FuzzName
```

For example:

```powershell
go test ./container -run=^$ -fuzz=FuzzStack -fuzztime=30s
```

---

## Benchmarks

The benchmark suite is located in:

```text
benchmarks/
```
Run benchmarks against the `benchmarks` package explicitly:

```powershell
go test -run="^$" -bench="." -benchmem ./benchmarks
```

### Run Benchmarks From Inside `benchmarks/`

```powershell
cd benchmarks
go test -run="^$" -bench="." -benchmem .
```

### Run a Specific Benchmark Group

| Group       | Command                                                        |
| ----------- | --------------------------------------------------------------- |
| Deque       | `go test -bench="BenchmarkDeque" -benchmem ./benchmarks`         |
| Queue       | `go test -bench="BenchmarkQueue" -benchmem ./benchmarks`         |
| Stack       | `go test -bench="BenchmarkStack" -benchmem ./benchmarks`         |
| Search      | `go test -bench="BenchmarkSearch" -benchmem ./benchmarks`        |
| Numeric     | `go test -bench="BenchmarkNumeric" -benchmem ./benchmarks`       |
| Math        | `go test -bench="BenchmarkMath" -benchmem ./benchmarks`          |
| Permutation | `go test -bench="BenchmarkPermutation" -benchmem ./benchmarks`   |
| Sort        | `go test -bench="BenchmarkSort" -benchmem ./benchmarks`          |



## Published Benchmark Results

The published benchmark results are based on 5 runs.
<img width="1207" height="761" alt="image" src="https://github.com/user-attachments/assets/de82cbb3-69c9-49d5-8395-ee118391823f" />



## Troubleshooting

### `no Go files in ...\stlgo`

If you see:

```text
no Go files in C:\...\stlgo
```

while trying to run benchmarks, make sure you are explicitly targeting the
benchmark package:

```powershell
go test -run="^$" -bench="." -benchmem ./benchmarks
```

Do **not** use the repository root as the benchmark package:

```powershell
go test -bench="." .
```

unless the repository root itself contains Go source files belonging to a Go
package.

This error is Go's standard message for "zero package arguments were
received," and it usually means the pattern (`./benchmarks`) was dropped or
corrupted somewhere in the command — for example by a copy/paste that
introduced a smart quote or invisible character. Retype the command instead
of pasting it if it recurs.

### Running Benchmarks From the `benchmarks` Directory

If your terminal is already inside `stlgo\benchmarks`, use:

```powershell
go test -run="^$" -bench="." -benchmem .
```

not:

```powershell
go test -run="^$" -bench="." -benchmem ./benchmarks
```

because `./benchmarks` from inside the `benchmarks` directory would refer to
a nested directory that normally does not exist.


## Requirements

- **Go 1.26+**

---

## Roadmap

### Current

- [x] Generic `Stack`
- [x] Generic `Queue`
- [x] Generic `Deque`
- [x] Complete `algo/search`
- [x] Complete `algo/numeric`
- [x] Complete `algo/sort`
- [x] Complete `algo/math`
- [x] `NextPermutation` / `PrevPermutation`
- [x] Unit test suites with edge-case coverage
- [x] Fuzz testing
- [x] Dedicated benchmark suite
- [x] Benchmark result tracking
- [x] Complexity documentation

### In Progress / Planned

- [ ] Generic `PriorityQueue` (Binary Heap)
- [ ] Generic `Set` (`map[T]struct{}` backed)
- [ ] Disjoint Set Union (DSU)
- [ ] Fenwick Tree (Binary Indexed Tree)
- [ ] Segment Tree
- [ ] Formal benchmark suite against standard library `container/heap`
- [ ] Additional cache-conscious data structures
- [ ] Additional algorithms inspired by the C++ STL

---

## Contributing

Contributions are welcome. Before opening a pull request, ensure that:

```powershell
go test ./...
```

passes successfully.

For concurrency-sensitive changes:

```powershell
go test -race ./...
```

For fuzz-sensitive code:

```powershell
go test ./path/to/package -run=^$ -fuzz=FuzzName -fuzztime=30s
```

For performance-sensitive changes:
```powershell
go test -run="^$" -bench="." -benchmem -count=10 ./benchmarks
```

When submitting performance-related changes, include benchmark results when
possible.

---

## License

This project is licensed under the MIT License. See the
[LICENSE](LICENSE) file for details.

---

## Credits

The image used in this project was sourced from
[MariaLetta/free-gophers-pack](https://github.com/MariaLetta/free-gophers-pack).

📷 Image by [MariaLetta](https://github.com/MariaLetta), used under the
[Creative Commons (CC0-1.0)](https://github.com/MariaLetta/free-gophers-pack?tab=CC0-1.0-1-ov-file)
license.

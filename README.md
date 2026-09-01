# stlgo

High-performance, generic data structures and algorithms for Go 1.23+.

`stlgo` is built for competitive programming and high-throughput systems engineering. It provides modern, type-safe data structures designed around Go generics, CPU cache locality, and Go's garbage collection model.

## Design Philosophy

* **Cache-Friendly Memory Layouts:** Modern CPUs benefit significantly from predictable memory access patterns. `stlgo` prioritizes contiguous slice-backed memory layouts over pointer-heavy data structures where practical, improving cache locality and reducing unnecessary cache misses.

* **Allocation-Conscious Operations:** Core operations are designed to minimize heap allocations. Where capacity is already available, operations such as `Push`, `Pop`, `Enqueue`, and `Dequeue` can execute without additional allocations.

* **Bring Your Own Concurrency (BYOC):** `stlgo` containers do not include internal locking by default. This avoids synchronization overhead in single-threaded and performance-critical workloads while allowing applications to implement concurrency strategies appropriate for their use case.

* **Memory Safe:** Removed elements are explicitly zeroed when necessary, preventing stale references from unnecessarily keeping objects alive in Go's garbage collector.

* **Consistent Generic API:** Containers use Go's native generics for compile-time type safety. Variadic constructors support convenient type inference when initial values are provided, while APIs follow consistent and predictable naming conventions.

* **Predictable Complexity:** APIs are designed with clear algorithmic guarantees and avoid hiding expensive operations behind seemingly simple method calls.

## Installation

```bash
go get github.com/mohitjoshi-hey/stlgo
```

## Supported Data Structures

### Stack

A fast, generic LIFO stack backed by contiguous slice storage.

### Queue

A generic FIFO queue backed by a slice-based sliding window with automated amortized memory compaction.

## CI/CD & Testing

`stlgo` is tested using Go's native testing tools.

The test suite verifies:

* Container construction
* Core operations
* FIFO and LIFO behavior
* Empty container behavior
* Generic type support
* Memory compaction behavior
* Internal state consistency

Run the complete test suite:

```bash
go test -v ./...
```

Run tests with the race detector:

```bash
go test -v -race ./...
```

## Project Goals

`stlgo` aims to provide:

* Modern generic data structures
* High-performance implementations
* Cache-conscious memory layouts
* Minimal allocation overhead
* Predictable algorithmic complexity
* Idiomatic Go APIs
* Zero external dependencies
* Comprehensive testing and benchmarks

## Roadmap

### Current

* [x] Generic Stack
* [x] Generic Queue
* [x] Unit tests

### Planned

* [ ] Deque
* [ ] Set
* [ ] Priority Queue
* [ ] BitSet
* [ ] TreeSet
* [ ] TreeMap
* [ ] Segment Tree
* [ ] Fenwick Tree
* [ ] Disjoint Set Union
* [ ] Search algorithms
* [ ] STL-inspired algorithms
* [ ] Benchmarks
* [ ] GitHub Actions CI

## Requirements

* Go **1.23+**

## License

This project is licensed under the **MIT License**.

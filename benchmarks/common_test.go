package benchmarks

// Package-level sinks to prevent compiler dead-code elimination
var (
	sinkInt  int
	sinkBool bool
	sinkInts []int
)

// Standard problem sizes matching competitive programming constraints
var benchmarkSizes = []int{100, 1_000, 10_000, 100_000}
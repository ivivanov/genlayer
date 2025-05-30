# Fragmentation

## Instructions
```
go test -v
go test -bench .
```

## Dependencies
Go 1.24+
- using Benchmark Loop introduced in go 1.24

## Explanation
Since each char in the input string can be represented as an int I am creating a formula to combine them all. The sum is represented as a binary string which has prepended zeros up to the required length. The multiplication with a prime number in the formula gives additional distribution across the binary representation reducing the probability of collisions.

## Assumptions
Rename example(functions, errors, etc.) to follow idiomatic Go

## TODO
missing fragments - not sure how to define missing fragment 
think of better way for sorting of the fragments map keys - push them in ordered DS
maybe extract assert functions as pkg on root level
add tests for getSortedKeys (good time to use benchmarking to test different sorting strategies)
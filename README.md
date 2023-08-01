# Go illustrated

A humble package aiming to illustrate Go patterns and language features.

## Description

The repo aims at illustrating Go patterns and language feautures with simple working examples. It has multiple modules tied together in a single workspace. The driverpgms folder has separate main modules to illustrate features. The features themselves are in individual packages.
Note that the code snippets may be contrived and not necessarily production ready.

## Getting Started

### Dependencies

go1.19.4 or above

### Executing program

Illustrate sliceworking - this execution will illustrate why it is better to provide a capacity for a slice whenever possible.
```
go test -bench Benchmark_Compare ./sliceworking/ -benchmem -cpu=1
```

Illustrate fanin - a simple example of the fanin concurrency pattern.
```
go run ./driverpgms/fanin/main.go
```

Illustrate fanin - a simple example of the fanout concurrency pattern.
```
go run ./driverpgms/fanout/main.go
```

## Authors

Contributors names and contact info

https://github.com/abhisheknaths

## Acknowledgments

Inspiration, code snippets, etc.
* [fanin-fanout](https://www.youtube.com/watch?v=x6vBvgKGvxU)


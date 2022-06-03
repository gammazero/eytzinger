# eytzinger

![GoDoc](https://pkg.go.dev/badge/github.com/gammazero/eytzinger)](https://pkg.go.dev/github.com/gammazero/eytzinger)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Eytzinger binary search, minimalistic implementation.

## Installation

```
$ go get github.com/gammazero/eytzinger
```

## Example
```go
a := make([]int, 100)
for i := 0; i < len(a); i++ {
    a[i] = i
}

// Sort slice into Eytzinger order.
eytzinger.Sort(a)

// Find some numbers.
for _, find := range []int{13, 17, 19, 23, 29, 37, 73} {
    index := eytzinger.Search(a, find)
    fmt.Println(find, "is at index", index)
}
```

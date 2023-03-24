# goutils
## Collection of implementations of useful functions for GoLang v.19+

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/egorgasay/gotils)
![GitHub issues](https://img.shields.io/github/issues/egorgasay/gotils)
![License](https://img.shields.io/badge/license-MIT-green)
# gotils

## Collection of implementations of useful functions for the Go language

## Quick start
```
go get github.com/egorgasay/gotils
```

# Reverse
Flips any array.
```go
arr := []int{1, 2, 3, 4, 5, 6}
gotils.Reverse(arr)
fmt.Println(arr) // [6 5 4 3 2 1]

arr := []any{"1", 56, struct{}{}}
gotils.Reverse(arr)
fmt.Println(arr) // [{} 56 1]
```

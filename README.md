# go-evmop

Package **evmop** provides tools for turning Ethereum Virtual Machine (EVM) OpCodes into bytecodes, for the Go programming language.

This might be useful to someone writing an AOT or JIT compiler targeting the Ethereum Virtual Machine (EVM).


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-evmop

[![GoDoc](https://godoc.org/github.com/reiver/go-evmop?status.svg)](https://godoc.org/github.com/reiver/go-evmop)


## Example

```go
var writer io.Writer

// ...

n, err = evmop.Push1{0x60}.WriteTo(writer)

// ...

n, err = evmop.Push1{0x40}.WriteTo(writer)


// ...

n, err = evmop.MStore{}.WriteTo(writer)

```

# GoMem

![Continuous Integration](https://github.com/jamesmoriarty/gomem/workflows/Continuous%20Integration/badge.svg?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/jamesmoriarty/gomem)](https://goreportcard.com/report/github.com/jamesmoriarty/gomem)

A Go library to manipulate Windows processes.

```go
import "github.com/jamesmoriarty/gomem"
...
process, err := gomem.GetFromProcessName(name)
process.Open()
process.Read(offsetPtr, bufferPtr, unsafe.Sizeof(value))
process.Write(offsetPtr, bufferPtr, unsafe.Sizeof(value))
```

## Build

```
go build
```

## Test

```
go test
```

## Docs

[godoc.org](https://godoc.org/github.com/jamesmoriarty/gomem)

## Examples

[gohack](https://github.com/jamesmoriarty/gohack)

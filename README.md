# GoMem

![Continuous Integration](https://github.com/jamesmoriarty/gomem/workflows/Continuous%20Integration/badge.svg?branch=master) ![Latest Tag](https://img.shields.io/github/v/tag/jamesmoriarty/gomem.svg?logo=github&label=latest) [![Go Report Card](https://goreportcard.com/badge/github.com/jamesmoriarty/gomem)](https://goreportcard.com/report/github.com/jamesmoriarty/gomem)

A Go package for manipulating Windows processes.

```go
import "github.com/jamesmoriarty/gomem"

# Open process with handle.
process, err  := gomem.GetOpenProcessFromName("example.exe")

# Read from process memory.
valuePtr, err := process.ReadUInt32(offsetPtr)

# Write to process memory.
process.WriteByte(valuePtr, value)
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

[package](https://pkg.go.dev/github.com/jamesmoriarty/gomem)

## Examples

[gohack](https://github.com/jamesmoriarty/gohack)

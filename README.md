# GoMem

![Continuous Integration](https://github.com/jamesmoriarty/gomem/workflows/Continuous%20Integration/badge.svg?branch=master)

A Go library to manipulate Windows processes.

```go
import "github.com/jamesmoriarty/gomem"
...
process, err := gomem.GetFromProcessName(name)
process.Open()
process.Read(valuePtr, bufferPtr, unsafe.Sizeof(value))
process.Write(valuePtr, bufferPtr, unsafe.Sizeof(value))
```

## Build

```
go build
```

## Test

```
go test
```
# GoMem

![Build Status](https://img.shields.io/github/workflow/status/jamesmoriarty/gomem/Continuous%20Integration)

A Go library to manipulate Windows processes.

```go
import "github.com/jamesmoriarty/gomem"

...

process, err := gomem.GetFromProcessName(name)
process.Open()
process.Read(offset, len(bytes))
process.Write(offset, bytes)
```

## Build

```
go build
```

## Test

```
go test -v -coverprofile cover.txt
```
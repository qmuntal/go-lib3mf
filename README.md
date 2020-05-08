# go-lib3mf
Golang wrapper for the [lib3mf](https://github.com/3MFConsortium/lib3mf) library

## How it works

This library uses the [Autodesk/AutomaticComponentToolkit](https://github.com/Autodesk/AutomaticComponentToolkit) and the [lib3mf IDL](https://github.com/3MFConsortium/lib3mf/blob/master/AutomaticComponentToolkit/lib3mf.xml) to automatically generate Go wrappers. The added value is that the resulting wrappers are published as Go module that can be imported without having to manually create and mantain the wrappers.

## Usage

A precompiled version of lib3mf is distributed together with the Go wrapper and the appropiate `cgo` flags are set up, so you only need to have a `cgo` toolchain to use this library as an ordinary go package.

When distributing an executable that uses `go-lib3mf` and dynamic linking make sure to distribute to apropiate `lib3mf` shared library and that it is accessible by the executable.

### Example

```go
package main

import (
  "github.com/qmuntal/go-lib3mf/v2"
)

func main() {
  flag.Parse()
  model, _ := lib3mf.CreateModel()
  reader, _ := model.QueryReader("3mf")
  reader.ReadFromFile(flag.Arg(0))
  box_, _ := model.GetOutbox()
  fmt.Println(box)
}

```
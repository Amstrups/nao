package main

import (
	"github.com/amstrups/nao/run"
	"github.com/davecgh/go-spew/spew"
)

func main() {
  spew.Config.DisablePointerAddresses = true
  run.RunFile("./examples/main.nao")
}

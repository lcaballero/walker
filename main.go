package main

import (
	"github.com/lcaballero/walker/gather"
)

func main() {
	gather.NewWalker().Walk(".files/inbound")
}

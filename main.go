package main

import (
	"fmt"
	"github.com/lcaballero/walker/start"
	"time"
)

func main() {
	start.Start()
}

func showTime() {
	nanos := time.Now().UnixNano()
	v := time.Duration(nanos) / time.Second
	fmt.Printf("%d\n", v)
}

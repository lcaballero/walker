package indexing

import (
	"github.com/lcaballero/walker/conf"
	"github.com/lcaballero/walker/gather"
	"fmt"
)

func RunIndexing(vals conf.ValueContext) {
	index := conf.LoadIndexing(vals)
	fmt.Println(index.ToJson())
	gather.NewWalker().Walk(".files/src/lucene-6.0.1")
}

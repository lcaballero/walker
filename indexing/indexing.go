package indexing

import (
	"fmt"
	"github.com/lcaballero/walker/conf"
	"github.com/lcaballero/walker/gather"
)

func RunIndexing(vals conf.ValueContext) {
	index := conf.LoadIndexing(vals)
	fmt.Println(index.ToJson())
	gather.NewWalker().Walk(".files/src/lucene-6.0.1")
}

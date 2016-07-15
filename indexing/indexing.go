package indexing

import (
	"github.com/lcaballero/walker/conf"
	"github.com/lcaballero/walker/gather"
)

func Indexing(conf *conf.Config) {
	gather.NewWalker().Walk(".files/inbound")
}

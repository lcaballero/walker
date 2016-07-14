package indexing

import (
	"github.com/lcaballero/walker/gather"
	"github.com/lcaballero/walker/conf"
)


func Indexing(conf *conf.Config) {
	gather.NewWalker().Walk(".files/inbound")
}
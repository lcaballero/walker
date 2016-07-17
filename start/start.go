package start

import (
	"fmt"
	"github.com/lcaballero/walker/cli"
	"github.com/lcaballero/walker/indexing"
	"github.com/lcaballero/walker/searching"
	"os"
)

func Start() {
	conf := cli.ParseArgs(os.Args...)
	switch conf.Command {
	case "indexing":
		fmt.Println("indexing...")
		indexing.Indexing(conf)
	case "searching":
		s, err := searching.Search(conf)
		if err != nil {
			panic(err)
		}
		s.Start()
		fmt.Println("searching...")
	default:
		fmt.Println("default indexing")
		indexing.Indexing(conf)
	}
}

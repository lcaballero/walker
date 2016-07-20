package start

import (
	"fmt"
	"github.com/lcaballero/walker/cli"
	"github.com/lcaballero/walker/indexing"
	"github.com/lcaballero/walker/searching"
	web "github.com/lcaballero/walker/web/start"
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
		if conf.HasQuery() {
			s.Query(os.Stdout, conf.Query)
		} else {
			s.Start()
			fmt.Println("searching...")
		}
	case "web":
		s, err := searching.Search(conf)
		if err != nil {
			panic(err)
		}
		web.Start(s, conf)

	default:
		fmt.Println("default indexing")
		indexing.Indexing(conf)
	}
}

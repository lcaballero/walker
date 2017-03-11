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
		indexing.Indexing(conf)

	case "searching":
		s, err := searching.Search(conf)
		if err != nil {
			panic(err)
		}
		if conf.HasQuery() {
			qr, err := s.Query(*conf)
			if err != nil {
				fmt.Println(err)
			}
			sr := qr.Render()
			fmt.Println(sr.Query)
			fmt.Println(sr.FilesAndHits)
			fmt.Println(sr.Summary)
		} else {
			s.Start()
			fmt.Println("searching...")
		}

	default:
		fmt.Println("default indexing")
		indexing.Indexing(conf)
	}
}

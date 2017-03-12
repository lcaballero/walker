package start

import (
	"fmt"
	"github.com/lcaballero/walker/cli"
	cmd "gopkg.in/urfave/cli.v2"
	"os"
	"github.com/lcaballero/walker/indexing"
	"github.com/lcaballero/walker/searching"
)

func Start() {
	proc := cli.Processing{
		IndexingAction: func(ctx *cmd.Context) error {
			fmt.Println("indexing")
			indexing.RunIndexing(ctx)
			return nil
		},
		SearchingAction: func(ctx *cmd.Context) error {
			s, err := searching.Search(ctx)
			if err != nil {
				return err
			}

			if s.HasQuery() {
				qr, err := s.Query()
				if err != nil {
					return err
				}
				sr := qr.Render()
				fmt.Println(sr.Query)
				fmt.Println(sr.FilesAndHits)
				fmt.Println(sr.Summary)
			}
			return nil
		},
		InteractiveAction: func(ctx *cmd.Context) error {
			fmt.Println("interactive")
			return nil
		},
	}
	cli.New(proc).Run(os.Args)
}

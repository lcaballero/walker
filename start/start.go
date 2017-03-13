package start

import (
	"fmt"
	"github.com/lcaballero/walker/cli"
	"github.com/lcaballero/walker/indexing"
	"github.com/lcaballero/walker/interactive"
	"github.com/lcaballero/walker/searching"
	cmd "gopkg.in/urfave/cli.v2"
	"os"
)

func Start() {
	proc := cli.Processing{
		IndexingAction: func(ctx *cmd.Context) error {
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
			interactive.Run(ctx)
			return nil
		},
	}
	err := cli.New(proc).Run(os.Args)
	if err != nil {
		panic(err)
	}
}

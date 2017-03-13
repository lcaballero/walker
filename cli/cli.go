package cli

import (
	"gopkg.in/urfave/cli.v2"
)

type Processing struct {
	IndexingAction    cli.ActionFunc
	SearchingAction   cli.ActionFunc
	InteractiveAction cli.ActionFunc
}

func New(proc Processing) *cli.App {
	app := &cli.App{
		Name:    "walker",
		Version: "0.0.1",
		Usage:   "A CLI to interface with github from inside a repository directory",
		Commands: []*cli.Command{
			&cli.Command{
				Name:   "indexing",
				Usage:  "Carry out indexing process and save results into file.",
				Action: proc.IndexingAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "out-file",
						Usage: "The name of the index file to write.",
						Value: "out.json",
					},
				},
			},
			&cli.Command{
				Name:   "searching",
				Usage:  "Load index file and carry out search query.",
				Action: proc.SearchingAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "query",
						Usage: "Executes the query without entering the repl",
					},
					&cli.StringFlag{
						Name:  "filename",
						Usage: "Load index file and start interactive query repl.",
					},
					&cli.IntFlag{
						Name:  "max-hits",
						Usage: "Limits the output to a maximum of hits (absolute max: 1000)",
						Value: 20,
					},
				},
			},
			&cli.Command{
				Name:   "interactive",
				Usage:  "Load index file and start interactive query repl.",
				Action: proc.InteractiveAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "filename",
						Usage: "The name of the index file to read.",
						Value: "out.json",
					},
					&cli.IntFlag{
						Name:  "max-hits",
						Usage: "Limits the output to a maximum of hits (absolute max: 1000)",
						Value: 20,
					},
				},
			},
		},
	}

	return app
}

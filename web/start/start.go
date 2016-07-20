package start

import (
	"github.com/lcaballero/hitman"
	"github.com/lcaballero/walker/conf"
	"github.com/lcaballero/walker/searching"
	"github.com/lcaballero/walker/web/server"
	"github.com/vrecan/death"
	"syscall"
)

func Start(searcher *searching.Searcher, conf *conf.Config) {
	targets := hitman.NewTargets()
	targets.AddOrPanic(server.NewWebServer(searcher, conf))

	death.NewDeath(syscall.SIGTERM, syscall.SIGINT).WaitForDeath(targets)
}

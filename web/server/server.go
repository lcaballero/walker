package server

import (
	"github.com/labstack/echo/engine/standard"
	"github.com/lcaballero/gel"
	"github.com/lcaballero/hitman"
	app "github.com/lcaballero/walker/web/context"
	"github.com/lcaballero/reverb/base"
	"github.com/lcaballero/walker/conf"
	"github.com/lcaballero/walker/searching"
	"log"
	"github.com/labstack/echo"
)

type WebServer struct {
	*conf.Config
	searcher         *searching.Searcher
	IncludesResolver func(string) gel.View
}

func NewWebServer(searcher *searching.Searcher, config *conf.Config) (*WebServer, error) {
	ws := &WebServer{
		Config:   config,
		searcher: searcher,
	}
	return ws, nil
}

func (w *WebServer) Start() hitman.KillChannel {
	done := hitman.NewKillChannel()
	go func() {
		go w.run()
		for {
			select {
			case cleaner := <-done:
				cleaner.WaitGroup.Done()
				return
			}
		}
	}()
	return done
}

func (w *WebServer) run() {
	log.Printf("finding templates at: %s", w.IncludesDir)
	log.Printf("using assets found at: %s", w.AssetsDir)

	ctx := app.NewContext(w.Config)

	r := base.NewRegister()
	r.Get("/asset/:kind/:hash/:file", ctx.ToHandler())
	r.Get("/searching", Index(ctx, w.searcher))

	log.Printf("starting web server at: %s", w.Ip)
	r.Echo.Run(standard.New(w.Ip))
}

func Index(ctx *app.Context, searcher *searching.Searcher) echo.HandlerFunc {
	return func(c echo.Context) error {
		q := c.QueryParam("q")
		return searcher.Query(c.Response().Writer(), q)
	}
}

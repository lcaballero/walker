package searching

import (
	"bufio"
	"fmt"
	"github.com/lcaballero/walker/conf"
	"github.com/lcaballero/walker/gather"
	"os"
	"regexp"
	"strings"
)

type Searcher struct {
	reduce *gather.Reduction
	prefs  Prefs
	conf   *conf.Config
}

func Search(conf *conf.Config) (*Searcher, error) {
	r, err := LoadReduction(conf.Filename)
	if err != nil {
		return nil, err
	}
	s := &Searcher{
		reduce: r,
		prefs:  DefaultPrefs,
		conf:   conf,
	}
	return s, nil
}

func (s *Searcher) Query(cf conf.Config) (*QueryResult, error) {
	qr := NewQueryResult(cf, *s.reduce.IndexInfo)

	re, err := regexp.Compile(cf.Query)
	if err != nil {
		return nil, err
	}

	for _, unit := range s.reduce.Units {
		qr.Searched++
		matches := re.FindAllIndex(unit.Content, -1)
		if matches == nil {
			continue
		}
		for _, match := range matches {
			qr.HitCount++
			hit := NewHitBounds(match[0], match[1], unit)
			if qr.HitCount <= cf.MaxHits {
				qr.Hits = append(qr.Hits, hit)
				qr.ResultCount++
			}
		}
	}

	qr.Timing.Stop()

	return qr, nil
}

func (s *Searcher) Start() {
	rd := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(s.prefs.Prompt)
		line, err := rd.ReadString('\n')
		if err != nil {
			fmt.Print(err)
			continue
		}

		cfg := *s.conf
		cfg.Query = strings.TrimRight(line, "\n")

		fmt.Printf("%s: %s\n", s.prefs.Echo, cfg)

		res, err := s.Query(cfg)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		os.Stdout.Write(res.Render().Bytes())
	}
}

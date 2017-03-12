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
	searching   *conf.Searching
}

func Search(vals conf.ValueContext) (*Searcher, error) {
	searching := conf.LoadSearching(vals)

	r, err := LoadReduction(searching.Filename)
	if err != nil {
		return nil, err
	}

	s := &Searcher{
		reduce: r,
		prefs:  DefaultPrefs,
		searching: &searching,
	}
	return s, nil
}

func (s *Searcher) HasQuery() bool {
	return s.searching.HasQuery()
}

func (s *Searcher) Query() (*QueryResult, error) {
	qr := NewQueryResult(*s.searching, *s.reduce.IndexInfo)

	re, err := regexp.Compile(s.searching.Query)
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
			if qr.HitCount <= s.searching.MaxHits {
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

		s.searching.Query = strings.TrimRight(line, "\n")

		fmt.Printf("%s: %s\n", s.prefs.Echo, s.searching)

		res, err := s.Query()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		os.Stdout.Write(res.Render().Bytes())
	}
}

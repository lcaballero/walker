package searching

import (
	"bufio"
	"fmt"
	"github.com/lcaballero/walker/conf"
	"github.com/lcaballero/walker/gather"
	"github.schq.secious.com/Logrhythm/GoDispatch/bench"
	"os"
	"regexp"
	"strings"
	"io"
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

func (s *Searcher) Query(out io.Writer, query string) error {
	tc := bench.Start()

	re, err := regexp.Compile(query)
	if err != nil {
		return err
	}

	searched, hitCount := 0, 0

	for _, unit := range s.reduce.Units {
		searched++
		matches := re.FindAllIndex(unit.Content, -1)
		if matches == nil {
			continue
		}
		hitCount++
		for _, match := range matches {
			hit := NewHitBounds(match[0], match[1], unit)
			hf := HitFormatter{
				Hit:    hit,
				Window: s.prefs.Width,
			}
			fmt.Fprintln(out, "in", hf.Hit.Unit.Path)
			fmt.Fprintln(out, hf.String())
		}
	}

	tc.Stop()

	if s.conf.NoStats {
		return nil
	}

	fmt.Fprintln(out)
	fmt.Fprintf(out, "Searched: %d, Hits: %d\n", searched, hitCount)
	fmt.Fprintln(out, tc.String())
	return nil
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

		line = strings.TrimRight(line, "\n")
		fmt.Printf("%s: %s\n", s.prefs.Echo, line)

		err = s.Query(os.Stdout, line)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
	}
}

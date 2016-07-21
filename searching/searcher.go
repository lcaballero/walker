package searching

import (
	"bufio"
	"fmt"
	"github.com/lcaballero/walker/conf"
	"github.com/lcaballero/walker/gather"
	"github.schq.secious.com/Logrhythm/GoDispatch/bench"
	"io"
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

func (s *Searcher) Query(out io.Writer, cf conf.Config) error {
	tc := bench.Start()

	re, err := regexp.Compile(cf.Query)
	if err != nil {
		return err
	}

	if cf.ShowQuery {
		fmt.Fprintf(out, "Search query: %s\n", cf.Query)
		fmt.Fprintln(out)
	}

	searched, hitCount, resultCount := 0, 0, 0

	for _, unit := range s.reduce.Units {
		showPath := true
		searched++
		matches := re.FindAllIndex(unit.Content, -1)
		if matches == nil {
			continue
		}
		for _, match := range matches {
			hitCount++
			hit := NewHitBounds(match[0], match[1], unit)
			hf := HitFormatter{
				Hit:    hit,
				Window: s.prefs.Width,
			}
			if hitCount <= s.conf.MaxHits {
				resultCount++
				if showPath {
					fmt.Fprintln(out, "in", hf.Hit.Unit.Path)
					showPath = false
				}
				fmt.Fprintln(out, hf.String())
			}
		}
	}

	tc.Stop()

	if s.conf.NoStats {
		return nil
	}

	fmt.Fprintln(out)
	fmt.Fprintf(out, "Searched: %d, Hits: %d of %d\n", searched, resultCount, hitCount)
	fmt.Fprintf(out, "Max Hits: %d\n", s.conf.MaxHits)
	fmt.Fprintln(out, tc.String())

	if cf.ShowFontLocks {
		s.outputFonts(out)
	}
	return nil
}

func (s *Searcher) outputFonts(out io.Writer) {
	fmt.Fprintln(out)
	fmt.Fprintln(out, `
warning                font-lock-warning-face
name                   font-lock-function-name-face
variable               font-lock-variable-name-face
keyword                font-lock-keyword-face
comment                font-lock-comment-face
comment_delimiter      font-lock-comment-delimiter-face
type                   font-lock-type-face
constant               font-lock-constant-face
builtin                font-lock-builtin-face
preprocessor           font-lock-preprocessor-face
string                 font-lock-string-face
doc                    font-lock-doc-face
negation-char          font-lock-negation-char-face
	`)
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

		err = s.Query(os.Stdout, cfg)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
	}
}

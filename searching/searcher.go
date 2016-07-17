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
}

func Search(conf *conf.Config) (*Searcher, error) {
	r, err := LoadReduction(conf.Filename)
	if err != nil {
		return nil, err
	}
	s := &Searcher{reduce: r}
	return s, nil
}

func (s *Searcher) Start() {
	rd := bufio.NewReader(os.Stdin)
	prompt := "%> "
	for {
		fmt.Print(prompt)
		line, err := rd.ReadString('\n')
		if err != nil {
			fmt.Print(err)
			continue
		}

		line = strings.TrimRight(line, " \n\r\f")
		fmt.Println("Searching for: ", line)

		re, err := regexp.Compile(line)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		for _, unit := range s.reduce.Units {
			matches := re.FindAllIndex(unit.Content, -1)
			if !unit.HasLines() {
				unit.FindLines()
			}
			if matches == nil {
				continue
			}
			for _, match := range matches {
				hit := NewHitBounds(match[0], match[1], unit)
				fmt.Println(hit)
			}
		}
	}
}

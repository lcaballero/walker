package gather

import (
	"encoding/json"
	"fmt"
	"github.com/lcaballero/time-capture/bench"
	"github.com/lcaballero/walker/shared/checks"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Reduction struct {
	ExtCount           map[string]int
	Units              []*Unit
	FilesRead          int
	PathsWalked        int
	FilePathsCollected int
	SkippedReading     int
	DirsFound          int
	ExtensionsSkipped  []string
	GoRountineCount    int
	CpuCount           int
	ReductionTime      *bench.TimeCapture

	IndexingRoot string
	AbsoluteRoot string
}

func (w *Reduction) Report() {
	fmt.Println("Num CPUs: ", w.CpuCount)
	fmt.Println("Num Goroutines: ", w.GoRountineCount)
	fmt.Println()
	fmt.Println("Skipped Extensions: ", strings.Join(w.ExtensionsSkipped, ", "))
	fmt.Println("Total Paths Found: ", w.PathsWalked)
	fmt.Println("Total Paths Collected: ", w.FilePathsCollected)
	fmt.Println("Total Files Skipped: ", w.SkippedReading)
	fmt.Println("Total Directories Found: ", w.DirsFound)
	fmt.Println()
	fmt.Println("Extensions Collected: ", len(w.ExtCount))
	//	fmt.Println()
	//	for k,v := range w.ExtCount {
	//		fmt.Printf("ext: %s, count: %d\n", k, v)
	//	}
	fmt.Println()

	w.ReductionTime.Out(os.Stdout)
	fmt.Println()
	fmt.Println()
	checks.GoroutinesLeaked()
}

func (w *Reduction) Write() {
	bb, err := json.Marshal(w)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("out.json", bb, 0666)
	if err != nil {
		panic(err)
	}
}

func (w *Reduction) Out(out io.Writer) {
	r := *w
	r.Units = nil

	bb, err := json.MarshalIndent(&r, "", "  ")
	if err != nil {
		fmt.Fprintln(out, err)
		return
	}
	fmt.Fprintln(out, string(bb))
	return
}

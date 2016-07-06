package gather

import (
	"fmt"
	"github.com/lcaballero/time-capture/bench"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var PartiallyAvailableReceiversError = fmt.Errorf("Less than all receivers available.")
var AllReceiversUnavailableError = fmt.Errorf("No recievers available.")

type Walker struct {
	walkedPaths chan *Unit
	loadedPaths chan *Unit
	wait        *sync.WaitGroup
	total       int
}

func NewWalker() *Walker {
	return &Walker{
		loadedPaths: make(chan *Unit, 100),
		wait:        &sync.WaitGroup{},
	}
}

func (w *Walker) Walk(root string) error {
	tc := bench.Start()
	fr := NewFileReader(runtime.NumCPU(), w.wait.Done, w.loadedPaths)
	w.walkedPaths = fr.Start()

	reducer := NewReducer(w.loadedPaths)
	wrapUp := reducer.Start()

	err := filepath.Walk(root, w.Walking)
	if err != nil {
		fmt.Println(err)
	}
	w.wait.Wait()
	wrapUp <- true

	fr.Close()
	tc.Stop()

	reducer.sums.PathsWalked = w.total
	reducer.sums.CpuCount = runtime.NumCPU()
	reducer.sums.GoRountineCount = fr.maxReaders
	reducer.sums.ReductionTime = tc
	reducer.sums.ExtensionsSkipped = []string{".tgz"}
	reducer.sums.Report()

	reducer.sums.Write()

	return nil
}

func (w *Walker) Handle(path string, info os.FileInfo) error {
	w.total++
	w.wait.Add(1)
	w.walkedPaths <- &Unit{
		Path:  path,
		IsDir: info.IsDir(),
		Ext:   filepath.Ext(path),
	}
	return nil
}

func (w *Walker) Walking(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	return w.Handle(path, info)
}

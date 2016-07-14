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

// A Walker traverses the file-system and produces Unit's which composed of
// fields that are common to Files.   These unit can be filtered out of the
// final result.  Unfilteerd Units are sent to a reducer that carries out
// further processing.
type Walker struct {
	walkedPaths chan *Unit
	loadedPaths chan *Unit
	wait        *sync.WaitGroup
	total       int
}

// NewWalker allocates a Walker instance with a channel to file with
// encountered files. A shared instance of
func NewWalker() *Walker {
	return &Walker{
		loadedPaths: make(chan *Unit, 100),
		wait:        &sync.WaitGroup{},
	}
}

// Walk traverses the file-system starting at the given root directory.  It
// returns an error if one is encountered during traversal.  It will collect
// files and then eventually write out a file.
func (w *Walker) Walk(root string) error {
	tc := bench.Start()

	reducer := NewReducer(w.loadedPaths)
	wrapUp := reducer.Start()

	fr := NewFileReader(runtime.NumCPU(), w.wait.Done, w.loadedPaths)
	w.walkedPaths = fr.Start()

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
	reducer.sums.Report() // output to standard out
	reducer.sums.Write()  // make file of the collected information.

	return nil
}

// Handle increments the total count for which it will wait once traversal is
// finished.  Handle also pipes information derived from the FileInfo to
// the walkedPaths.
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

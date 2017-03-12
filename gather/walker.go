package gather

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var PartiallyAvailableReceiversError = fmt.Errorf("Less than all receivers available.")
var AllReceiversUnavailableError = fmt.Errorf("No recievers available.")
var DefaultFilteredExtensions = []string{
	".tgz", ".zip", ".png", ".gif", ".bz2",
	".jpg", ".jpeg", ".bmp", ".tar", ".gz",
}

// A Walker traverses the file-system and produces Unit's which composed of
// fields that are common to Files.   These unit can be filtered out of the
// final result.  Unfilteerd Units are sent to a reducer that carries out
// further processing.
type Walker struct {
	loadedPaths chan *Unit
	wait        *sync.WaitGroup
	total       int
	fileReader  *FileReader
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
	reducer := NewReducer(w.loadedPaths)
	reducer.Start()

	w.fileReader = NewFileReader(runtime.NumCPU(), w.wait.Done, w.loadedPaths)
	w.fileReader.filteredExt = DefaultFilteredExtensions
	w.fileReader.Start()

	err := filepath.Walk(root, w.Walking)
	if err != nil {
		fmt.Println(err)
	}
	w.wait.Wait()
	reducer.Close()

	w.fileReader.Close()

	cwd, err := os.Getwd()
	reducer.sums.AbsoluteRoot = cwd
	reducer.sums.IndexingRoot = root
	reducer.sums.PathsWalked = w.total
	reducer.sums.CpuCount = runtime.NumCPU()
	reducer.sums.GoRountineCount = w.fileReader.maxReaders
	reducer.sums.ExtensionsSkipped = w.fileReader.filteredExt
	reducer.sums.Report() // output to standard out
	reducer.sums.Out(os.Stdout)
	reducer.sums.Write() // make file of the collected information.

	return err
}

// Handle increments the total count for which it will wait once traversal is
// finished.  Handle also pipes information derived from the FileInfo to
// the walkedPaths.
func (w *Walker) Handle(path string, info os.FileInfo) error {
	w.total++
	w.wait.Add(1)

	w.fileReader.Add(&Unit{
		Path:  path,
		IsDir: info.IsDir(),
		Ext:   filepath.Ext(path),
		Size:  info.Size(),
	})
	return nil
}

func (w *Walker) Walking(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	return w.Handle(path, info)
}

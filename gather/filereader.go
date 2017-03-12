package gather

import (
	"fmt"
	"io/ioutil"
	"sync"
)

type FileReader struct {
	inbound     chan *Unit
	maxReaders  int
	outbound    chan *Unit
	done        func()
	closed      func()
	filteredExt []string
}

func NewFileReader(
	maxReaders int,
	done func(),
	outbound chan *Unit) *FileReader {

	inbound := make(chan *Unit, 100)
	fr := &FileReader{
		inbound:    inbound,
		maxReaders: maxReaders,
		outbound:   outbound,
		done:       done,
	}
	return fr
}

func (fr *FileReader) Add(unit *Unit) {
	fr.inbound <- unit
}

func (fr *FileReader) Start() {
	for i := 0; i < fr.maxReaders; i++ {
		go fr.readFile(fr.inbound, fr.outbound)
	}
}

func (fr *FileReader) filter(unit *Unit) bool {
	for _, ext := range fr.filteredExt {
		if ext == unit.Ext {
			return true
		}
	}
	return false
}

// readFile takes the id (number) of the thread, a inbound channel to receive
// new Units and an outbound chanel to send processed Units on.
func (fr *FileReader) readFile(inbound <-chan *Unit, outbound chan<- *Unit) {

	for unit := range inbound {
		if unit.IsDir {
			outbound <- unit
			fr.done()
			continue
		}

		if fr.filter(unit) {
			unit.IsSkipped = true
			outbound <- unit
			fr.done()
			continue
		}

		bb, err := ioutil.ReadFile(unit.Path)

		if err != nil {
			fmt.Println(err)
			outbound <- unit
			fr.done()
			continue
		}

		unit.Content = bb
		outbound <- unit
		fr.done()
	}
	fr.finished()
}

// finished is called to close down the inbound channel exiting for loop.
func (fr *FileReader) finished() {
	fr.closed()
}

// Close signals closed on all routinest
func (fr *FileReader) Close() error {
	wg := sync.WaitGroup{}
	wg.Add(fr.maxReaders)
	fr.closed = wg.Done
	close(fr.inbound)
	wg.Wait()
	return nil
}

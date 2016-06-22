package gather

import (
	"io/ioutil"
	"fmt"
	"sync"
)

type FileReader struct {
	inbound chan *Unit
	maxReaders int
	outbound chan *Unit
	done func()
	closed func()
}

func NewFileReader(
	maxReaders int,
	done func(),
	outbound chan *Unit) *FileReader {

	fr := &FileReader{
		maxReaders: maxReaders,
		outbound: outbound,
		done: done,
	}
	return fr
}

func (fr *FileReader) Start() chan *Unit {
	fr.inbound = make(chan *Unit, 100)
	for i := 0; i < fr.maxReaders; i++ {
		go fr.readFile(i, fr.inbound, fr.outbound)
	}
	return fr.inbound
}

func (fr *FileReader) filter(unit *Unit) bool {
	switch unit.Ext {
	case ".tgz":
		return true
	default:
		return false
	}
}

func (fr *FileReader) readFile(id int, inbound <-chan *Unit, outbound chan<- *Unit) {
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
	fr.finished(id)
}

func (fr *FileReader) finished(id int) {
	fr.closed()
}

func (fr *FileReader) Close() error {
	wg := sync.WaitGroup{}
	wg.Add(fr.maxReaders)
	fr.closed = wg.Done
	close(fr.inbound)
	wg.Wait()
	return nil
}

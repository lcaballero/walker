package searching

import (
	"encoding/json"
	"fmt"
	"github.com/lcaballero/walker/gather"
	"io/ioutil"
	"os"
)

func LoadReduction(filename string) (*gather.Reduction, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil, err
	}
	if info.IsDir() {
		return nil, fmt.Errorf("Cannot load directory (not an index): %s", filename)
	}
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	filebytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	aggr := &gather.Reduction{}
	err = json.Unmarshal(filebytes, aggr)
	if err != nil {
		return nil, err
	}
	return aggr, err
}

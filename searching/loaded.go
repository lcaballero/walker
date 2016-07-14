package searching

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Loaded struct {
}

func NewLoaded(filename string) (*Loaded, error) {
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
	//	json.Unmarshal(filebytes, )
	fmt.Println(filebytes)
	return nil, err
}

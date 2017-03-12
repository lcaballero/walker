package conf

import (
	"encoding/json"
)

type Indexing struct {
	OutFile string
}

func LoadIndexing(v ValueContext) Indexing {
	c := ContextLoader{v}
	in := Indexing{}
	c.String("out-file", &in.OutFile)
	return in
}

func (c *Indexing) ToJson() string {
	bin, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return ""
	}
	return string(bin)
}

package conf

import (
	"encoding/json"
)

type Interactive struct {
	Filename string
}

func LoadInteractive(v ValueContext) Interactive {
	c := ContextLoader{v}
	in := Interactive{}
	c.String("filename", &in.Filename)
	return in
}

func (c *Interactive) ToJson() string {
	bin, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return ""
	}
	return string(bin)
}

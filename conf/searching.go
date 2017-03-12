package conf

import "encoding/json"

type Searching struct {
	Query         string
	MaxHits       int
	Filename   string
}

func LoadSearching(v ValueContext) Searching {
	c := ContextLoader{v}
	s := Searching{}
	c.String("query", &s.Query)
	c.String("filename", &s.Filename)
	c.Int("max-hits", &s.MaxHits)
	return s
}

func (c *Searching) HasQuery() bool {
	return c.Query != ""
}

func (c *Searching) ToJson() string {
	bin, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return ""
	}
	return string(bin)
}

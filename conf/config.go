package conf

import "encoding/json"

//Root        string `long:"root" description:"When indexing the root directory to start traversal."`
type Config struct {
	Command       string `long:"cmd" description:"'indexing' or 'searching'"`
	Filename      string `long:"file" description:"The name of the index file to read/write." default:"out.json"`
	Query         string `long:"query" description:"Executes the query without entering the repl"`
	MaxHits       int    `long:"max-hits" description:"Limits the output to a maximum of hits (absolute max: 1000)" default:"20"`
}

func (c *Config) HasQuery() bool {
	return c.Query != ""
}
func (c *Config) String() string {
	bin, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return ""
	}
	return string(bin)
}

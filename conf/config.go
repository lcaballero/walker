package conf

import "encoding/json"

//Root        string `long:"root" description:"When indexing the root directory to start traversal."`
type Config struct {
	Command       string `long:"cmd" description:"'indexing' or 'searching'"`
	Filename      string `long:"file" description:"The name of the index file to read/write." default:"out.json"`
	Query         string `long:"query" description:"Executes the query without entering the repl"`
	MaxHits       int    `long:"max-hits" description:"Limits the output to a maximum of hits (absolute max: 1000)" default:"20"`
	NoStats       bool   `long:"no-stats" description:"Skips output of final stats when present on the command line."`
	Ip            string `long:"ip" description:"The ip:host which the server should bind to." default:"127.0.0.1:4000"`
	AssetsDir     string `short:"a" long:"assets" description:"Directory from which to server static assets."`
	IncludesDir   string `long:"includes" description:"Directory from which include templates."`
	ShowFontLocks bool   `long:"show-locks" description:"Show the list of font-locks so that help design syntax highlighting."`
	ShowQuery     bool   `long:"show-query" descriptoin:"Show the query that was submitted at the top of the result."`
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

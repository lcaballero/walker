package conf

type Config struct {
	Command     string `long:"cmd" description:"'indexing' or 'searching'"`
	Filename    string `long:"file" description:"The name of the index file to read/write." default:"out.json"`
	Root        string `long:"root" description:"When indexing the root directory to start traversal." default:".files/inbound"`
	Query       string `long:"query" description:"Executes the query without entering the repl"`
	NoStats     bool   `long:"no-stats" description:"Skips output of final stats when present on the command line."`
	Ip          string `long:"ip" description:"The ip:host which the server should bind to." default:"127.0.0.1:4000"`
	AssetsDir   string `short:"a" long:"assets" description:"Directory from which to server static assets."`
	IncludesDir string `long:"includes" description:"Directory from which include templates."`
}

func (c *Config) HasQuery() bool {
	return c.Query != ""
}

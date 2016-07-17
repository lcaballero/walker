package conf

type Config struct {
	Command  string `long:"cmd" description:"'indexing' or 'searching'"`
	Filename string `long:"file" description:"The name of the index file to read/write." default:"out.json"`
	Root     string `long:"root" description:"When indexing the root directory to start traversal." default:".files/inbound"`
}

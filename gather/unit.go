package gather


type Unit struct {
	Path string
	Ext string
	Content []byte
	IsDir bool
	Error error
	IsSkipped bool
}

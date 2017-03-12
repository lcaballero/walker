package searching

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lcaballero/time-capture/bench"
	"github.com/lcaballero/walker/conf"
	"github.com/lcaballero/walker/gather"
	"io"
)

type QueryResult struct {
	Searched    int
	HitCount    int
	Hits        []HitBounds
	ResultCount int
	MaxResults  int
	Timing      *bench.TimeCapture
	Config      conf.Searching
	IndexInfo   gather.IndexInfo
}

func NewQueryResult(cf conf.Searching, info gather.IndexInfo) *QueryResult {
	qr := &QueryResult{
		Timing:     bench.Start(),
		Hits:       make([]HitBounds, 0),
		Config:     cf,
		IndexInfo:  info,
		MaxResults: cf.MaxHits,
	}
	return qr
}

type Writer func(io.Writer)

func (qr *QueryResult) Render() SearchResult {
	sr := SearchResult{}
	fn := func(wr Writer) string {
		out := bytes.NewBuffer([]byte{})
		wr(out)
		return out.String()
	}

	sr.Query = fn(qr.ShowQuery)
	sr.Fonts = fn(qr.ShowFonts)
	sr.FilesAndHits = fn(qr.ShowFilesAndHits)
	sr.Summary = fn(qr.ShowSummary)

	sr.IndexInfo = qr.IndexInfo
	sr.MaxResults = qr.MaxResults
	return sr
}

func (qr *QueryResult) ShowFilesAndHits(out io.Writer) {
	prevFile := ""
	for i := 0; i < len(qr.Hits); i++ {
		hit := qr.Hits[i]
		if hit.Unit.Path != prevFile {
			fmt.Fprintf(out, "in %s", hit.Unit.Path)
			fmt.Fprintln(out)
			prevFile = hit.Unit.Path
		}
		hf := HitFormatter{
			Window: 30,
			Hit:    hit,
		}
		fmt.Fprintln(out, hf.String())
	}
}

func (qr *QueryResult) ShowQuery(out io.Writer) {
	fmt.Fprintf(out, "Search query: %s", qr.Config.Query)
	fmt.Fprintln(out)
}

func (qr *QueryResult) ShowSummary(out io.Writer) {
	fmt.Fprintf(out, "Searched: %d files", qr.Searched)
	fmt.Fprintf(out, "Hits: %d of %d\n", qr.ResultCount, qr.HitCount)
	fmt.Fprintf(out, "Max Result Count: %d\n", qr.Config.MaxHits)
	fmt.Fprintln(out, qr.Timing.String())
}

func (qr *QueryResult) ShowFonts(out io.Writer) {
	fmt.Fprintln(out, `
warning                font-lock-warning-face
name                   font-lock-function-name-face
variable               font-lock-variable-name-face
keyword                font-lock-keyword-face
comment                font-lock-comment-face
comment_delimiter      font-lock-comment-delimiter-face
type                   font-lock-type-face
constant               font-lock-constant-face
builtin                font-lock-builtin-face
preprocessor           font-lock-preprocessor-face
string                 font-lock-string-face
doc                    font-lock-doc-face
negation-char          font-lock-negation-char-face
	`)
}

type SearchResult struct {
	Query        string
	FilesAndHits string
	Summary      string
	Fonts        string
	IndexInfo    gather.IndexInfo
	MaxResults   int
}

func (sr SearchResult) Bytes() []byte {
	bin, err := json.MarshalIndent(&sr, "", " ")
	if err != nil {
		return []byte{}
	}
	return bin
}

package interactive

import (
	"github.com/lcaballero/walker/searching"
	cmd "gopkg.in/urfave/cli.v2"
	"io"
	"os"
	"bufio"
	"fmt"
)

func Run(ctx *cmd.Context) error {

	var err error
	in := bufio.NewReader(os.Stdin)

	for err == nil {
		_, err = ClearScreen(os.Stdout)
		if err != nil {
			return err
		}

		_, err = MoveCursorToOrigin(os.Stdout)
		if err != nil {
			return err
		}

		s, err := searching.Search(ctx)
		if err != nil {
			return err
		}

		fmt.Print("query> ")

		line, _, err := in.ReadLine()
		if err != nil {
			return err
		}

		qr, err := s.NewQuery(string(line))
		if err != nil {
			return err
		}

		sr := qr.Render()

		os.Stdout.Write([]byte{'\n'})
		os.Stdout.Write([]byte(sr.Query))
		os.Stdout.Write([]byte(sr.FilesAndHits))
		os.Stdout.Write([]byte(sr.Summary))
		os.Stdout.Write([]byte{'\n'})

		_, _, err = in.ReadRune()
		if err != nil {
			return err
		}
	}

	return nil
}

// ClearScreen clears the console screen
func ClearScreen(w io.Writer) (int, error) {
	return w.Write([]byte("\033[2J"))
}

func MoveCursorToOrigin(w io.Writer) (int, error) {
	return w.Write([]byte("\033[1;1H"))
}

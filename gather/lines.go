package gather

import "errors"

type Lines []Line

func (k Lines) Len() int {
	return len(k)
}

func (lines Lines) Apply(c []byte, n int) ([]byte, error) {
	line, err := lines.Line(n)
	return c[line.Start:line.End], err
}

func (lines Lines) Line(n int) (Line, error) {
	if n < 0 || n >= lines.Len() {
		return Line{}, errors.New("Cannot apply line outside of line bounds")
	}
	line := lines[n]
	return line, nil
}

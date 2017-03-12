package searching

import (
	"fmt"
	"github.com/lcaballero/walker/gather"
	"strings"
)

type HitFormatter struct {
	Hit    HitBounds
	Window int
}

func (h HitFormatter) String() string {
	line, err := h.Hit.FindLine()
	if err != nil {
		return ""
	}

	left := Max(line.Start, h.Hit.Start-h.Window)
	right := Min(line.End, h.Hit.End+h.Window)

	result := fmt.Sprintf("%s %s",
		h.lead(h.Hit.Start, Min(right, h.Hit.End), line),
		h.window(line.Number, left, right))

	return result
}

func (h HitFormatter) match() string {
	return string(h.Hit.Unit.Content[h.Hit.Start:h.Hit.End])
}

func (h HitFormatter) window(line, left, right int) string {
	s := string(h.Hit.Unit.Content[left:right])
	s = strings.TrimRight(s, " \n\r\t")
	return s
}

func (h HitFormatter) lead(start, end int, line gather.Line) string {
	a := start - line.Start
	b := end - line.Start
	return fmt.Sprintf("%6d %2d:%2d | ", line.Number, a, b)
}

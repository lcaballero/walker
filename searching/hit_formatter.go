package searching

import (
	"fmt"
	"github.com/lcaballero/walker/gather"
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

	if h.Hit.Unit.Content[left] == '\n' {
		left += 1
	}
	if h.Hit.Unit.Content[right-1] == '\n' {
		right -= 1
	}

	result := fmt.Sprintf("%s %s",
		h.lead(h.Hit.Start, Min(right, h.Hit.End), line),
		h.window(line.Number, left, right))

	return result
}

func (h HitFormatter) match() string {
	return string(h.Hit.Unit.Content[h.Hit.Start:h.Hit.End])
}

func (h HitFormatter) window(line, left, right int) string {
	return string(h.Hit.Unit.Content[left:right])
}

func (h HitFormatter) lead(start, end int, line gather.Line) string {
	a := start - line.Start
	b := end - line.Start
	return fmt.Sprintf("%6d %2d:%2d", line.Number, a, b)
}

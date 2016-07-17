package searching

import (
	"errors"
	"fmt"
	"github.com/lcaballero/walker/gather"
)

var IndexNotFoundInLines = errors.New("Couldn't find line")

type HitBounds struct {
	Start     int
	End       int
	Width     int
	Unit      *gather.Unit
}

func NewHitBounds(start, end int, unit *gather.Unit) HitBounds {
	hit := HitBounds{
		Start: start,
		End:   end,
		Unit:  unit,
		Width: 30,
	}
	return hit
}

func (h HitBounds) String() string {
	line, err := h.FindLine()
	if err != nil {
		return ""
	}

	left := Max(line.Start, h.Start-h.Width)
	right := Min(line.End, h.End+h.Width)

	if h.Unit.Content[left] == '\n' {
		left += 1
	}
	if h.Unit.Content[right-1] == '\n' {
		right -= 1
	}

	result := fmt.Sprintf("%s %s",
		h.lead(line.Number),
		h.window(line.Number, left, right))

	return result
}

func (h HitBounds) match() string {
	return string(h.Unit.Content[h.Start:h.End])
}

func (h HitBounds) window(line, left, right int) string {
	return string(h.Unit.Content[left:right])
}

func (h HitBounds) lead(line int) string {
	return fmt.Sprintf("[%d:%d,%d]", line, h.Start, h.End)
}

func (h HitBounds) FindLine() (gather.Line, error) {
	a := 0;
	b := h.Unit.LineCount();
	mid := (b - a) / 2

	// Case: we have only one line.
	if b == 1 {
		line, err := h.Unit.Line(0)
		if err != nil {
			return gather.Line{}, err
		}
		if line.HasIndex(h.Start) {
			return gather.Line{Number:1}, nil
		} else {
			return gather.Line{}, errors.New("Couldn't locate line based on hit start point")
		}
	}

	return h.findLine(h.Start, a, b, mid, h.Unit.Lines())
}

func (h HitBounds) findLine(pt, a, b, mid int, lines gather.Lines) (gather.Line, error) {
	line := lines[mid]

	if line.HasIndex(pt) {
		return line, nil
	}

	if pt < line.Start {
		m := (mid - a) / 2
		return h.findLine(pt, a, mid, m, lines)
	}

	if pt > line.End {
		m := (b - mid) / 2
		return h.findLine(pt, mid, b, mid+m, lines)
	}

	return gather.Line{}, IndexNotFoundInLines
}

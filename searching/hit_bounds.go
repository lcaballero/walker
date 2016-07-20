package searching

import (
	"errors"
	"github.com/lcaballero/walker/gather"
)

var IndexNotFoundInLines = errors.New("Couldn't find line")

type HitBounds struct {
	Start int
	End   int
	Unit  *gather.Unit
}

func NewHitBounds(start, end int, unit *gather.Unit) HitBounds {
	hit := HitBounds{
		Start: start,
		End:   end,
		Unit:  unit,
	}
	return hit
}

func (h HitBounds) FindLine() (gather.Line, error) {
	a := 0
	b := h.Unit.LineCount()
	mid := (b - a) / 2

	// Case: we have only one line.
	if b == 1 {
		line, err := h.Unit.Line(0)
		if err != nil {
			return gather.Line{}, err
		}
		if line.HasIndex(h.Start) {
			return gather.Line{Number: 1}, nil
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

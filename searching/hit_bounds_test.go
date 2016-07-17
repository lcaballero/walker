package searching
import (
	"testing"
	"github.com/lcaballero/walker/gather"
	. "github.com/lcaballero/exam/assert"
)


func NewUnit(content string) *gather.Unit {
	u := &gather.Unit{
		Content: []byte(content),
	}
	return u
}

func NewBounds(a, b int, text string) *HitBounds {
	return &HitBounds{
		Start: a,
		End: b,
		Unit: NewUnit(text),
	}
}

const BaseText = `a
b
c
10
  11
    12
      13
    14

  15

`


func Test_HitBound_002(t *testing.T) {
	h := NewBounds(11, 13, BaseText)

	line, _ := h.FindLine()

	IsNotNil(t, line)
	IsEqStrings(t, h.match(), "11")
	IsEqInt(t, line.Start, 9)
	IsEqInt(t, line.End, 14)
	IsEqInt(t, line.Number, 5)
}

func Test_HitBounds_001(t *testing.T) {
	h := HitBounds{
		Start: 0,
		End: 0,
		Unit: NewUnit(""),
	}
	line, err := h.FindLine()

	IsNil(t, err)
	IsNotNil(t, line)
	IsZero(t, line.Start)
	IsZero(t, line.End)
	IsEqInt(t, line.Number, 1)
}

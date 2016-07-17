package gather

import (
	. "github.com/lcaballero/exam/assert"
	"testing"
)

const test_file_content = `

Here's some text.

    Some more text.

	Line starts with a tab.
\t

`

func IsPair(t *testing.T, pair Line, a1, b1 int) {
	a2 := pair.Start
	b2 := pair.End
	if a1 != a2 || b1 != b2 {
		t.Logf("Pair value for (%d,%d) doesn't match (%d,%d)", a1, b1, a2, b2)
		t.Fail()
	}
}

func Test_Unit_004(t *testing.T) {
	t.Log("Parsing file content should recognize empty lines")
	u := &Unit{
		Content: []byte(test_file_content),
	}
	lines, err := u.FindLines()

	IsNil(t, err)
	IsNotNil(t, lines)
	IsTrue(t, u.HasLines())
	IsFalse(t, u.HasError())
	IsEqInt(t, lines.Len(), 9)

	expected := []struct {
		line string
		a, b int
	}{
		{"\n", 0, 1},
		{"\n", 1, 2},
		{"Here's some text.\n", 2, 20},
		{"\n", 20, 21},
		{"    Some more text.\n", 21, 41},
		{"\n", 41, 42},
		{"\tLine starts with a tab.\n", 42, 67},
		{"\\t\n", 67, 70},
		{"\n", 70, 71},
	}
	for i, e := range expected {
		IsEqStrings(t, e.line, u.LineString(i))
		IsPair(t, lines[i], e.a, e.b)
	}
}

func Test_Unit_003(t *testing.T) {
	t.Log("Parsing content with non-line endings should have 1 line")
	u := &Unit{
		Content: []byte(" \t abcdefg 12345"),
	}
	lines := u.Lines()

	IsNotNil(t, lines)
	IsTrue(t, u.HasLines())
	IsEqInt(t, len(lines), 1)
}

func Test_Unit_002(t *testing.T) {
	t.Log("Parsing an empty string should produce 1 line")
	u := &Unit{
		Content: []byte(""),
	}

	lines := u.Lines()

	IsNotNil(t, lines)
	IsEqInt(t, len(lines), 1)
	IsTrue(t, u.HasLines())
}

func Test_Unit_001c(t *testing.T) {
	t.Log("Newly created Unit should have 0 length")
	u := &Unit{}
	IsZero(t, u.Len())
}

func Test_Unit_001b(t *testing.T) {
	t.Log("Newly created Unit shouldn't have lines")
	u := &Unit{}
	IsFalse(t, u.HasLines())
}

func Test_Unit_001(t *testing.T) {
	t.Log("Newly created Unit shouldn't have errors")
	u := &Unit{}
	IsFalse(t, u.HasError())
}

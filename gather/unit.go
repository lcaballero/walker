package gather

import (
	"errors"
)

// A Unit represents the information about a given file found while walking
// the directory structure.  An error will be set if one occurs while
// traversing the tree and IsSkipped will also be set if the Unit was marked
// Skipped.
type Unit struct {
	Path      string
	Ext       string
	Content   []byte
	IsDir     bool
	Error     error
	IsSkipped bool
	lines     Lines // Offset to line starts in the Content
}

func (u *Unit) HasError() bool {
	return u.Error != nil
}

func (u *Unit) FindLines() (Lines, error) {
	lines := make([]Line, 0)
	if u.Content == nil {
		return lines, errors.New("Cannot determine the lines of nil content")
	}

	if len(u.Content) == 0 {
		lines = append(lines, Line{0, 0, 1})
		return lines, nil
	}

	// Case: File has minimum of 1 line.
	start := 0
	i := 0
	n := 1

	if u.Content[i] == '\n' {
		lines = append(lines, Line{start, i + 1, n})
		i++
		n++
		start = i
	}

	for ; i < len(u.Content); i++ {
		if u.Content[i] == '\n' {
			lines = append(lines, Line{start, i + 1, n})
			n++
			start = i + 1
			continue
		}
	}

	// Case: File doesn't end with a newline.
	if i >= len(u.Content) && start < i {
		lines = append(lines, Line{start, i, n})
	}

	return lines, nil
}

func (u *Unit) lazyLines() Lines {
	if u.lines == nil {
		lines, err := u.FindLines()
		if err == nil {
			u.lines = lines
		} else {
			u.lines = make(Lines, 0)
		}
	}
	return u.lines
}

func (u *Unit) Lines() Lines {
	return u.lazyLines()
}

func (u *Unit) HasLines() bool {
	return u.LineCount() > 0
}

func (u *Unit) Line(n int) (Line, error) {
	return u.lazyLines().Line(n)
}

func (u *Unit) Len() int {
	return len(u.Content)
}

func (u *Unit) LineCount() int {
	return u.lazyLines().Len()
}

func (u *Unit) LineString(n int) string {
	if !u.HasLines() {
		return ""
	}
	line, err := u.lazyLines().Apply(u.Content, n)
	if err != nil {
		return ""
	}
	return string(line)
}

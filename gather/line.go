package gather

type Line struct {
	Start int
	End int
	Number int
}

func (k Line) HasIndex(n int) bool {
	return k.Start <= n && n <= k.End
}

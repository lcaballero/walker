package conf

type ValueContext interface {
	String(string) string
	Int(string) int
	Bool(string) bool
	IsSet(string) bool
}

type ContextLoader struct {
	ctx ValueContext
}

func (c ContextLoader) String(name string, dest *string) {
	*dest = c.ctx.String(name)
}

func (c ContextLoader) Int(name string, dest *int) {
	*dest = c.ctx.Int(name)
}

func (c ContextLoader) Bool(name string, dest *bool) {
	*dest = c.ctx.Bool(name)
}


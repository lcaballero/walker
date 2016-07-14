package gather

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
}

func (u Unit) HasError() bool {
	return u.Error != nil
}

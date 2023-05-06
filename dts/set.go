package dts

// Simple Set implementation
type Set interface {
	Add(string)
	Exists(string) bool
}

type DefaultSet struct {
	set map[string]struct{}
}

func (d *DefaultSet) Add(s string) {
	d.set[s] = struct{}{}
}

func (d *DefaultSet) Exists(s string) bool {
	_, ok := d.set[s]

	return ok
}

func NewSet() Set {
	return &DefaultSet{make(map[string]struct{})}
}

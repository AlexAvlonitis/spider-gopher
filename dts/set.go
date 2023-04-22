package dts

// Simple Set implementation
type Set interface {
	Add(string)
	Exists(string) bool
}

type DefaultSet struct {
	set map[string]bool
}

func (d *DefaultSet) Add(s string) {
	d.set[s] = true
}

func (d *DefaultSet) Exists(s string) bool {
	return d.set[s]
}

func NewSet() Set {
	return &DefaultSet{set: make(map[string]bool)}
}

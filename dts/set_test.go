package dts

import "testing"

func TestAdd(t *testing.T) {
	d := NewSet()
	d.Add("test")
	ok := d.Exists("test")
	if !ok {
		t.Error("String has not been saved in the set")
	}
}

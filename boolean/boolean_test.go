package boolean

import (
	"testing"
)

func TestBooleanPointer(t *testing.T) {
	v := true
	got := IsPTrue(&v)
	if got != true {
		t.Errorf("unexpected boolean: expected=%v, got=%v", v, got)
	}

	type BoolPointer *bool
	v = false
	b := BoolPointer(&v)
	got2 := IsPTrue(b)
	if got2 != false {
		t.Errorf("unexpected boolean: expected=%v, get=%v", b, got2)
	}

	var p *bool = nil
	got3 := IsPTrue(p)
	if got3 != false {
		t.Errorf("unexpected boolean: expected=%v, get=%v", true, got3)
	}
}

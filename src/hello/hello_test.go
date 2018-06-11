package hello

import "testing"

func TestGreet(t *testing.T) {
	if s := Greet(); s != "Hello" {
		t.Errorf("expected hello, got %v", s)
	}
}

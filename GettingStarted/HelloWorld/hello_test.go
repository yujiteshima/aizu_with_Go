package hello

import "testing"

func TestHello(t *testing.T) {
	expect := "Hello World"
	actual := hello()

	if actual != expect {
		t.Errorf("actual: %v, expected: %v", actual, expect)
	}
}

package hello

import "testing"

func TestHello(t *testing.T) {
	msg := GetMessage()
	if msg != "Hello, World! SUI!" {
		t.Errorf("expected \"Hello, World! SUI!\", but got %v", msg)
	}
}

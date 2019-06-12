package cli

import (
	"testing"
)

func TestHandlerFunc(t *testing.T) {

	var handler Handler = HandlerFunc(nil) // THIS IS WHAT ACTUALLY MATTERS

	if nil == handler {
		t.Errorf("This should never happen.")
	}
}

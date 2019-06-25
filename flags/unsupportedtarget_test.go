package cliflags

import (
	"testing"
)

func TestInternalUnsupportedTargetAsError(t *testing.T) {
	var err error = internalUnsupportedTarget{} // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == err {
		t.Errorf("This should never happen.")
		return
	}
}

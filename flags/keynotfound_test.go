package cliflags

import (
	"testing"
)

func TestInternalKeyNotFoundAsError(t *testing.T) {
	var err error = internalKeyNotFound{} // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == err {
		t.Errorf("This should never happen.")
		return
	}
}

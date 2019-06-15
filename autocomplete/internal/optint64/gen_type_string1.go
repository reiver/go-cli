package optint64

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

import (
	"fmt"
)

func (receiver Type) String() string {
	if None() == receiver {
		return "optint64.None()"
	}

	return fmt.Sprintf("optint64.Some(%d)", receiver.value)
}

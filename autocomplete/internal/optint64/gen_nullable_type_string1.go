package optint64

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

import (
	"fmt"
)

func (receiver NullableType) String() string {
	if NoneNullable() == receiver {
		return "optint64.NoneNullable()"
	}
	if Null() == receiver {
		return "optint64.Null()"
	}

	return fmt.Sprintf("optint64.SomeNullable(%d)", receiver.value)
}

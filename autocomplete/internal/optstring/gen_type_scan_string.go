package optstring

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

import (
	"fmt"
)

func (receiver *Type) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch t := src.(type) {
	case Type:
		*receiver = t
		return nil
	case string:
		*receiver = Some(t)
		return nil
	case []byte:
		s := string(t)
		*receiver = Some(s)
		return nil
	default:
		return fmt.Errorf("Cannot scan something of type %T into an %T.", src, *receiver)
	}
}

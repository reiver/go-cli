package optstring

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

import (
	"fmt"
)

func (receiver *NullableType) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == src {
		*receiver = Null()
		return nil
	}

	switch t := src.(type) {
	case NullableType:
		*receiver = t
		return nil
	case Type:
		switch t {
		case None():
			*receiver = NoneNullable()
		default:
			datum, err := t.String()
			if nil != err {
				return fmt.Errorf("Problem unwrapping %T: (%T) %v", t, err, err)
			}
			*receiver = SomeNullable(datum)
		}
		return nil
	case string:
		*receiver = SomeNullable(t)
		return nil
	case []byte:
		s := string(t)
		*receiver = SomeNullable(s)
		return nil
	default:
		return fmt.Errorf("Cannot scan something of type %T into an %T.", src, *receiver)
	}
}

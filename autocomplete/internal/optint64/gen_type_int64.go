package optint64

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */


func (receiver Type) Int64() (int64, error) {
	if None() == receiver {
		return 0, errNone
	}

	return receiver.value, nil
}

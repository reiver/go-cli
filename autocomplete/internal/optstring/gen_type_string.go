package optstring

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */


func (receiver Type) String() (string, error) {
	if None() == receiver {
		return "", errNone
	}

	return receiver.value, nil
}

package optstring

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */


func (receiver NullableType) String() (string, error) {
	if NoneNullable() == receiver {
		return "", errNoneNullable
	}
	if Null() == receiver {
		return "", errNull
	}

	return receiver.value, nil
}

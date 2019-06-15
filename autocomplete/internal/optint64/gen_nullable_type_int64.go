package optint64

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */


func (receiver NullableType) Int64() (int64, error) {
	if NoneNullable() == receiver {
		return 0, errNoneNullable
	}
	if Null() == receiver {
		return 0, errNull
	}

	return receiver.value, nil
}

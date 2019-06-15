package optstring

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
)

type NullableType struct {
	loaded bool
	null   bool
	value  string
}

func (receiver NullableType) MarshalJSON() ([]byte, error) {
	if NoneNullable() == receiver {
		return nil, errNoneNullable
	}
	if Null() == receiver {
		return json.Marshal(nil)
	}

	return json.Marshal(receiver.value)
}

func (receiver NullableType) WhenNone(fn func()) {
	if NoneNullable() == receiver {
		fn()
	}
}

func (receiver NullableType) WhenNull(fn func()) {
	if Null() == receiver {
		fn()
	}
}

func (receiver NullableType) WhenSome(fn func(string)) {
	if NoneNullable() != receiver && Null() != receiver {
		fn(receiver.value)
	}
}

func (receiver *NullableType) UnmarshalJSON(b []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if 0 == bytes.Compare(b, []byte{'n','u','l','l'}) {
		*receiver = Null()
		return nil
	}

	var target string

	if err := json.Unmarshal(b, &target); nil != err {
		return err
	}

	*receiver = SomeNullable(target)

	return nil
}

func (receiver NullableType) Value() (driver.Value, error) {
	if NoneNullable() == receiver {
		return nil, errNoneNullable
	}
	if Null() == receiver {
		return nil, nil
	}

	return receiver.value, nil
}

func NoneNullable() NullableType {
	return NullableType{}
}

func SomeNullable(value string) NullableType {
	return NullableType{
		value:  value,
		loaded: true,
	}
}

func Null() NullableType {
	return NullableType{
		null:   true,
		loaded: true,
	}
}

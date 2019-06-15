package optstring

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Type struct {
	loaded bool
	value  string
}

func (receiver Type) MarshalJSON() ([]byte, error) {
	if None() == receiver {
		return nil, errNone
	}

	return json.Marshal(receiver.value)
}

func (receiver *Type) UnmarshalJSON(b []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if 0 == bytes.Compare(b, []byte{'n','u','l','l'}) {
		return fmt.Errorf("Cannot unmarshal %q into %T.", string(b), *receiver)
	}

	var target string

	if err := json.Unmarshal(b, &target); nil != err {
		return err
	}

	*receiver = Some(target)

	return nil
}

func (receiver Type) WhenNone(fn func()) {
	if None() == receiver {
		fn()
	}
}

func (receiver Type) WhenSome(fn func(string)) {
	if None() != receiver {
		fn(receiver.value)
	}
}

func (receiver Type) Value() (driver.Value, error) {
	if None() == receiver {
		return nil, errNone
	}

	return receiver.value, nil
}

func None() Type {
	return Type{}
}

func Some(value string) Type {
	return Type{
		value:  value,
		loaded: true,
	}
}

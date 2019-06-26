package cliflags

import (
	"fmt"
)

type Value struct {
	loaded bool
	datum string
}

func SomeValue(v string) Value {
	return Value{
		loaded:true,
		datum:v,
	}
}

func NoValue() Value {
	return Value{}
}

func (receiver Value) Else(value string) Value {
	if receiver.loaded {
		return receiver
	}

	return SomeValue(value)
}

func (receiver Value) ElseUnwrap(value string) string {
	if receiver.loaded {
		return receiver.datum
	}

	return value
}

func (receiver Value) GoString() string {
	switch receiver.loaded {
	case true:
		return fmt.Sprintf("cliflags.SomeValue(%q)", receiver.datum)
	default:
		return "cliflags.NoValue()"
	}
}

func (receiver Value) Map(fn func(string)string) Value {
	if NoValue() == receiver {
		return receiver
	}

	var value string = fn(receiver.datum)

	return SomeValue(value)
}

func (receiver Value) Then(fn func(string)Value) Value {
	if NoValue() == receiver {
		return receiver
	}

	return fn(receiver.datum)
}

func (receiver Value) Unwrap() (string, bool) {
	if NoValue() == receiver {
		return "", false
	}

	return receiver.datum, true
}

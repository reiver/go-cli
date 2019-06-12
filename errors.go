package cli

import (
	"errors"
)

var (
	errNilReceiver = errors.New("cli: Nil Receiver")
	errNilHandler  = errors.New("cli: Nil Handler")
	errInvalidUTF8 = errors.New("cli: Invalid UTF-8")
)

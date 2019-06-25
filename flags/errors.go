package cliflags

import (
	"errors"
)

var (
	errInternalError = errors.New("cli: Internal Error")
	errNilReceiver   = errors.New("cli: Nil Receiver")
	errNilTarget     = errors.New("cli: Nil Target")
)

package cliflag

import (
	"errors"
)

var (
	errFound             = errors.New("cli: Found")
	errLoaded            = errors.New("cli: Loaded")
	errNilReceiver       = errors.New("cli: Nil Receiver")
	errNilTarget         = errors.New("cli: Nil Target")
	errNoTokens          = errors.New("cli: No Tokens")
	errNotLoaded         = errors.New("cli: Not Loaded")
	errUnsupportedTarget = errors.New("cli: Unsupported Target")
)

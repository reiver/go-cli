package clitest

import (
	"errors"
)

var (
	errClosed      = errors.New("Closed")
	errNilReceiver = errors.New("Nil Receiver")
)

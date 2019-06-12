package cli

import (
	"io"
)

type HandlerFunc func(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) ExitCode

func (receiver HandlerFunc) Run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) ExitCode {
	return receiver(stdin, stdout, stderr, args...)
}

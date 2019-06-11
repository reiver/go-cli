package cli

import (
	"io"
)

// A Handler responds to a CLI command.
type Handler interface {
	Run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) ExitCode
}

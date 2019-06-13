package cli

import (
	"io"

	"testing"
)

func TestMuxHandleFunc(t *testing.T) {
	fn := func(io.ReadCloser, io.WriteCloser, io.WriteCloser, ...string) ExitCode {
		return ExitCodeOK
	}

	var mux Mux

	if err := mux.HandleFunc(fn, "some", "command"); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %q", err, err)
		return
	}
}

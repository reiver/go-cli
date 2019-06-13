package cli_test

import (
	"io"
)

var (
	fakeStdin  io.ReadCloser  = nopReadCloser{}
	fakeStdout io.WriteCloser = nopWriteCloser{}
	fakeStderr io.WriteCloser = nopWriteCloser{}
)


type nopReadCloser struct{}

func (nopReadCloser) Close() error {
	return nil
}

func (nopReadCloser) Read(p []byte) (n int, err error) {
	return 0, nil
}


type nopWriteCloser struct{}

func (nopWriteCloser) Close() error {
	return nil
}

func (nopWriteCloser) Write(data []byte) (n int, err error) {
	return 0, nil
}

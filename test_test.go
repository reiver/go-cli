package cli_test

import (
	"bytes"
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


type writeCloseBuffer struct {
	buffer bytes.Buffer
}

func (writeCloseBuffer) Close() error {
	return nil
}

func (receiver writeCloseBuffer) String() string {
	return receiver.buffer.String()
}

func (receiver *writeCloseBuffer) Write(p []byte) (n int, err error) {
	return receiver.buffer.Write(p)
}


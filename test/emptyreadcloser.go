package clitest

import (
	"io"
)

type EmptyReadCloser struct{}

func (receiver EmptyReadCloser) Close() error {
	return nil
}

func (receiver EmptyReadCloser) String() string {
	return ""
}

func (receiver EmptyReadCloser) Read(p []byte) (int, error) {
	return 0, io.EOF
}

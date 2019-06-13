package clitest

// WriteCloser provides a io.WriteCloser which stores the data written to it, and returns it via its .String() method.
type WriteCloser struct {
	buffer []byte
	closed bool
}

// Close makes clitest.WriteCloser fit the io.Closer interface (which is a necessary condition to fit the io.WriteCloser interface).
func (receiver *WriteCloser) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.closed = true

	return nil
}

// String returns what was written to the io.WriteCloser
func (receiver WriteCloser) String() string {
	return string(receiver.buffer)
}

// Write makes clitest.WriteCloser fit the io.Writer interface (which is a necessary condition to fit the io.WriteCloser interface).
func (receiver *WriteCloser) Write(p []byte) (int, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	if receiver.closed {
		return 0, errClosed
	}

	receiver.buffer = append(receiver.buffer, p...)

	return len(p), nil
}

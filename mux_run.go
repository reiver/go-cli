package cli

import (
	"io"
)

func (receiver *Mux) Run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	if nil == receiver {
		return ExitCodeInternalError
	}

	commandName, commandParams := commandSplit(command...)

	var name   []string = append([]string(nil), commandName...)
	var params []string = append([]string(nil), commandParams...)

	var handler Handler
	for {
		serialized, err := serialize(name...)
		if nil != err {
			return ExitCodeInternalError
		}

		receiver.mutex.RLock()
		var found bool
		handler, found = receiver.handlers[serialized]
		receiver.mutex.RUnlock()
		if found {
			break
		}

		if 1 > len(name) {
			return ExitCodeBadRequest
		}

		popped := name[len(name)-1]
		name = name[:len(name)-1]
		params = append([]string{popped}, params...)
	}

	return handler.Run(stdin, stdout, stderr, params...)
}

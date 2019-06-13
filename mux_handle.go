package cli

import (
	"io"
)

// Handle registers a new handler for a command.
//
// Example
//
//	var mux cli.Mux
//	
//	// ...
//	
//	err = mux.Handle(initHandler, "init")
//	
//	err = mux.Handle(launchHandler, "launch")
//	
//	err = mux.Handle(launchWebHandler, "launch", "web")
//	
//	err = mux.Handle(launchWebHandler, "backup", "--file={fileName}")
func (receiver *Mux) Handle(handler Handler, command ...string) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == handler {
		return errNilHandler
	}

	commandName, _ := commandSplit(command...)
//@TODO: Should we do something about the ‘command params’ returned from commandSplit()

	serialized, err := serialize(commandName...)
	if nil != err {
		return err
	}

	receiver.mutex.Lock()
	if nil == receiver.handlers {
		receiver.handlers = map[string]Handler{}
	}
	receiver.handlers[serialized] = handler
	receiver.mutex.Unlock()

	return nil
}

func (receiver *Mux) HandleFunc(fn func(io.ReadCloser, io.WriteCloser, io.WriteCloser, ...string)ExitCode, command ...string) error {
	var handler Handler = HandlerFunc(fn)

	return receiver.Handle(handler, command...)
}

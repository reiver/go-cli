package cli_test

import (
	"github.com/reiver/go-cli"
	"github.com/reiver/go-cli/test"

	"fmt"
	"io"
	"os"
)

// GoodbyeMiddlewareHandler is a cli.Handler "middleware".
//
// It ‘wraps’ another cli.Handler.
//
// Usage
//
//	var innerHandler cli.Handler
//	
//	// ...
//	
//	var middleware GoodbyeMiddlewareHandler
//	
//	middleware.Handler = innerHandler
//	
//	// ...
//	
//	cii.RunAndTheExit(middleware)
type GoodbyeMiddlewareHandler struct {
	Handler cli.Handler
}

func (receiver GoodbyeMiddlewareHandler) Run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) cli.ExitCode {

	handler := receiver.Handler
	if nil == handler {
		return cli.ExitCodeInternalError
	}

	exitcode := handler.Run(stdin, stdout, stderr, command...)

	fmt.Fprintln(stdout, "Goodbye · Khodafez · 안녕 · 再見")

	return exitcode
}

func ExampleHandler_middleware() {

	handlerFunc := func(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) cli.ExitCode {
		fmt.Fprintln(stderr, "Oh no!")

		return cli.ExitCodeOK
	}

	var mux cli.Mux

	err := mux.HandleFunc(handlerFunc)
	if nil != err {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		return
	}

	// ...

	var handler cli.Handler = GoodbyeMiddlewareHandler{
		Handler: &mux,
	}

	// Normally we would call something such as:
	//
	//	cli.RunAndTheExit(handler)
	//
	// But for the sake of this example, we call the handler's .Run() method
	// ourselves, so we can capture the Stdout, and the Exit Code.
	var stdin  clitest.EmptyReadCloser
	var stdout clitest.WriteCloser
	var stderr clitest.WriteCloser
	exitcode := handler.Run(stdin, &stdout, &stderr, "one", "two", "three")

	fmt.Printf("exitcode = %q\n", exitcode)
	fmt.Printf("stdout   = %q\n", stdout.String())
	fmt.Printf("stderr   = %q\n", stderr.String())

	// Output:
	// exitcode = "OK"
	// stdout   = "Goodbye · Khodafez · 안녕 · 再見\n"
	// stderr   = "Oh no!\n"
}

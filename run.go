package cli

import (
	"io"
	"os"
)

// Run runs a cli.Handler.
//
// Calling cli.Run() as in:
//
//	var handler cli.Handler
//	
//	// ...
//	
//	exticode := cli.Run(handler)
//
// Is the equivalent of doing:
//
//	var handler cli.Handler
//	
//	// ...
//	
//	exticode := handler.Run(os.Stdin, os.Stdout, os.Stderr, os.Args[1:]...)
func Run(handler Handler) ExitCode {
	return handler.Run(os.Stdin, os.Stdout, os.Stderr, os.Args[1:]...)
}

// RunFunc is similar to Run, expect instead of taking a cli.Handler, it takes a func
// with the same signature as the Run() method in the cli.Handler.
func RunFunc(fn func(io.ReadCloser, io.WriteCloser, io.WriteCloser, ...string)ExitCode) ExitCode {
	return Run(HandlerFunc(fn))
}

// RunAndThenExit runs a cli.Handler, and then exits using the ‘exit code’ returned from the handler's Run() method.
//
// Calling cli.RunAndThenExit() as in:
//
//	var handler cli.Handler
//	
//	// ...
//	
//	cli.RunAndThenExit(handler)
//
// Is the equivalent of doing:
//
//	var handler cli.Handler
//	
//	// ...
//	
//	exticode := handler.Run(os.Stdin, os.Stdout, os.Stderr, os.Args[1:]...)
//	
//	os.Exit( int(exitcode) )
func RunAndThenExit(handler Handler) {
	os.Exit( int( Run(handler) ) )
}

// RunFuncAndThenExit is similar to RunAndThenExit, expect instead of taking a cli.Handler, it takes a func
// with the same signature as the Run() method in the cli.Handler.
func RunFuncAndThenExit(fn func(io.ReadCloser, io.WriteCloser, io.WriteCloser, ...string)ExitCode) {
	RunAndThenExit(HandlerFunc(fn))
}

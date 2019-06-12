package cli

import (
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

package cli

import (
	"sync"
)

// Mux registers (a number of) handlers to run (a number of) commands.
//
// Example
//
// Typical usage of cli.Mux might be:
//
//	import "github.com/reiver/go-mux"
//	
//	// ...
//	
//	var mux cli.Mux
//	
//	// ...
//	
//	// myprogram init
//	err := mux.Handle(initHandler, "init")
//	
//	// ...
//	
//	// myprogram launch
//	err := mux.Handle(launchHandler, "launch")
//	// ...
//	
//	// myprogram launch web
//	err := mux.Handle(launchWebHandler, "launch", "web")
//	
//	// ...
//	
//	cli.RunAndThenExit(mux)
//
// Handler
//
// cli.Mux is itself also a cli.Handler.
//
// So it has a ‘.Run()’ method.
//
// And thus a cli.Mux can be called from ‘cli.Run()’, and ‘cli.RunAndThenExit()’.
//
// As in, for example:
//
//	var mux cli.Mux
//	
//	// ...
//	
//	cli.RunAndThenExit(mux)
//
// Middleware
//
// Because cli.Mux is also a cli.Handler, it can be wrapped in "middleware".
//
// As in:
//
//	var mux cli.Mux
//	
//	// ...
//	
//	var handler cli.Handler
//
//	handler = AlsoSaveStderrToFileMiddleware(mux)
//	
//	// ...
//	
//	cli.RunAndThenExit(handler)
type Mux struct {
	mutex    sync.RWMutex
	handlers map[string]Handler
}

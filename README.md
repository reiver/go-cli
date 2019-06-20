# go-cli

Package **cli** provides a way to creating **command line interface** (**CLI**) programs, for the Go programming language, in a style similar to the "net/http" library that is part of the built-in Go standard library, including support for "middleware".

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-cli

[![GoDoc](https://godoc.org/github.com/reiver/go-cli?status.svg)](https://godoc.org/github.com/reiver/go-cli)

## Example
```go
import "github.com/reiver/go-cli"
import "github.com/reiver/go-cli/flag"

// ...

var mux cli.Mux

// ...

func main() {
	cli.RunAndThenExit(mux)
}

// ...

// This handler will run with this program is called with with sub-command argument “apple”:
//
//	myprogram apple
//
func init() {

	var handler cli.Handler = new(appleHandler)

	var command []string = []string{"apple"}

	if err := mux.Handle(handler, command...); nil != err {
		panic(err)
	}
}

// ...

// This handler will run with this program is called with with command argument “banana”:
//
//	myprogram banana
//
func init() {

	var handler cli.Handler = new(bananaHandler)

	var command []string = []string{"banana"}

	if err := mux.Handle(handler, command...); nil != err {
		panic(err)
	}
}

// ...

// This handler will run with this program is called with with command, and sub-command argument “banana”, “bread”:
//
//	myprogram banana bread
//
// NOTE THAT THIS HAS BOTH A COMMAND, AND A SUB-COMMAND!
//
func init() {

	var handler cli.Handler = new(bananaBreadHandler)

	var command []string = []string{"banana", bread}

	if err := mux.Handle(handler, command...); nil != err {
		panic(err)
	}
}

// ...

// This handler will run with this program is called with with sub-command argument “cherry”:
//
//	myprogram banana
//
func init() {

	var handler cli.Handler = new(cherryHandler)

	var command []string = []string{"cherry"}

	if err := mux.Handle(handler, command...); nil != err {
		panic(err)
	}
}

// ...

type bananaBreadHandler struct{}

func Run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) cli.ExitCode {

	//@TODO

}
```

/*
Package cli provides a way to creating command line interface (CLI) programs, for the Go programming language, in a style similar to the "net/http" library that is part of the built-in Go standard library, including support for "middleware".

Hello World Handler

WARING: The ‘Hello World Handler’ example you are about to see is meant to give you a sense of how the cli.Handler works.
If does not use cli.ServerMux. It would be pointless to use pakage cli without using cli.ServerMux.
Later on we show you a more realistic example using cli.ServerMux. But anyway....

Similar to http.Handler, that is part of the built-in Go "net/http" library, package cli has a cli.Handler.

The definition of cli.Handler is:

	type cli.Handler interface {
		Run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) cli.ExitCode
	}

So, a ‘hello world’ cli.Handler might be:

	import "github.com/reiver/go-cli"
	
	// ...
	
	type HelloWorldCLIHandler struc{}
	
	func (HelloWorldCLIHandler) Run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) cli.ExitCode {
		fmt.Fprintln(stdout, "Hello world!")

		return cli.ExitCodeOK
	}

And from your main() function, you might use do something such as:

	import "github.com/reiver/go-cli"
	
	// ...
	
	func main() {
		var handler cli.Handler = new(HelloWorldCLIHandler)

		cli.RunAndThenExit(handler)
	}

Of course, if this is all you did, you would never use package cli.

The power of package cli is when you use cli.ServerMux.

Exit Codes

Note that in the example handler we just showed:

	import "github.com/reiver/go-cli"
	
	// ...
	
	type HelloWorldCLIHandler struc{}
	
	func (HelloWorldCLIHandler) Run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) cli.ExitCode {
		fmt.Fprintln(stdout, "Hello world!")

		return cli.ExitCodeOK
	}

That is returned ‘cli.ExitCodeOK’.

‘cli.ExitCodeOK’ is an ‘exit code’.

There are in fact a number of ‘exit codes’ for you to choose from.

And you should return the appropriate ‘exit code’ when writing your handlers.

Other ‘exit codes’ include:

• cli.ExitCodeOK

• cli.ExitCodeError

• cli.ExitCodeBadRequest

• cli.ExitCodeBadInput

• cli.ExitCodeNoInput

• cli.ExitCodeInternalError

• cli.ExitCodeOSError

• cli.ExitCodeOSFileError

See the documentation on each of these for when it is appropriate to use each of these.
*/
package cli

/*
Package cli provides a way to creating command line interface (CLI) programs, for the Go programming language, in a style similar to the "net/http" library that is part of the built-in Go standard library, including support for "middleware".

Hello World Handler

WARING: The ‘Hello World Handler’ example you are about to see is meant to give you a sense of how the cli.Handler works.
If does not use cli.Mux. It would be pointless to use pakage cli without using cli.Mux.
Later on we show you a more realistic example using cli.Mux. But anyway....

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

And from your main() function, you might do something such as:

	import "github.com/reiver/go-cli"
	
	// ...
	
	func main() {
		var handler cli.Handler = new(HelloWorldCLIHandler)

		cli.RunAndThenExit(handler)
	}

Of course, if this is all you did, you would never use package cli.

The power of package cli is when you use cli.Mux.

Exit Codes

Note that in the example handler we just showed:

	import "github.com/reiver/go-cli"
	
	// ...
	
	type HelloWorldCLIHandler struct{}
	
	func (HelloWorldCLIHandler) Run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) cli.ExitCode {
		fmt.Fprintln(stdout, "Hello world!")

		return cli.ExitCodeOK // <--------------- Exit Code used here.
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

• cli.ExitCodeNoUser

• cli.ExitCodeNoHost

• cli.ExitCodeUnavailable

• cli.ExitCodeInternalError

• cli.ExitCodeOSError

• cli.ExitCodeOSFileError

So, for another example:

	import "github.com/reiver/go-cli"
	
	// ...
	
	type MyCLIHandler struct{}
	
	func (MyCLIHandler) Run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) cli.ExitCode {
		if 1 > len(command) {
			return cli.ExitCodeBadRequest // <--------------- Exit Code used here.
		}

		err := computeSomething()
		if nil != err {
			fmt.FPrintf(stderr, "ERROR: %s\n", err)

			return cli.ExitCodeInternalError // <--------------- Exit Code used here.
		}
		fmt.Fprintln(stdout, "Computation Complete!")

		return cli.ExitCodeOK // <--------------- Exit Code used here.
	}

See the documentation on each of these ‘exit codes’ for when it is appropriate to use each of these.

Stdout And Stderr

In the initial example handler we showed:

	import "github.com/reiver/go-cli"
	
	// ...
	
	type HelloWorldCLIHandler struc{}
	
	func (HelloWorldCLIHandler) Run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) cli.ExitCode {
		fmt.Fprintln(stdout, "Hello world!") // <--------------- Stdout used here.

		return cli.ExitCodeOK
	}

We used the ‘stdout’ parameter to output something to the user.

We did that with the code:

	fmt.Fprintln(stdout, "Hello world!")

And in the second example handler we showed:

	func (MyCLIHandler) Run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) cli.ExitCode {
		if 1 > len(command) {
			return cli.ExitCodeBadRequest
		}

		err := computeSomething()
		if nil != err {
			fmt.FPrintf(stderr, "ERROR: %s\n", err) // <--------------- Stderr used here.

			return cli.ExitCodeInternalError
		}
		fmt.Fprintln(stdout, "Computation Complete!")

		return cli.ExitCodeOK
	}
We used the ‘stderr’ to output an error message to the user.

We did that with the code:

	fmt.FPrintf(stderr, "ERROR: %s\n", err)

‘stdout’, and ‘stdout’ are 2 different ways outputting to the user.

‘stdout’ is for error messages.
And ‘stdout’ is for everything else (that isn't an error message).

Since both ‘stdout’, and ‘stdout’ are ‘io.Writer’, you can use funcs such as ‘fmt.Fprint()’, ‘fmt.Fprintf()’, and ‘fmt.Fprintln()’ to output to them.
In addition to being able to just call each of their ‘.Write()’ methods directly.

*/
package cli

/*
Package cliflags provides tools for parsing the command line, and interpreting it as a command line flag.

Basic Example

Here is a basic example that uses ‘cliflag.KeyValues’ as the target for cliflags.Parse():…

	import "github.com/reiver/go-cliflags"
	
	// ...
	
	var flags cli.KeyValues
	
	args, err := cliflags.Parse(&flags, tokens...)
	
	// --http=?
	// default to "8080"
	httpPort := flags.Load("http").ElseUnwrap("8080")

	// --level=?
	// default to "5"
	level    := flags.Load("level").ElseUnwrap("5")

	fileName := args[0]
*/
package cliflags

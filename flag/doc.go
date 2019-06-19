/*
Package cliflag provides tools for parsing a single token from the command line, and interpreting it as a command line flag.

For most people, it would make more sense to use the ‘cliflags’ (with an ‘s’ at the end) package, rather than this package ‘cliflag’.

(I.e., For most people, it would make more sense to use the "github.com/reiver/go-cli/‘cliflags" package,
rather than the "github.com/reiver/go-cli/‘cliflag" package.)

This package is exposed for those wishing for a more low-level API.

Example

Here is a basic example that uses ‘cliflag.KeyValue’ as the target for cliflag.Parse():…

	import "github.com/reiver/go-cliflag"
	
	// ...
	
	var keyvalue cli.KeyValue
	
	remainingTokens, err := cliflag.Parse(&keyvalue, tokens)
	
	key, value, err := keyvalue.Unwrap()

Example Map String String

Here is a basic example that uses a ‘map[string]string’ as the target for cliflag.Parse():…

	import "github.com/reiver/go-cliflag"
	
	// ...
	
	var target map[string]string
	
	remainingTokens, err := cliflag.Parse(&target, tokens)

Example Map String String

Here is a basic example that uses a ‘map[string]interface{}’ as the target for cliflag.Parse():…

	import "github.com/reiver/go-cliflag"
	
	// ...
	
	var target map[string]interface{}
	
	remainingTokens, err := cliflag.Parse(&target, tokens)

*/
package cliflag

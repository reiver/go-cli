package cliflag

import (
	"fmt"
)

// NotFlag is the error returned from cliflag.Parse() when the token it parsed is not a flag.
//
// For example:
//
//	// Imagine that:
//	//
//	//	token := []string{"----", "-m", "This is a message", "filename.txt"}
//	//
//	// Note that the first element in the ‘[]string’ is "----".
//	// Which is not a (valid) flag.
//	
//	remainingTokens, err := cliflag.Parse(&target, tokens...)
//	
//	if nil != err {
//		
//		switch casted := err.(type) {
//		case cliflag.NotFlag:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
//
// Note that here we used a type-switch to determine if the error returned from cliflag.Parse()
// is a cliflag.NotFlag.
//
// Note that you can get the offending token from a cliflag.NotFlag via the .Token() method.
// As in, for example:
//
//	remainingTokens, err := cliflag.Parse(&target, tokens...)
//	
//	if nil != err {
//		
//		switch casted := err.(type) {
//		case cliflag.NotFlag:
//			
//			var token string = casted.Token()
//			
//			fmt.Fprintf(os.Stderr, "CLI SYNTAX ERROR: %q is not a flag.", token)
//			
//		default:
//			
//			fmt.Fprintf(os.Stderr, "ERROR: something bad happened: %s", err)
//			
//		}
//	}
type NotFlag interface {
	error
	NotFlag()
	Token() string
}

type internalNotFlag struct {
	token string
}

func (receiver internalNotFlag) Error() string {
	return fmt.Sprintf("cli: %q is not a flag", receiver.token)
}

func (receiver internalNotFlag) Token() string {
	return receiver.token
}

func (internalNotFlag) NotFlag() {
	// Nothing here.
}

package cliflag

import (
	"fmt"
)

// EndOfFlags is the error returned from cliflag.Parse() when there are no more flags.
// This also means that if there are any remaining tokens, they are (non-flag) arguments.
//
// For example:
//
//	// Imagine that:
//	//
//	//	token := []string{"filename.txt", "abc.ini"}
//	//
//	// Note that the first element in the ‘[]string’ is "filename.txt".
//	// Which is not a (valid) flag.
//	
//	remainingTokens, err := cliflag.Parse(&target, tokens...)
//	
//	if nil != err {
//		
//		switch casted := err.(type) {
//		case cliflag.EndOfFlags:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
//
// Note that here we used a type-switch to determine if the error returned from cliflag.Parse()
// is a cliflag.EndOfFlags.
//
// Note that you can get the offending token from a cliflag.EndOfFlags via the .Token() method.
// As in, for example:
//
//	remainingTokens, err := cliflag.Parse(&target, tokens...)
//	
//	if nil != err {
//		
//		switch casted := err.(type) {
//		case cliflag.EndOfFlags:
//			
//			var token string = casted.Token()
//			
//			fmt.Fprintf(os.Stderr, "END OF FLAGS: %q is not a flag.", token)
//			
//		default:
//			
//			fmt.Fprintf(os.Stderr, "ERROR: something bad happened: %s", err)
//			
//		}
//	}
type EndOfFlags interface {
	error
	EndOfFlags()
	Token() string
}

type internalEndOfFlags struct {
	token string
}

func (receiver internalEndOfFlags) Error() string {
	return fmt.Sprintf("cli: end of flags: %q is not a flag", receiver.token)
}

func (receiver internalEndOfFlags) Token() string {
	return receiver.token
}

func (internalEndOfFlags) EndOfFlags() {
	// Nothing here.
}

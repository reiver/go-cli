package cliflags

import (
	"fmt"
)

// UnsupportedTarget is the error returned from cliflag.Parse() when the type of the target
// is not supported.
//
// For example:
//
//	var target int64 // <---- Note that the type is int64, which cliflags.Parse() does not support. So it will return an error.
//	
//	// ...
//	
//	args, err := cliflags.Parse(&target, tokens)
//	
//	if nil != err {
//		switch err.(type) {
//		case cliflags.UnsupportedTarget:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
type UnsupportedTarget interface {
	error

	UnsupportedTarget()

	Target() interface{}
}

type internalUnsupportedTarget struct {
	target interface{}
}

func (receiver internalUnsupportedTarget) Error() string {
	return fmt.Sprintf("cli: unsupported target: %T", receiver.target)
}

func (receiver internalUnsupportedTarget) Target() interface{} {
	return receiver.target
}

func (internalUnsupportedTarget) UnsupportedTarget() {
	// Nothing here.
}

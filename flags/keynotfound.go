package cliflags

import (
	"fmt"
)

// KeyNotFound is the error returned from cliflags.Parse() when Load is called and there is
// no value stored for that key.
//
// For example:
//
//	var keyvalues cliflags.KeyValues
//	
//	// ...
//	
//	err := keyvalues.Store(key, value)
//	
//	if nil != err {
//		switch err.(type) {
//		case cliflags.KeyNotFound:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
type KeyNotFound interface {
	error
	KeyNotFound()

	// Key returns the key.
	Key() string
}

type internalKeyNotFound struct {
	key string
}

func (receiver internalKeyNotFound) Error() string {
	return fmt.Sprintf("cli: key %q not found", receiver.key)
}

func (receiver internalKeyNotFound) Key() string {
	return receiver.key
}

func (internalKeyNotFound) KeyNotFound() {
	// Nothing here.
}

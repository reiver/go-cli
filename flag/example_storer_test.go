package cliflag_test

import (
	"github.com/reiver/go-cli/flag"

	"errors"
	"fmt"
)

type MyFlag struct {
	Key string
	Value interface{}
}

func (receiver *MyFlag) Store(key string, value interface{}) error {
	if nil == receiver {
		return errors.New("Nil Receiver")
	}

	receiver.Key   = key
	receiver.Value = value

	return nil
}

func ExampleStorer() {

	// This is the equivalent to having a command line with something such as:
	//
	//	theprogram --when=now -n5 -m 'This is a message' filename.txt
	//
	// Or, if you want to line it up with the code below:
	//
	//                   theprogram --when=now    -n5    -m   'This is a message'   filename.txt
	var tokens []string = []string{"--when=now", "-n5", "-m", "This is a message", "filename.txt"}


	var flag MyFlag // <---- Here is the custom type we created.

	_, err := cliflag.Parse(&flag, tokens...) // <---- This will assign a value to ‘flag’.
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("KEY:   %q\n", flag.Key)
	fmt.Printf("VALUE: %q\n", flag.Value)

	// Output:
	// KEY:   "when"
	// VALUE: "now"
}

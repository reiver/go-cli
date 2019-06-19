package cliflag_test

import (
	"github.com/reiver/go-cli/flag"

	"fmt"
)

func ExampleKeyValue() {

	// This is the equivalent to having a command line with something such as:
	//
	//	theprogram --when=now -n5 -m 'This is a message' filename.txt
	//
	// Or, if you want to line it up with the code below:
	//
	//                   theprogram --when=now    -n5    -m   'This is a message'   filename.txt
	var tokens []string = []string{"--when=now", "-n5", "-m", "This is a message", "filename.txt"}


	var keyvalue cliflag.KeyValue // <---- Here is the key-value.

	_, err := cliflag.Parse(&keyvalue, tokens...) // <---- This will assign a value to ‘keyvalue’.
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	key, value, err := keyvalue.Unwrap() // <--- Here we get the ‘key’, and the ‘value’, out of ‘keyvalue’
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("KEY:   %q\n", key)
	fmt.Printf("VALUE: %q\n", value)

	// Output:
	// KEY:   "when"
	// VALUE: "now"
}

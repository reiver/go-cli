package cliflag_test

import (
	"errors"
	"fmt"

	"github.com/reiver/go-cli/flag"
)

// KeyValues stores many key-value pairs.
type KeyValues struct {
	data []cliflag.KeyValue
}

// Store makes &KeyValues fit the cli.Storer interface, which means we can pass
// &KeyValues to cliflag.Parse() as the target.
func (receiver *KeyValues) Store(key string, value string) error {
	if nil == receiver {
		return errors.New("Nil Receiver")
	}

	var datum cliflag.KeyValue

	if err := datum.Store(key, value); nil != err {
		return err
	}

	receiver.data = append(receiver.data, datum)

	return nil
}

// For lets us iterate over all the key-value pairs inside of KeyValues.
//
// For example:
//
//	var data KeyValues
//	
//	// ...
//	
//	data.For(func(key string, value string){
//		fmt.Printf("KEY:   %q\n", key)
//		fmt.Printf("VALUE: %q\n", value)
//	})
func (receiver KeyValues) For(fn func(string, string)) {
	if nil == fn {
		return
	}

	for _, keyvalue := range receiver.data {
		key, value, err := keyvalue.Unwrap()
		if nil != err {
			continue
		}

		fn(key, value)
	}
}

func ExampleParse_loop() {

	// This is the equivalent to having a command line with something such as:
	//
	//	theprogram --when=now -n5 -m 'This is a message' --apple=one --banana=two -- --cherry=three filename.txt
	//
	// Or, if you want to line it up with the code below:
	//
	//                   theprogram --when=now    -n5    -m   'This is a message'   --apple=one    --banana=two    --    --cherry=three    filename.txt
	var tokens []string = []string{"--when=now", "-n5", "-m", "This is a message", "--apple=one", "--banana=two", "--", "--cherry=three", "filename.txt"}


	var flags KeyValues         // <---- We will store all the flags we find in ‘tokens’ in here in ‘flags’.

	var args []string = tokens  // <---- We will store the remaining (non-flag) args in here. BUT note that
	                            //       we initialize it to ‘tokens’; BUT (also note that) the loop below
	                            //       will change the value of ‘args’.
	for {
		var err error

		args, err = cliflag.Parse(&flags, args...) // <--- both ‘flags’, and ‘args’ are being set (and re-set)  here.
		if nil != err {
			if _, casted := err.(cliflag.EndOfFlags); casted { // <---- If we have reach the ‘end of flags’, then we break out of the loop.
				break
			}

			fmt.Printf("ERROR: %s\n", err)
			return
		}

		if 1 > len(args) { // <---- If there are no more ‘args’, then we break out of the loop.
			break
		}
	}


	flags.For(func(key string, value string){ // <---- We iterate over all our key-value pairs.
		fmt.Printf("KEY:   %q\n", key)
		fmt.Printf("VALUE: %q\n", value)
		fmt.Println()
	})

	fmt.Printf("args = %#v\n", args)

	// Output:
	// KEY:   "when"
	// VALUE: "now"
	//
	// KEY:   "n"
	// VALUE: "5"
	//
	// KEY:   "m"
	// VALUE: "This is a message"
	//
	// KEY:   "apple"
	// VALUE: "one"
	//
	// KEY:   "banana"
	// VALUE: "two"
	//
	// args = []string{"--", "--cherry=three", "filename.txt"}
}

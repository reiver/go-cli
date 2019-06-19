package cliflag_test

import (
	"fmt"

	"github.com/reiver/go-cli/flag"
)

func ExampleParse() {

	var tokens []string = []string{"--when=now", "-n5", "-m", "This is a message", "filename.txt"}


	var flags map[string]string = map[string]string{}

	_, err := cliflag.Parse(&flags, tokens...)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	for key, value := range flags {
		fmt.Printf("KEY:   %q\n", key)
		fmt.Printf("VALUE: %q\n", value)
	}

	// Output:
	// KEY:   "when"
	// VALUE: "now"
}

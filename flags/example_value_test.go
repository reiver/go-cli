package cliflags_test

import (
	"github.com/reiver/go-cli/flags"

	"fmt"
	"os"
)

func ExampleValue() {

	// Pretend that these are the tokens a cli.Handler receives.
	var tokens []string = []string{"-v", "--http=7777", "-m", "This is a message", "filename.txt"}


	var flags cliflags.KeyValues

	_, err := cliflags.Parse(&flags, tokens...)
	if nil != err {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		return
	}


	// We get the value of the --http flag here; if it was not set we default it to "8080".
	httpPort := flags.Load("http").ElseUnwrap("8080")

	// We get the value of the --level flag here; if it was not set we default it to "5".
	level := flags.Load("level").ElseUnwrap("5")


	fmt.Printf("HTTP port: %s\n", httpPort)
	fmt.Printf("level: %s\n", level)

	// Output:
	// HTTP port: 7777
	// level: 5
}

package cli_test

import (
	"github.com/reiver/go-cli"

	"testing"
)

func TestRun(t *testing.T) {

	var handler cli.Handler = nil

	exitcode := cli.Run(handler)

	if expected, actual := cli.ExitCodeInternalError, exitcode; expected != actual {
		t.Errorf("The actual exit code returned from Run() is not what was expected.")
		t.Logf("EXPECTED: %s", expected)
		t.Logf("ACTUAL:   %s", actual)
	}
}

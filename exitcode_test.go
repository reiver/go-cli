package cli

import (
	"fmt"

	"testing"
)

func TestExitCodeCode(t *testing.T) {

	for testNumber:=0; testNumber<=255; testNumber++ {

		var exitcode ExitCode = ExitCode(testNumber)

		if expected, actual := uint8(testNumber), exitcode.Code(); expected != actual {
			t.Errorf("For test #%d (exitcode: %q), expected %d, but actually got %d.", testNumber, exitcode.String(), expected, actual)
			continue
		}
	}
}

func TestExitCodeGoString(t *testing.T) {

	for testNumber:=0; testNumber<=255; testNumber++ {

		var exitcode ExitCode = ExitCode(testNumber)

		if expected, actual := fmt.Sprintf("cli.ExitCode(%d)", exitcode.Code()), exitcode.GoString(); expected != actual {
			t.Errorf("For test #%d (exitcode: %q), expected %q, but actually got %q.", testNumber, exitcode.String(), expected, actual)
			continue
		}
	}
}

func TestExitCodeString(t *testing.T) {

	tests := []struct{
		ExitCode ExitCode
		Expected string
	}{
		{
			ExitCode: ExitCodeOK,
			Expected: "OK",
		},
		{
			ExitCode: ExitCodeError,
			Expected: "Error",
		},

		{
			ExitCode: ExitCodeBadRequest,
			Expected: "Bad Request",
		},
		{
			ExitCode: ExitCodeBadInput,
			Expected: "Bad Input",
		},
		{
			ExitCode: ExitCodeNoInput,
			Expected: "No Input",
		},
		{
			ExitCode: ExitCodeNoUser,
			Expected: "No User",
		},
		{
			ExitCode: ExitCodeNoHost,
			Expected: "No Host",
		},
		{
			ExitCode: ExitCodeUnavailable,
			Expected: "Unavailable",
		},

		{
			ExitCode: ExitCodeInternalError,
			Expected: "Internal Error",
		},
		{
			ExitCode: ExitCodeOSError,
			Expected: "OS Error",
		},
		{
			ExitCode: ExitCodeOSFileError,
			Expected: "OS File Error",
		},
		{
			ExitCode: ExitCodeIOError,
			Expected: "I/O Error",
		},
	}

	for testNumber, test := range tests {

		if expected, actual := test.Expected, test.ExitCode.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}

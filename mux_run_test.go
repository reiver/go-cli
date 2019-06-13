package cli_test

import (
	"github.com/reiver/go-cli"

	"bytes"
	"fmt"
	"io"

	"testing"
)

func TestMuxRun(t *testing.T) {

	fn := func(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) cli.ExitCode {
		fmt.Fprintf(stdout, "%d:{", len(command))
		for i, token := range command {
			if 0 != i {
				fmt.Fprint(stdout, ";")
			}

			fmt.Fprintf(stdout, "%q", token)
		}
		fmt.Fprint(stdout, "}")

		return cli.ExitCodeOK
	}

	var mux cli.Mux

	if err := mux.HandleFunc(fn); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %q", err, err)
		return
	}
	if err := mux.HandleFunc(fn, "apple"); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %q", err, err)
		return
	}
	if err := mux.HandleFunc(fn, "apple", "banana"); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %q", err, err)
		return
	}
	if err := mux.HandleFunc(fn, "apple", "banana", "cherry"); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %q", err, err)
		return
	}

	tests := []struct{
		Command []string
		ExpectedParams []string
	}{
		{
			Command: []string{},
			ExpectedParams: []string{},
		},



		{
			Command:        []string{"-a"},
			ExpectedParams: []string{"-a"},
		},
		{
			Command:        []string{"-vvvv"},
			ExpectedParams: []string{"-vvvv"},
		},
		{
			Command:        []string{"--debug"},
			ExpectedParams: []string{"--debug"},
		},
		{
			Command:        []string{"file.txt"},
			ExpectedParams: []string{"file.txt"},
		},
		{
			Command:        []string{"/path/to/directory"},
			ExpectedParams: []string{"/path/to/directory"},
		},



		{
			Command:        []string{"-a", "-b"},
			ExpectedParams: []string{"-a", "-b"},
		},
		{
			Command:        []string{"-vvvv", "-b"},
			ExpectedParams: []string{"-vvvv", "-b"},
		},
		{
			Command:        []string{"--debug", "-b"},
			ExpectedParams: []string{"--debug", "-b"},
		},
		{
			Command:        []string{"file.txt", "-b"},
			ExpectedParams: []string{"file.txt", "-b"},
		},
		{
			Command:        []string{"/path/to/directory", "-b"},
			ExpectedParams: []string{"/path/to/directory", "-b"},
		},



		{
			Command: []string{"apple"},
			ExpectedParams: []string{},
		},
		{
			Command: []string{"apple", "-a"},
			ExpectedParams:   []string{"-a"},
		},
		{
			Command: []string{"apple", "-vvvv"},
			ExpectedParams:   []string{"-vvvv"},
		},
		{
			Command: []string{"apple", "--debug"},
			ExpectedParams:   []string{"--debug"},
		},
		{
			Command: []string{"apple", "file.txt"},
			ExpectedParams:   []string{"file.txt"},
		},
		{
			Command: []string{"apple", "/path/to/directory"},
			ExpectedParams:   []string{"/path/to/directory"},
		},



		{
			Command: []string{"apple", "-a", "-b"},
			ExpectedParams:   []string{"-a", "-b"},
		},
		{
			Command: []string{"apple", "-vvvv", "-b"},
			ExpectedParams:   []string{"-vvvv", "-b"},
		},
		{
			Command: []string{"apple", "--debug", "-b"},
			ExpectedParams:   []string{"--debug", "-b"},
		},
		{
			Command: []string{"apple", "file.txt", "-b"},
			ExpectedParams:   []string{"file.txt", "-b"},
		},
		{
			Command: []string{"apple", "/path/to/directory", "-b"},
			ExpectedParams:   []string{"/path/to/directory", "-b"},
		},



		{
			Command: []string{"apple", "banana"},
			ExpectedParams: []string{},
		},
		{
			Command: []string{"apple", "banana", "-a"},
			ExpectedParams:             []string{"-a"},
		},
		{
			Command: []string{"apple", "banana", "-vvvv"},
			ExpectedParams:             []string{"-vvvv"},
		},
		{
			Command: []string{"apple", "banana", "--debug"},
			ExpectedParams:             []string{"--debug"},
		},
		{
			Command: []string{"apple", "banana", "file.txt"},
			ExpectedParams:             []string{"file.txt"},
		},
		{
			Command: []string{"apple", "banana", "/path/to/directory"},
			ExpectedParams:             []string{"/path/to/directory"},
		},



		{
			Command: []string{"apple", "banana", "-a", "-b"},
			ExpectedParams:             []string{"-a", "-b"},
		},
		{
			Command: []string{"apple", "banana", "-vvvv", "-b"},
			ExpectedParams:             []string{"-vvvv", "-b"},
		},
		{
			Command: []string{"apple", "banana", "--debug", "-b"},
			ExpectedParams:             []string{"--debug", "-b"},
		},
		{
			Command: []string{"apple", "banana", "file.txt", "-b"},
			ExpectedParams:             []string{"file.txt", "-b"},
		},
		{
			Command: []string{"apple", "banana", "/path/to/directory", "-b"},
			ExpectedParams:             []string{"/path/to/directory", "-b"},
		},



		{
			Command: []string{"apple", "banana", "cherry"},
			ExpectedParams: []string{},
		},
		{
			Command: []string{"apple", "banana", "cherry", "-a"},
			ExpectedParams:                       []string{"-a"},
		},
		{
			Command: []string{"apple", "banana", "cherry", "-vvvv"},
			ExpectedParams:                       []string{"-vvvv"},
		},
		{
			Command: []string{"apple", "banana", "cherry", "--debug"},
			ExpectedParams:                       []string{"--debug"},
		},
		{
			Command: []string{"apple", "banana", "cherry", "file.txt"},
			ExpectedParams:                       []string{"file.txt"},
		},
		{
			Command: []string{"apple", "banana", "cherry", "/path/to/directory"},
			ExpectedParams:                       []string{"/path/to/directory"},
		},



		{
			Command: []string{"apple", "banana", "cherry", "-a", "-b"},
			ExpectedParams:                       []string{"-a", "-b"},
		},
		{
			Command: []string{"apple", "banana", "cherry", "-vvvv", "-b"},
			ExpectedParams:                       []string{"-vvvv", "-b"},
		},
		{
			Command: []string{"apple", "banana", "cherry", "--debug", "-b"},
			ExpectedParams:                       []string{"--debug", "-b"},
		},
		{
			Command: []string{"apple", "banana", "cherry", "file.txt", "-b"},
			ExpectedParams:                       []string{"file.txt", "-b"},
		},
		{
			Command: []string{"apple", "banana", "cherry", "/path/to/directory", "-b"},
			ExpectedParams:                       []string{"/path/to/directory", "-b"},
		},



		{
			Command: []string{"apple", "banana", "what"},
			ExpectedParams:             []string{"what",},
		},
		{
			Command: []string{"apple", "banana", "what", "-a"},
			ExpectedParams:             []string{"what", "-a"},
		},
		{
			Command: []string{"apple", "banana", "what", "-vvvv"},
			ExpectedParams:             []string{"what", "-vvvv"},
		},
		{
			Command: []string{"apple", "banana", "what", "--debug"},
			ExpectedParams:             []string{"what", "--debug"},
		},
		{
			Command: []string{"apple", "banana", "what", "file.txt"},
			ExpectedParams:             []string{"what", "file.txt"},
		},
		{
			Command: []string{"apple", "banana", "what", "/path/to/directory"},
			ExpectedParams:             []string{"what", "/path/to/directory"},
		},



		{
			Command: []string{"apple", "banana", "what", "-a", "-b"},
			ExpectedParams:             []string{"what", "-a", "-b"},
		},
		{
			Command: []string{"apple", "banana", "what", "-vvvv", "-b"},
			ExpectedParams:             []string{"what", "-vvvv", "-b"},
		},
		{
			Command: []string{"apple", "banana", "what", "--debug", "-b"},
			ExpectedParams:             []string{"what", "--debug", "-b"},
		},
		{
			Command: []string{"apple", "banana", "what", "file.txt", "-b"},
			ExpectedParams:             []string{"what", "file.txt", "-b"},
		},
		{
			Command: []string{"apple", "banana", "what", "/path/to/directory", "-b"},
			ExpectedParams:             []string{"what", "/path/to/directory", "-b"},
		},
	}

	for testNumber, test := range tests {

		var stdout writeCloseBuffer

		if expected, actual := cli.ExitCodeOK, mux.Run(fakeStdin, &stdout, fakeStderr, test.Command...); expected != actual {
			t.Errorf("For test #%d, expected exit code %q, but actually got %q.", testNumber, expected, actual)
			continue
		}

		{
			var expectedBuffer bytes.Buffer
			fmt.Fprintf(&expectedBuffer, "%d:{", len(test.ExpectedParams))
			for i, token := range test.ExpectedParams {
				if 0 != i {
					fmt.Fprint(&expectedBuffer, ";")
				}

				fmt.Fprintf(&expectedBuffer, "%q", token)
			}
			fmt.Fprint(&expectedBuffer, "}")

			if expected, actual := expectedBuffer.String(), stdout.String(); expected != actual {
				t.Errorf("For #%d...", testNumber)
				t.Errorf("\tEXPECTED: %q", expected)
				t.Errorf("\tACTUAL:   %q", actual)
			}
		}
	}
}

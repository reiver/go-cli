package cli_test

import (
	"github.com/reiver/go-cli"

	"reflect"
	"io"

	"testing"
)

func TestMuxCommands(t *testing.T) {

	tests := []struct{
		Commands [][]string
		Filter []string
		Expected [][]string
	}{
		{
			Commands: [][]string{
				[]string{"apple"},
				[]string{"apple","banana"},
				[]string{"apple","banana","cherry"},
				[]string{"hi"},
				[]string{"hello"},
				[]string{"hello","people"},
			},
			Filter: []string{},
			Expected: [][]string{
				[]string{"apple"},
				[]string{"apple","banana"},
				[]string{"apple","banana","cherry"},
				[]string{"hello"},
				[]string{"hello","people"},
				[]string{"hi"},
			},
		},

		{
			Commands: [][]string{
				[]string{"apple"},
				[]string{"apple","banana"},
				[]string{"apple","banana","cherry"},
				[]string{"hi"},
				[]string{"hello"},
				[]string{"hello","people"},
			},
			Filter: []string{"h"},
			Expected: [][]string{
				[]string{"hello"},
				[]string{"hello","people"},
				[]string{"hi"},
			},
		},

		{
			Commands: [][]string{
				[]string{"apple"},
				[]string{"apple","banana"},
				[]string{"apple","banana","cherry"},
				[]string{"hi"},
				[]string{"hello"},
				[]string{"hello","people"},
			},
			Filter: []string{"he"},
			Expected: [][]string{
				[]string{"hello"},
				[]string{"hello","people"},
			},
		},

		{
			Commands: [][]string{
				[]string{"apple"},
				[]string{"apple","banana"},
				[]string{"apple","banana","cherry"},
				[]string{"hi"},
				[]string{"hello"},
				[]string{"hello","people"},
			},
			Filter: []string{"hea"},
			Expected: [][]string{},
		},

		{
			Commands: [][]string{
				[]string{"apple"},
				[]string{"apple","banana"},
				[]string{"apple","banana","cherry"},
				[]string{"hi"},
				[]string{"hello"},
				[]string{"hello","people"},
			},
			Filter: []string{"heat"},
			Expected: [][]string{},
		},

		{
			Commands: [][]string{
				[]string{"apple"},
				[]string{"apple","banana"},
				[]string{"apple","banana","cherry"},
				[]string{"hi"},
				[]string{"hello"},
				[]string{"hello","people"},
			},
			Filter: []string{"b"},
			Expected: [][]string{},
		},

		{
			Commands: [][]string{
				[]string{"apple"},
				[]string{"apple","banana"},
				[]string{"apple","banana","cherry"},
				[]string{"hi"},
				[]string{"hello"},
				[]string{"hello","people"},
			},
			Filter: []string{"a"},
			Expected: [][]string{
				[]string{"apple"},
				[]string{"apple","banana"},
				[]string{"apple","banana","cherry"},
			},
		},
		{
			Commands: [][]string{
				[]string{"apple"},
				[]string{"apple","banana"},
				[]string{"apple","banana","cherry"},
				[]string{"hi"},
				[]string{"hello"},
				[]string{"hello","people"},
			},
			Filter: []string{"ap"},
			Expected: [][]string{
				[]string{"apple"},
				[]string{"apple","banana"},
				[]string{"apple","banana","cherry"},
			},
		},
		{
			Commands: [][]string{
				[]string{"apple"},
				[]string{"apple","banana"},
				[]string{"apple","banana","cherry"},
				[]string{"hi"},
				[]string{"hello"},
				[]string{"hello","people"},
			},
			Filter: []string{"app"},
			Expected: [][]string{
				[]string{"apple"},
				[]string{"apple","banana"},
				[]string{"apple","banana","cherry"},
			},
		},
		{
			Commands: [][]string{
				[]string{"apple"},
				[]string{"apple","banana"},
				[]string{"apple","banana","cherry"},
				[]string{"hi"},
				[]string{"hello"},
				[]string{"hello","people"},
			},
			Filter: []string{"app2"},
			Expected: [][]string{},
		},
	}

	Loop: for testNumber, test := range tests {
		var mux cli.Mux

		for _, command := range test.Commands {
			handlerFunc := func(io.ReadCloser, io.WriteCloser, io.WriteCloser, ...string) cli.ExitCode {
				return cli.ExitCodeOK
			}

			if err := mux.HandleFunc(handlerFunc, command...); nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
				continue Loop
			}
		}

		commands, err := mux.Commands(test.Filter...)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, commands; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d...", testNumber)
			t.Errorf("\tCOMMANDS:...")
			for _, command := range test.Commands {
				t.Errorf("\t\t%#v", command)
			}
			t.Errorf("\tFILTER:...")
			t.Errorf("\t\t%#v", test.Filter)
			t.Errorf("\tEXPECTED:...")
			for _, command := range expected {
				t.Errorf("\t\t%#v", command)
			}
			t.Errorf("\tACTUAL:...")
			for _, command := range actual {
				t.Errorf("\t\t%#v", command)
			}
			continue
		}
	}
}

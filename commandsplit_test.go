package cli

import (
	"reflect"

	"testing"
)

func TestCommandSplit(t *testing.T) {

	tests := []struct{
		Command              []string
		ExpectedCommandName  []string
		ExpectedCommandParams []string
	}{
		{
			Command:             []string{},
			ExpectedCommandName: []string(nil),
			ExpectedCommandParams: []string(nil),
		},



		{
			Command:             []string{"apple"},
			ExpectedCommandName: []string{"apple"},
			ExpectedCommandParams: []string(nil),
		},
		{
			Command:             []string{"apple", "--debug"},
			ExpectedCommandName: []string{"apple"},
			ExpectedCommandParams:        []string{"--debug"},
		},
		{
			Command:             []string{"apple", "-a"},
			ExpectedCommandName: []string{"apple"},
			ExpectedCommandParams:        []string{"-a"},
		},
		{
			Command:             []string{"apple", "--http=8080", "-b"},
			ExpectedCommandName: []string{"apple"},
			ExpectedCommandParams:        []string{"--http=8080", "-b"},
		},



		{
			Command:             []string{"apple", "banana"},
			ExpectedCommandName: []string{"apple", "banana"},
			ExpectedCommandParams: []string(nil),
		},
		{
			Command:             []string{"apple", "banana", "--debug"},
			ExpectedCommandName: []string{"apple", "banana"},
			ExpectedCommandParams:                  []string{"--debug"},
		},
		{
			Command:             []string{"apple", "banana", "-a"},
			ExpectedCommandName: []string{"apple", "banana"},
			ExpectedCommandParams:                  []string{"-a"},
		},
		{
			Command:             []string{"apple", "banana", "--http=8080", "-b"},
			ExpectedCommandName: []string{"apple", "banana"},
			ExpectedCommandParams:                  []string{"--http=8080", "-b"},
		},



		{
			Command:             []string{"apple", "banana", "cherry"},
			ExpectedCommandName: []string{"apple", "banana", "cherry"},
			ExpectedCommandParams: []string(nil),
		},
		{
			Command:             []string{"apple", "banana", "cherry", "--debug"},
			ExpectedCommandName: []string{"apple", "banana", "cherry"},
			ExpectedCommandParams:                            []string{"--debug"},
		},
		{
			Command:             []string{"apple", "banana", "cherry", "-a"},
			ExpectedCommandName: []string{"apple", "banana", "cherry"},
			ExpectedCommandParams:                            []string{"-a"},
		},
		{
			Command:             []string{"apple", "banana", "cherry", "--http=8080", "-b"},
			ExpectedCommandName: []string{"apple", "banana", "cherry"},
			ExpectedCommandParams:                            []string{"--http=8080", "-b"},
		},



		{
			Command:             []string{"hello", "world", "file.txt"},
			ExpectedCommandName: []string{"hello", "world", "file.txt"},
			ExpectedCommandParams: []string(nil),
		},
		{
			Command:             []string{"hello", "world", "file.txt", "--debug"},
			ExpectedCommandName: []string{"hello", "world", "file.txt"},
			ExpectedCommandParams:                             []string{"--debug"},
		},
		{
			Command:             []string{"hello", "world", "file.txt", "-a"},
			ExpectedCommandName: []string{"hello", "world", "file.txt"},
			ExpectedCommandParams:                             []string{"-a"},
		},
		{
			Command:             []string{"hello", "world", "file.txt", "--http=8080", "-b"},
			ExpectedCommandName: []string{"hello", "world", "file.txt"},
			ExpectedCommandParams:                             []string{"--http=8080", "-b"},
		},



		{
			Command:             []string{"one", "two", "three", "/path/to/directory"},
			ExpectedCommandName: []string{"one", "two", "three", "/path/to/directory"},
			ExpectedCommandParams: []string(nil),
		},
		{
			Command:             []string{"one", "two", "three", "/path/to/directory", "--debug"},
			ExpectedCommandName: []string{"one", "two", "three", "/path/to/directory"},
			ExpectedCommandParams:                                            []string{"--debug"},
		},
		{
			Command:             []string{"one", "two", "three", "/path/to/directory", "-a"},
			ExpectedCommandName: []string{"one", "two", "three", "/path/to/directory"},
			ExpectedCommandParams:                                            []string{"-a"},
		},
		{
			Command:             []string{"one", "two", "three", "/path/to/directory", "--http=8080", "-b"},
			ExpectedCommandName: []string{"one", "two", "three", "/path/to/directory"},
			ExpectedCommandParams:                                            []string{"--http=8080", "-b"},
		},
	}

	for testNumber, test := range tests {

		actualName, actualParams := commandSplit(test.Command...)

		if expected, actual := test.ExpectedCommandName, actualName; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d...", testNumber)
			t.Errorf("\tCOMMAND:  %#v", test.Command)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}

		if expected, actual := test.ExpectedCommandParams, actualParams; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d...", testNumber)
			t.Errorf("\tCOMMAND:  %#v", test.Command)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}
	}
}

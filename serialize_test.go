package cli

import (
	"testing"
)

func TestSerialize(t *testing.T) {

	tests := []struct{
		Command            []string
		Expected             string
	}{
		{
			Command:  []string(nil),
			Expected: "",
		},
		{
			Command:  []string{},
			Expected: "",
		},



		{
			Command:  []string{"\x1B"},
			Expected: "\x1B\x1B",
		},
		{
			Command:  []string{"\x1B\x1B"},
			Expected: "\x1B\x1B\x1B\x1B",
		},
		{
			Command:  []string{"\x1B\x1B\x1B"},
			Expected: "\x1B\x1B\x1B\x1B\x1B\x1B",
		},
		{
			Command:  []string{"\x1B\x1B\x1B\x1B"},
			Expected: "\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B",
		},



		{
			Command:  []string{"\x1B", "\x1B"},
			Expected: "\x1B\x1B"+"\u001F"+"\x1B\x1B",
		},
		{
			Command:  []string{"\x1B\x1B", "\x1B"},
			Expected: "\x1B\x1B\x1B\x1B"+"\u001F"+"\x1B\x1B",
		},
		{
			Command:  []string{"\x1B\x1B\x1B", "\x1B"},
			Expected: "\x1B\x1B\x1B\x1B\x1B\x1B"+"\u001F"+"\x1B\x1B",
		},
		{
			Command:  []string{"\x1B\x1B\x1B\x1B", "\x1B"},
			Expected: "\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B"+"\u001F"+"\x1B\x1B",
		},



		{
			Command:  []string{"\x1B", "\x1B\x1B"},
			Expected: "\x1B\x1B"+"\u001F"+"\x1B\x1B\x1B\x1B",
		},
		{
			Command:  []string{"\x1B\x1B", "\x1B\x1B"},
			Expected: "\x1B\x1B\x1B\x1B"+"\u001F"+"\x1B\x1B\x1B\x1B",
		},
		{
			Command:  []string{"\x1B\x1B\x1B", "\x1B\x1B"},
			Expected: "\x1B\x1B\x1B\x1B\x1B\x1B"+"\u001F"+"\x1B\x1B\x1B\x1B",
		},
		{
			Command:  []string{"\x1B\x1B\x1B\x1B", "\x1B\x1B"},
			Expected: "\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B"+"\u001F"+"\x1B\x1B\x1B\x1B",
		},



		{
			Command:  []string{"\x1B", "\x1B\x1B\x1B"},
			Expected: "\x1B\x1B"+"\u001F"+"\x1B\x1B\x1B\x1B\x1B\x1B",
		},
		{
			Command:  []string{"\x1B\x1B", "\x1B\x1B\x1B"},
			Expected: "\x1B\x1B\x1B\x1B"+"\u001F"+"\x1B\x1B\x1B\x1B\x1B\x1B",
		},
		{
			Command:  []string{"\x1B\x1B\x1B", "\x1B\x1B\x1B"},
			Expected: "\x1B\x1B\x1B\x1B\x1B\x1B"+"\u001F"+"\x1B\x1B\x1B\x1B\x1B\x1B",
		},
		{
			Command:  []string{"\x1B\x1B\x1B\x1B", "\x1B\x1B\x1B"},
			Expected: "\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B"+"\u001F"+"\x1B\x1B\x1B\x1B\x1B\x1B",
		},



		{
			Command:  []string{"\x1B", "\x1B\x1B\x1B\x1B"},
			Expected: "\x1B\x1B"+"\u001F"+"\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B",
		},
		{
			Command:  []string{"\x1B\x1B", "\x1B\x1B\x1B\x1B"},
			Expected: "\x1B\x1B\x1B\x1B"+"\u001F"+"\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B",
		},
		{
			Command:  []string{"\x1B\x1B\x1B", "\x1B\x1B\x1B\x1B"},
			Expected: "\x1B\x1B\x1B\x1B\x1B\x1B"+"\u001F"+"\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B",
		},
		{
			Command:  []string{"\x1B\x1B\x1B\x1B", "\x1B\x1B\x1B\x1B"},
			Expected: "\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B"+"\u001F"+"\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B",
		},



		{
			Command:  []string{"\x1B", "\x1B\x1B", "\x1B\x1B\x1B", "\x1B\x1B\x1B\x1B"},
			Expected: "\x1B\x1B"+"\u001F"+"\x1B\x1B\x1B\x1B"+"\u001F"+"\x1B\x1B\x1B\x1B\x1B\x1B"+"\u001F"+"\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B",
		},
		{
			Command:  []string{"\x1B\x1B\x1B\x1B", "\x1B\x1B\x1B", "\x1B\x1B", "\x1B"},
			Expected: "\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B"+"\u001F"+"\x1B\x1B\x1B\x1B\x1B\x1B"+"\u001F"+"\x1B\x1B\x1B\x1B"+"\u001F"+"\x1B\x1B",
		},



		{
			Command:  []string{"\u001F"},
			Expected: "\x1B\u001F",
		},



		{
			Command:  []string{"apple"},
			Expected: "apple",
		},
		{
			Command:  []string{"apple", "banana"},
			Expected: "apple\u001Fbanana",
		},
		{
			Command:  []string{"apple", "banana", "cherry"},
			Expected: "apple\u001Fbanana\u001Fcherry",
		},




		{
			Command:  []string{"hello"},
			Expected: "hello",
		},
		{
			Command:  []string{"hello", "world"},
			Expected: "hello\u001Fworld",
		},



		{
			Command:  []string{"خُدا"},
			Expected: "خُدا",
		},
		{
			Command:  []string{"خُدا","حافِظ"},
			Expected: "خُدا\u001Fحافِظ",
		},



		{
			Command:  []string{"😀"},
			Expected: "😀",
		},
		{
			Command:  []string{"😀", "😁"},
			Expected: "😀\u001F😁",
		},
		{
			Command:  []string{"😀", "😁", "😆"},
			Expected: "😀\u001F😁\u001F😆",
		},
		{
			Command:  []string{"😀", "😁", "😆", "😅"},
			Expected: "😀\u001F😁\u001F😆\u001F😅",
		},
		{
			Command:  []string{"😀", "😁", "😆", "😅", "😲"},
			Expected: "😀\u001F😁\u001F😆\u001F😅\u001F😲",
		},



		{
			Command:  []string{"🐸🐒"},
			Expected: "🐸🐒",
		},
		{
			Command:  []string{"🐸🐒", "🍉🍑🍌"},
			Expected: "🐸🐒\u001F🍉🍑🍌",
		},
		{
			Command:  []string{"🐸🐒", "🍉🍑🍌", "👽👻👾😉"},
			Expected: "🐸🐒\u001F🍉🍑🍌\u001F👽👻👾😉",
		},



		{
			Command:  []string{"help", "init"},
			Expected: "help\u001Finit",
		},
		{
			Command:  []string{"help", "init", "repo"},
			Expected: "help\u001Finit\u001Frepo",
		},
		{
			Command:  []string{"help", "backup"},
			Expected: "help\u001Fbackup",
		},
		{
			Command:  []string{"help", "launch"},
			Expected: "help\u001Flaunch",
		},



		{
			Command:  []string{"help", "init", "-vvvv"},
			Expected: "help\u001Finit\u001F-vvvv",
		},
		{
			Command:  []string{"help", "init", "repo", "-vvvv"},
			Expected: "help\u001Finit\u001Frepo\u001F-vvvv",
		},
		{
			Command:  []string{"help", "backup", "-vvvv"},
			Expected: "help\u001Fbackup\u001F-vvvv",
		},
		{
			Command:  []string{"help", "launch", "-vvvv"},
			Expected: "help\u001Flaunch\u001F-vvvv",
		},



		{
			Command:  []string{"help", "init", "-vvvv"},
			Expected: "help\u001Finit\u001F-vvvv",
		},
		{
			Command:  []string{"help", "init", "repo", "--path", "/path/to/where/i/want/it", "-vvvv"},
			Expected: "help\u001Finit\u001Frepo\u001F--path\u001F/path/to/where/i/want/it\u001F-vvvv",
		},
		{
			Command:  []string{"help", "backup", "-a", "-vvvv"},
			Expected: "help\u001Fbackup\u001F-a\u001F-vvvv",
		},
		{
			Command:  []string{"help", "launch", "--http=8080", "-vvvv"},
			Expected: "help\u001Flaunch\u001F--http=8080\u001F-vvvv",
		},
	}

	for testNumber, test := range tests {
		actual, err := serialize(test.Command...)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d...", testNumber)
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}
	}
}

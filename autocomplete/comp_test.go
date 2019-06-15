package cliautocomplete

import (
	"github.com/reiver/go-cli/autocomplete/internal/optint64"
	"github.com/reiver/go-cli/autocomplete/internal/optstring"

	"testing"
)

func TestCompParse(t *testing.T) {

	tests := []struct{
		Data []string
		Expected_COMP_LINE  optstring.Type
		Expected_COMP_POINT optint64.Type
		Expected_COMP_TYPE  optstring.Type
		Expected_COMP_KEY   optstring.Type
	}{
		{ // 0
			Data: []string{},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.None(),
		},



		{ // 1
			Data: []string{"COMP"},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.None(),
		},
		{ // 2
			Data: []string{"COMP_"},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.None(),
		},



		{ // 3
			Data: []string{"COMP=hello"},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.None(),
		},
		{ // 4
			Data: []string{"COMP_=hello"},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.None(),
		},



		{ // 5
			Data: []string{"COMP_LINE"},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.None(),
		},
		{ // 6
			Data: []string{"COMP_POINT"},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.None(),
		},
		{ // 7
			Data: []string{"COMP_TYPE"},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.None(),
		},
		{ // 8
			Data: []string{"COMP_KEY"},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.None(),
		},



		{ // 9
			Data: []string{"COMP_LINE="},
			Expected_COMP_LINE: optstring.Some(""),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.None(),
		},
		{ // 10
			Data: []string{"COMP_POINT="},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.None(),
		},
		{ // 11
			Data: []string{"COMP_TYPE="},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.Some(""),
			Expected_COMP_KEY:  optstring.None(),
		},
		{ // 12
			Data: []string{"COMP_KEY="},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.Some(""),
		},



		{ // 13
			Data: []string{"COMP_LINEapple"},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.None(),
		},
		{ // 14
			Data: []string{"COMP_POINT123"},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.Some(123),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.None(),
		},
		{ // 15
			Data: []string{"COMP_TYPEbanana"},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.None(),
		},
		{ // 16
			Data: []string{"COMP_KEYcherry"},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.None(),
		},



		{ // 17
			Data: []string{"COMP_LINE=apple"},
			Expected_COMP_LINE: optstring.Some("apple"),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.None(),
		},
		{ // 18
			Data: []string{"COMP_POINT=123"},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.Some(123),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.None(),
		},
		{ // 19
			Data: []string{"COMP_TYPE=banana"},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.Some("banana"),
			Expected_COMP_KEY:  optstring.None(),
		},
		{ // 20
			Data: []string{"COMP_KEY=cherry"},
			Expected_COMP_LINE: optstring.None(),
			Expected_COMP_POINT: optint64.None(),
			Expected_COMP_TYPE: optstring.None(),
			Expected_COMP_KEY:  optstring.Some("cherry"),
		},
	}

	for testNumber, test := range tests {

		result := compParse(test.Data)

		if expected, actual := test.Expected_COMP_LINE, result.COMP_LINE; expected != actual {
			t.Errorf("For test #%d (COMP_LINE)...", testNumber)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}

		if expected, actual := test.Expected_COMP_TYPE, result.COMP_TYPE; expected != actual {
			t.Errorf("For test #%d (COMP_TYPE)...", testNumber)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}

		if expected, actual := test.Expected_COMP_KEY, result.COMP_KEY; expected != actual {
			t.Errorf("For test #%d (COMP_KEY)...", testNumber)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}
	}
}

package cliflags

import (
	"reflect"

	"testing"
)

func TestStorerMapStringString(t *testing.T) {

	tests := []struct{
		Tokens []string
		ExpectedKey   string
		ExpectedValue string
		ExpectedRemainingTokens []string
	}{
		{
			Tokens:        []string{"--size", "-m", "This is a message", "filename.txt"},
			ExpectedKey:              "size",
			ExpectedValue:                 "",
			ExpectedRemainingTokens: []string{"-m", "This is a message", "filename.txt"},
		},



		{
			Tokens:       []string{"--size=", "-m", "This is a message", "filename.txt"},
			ExpectedKey:             "size",
			ExpectedValue:                "",
			ExpectedRemainingTokens: []string{"-m", "This is a message", "filename.txt"},
		},



		{
			Tokens:       []string{"--size:", "-m", "This is a message", "filename.txt"},
			ExpectedKey:             "size",
			ExpectedValue:                "",
			ExpectedRemainingTokens: []string{"-m", "This is a message", "filename.txt"},
		},



		{
			Tokens:  []string{"--size=hello", "-m", "This is a message", "filename.txt"},
			ExpectedKey:        "size",
			ExpectedValue:           "hello",
			ExpectedRemainingTokens: []string{"-m", "This is a message", "filename.txt"},
		},



		{
			Tokens:  []string{"--size:hello", "-m", "This is a message", "filename.txt"},
			ExpectedKey:        "size",
			ExpectedValue:           "hello",
			ExpectedRemainingTokens: []string{"-m", "This is a message", "filename.txt"},
		},



		{
			Tokens: []string{"-m", "This is a message", "filename.txt"},
			ExpectedKey:     "m",
			ExpectedValue:         "This is a message",
			ExpectedRemainingTokens:           []string{"filename.txt"},
		},



		{
			Tokens:                  []string{"-m"},
			ExpectedKey:                       "m",
			ExpectedValue:                          "",
			ExpectedRemainingTokens: []string{},
		},



		{
			Tokens:                  []string{"-m", "-t"},
			ExpectedKey:                       "m",
			ExpectedValue:                          "",
			ExpectedRemainingTokens:       []string{"-t"},
		},



		{
			Tokens:                  []string{"-m", "--wow=ok"},
			ExpectedKey:                       "m",
			ExpectedValue:                          "",
			ExpectedRemainingTokens:       []string{"--wow=ok"},
		},
	}

	for testNumber, test := range tests {

		var data map[string]string
		if 0 != len(data) {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tlen(data) == %d", len(data))
			t.Errorf("\tdata = %#v", data)
			continue
		}

		var storer Storer = &storerMapStringString{
			p: &data,
		}

		remainingTokens, err := Parse(storer, test.Tokens...)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		if 1 != len(data) {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tlen(data) == %d", len(data))
			t.Errorf("\tdata = %#v", data)
			continue
		}

		if expected, actual := test.ExpectedRemainingTokens, remainingTokens; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tEXPECTED remaining tokens: %#v", expected)
			t.Errorf("\tACTUAL   remaining tokens: %#v", actual)
			continue
		}

		{
			_, found := data[test.ExpectedKey]
			if !found {
				t.Errorf("For test #%d ...", testNumber)
				t.Errorf("\tEXPECTED key: %#v", test.ExpectedKey)
				t.Errorf("\tdata = %#v", data)
				continue
			}
		}

		{
			actualValue := data[test.ExpectedKey]
			if expected, actual := test.ExpectedValue, actualValue; expected != actual {
				t.Errorf("For test #%d ...", testNumber)
				t.Errorf("\tEXPECTED value: %#v", expected)
				t.Errorf("\tACTUAL   value: %#v", actual)
				continue
			}
		}
	}
}

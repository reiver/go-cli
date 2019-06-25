package cliflags

import (
	"reflect"

	"testing"
)

func TestStorerMapStringString(t *testing.T) {

	tests := []struct{
		Tokens []string
		ExpectedFlags map[string]string
		ExpectedArgs []string
	}{
		{
			Tokens: []string{"--size", "-m", "This is a message", "filename.txt"},
			ExpectedFlags: map[string]string{
			                   "size":"",
			                            "m": "This is a message",
			},
			ExpectedArgs:                                []string{"filename.txt"},
		},



		{
			Tokens: []string{"--size=", "-m", "This is a message", "filename.txt"},
			ExpectedFlags: map[string]string{
			                   "size":"",
			                             "m": "This is a message",
			},
			ExpectedArgs:                                 []string{"filename.txt"},
		},



		{
			Tokens: []string{"--size:", "-m", "This is a message", "filename.txt"},
			ExpectedFlags: map[string]string{
			                   "size":"",
			                             "m": "This is a message",
			},
			ExpectedArgs:                                 []string{"filename.txt"},
		},



		{
			Tokens: []string{"--size=hello", "-m", "This is a message", "filename.txt"},
			ExpectedFlags: map[string]string{
			                   "size":"hello",
			                                  "m": "This is a message",
			},
			ExpectedArgs:                                      []string{"filename.txt"},
		},



		{
			Tokens: []string{"--size:hello", "-m", "This is a message", "filename.txt"},
			ExpectedFlags: map[string]string{
			                   "size":"hello",
			                                  "m":"This is a message",
			},
			ExpectedArgs:                                      []string{"filename.txt"},
		},



		{
			Tokens: []string{"-m", "This is a message", "filename.txt"},
			ExpectedFlags: map[string]string{
			                  "m": "This is a message",
			},
			ExpectedArgs:                      []string{"filename.txt"},
		},



		{
			Tokens: []string{"-m"},
			ExpectedFlags: map[string]string{
			                  "m": "",
			},
			ExpectedArgs:            []string{},
		},



		{
			Tokens: []string{"-m", "-t"},
			ExpectedFlags: map[string]string{
			                  "m":"",
			                        "t":"",
			},
			ExpectedArgs:            []string{},
		},



		{
			Tokens: []string{"-m", "--wow=ok"},
			ExpectedFlags: map[string]string{
			                  "m":"",
			                         "wow":"ok",
			},
			ExpectedArgs:            []string{},
		},









		{
			Tokens: []string{"--"},
			ExpectedFlags: map[string]string(nil),
			ExpectedArgs: []string{},
		},
		{
			Tokens: []string{"--", "--"},
			ExpectedFlags: map[string]string(nil),
			ExpectedArgs: []string{"--"},
		},
		{
			Tokens: []string{"--", "--", "--"},
			ExpectedFlags: map[string]string(nil),
			ExpectedArgs: []string{"--", "--"},
		},
		{
			Tokens: []string{"--", "--", "--", "--"},
			ExpectedFlags: map[string]string(nil),
			ExpectedArgs: []string{"--", "--", "--"},
		},
		{
			Tokens: []string{"--", "--", "--", "--", "--"},
			ExpectedFlags: map[string]string(nil),
			ExpectedArgs: []string{"--", "--", "--", "--"},
		},



		{
			Tokens: []string{"--apple", "--banana=two", "--", "--cherry=three", "filename.txt"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			                              "banana":"two",
			},
			ExpectedArgs:                            []string{"--cherry=three", "filename.txt"},
		},



		{
			Tokens: []string{"-v", "-m", "This is a message", "-n" ,"--", "-3", "-2+1", "--", "func"},
			ExpectedFlags: map[string]string{
			                  "v":"",
			                        "m": "This is a message",
			                                                   "n":       "-3",
			},
			ExpectedArgs:                                              []string{"-2+1", "--", "func"},
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

		actualArgs, err := Parse(storer, test.Tokens...)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected, actual := len(test.ExpectedFlags), len(data); expected != actual {
			t.Errorf("For test #%d, expected %d flags, but actually got %d flags.", testNumber, expected, actual)
			t.Errorf("\tTOKENS: %#v", test.Tokens)
			t.Errorf("\tEXPECTED FLAGS: %#v", test.ExpectedFlags)
			t.Errorf("\tACTUAL   FLAGS: %#v", data)
			t.Errorf("\tEXPECTED ARGS: %#v", test.ExpectedArgs)
			t.Errorf("\tACTUAL   ARGS: %#v", actualArgs)
			continue
		}

		if expected, actual := test.ExpectedFlags, data; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, expected certain flags, but actually got different flags.", testNumber)
			t.Errorf("\tTOKENS: %#v", test.Tokens)
			t.Errorf("\tEXPECTED FLAGS: %#v", test.ExpectedFlags)
			t.Errorf("\tACTUAL   FLAGS: %#v", data)
			t.Errorf("\tEXPECTED ARGS: %#v", test.ExpectedArgs)
			t.Errorf("\tACTUAL   ARGS: %#v", actualArgs)
			continue
		}

		if expected, actual := len(test.ExpectedArgs), len(actualArgs); expected != actual {
			t.Errorf("For test #%d, expected %d args, but actually got %d args.", testNumber, expected, actual)
			t.Errorf("\tTOKENS: %#v", test.Tokens)
			t.Errorf("\tEXPECTED FLAGS: %#v", test.ExpectedFlags)
			t.Errorf("\tACTUAL   FLAGS: %#v", data)
			t.Errorf("\tEXPECTED ARGS: %#v", test.ExpectedArgs)
			t.Errorf("\tACTUAL   ARGS: %#v", actualArgs)
			continue
		}

		if expected, actual := test.ExpectedArgs, actualArgs; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, expected certain args, but actually got different args.", testNumber)
			t.Errorf("\tTOKENS: %#v", test.Tokens)
			t.Errorf("\tEXPECTED FLAGS: %#v", test.ExpectedFlags)
			t.Errorf("\tACTUAL   FLAGS: %#v", data)
			t.Errorf("\tEXPECTED ARGS: %#v", test.ExpectedArgs)
			t.Errorf("\tACTUAL   ARGS: %#v", actualArgs)
			continue
		}
	}
}

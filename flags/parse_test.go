package cliflags_test

import (
	"github.com/reiver/go-cli/flags"

	"reflect"

	"testing"
)

func TestParse(t *testing.T) {

	tests := []struct{
		Tokens []string
		ExpectedFlags map[string]string
		ExpectedArgs []string
	}{
		{
			Tokens: []string(nil),
			ExpectedFlags: map[string]string{},
			ExpectedArgs: []string(nil),
		},
		{
			Tokens: []string{},
			ExpectedFlags: map[string]string{},
			ExpectedArgs: []string{},
		},



		{
			Tokens: []string{"--"},
			ExpectedFlags: map[string]string{},
			ExpectedArgs: []string{},
		},
		{
			Tokens: []string{"--", "--"},
			ExpectedFlags: map[string]string{},
			ExpectedArgs: []string{"--"},
		},
		{
			Tokens: []string{"--", "--", "--"},
			ExpectedFlags: map[string]string{},
			ExpectedArgs: []string{"--", "--"},
		},
		{
			Tokens: []string{"--", "--", "--", "--"},
			ExpectedFlags: map[string]string{},
			ExpectedArgs: []string{"--", "--", "--"},
		},



		{
			Tokens: []string{"--", "-m", "-n", "-o", "filename.txt"},
			ExpectedFlags: map[string]string{},
			ExpectedArgs: []string{"-m", "-n", "-o", "filename.txt"},
		},



		{
			Tokens: []string{"--apple"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			},
			ExpectedArgs: []string{},
		},
		{
			Tokens: []string{"--apple", "--banana"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			                            "banana":"",
			},
			ExpectedArgs: []string{},
		},
		{
			Tokens: []string{"--apple", "--banana", "--cherry"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			                            "banana":"",
			                                          "cherry":"",
			},
			ExpectedArgs: []string{},
		},



		{
			Tokens: []string{"--apple", "--banana", "--cherry", "one"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			                            "banana":"",
			                                          "cherry":"",
			},
			ExpectedArgs:                              []string{"one"},
		},
		{
			Tokens: []string{"--apple", "--banana", "--cherry", "one", "two"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			                            "banana":"",
			                                          "cherry":"",
			},
			ExpectedArgs:                              []string{"one", "two"},
		},
		{
			Tokens: []string{"--apple", "--banana", "--cherry", "one", "two", "three"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			                            "banana":"",
			                                          "cherry":"",
			},
			ExpectedArgs:                              []string{"one", "two", "three"},
		},



		{
			Tokens: []string{"--", "--apple", "--banana", "--cherry", "one"},
			ExpectedFlags: map[string]string{},
			ExpectedArgs: []string{"--apple", "--banana", "--cherry", "one"},
		},
		{
			Tokens: []string{"--", "--apple", "--banana", "--cherry", "one", "two"},
			ExpectedFlags: map[string]string{},
			ExpectedArgs: []string{"--apple", "--banana", "--cherry", "one", "two"},
		},
		{
			Tokens: []string{"--", "--apple", "--banana", "--cherry", "one", "two", "three"},
			ExpectedFlags: map[string]string{},
			ExpectedArgs: []string{"--apple", "--banana", "--cherry", "one", "two", "three"},
		},



		{
			Tokens: []string{"--apple", "--", "--banana", "--cherry", "one"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			},
			ExpectedArgs:            []string{"--banana", "--cherry", "one"},
		},
		{
			Tokens: []string{"--apple", "--", "--banana", "--cherry", "one", "two"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			},
			ExpectedArgs:            []string{"--banana", "--cherry", "one", "two"},
		},
		{
			Tokens: []string{"--apple", "--", "--banana", "--cherry", "one", "two", "three"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			},
			ExpectedArgs:            []string{"--banana", "--cherry", "one", "two", "three"},
		},



		{
			Tokens: []string{"--apple", "--banana", "--", "--cherry", "one"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			                              "banana":"",
			},
			ExpectedArgs:                        []string{"--cherry", "one"},
		},
		{
			Tokens: []string{"--apple", "--banana", "--", "--cherry", "one", "two"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			                              "banana":"",
			},
			ExpectedArgs:                        []string{"--cherry", "one", "two"},
		},
		{
			Tokens: []string{"--apple", "--banana", "--", "--cherry", "one", "two", "three"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			                              "banana":"",
			},
			ExpectedArgs:                        []string{"--cherry", "one", "two", "three"},
		},



		{
			Tokens: []string{"--apple", "--banana", "--cherry", "--", "one"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			                              "banana":"",
			                                          "cherry":"",
			},
			ExpectedArgs:                                    []string{"one"},
		},
		{
			Tokens: []string{"--apple", "--banana", "--cherry", "--", "one", "two"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			                              "banana":"",
			                                          "cherry":"",
			},
			ExpectedArgs:                                    []string{"one", "two"},
		},
		{
			Tokens: []string{"--apple", "--banana", "--cherry", "--", "one", "two", "three"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			                              "banana":"",
			                                          "cherry":"",
			},
			ExpectedArgs:                                    []string{"one", "two", "three"},
		},



		{
			Tokens: []string{"--apple", "--banana", "--cherry", "one", "--"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			                              "banana":"",
			                                          "cherry":"",
			},
			ExpectedArgs:                              []string{"one", "--"},
		},
		{
			Tokens: []string{"--apple", "--banana", "--cherry", "one", "--", "two"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			                              "banana":"",
			                                          "cherry":"",
			},
			ExpectedArgs:                              []string{"one", "--", "two"},
		},
		{
			Tokens: []string{"--apple", "--banana", "--cherry", "one", "--", "two", "three"},
			ExpectedFlags: map[string]string{
			                   "apple":"",
			                              "banana":"",
			                                          "cherry":"",
			},
			ExpectedArgs:                              []string{"one", "--", "two", "three"},
		},
	}

	for testNumber, test := range tests {

		var keyvalues cliflags.KeyValues

		actualArgs, err := cliflags.Parse(&keyvalues, test.Tokens...)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		var actualFlags map[string]string = map[string]string{}
		keyvalues.For(func(key string, value string){
			actualFlags[key] = value
		})

		if expected, actual := test.ExpectedFlags, actualFlags; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, expected certain flags, but actually got different flags.", testNumber)
			t.Errorf("\tTOKENS: %#v", test.Tokens)
			t.Errorf("\tEXPECTED FLAGS: %#v", expected)
			t.Errorf("\tACTUAL   FLAGS: %#v", actual)
			t.Errorf("\tEXPECTED ARGS: %#v", test.ExpectedArgs)
			t.Errorf("\tACTUAL   ARGS: %#v", actualArgs)
			continue
		}

		if expected, actual := test.ExpectedArgs, actualArgs; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, expected certain args, but actually got different args.", testNumber)
			t.Errorf("\tTOKENS: %#v", test.Tokens)
			t.Errorf("\tEXPECTED FLAGS: %#v", test.ExpectedFlags)
			t.Errorf("\tACTUAL   FLAGS: %#v", actualFlags)
			t.Errorf("\tEXPECTED ARGS: %#v", expected)
			t.Errorf("\tACTUAL   ARGS: %#v", actual)
			continue
		}
	}
}

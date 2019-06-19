package cliflag_test

import (
	"github.com/reiver/go-cli/flag"

	"reflect"

	"testing"
)

func TestParse(t *testing.T) {

	tests := []struct{
		Tokens []string
		ExpectedKey   string
		ExpectedValue string
		ExpectedTokens []string
	}{
		{
			Tokens: []string{"--apple"},
			ExpectedKey:       "apple",
			ExpectedValue:           "",
			ExpectedTokens:    []string{},
		},
		{
			Tokens: []string{"--apple", "--banana"},
			ExpectedKey:       "apple",
			ExpectedValue:           "",
			ExpectedTokens:    []string{"--banana"},
		},
		{
			Tokens: []string{"--apple", "--banana", "--cherry"},
			ExpectedKey:       "apple",
			ExpectedValue:           "",
			ExpectedTokens:    []string{"--banana", "--cherry"},
		},




		{
			Tokens: []string{"--apple=one"},
			ExpectedKey:       "apple",
			ExpectedValue:           "one",
			ExpectedTokens:        []string{},
		},
		{
			Tokens: []string{"--apple=one", "--banana=two"},
			ExpectedKey:       "apple",
			ExpectedValue:           "one",
			ExpectedTokens:        []string{"--banana=two"},
		},
		{
			Tokens: []string{"--apple=one", "--banana=two", "--cherry=three"},
			ExpectedKey:       "apple",
			ExpectedValue:           "one",
			ExpectedTokens:        []string{"--banana=two", "--cherry=three"},
		},



		{
			Tokens: []string{"--apple=o"},
			ExpectedKey:       "apple",
			ExpectedValue:           "o",
			ExpectedTokens:        []string{},
		},
		{
			Tokens: []string{"--apple=o", "--banana=t"},
			ExpectedKey:       "apple",
			ExpectedValue:           "o",
			ExpectedTokens:      []string{"--banana=t"},
		},
		{
			Tokens: []string{"--apple=o", "--banana=t", "--cherry=t"},
			ExpectedKey:       "apple",
			ExpectedValue:           "o",
			ExpectedTokens:      []string{"--banana=t", "--cherry=t"},
		},



		{
			Tokens: []string{"--apple="},
			ExpectedKey:       "apple",
			ExpectedValue:           "",
			ExpectedTokens:        []string{},
		},
		{
			Tokens: []string{"--apple=", "--banana="},
			ExpectedKey:       "apple",
			ExpectedValue:           "",
			ExpectedTokens:     []string{"--banana="},
		},
		{
			Tokens: []string{"--apple=", "--banana=", "--cherry="},
			ExpectedKey:       "apple",
			ExpectedValue:           "",
			ExpectedTokens:     []string{"--banana=", "--cherry="},
		},



		{
			Tokens: []string{"--apple:one"},
			ExpectedKey:       "apple",
			ExpectedValue:           "one",
			ExpectedTokens:        []string{},
		},
		{
			Tokens: []string{"--apple:one", "--banana:two"},
			ExpectedKey:       "apple",
			ExpectedValue:           "one",
			ExpectedTokens:        []string{"--banana:two"},
		},
		{
			Tokens: []string{"--apple:one", "--banana:two", "--cherry:three"},
			ExpectedKey:       "apple",
			ExpectedValue:           "one",
			ExpectedTokens:        []string{"--banana:two", "--cherry:three"},
		},



		{
			Tokens: []string{"--apple:o"},
			ExpectedKey:       "apple",
			ExpectedValue:           "o",
			ExpectedTokens:        []string{},
		},
		{
			Tokens: []string{"--apple:o", "--banana:t"},
			ExpectedKey:       "apple",
			ExpectedValue:           "o",
			ExpectedTokens:      []string{"--banana:t"},
		},
		{
			Tokens: []string{"--apple:o", "--banana:t", "--cherry:t"},
			ExpectedKey:       "apple",
			ExpectedValue:           "o",
			ExpectedTokens:      []string{"--banana:t", "--cherry:t"},
		},



		{
			Tokens: []string{"--apple:"},
			ExpectedKey:       "apple",
			ExpectedValue:           "",
			ExpectedTokens:        []string{},
		},
		{
			Tokens: []string{"--apple:", "--banana:"},
			ExpectedKey:       "apple",
			ExpectedValue:           "",
			ExpectedTokens:     []string{"--banana:"},
		},
		{
			Tokens: []string{"--apple:", "--banana:", "--cherry:"},
			ExpectedKey:       "apple",
			ExpectedValue:           "",
			ExpectedTokens:     []string{"--banana:", "--cherry:"},
		},









		{
			Tokens:  []string{"--a"},
			ExpectedKey:        "a",
			ExpectedValue:         "",
			ExpectedTokens: []string{},
		},
		{
			Tokens:  []string{"--a", "--b"},
			ExpectedKey:        "a",
			ExpectedValue:         "",
			ExpectedTokens: []string{"--b"},
		},
		{
			Tokens:  []string{"--a", "--b", "--c"},
			ExpectedKey:        "a",
			ExpectedValue:         "",
			ExpectedTokens: []string{"--b", "--c"},
		},




		{
			Tokens: []string{"--a=one"},
			ExpectedKey:       "a",
			ExpectedValue:       "one",
			ExpectedTokens:      []string{},
		},
		{
			Tokens: []string{"--a=one", "--b=two"},
			ExpectedKey:       "a",
			ExpectedValue:       "one",
			ExpectedTokens:      []string{"--b=two"},
		},
		{
			Tokens: []string{"--a=one", "--b=two", "--c=three"},
			ExpectedKey:       "a",
			ExpectedValue:       "one",
			ExpectedTokens:        []string{"--b=two", "--c=three"},
		},



		{
			Tokens: []string{"--a=o"},
			ExpectedKey:       "a",
			ExpectedValue:       "o",
			ExpectedTokens:        []string{},
		},
		{
			Tokens: []string{"--a=o", "--b=t"},
			ExpectedKey:       "a",
			ExpectedValue:       "o",
			ExpectedTokens:      []string{"--b=t"},
		},
		{
			Tokens: []string{"--a=o", "--b=t", "--c=t"},
			ExpectedKey:       "a",
			ExpectedValue:       "o",
			ExpectedTokens:      []string{"--b=t", "--c=t"},
		},



		{
			Tokens: []string{"--a="},
			ExpectedKey:       "a",
			ExpectedValue:       "",
			ExpectedTokens:        []string{},
		},
		{
			Tokens: []string{"--a=", "--b="},
			ExpectedKey:       "a",
			ExpectedValue:           "",
			ExpectedTokens:     []string{"--b="},
		},
		{
			Tokens: []string{"--a=", "--b=", "--c="},
			ExpectedKey:       "a",
			ExpectedValue:           "",
			ExpectedTokens:     []string{"--b=", "--c="},
		},



		{
			Tokens: []string{"--a:one"},
			ExpectedKey:       "a",
			ExpectedValue:           "one",
			ExpectedTokens:        []string{},
		},
		{
			Tokens: []string{"--a:one", "--b:two"},
			ExpectedKey:       "a",
			ExpectedValue:           "one",
			ExpectedTokens:        []string{"--b:two"},
		},
		{
			Tokens: []string{"--a:one", "--b:two", "--c:three"},
			ExpectedKey:       "a",
			ExpectedValue:           "one",
			ExpectedTokens:        []string{"--b:two", "--c:three"},
		},



		{
			Tokens: []string{"--a:o"},
			ExpectedKey:       "a",
			ExpectedValue:           "o",
			ExpectedTokens:        []string{},
		},
		{
			Tokens: []string{"--a:o", "--b:t"},
			ExpectedKey:       "a",
			ExpectedValue:           "o",
			ExpectedTokens:      []string{"--b:t"},
		},
		{
			Tokens: []string{"--a:o", "--b:t", "--c:t"},
			ExpectedKey:       "a",
			ExpectedValue:           "o",
			ExpectedTokens:      []string{"--b:t", "--c:t"},
		},



		{
			Tokens: []string{"--a:"},
			ExpectedKey:       "a",
			ExpectedValue:           "",
			ExpectedTokens:        []string{},
		},
		{
			Tokens: []string{"--a:", "--b:"},
			ExpectedKey:       "a",
			ExpectedValue:           "",
			ExpectedTokens:     []string{"--b:"},
		},
		{
			Tokens: []string{"--a:", "--b:", "--c:"},
			ExpectedKey:       "a",
			ExpectedValue:           "",
			ExpectedTokens:     []string{"--b:", "--c:"},
		},









		{
			Tokens: []string{"+apple"},
			ExpectedKey:      "apple",
			ExpectedValue:           "",
			ExpectedTokens:       []string{},
		},
		{
			Tokens: []string{"+apple", "+banana"},
			ExpectedKey:      "apple",
			ExpectedValue:           "",
			ExpectedTokens:   []string{"+banana"},
		},
		{
			Tokens: []string{"+apple", "+banana", "+cherry"},
			ExpectedKey:      "apple",
			ExpectedValue:           "",
			ExpectedTokens:   []string{"+banana", "+cherry"},
		},



		{
			Tokens: []string{"+apple=one"},
			ExpectedKey:      "apple",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{},
		},
		{
			Tokens: []string{"+apple=one", "+banana=two"},
			ExpectedKey:      "apple",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{"+banana=two"},
		},
		{
			Tokens: []string{"+apple=one", "+banana=two", "+cherry=three"},
			ExpectedKey:      "apple",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{"+banana=two", "+cherry=three"},
		},



		{
			Tokens: []string{"+apple=o"},
			ExpectedKey:      "apple",
			ExpectedValue:          "o",
			ExpectedTokens:       []string{},
		},
		{
			Tokens: []string{"+apple=o", "+banana=t"},
			ExpectedKey:      "apple",
			ExpectedValue:          "o",
			ExpectedTokens:     []string{"+banana=t"},
		},
		{
			Tokens: []string{"+apple=o", "+banana=t", "+cherry=t"},
			ExpectedKey:      "apple",
			ExpectedValue:          "o",
			ExpectedTokens:     []string{"+banana=t", "+cherry=t"},
		},



		{
			Tokens: []string{"+apple="},
			ExpectedKey:      "apple",
			ExpectedValue:          "",
			ExpectedTokens:    []string{},
		},
		{
			Tokens: []string{"+apple=", "+banana="},
			ExpectedKey:      "apple",
			ExpectedValue:          "",
			ExpectedTokens:    []string{"+banana="},
		},
		{
			Tokens: []string{"+apple=", "+banana=", "+cherry="},
			ExpectedKey:      "apple",
			ExpectedValue:          "",
			ExpectedTokens:    []string{"+banana=", "+cherry="},
		},



		{
			Tokens: []string{"+apple:one"},
			ExpectedKey:      "apple",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{},
		},
		{
			Tokens: []string{"+apple:one", "+banana:two"},
			ExpectedKey:      "apple",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{"+banana:two"},
		},
		{
			Tokens: []string{"+apple:one", "+banana:two", "+cherry:three"},
			ExpectedKey:      "apple",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{"+banana:two", "+cherry:three"},
		},



		{
			Tokens: []string{"+apple:o"},
			ExpectedKey:      "apple",
			ExpectedValue:          "o",
			ExpectedTokens:     []string{},
		},
		{
			Tokens: []string{"+apple:o", "+banana:t"},
			ExpectedKey:      "apple",
			ExpectedValue:          "o",
			ExpectedTokens:     []string{"+banana:t"},
		},
		{
			Tokens: []string{"+apple:o", "+banana:t", "+cherry:t"},
			ExpectedKey:      "apple",
			ExpectedValue:          "o",
			ExpectedTokens:     []string{"+banana:t", "+cherry:t"},
		},



		{
			Tokens: []string{"+apple:"},
			ExpectedKey:      "apple",
			ExpectedValue:          "",
			ExpectedTokens:    []string{},
		},
		{
			Tokens: []string{"+apple:", "+banana:"},
			ExpectedKey:      "apple",
			ExpectedValue:          "",
			ExpectedTokens:    []string{"+banana:"},
		},
		{
			Tokens: []string{"+apple:", "+banana:", "+cherry:"},
			ExpectedKey:      "apple",
			ExpectedValue:          "",
			ExpectedTokens:    []string{"+banana:", "+cherry:"},
		},









		{
			Tokens: []string{"+a"},
			ExpectedKey:      "a",
			ExpectedValue:           "",
			ExpectedTokens:       []string{},
		},
		{
			Tokens: []string{"+a", "+b"},
			ExpectedKey:      "a",
			ExpectedValue:           "",
			ExpectedTokens:   []string{"+b"},
		},
		{
			Tokens: []string{"+a", "+b", "+c"},
			ExpectedKey:      "a",
			ExpectedValue:           "",
			ExpectedTokens:   []string{"+b", "+c"},
		},



		{
			Tokens: []string{"+a=one"},
			ExpectedKey:      "a",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{},
		},
		{
			Tokens: []string{"+a=one", "+b=two"},
			ExpectedKey:      "a",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{"+b=two"},
		},
		{
			Tokens: []string{"+a=one", "+b=two", "+c=three"},
			ExpectedKey:      "a",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{"+b=two", "+c=three"},
		},



		{
			Tokens: []string{"+a=o"},
			ExpectedKey:      "a",
			ExpectedValue:          "o",
			ExpectedTokens:       []string{},
		},
		{
			Tokens: []string{"+a=o", "+b=t"},
			ExpectedKey:      "a",
			ExpectedValue:          "o",
			ExpectedTokens:     []string{"+b=t"},
		},
		{
			Tokens: []string{"+a=o", "+b=t", "+c=t"},
			ExpectedKey:      "a",
			ExpectedValue:          "o",
			ExpectedTokens:     []string{"+b=t", "+c=t"},
		},



		{
			Tokens: []string{"+a="},
			ExpectedKey:      "a",
			ExpectedValue:          "",
			ExpectedTokens:    []string{},
		},
		{
			Tokens: []string{"+a=", "+b="},
			ExpectedKey:      "a",
			ExpectedValue:          "",
			ExpectedTokens:    []string{"+b="},
		},
		{
			Tokens: []string{"+a=", "+b=", "+c="},
			ExpectedKey:      "a",
			ExpectedValue:          "",
			ExpectedTokens:    []string{"+b=", "+c="},
		},



		{
			Tokens: []string{"+a:one"},
			ExpectedKey:      "a",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{},
		},
		{
			Tokens: []string{"+a:one", "+b:two"},
			ExpectedKey:      "a",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{"+b:two"},
		},
		{
			Tokens: []string{"+a:one", "+b:two", "+c:three"},
			ExpectedKey:      "a",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{"+b:two", "+c:three"},
		},



		{
			Tokens: []string{"+a:o"},
			ExpectedKey:      "a",
			ExpectedValue:          "o",
			ExpectedTokens:     []string{},
		},
		{
			Tokens: []string{"+a:o", "+b:t"},
			ExpectedKey:      "a",
			ExpectedValue:          "o",
			ExpectedTokens:     []string{"+b:t"},
		},
		{
			Tokens: []string{"+a:o", "+b:t", "+c:t"},
			ExpectedKey:      "a",
			ExpectedValue:          "o",
			ExpectedTokens:     []string{"+b:t", "+c:t"},
		},



		{
			Tokens: []string{"+a:"},
			ExpectedKey:      "a",
			ExpectedValue:          "",
			ExpectedTokens:    []string{},
		},
		{
			Tokens: []string{"+a:", "+b:"},
			ExpectedKey:      "a",
			ExpectedValue:          "",
			ExpectedTokens:    []string{"+b:"},
		},
		{
			Tokens: []string{"+a:", "+b:", "+c:"},
			ExpectedKey:      "a",
			ExpectedValue:          "",
			ExpectedTokens:    []string{"+b:", "+c:"},
		},









		{
			Tokens: []string{"-a"},
			ExpectedKey:      "a",
			ExpectedValue:           "",
			ExpectedTokens:       []string{},
		},
		{
			Tokens: []string{"-a", "-b"},
			ExpectedKey:      "a",
			ExpectedValue:           "",
			ExpectedTokens:   []string{"-b"},
		},
		{
			Tokens: []string{"-a", "-b", "-c"},
			ExpectedKey:      "a",
			ExpectedValue:           "",
			ExpectedTokens:   []string{"-b", "-c"},
		},



		{
			Tokens: []string{"-a=one"},
			ExpectedKey:      "a",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{},
		},
		{
			Tokens: []string{"-a=one", "-b=two"},
			ExpectedKey:      "a",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{"-b=two"},
		},
		{
			Tokens: []string{"-a=one", "-b=two", "-c=three"},
			ExpectedKey:      "a",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{"-b=two", "-c=three"},
		},



		{
			Tokens: []string{"-a=o"},
			ExpectedKey:      "a",
			ExpectedValue:          "o",
			ExpectedTokens:       []string{},
		},
		{
			Tokens: []string{"-a=o", "-b=t"},
			ExpectedKey:      "a",
			ExpectedValue:          "o",
			ExpectedTokens:     []string{"-b=t"},
		},
		{
			Tokens: []string{"-a=o", "-b=t", "-c=t"},
			ExpectedKey:      "a",
			ExpectedValue:          "o",
			ExpectedTokens:     []string{"-b=t", "-c=t"},
		},



		{
			Tokens: []string{"-a="},
			ExpectedKey:      "a",
			ExpectedValue:          "",
			ExpectedTokens:    []string{},
		},
		{
			Tokens: []string{"-a=", "-b="},
			ExpectedKey:      "a",
			ExpectedValue:          "",
			ExpectedTokens:    []string{"-b="},
		},
		{
			Tokens: []string{"-a=", "-b=", "-c="},
			ExpectedKey:      "a",
			ExpectedValue:          "",
			ExpectedTokens:    []string{"-b=", "-c="},
		},



		{
			Tokens: []string{"-a:one"},
			ExpectedKey:      "a",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{},
		},
		{
			Tokens: []string{"-a:one", "-b:two"},
			ExpectedKey:      "a",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{"-b:two"},
		},
		{
			Tokens: []string{"-a:one", "-b:two", "-c:three"},
			ExpectedKey:      "a",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{"-b:two", "-c:three"},
		},



		{
			Tokens: []string{"-a:o"},
			ExpectedKey:      "a",
			ExpectedValue:          "o",
			ExpectedTokens:     []string{},
		},
		{
			Tokens: []string{"-a:o", "-b:t"},
			ExpectedKey:      "a",
			ExpectedValue:          "o",
			ExpectedTokens:     []string{"-b:t"},
		},
		{
			Tokens: []string{"-a:o", "-b:t", "-c:t"},
			ExpectedKey:      "a",
			ExpectedValue:          "o",
			ExpectedTokens:     []string{"-b:t", "-c:t"},
		},



		{
			Tokens: []string{"-a:"},
			ExpectedKey:      "a",
			ExpectedValue:          "",
			ExpectedTokens:    []string{},
		},
		{
			Tokens: []string{"-a:", "-b:"},
			ExpectedKey:      "a",
			ExpectedValue:          "",
			ExpectedTokens:    []string{"-b:"},
		},
		{
			Tokens: []string{"-a:", "-b:", "-c:"},
			ExpectedKey:      "a",
			ExpectedValue:          "",
			ExpectedTokens:    []string{"-b:", "-c:"},
		},



		{
			Tokens: []string{"-aone"},
			ExpectedKey:      "a",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{},
		},
		{
			Tokens: []string{"-aone", "-btwo"},
			ExpectedKey:      "a",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{"-btwo"},
		},
		{
			Tokens: []string{"-aone", "-btwo", "-cthree"},
			ExpectedKey:      "a",
			ExpectedValue:          "one",
			ExpectedTokens:       []string{"-btwo", "-cthree"},
		},



		{
			Tokens: []string{"-ao"},
			ExpectedKey:      "a",
			ExpectedValue:          "o",
			ExpectedTokens:       []string{},
		},
		{
			Tokens: []string{"-ao", "-bt"},
			ExpectedKey:      "a",
			ExpectedValue:          "o",
			ExpectedTokens:     []string{"-bt"},
		},
		{
			Tokens: []string{"-ao", "-bt", "-ct"},
			ExpectedKey:      "a",
			ExpectedValue:          "o",
			ExpectedTokens:     []string{"-bt", "-ct"},
		},









		{
			Tokens: []string{"-a", "one"},
			ExpectedKey:      "a",
			ExpectedValue:         "one",
			ExpectedTokens:      []string{},
		},
		{
			Tokens: []string{"-a", "one", "-b", "two"},
			ExpectedKey:      "a",
			ExpectedValue:         "one",
			ExpectedTokens:      []string{"-b", "two"},
		},
		{
			Tokens: []string{"-a", "one", "-b", "two", "-c", "three"},
			ExpectedKey:      "a",
			ExpectedValue:         "one",
			ExpectedTokens:      []string{"-b", "two", "-c", "three"},
		},









		{
			Tokens: []string{"-a", "--", "one"},
			ExpectedKey:      "a",
			ExpectedValue:               "one",
			ExpectedTokens:      []string{},
		},









		{
			Tokens:   []string{"-a"},
			ExpectedKey:        "a",
			ExpectedValue:       "",
			ExpectedTokens: []string{},
		},
		{
			Tokens:   []string{"-a", "-b"},
			ExpectedKey:        "a",
			ExpectedValue:       "",
			ExpectedTokens: []string{"-b"},
		},
		{
			Tokens:   []string{"-a", "-b", "-c"},
			ExpectedKey:        "a",
			ExpectedValue:       "",
			ExpectedTokens: []string{"-b", "-c"},
		},



		{
			Tokens:   []string{"-m", "This is a message", "-n5", "-f", "filename.txt"},
			ExpectedKey:        "m",
			ExpectedValue:           "This is a message",
			ExpectedTokens:                      []string{"-n5", "-f", "filename.txt"},
		},
		{
			Tokens:   []string{"-m", "This is a message", "-n5", "filename.txt"},
			ExpectedKey:        "m",
			ExpectedValue:           "This is a message",
			ExpectedTokens:                      []string{"-n5", "filename.txt"},
		},
		{
			Tokens:   []string{"-m", "This is a message", "filename.txt"},
			ExpectedKey:        "m",
			ExpectedValue:           "This is a message",
			ExpectedTokens:                      []string{"filename.txt"},
		},









		{
			Tokens: []string{"--op==", "filename.txt"},
			ExpectedKey:       "op",
			ExpectedValue:        "=",
			ExpectedTokens:   []string{"filename.txt"},
		},
		{
			Tokens: []string{"--op=:", "filename.txt"},
			ExpectedKey:       "op",
			ExpectedValue:        ":",
			ExpectedTokens:   []string{"filename.txt"},
		},
		{
			Tokens: []string{"--op:=", "filename.txt"},
			ExpectedKey:       "op",
			ExpectedValue:        "=",
			ExpectedTokens:   []string{"filename.txt"},
		},
		{
			Tokens: []string{"--op::", "filename.txt"},
			ExpectedKey:       "op",
			ExpectedValue:        ":",
			ExpectedTokens:   []string{"filename.txt"},
		},









		{
			Tokens:   []string{"-m", "--", "-3"},
			ExpectedKey:        "m",
			ExpectedValue:                 "-3",
			ExpectedTokens:             []string{},
		},
		{
			Tokens:   []string{"-m", "--", "-3", "filename.txt"},
			ExpectedKey:        "m",
			ExpectedValue:                 "-3",
			ExpectedTokens:             []string{"filename.txt"},
		},
	}

	for testNumber, test := range tests {

		{

			var keyvalue cliflag.KeyValue

			{
				remainingTokens, err := cliflag.Parse(&keyvalue, test.Tokens...)
				if nil != err {
					t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
					continue
				}

				if expected, actual := test.ExpectedTokens, remainingTokens; !reflect.DeepEqual(expected, actual) {
					t.Errorf("For test #%d ...", testNumber)
					t.Errorf("\tEXPECTED REMAINING TOKENS: %#v", expected)
					t.Errorf("\tACTUAL REMAINING TOKENS:   %#v", actual)
					t.Errorf("\t\ttokens: %#v", test.Tokens)
					t.Errorf("\t\texpected key: %q", test.ExpectedKey)
					t.Errorf("\t\texpected value: %q", test.ExpectedValue)
					continue
				}
			}

			actualKey, actualValue, err := keyvalue.Unwrap()
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
				continue
			}

			if expected, actual := test.ExpectedKey, actualKey; expected != actual {
				t.Errorf("For test #%d ...", testNumber)
				t.Errorf("\tEXPECTED KEY: %q", expected)
				t.Errorf("\tACTUAL KEY:   %q", actual)
				t.Errorf("\t\ttokens: %#v", test.Tokens)
				t.Errorf("\t\texpected tokens: %#v", test.ExpectedTokens)
				t.Errorf("\t\texpected value: %q", test.ExpectedValue)
				continue
			}

			if expected, actual := test.ExpectedValue, actualValue; expected != actual {
				t.Errorf("For test #%d ...", testNumber)
				t.Errorf("\tEXPECTED VALUE: %q", expected)
				t.Errorf("\tACTUAL VALUE:   %q", actual)
				t.Errorf("\t\ttokens: %#v", test.Tokens)
				t.Errorf("\t\texpected tokens: %#v", test.ExpectedTokens)
				t.Errorf("\t\texpected key: %q", test.ExpectedKey)
				continue
			}
		}



		{
			var mapStringString map[string]string

			{
				remainingTokens, err := cliflag.Parse(&mapStringString, test.Tokens...)
				if nil != err {
					t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
					continue
				}

				if expected, actual := test.ExpectedTokens, remainingTokens; !reflect.DeepEqual(expected, actual) {
					t.Errorf("For test #%d ...", testNumber)
					t.Errorf("\tEXPECTED REMAINING TOKENS: %#v", expected)
					t.Errorf("\tACTUAL REMAINING TOKENS:   %#v", actual)
					t.Errorf("\t\ttokens: %#v", test.Tokens)
					t.Errorf("\t\texpected key: %q", test.ExpectedKey)
					t.Errorf("\t\texpected value: %q", test.ExpectedValue)
					continue
				}

				if expected, actual := 1, len(mapStringString); expected != actual {
					t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
					continue
				}

				{
					_, found := mapStringString[test.ExpectedKey]
					if expected := test.ExpectedKey; !found {
						t.Errorf("For test #%d, expected key %q in map, but was not found.", testNumber, expected)
						continue
					}
				}

				{
					actualValue := mapStringString[test.ExpectedKey]
					if expected, actual := test.ExpectedValue, actualValue; expected != actual {
						t.Errorf("For test #%d ...", testNumber)
						t.Errorf("\tEXPECTED: %q", expected)
						t.Errorf("\tACTUAL:   %q", actual)
						continue
					}
				}
			}

		}



		{
			var mapStringInterface map[string]interface{}

			{
				remainingTokens, err := cliflag.Parse(&mapStringInterface, test.Tokens...)
				if nil != err {
					t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
					continue
				}

				if expected, actual := test.ExpectedTokens, remainingTokens; !reflect.DeepEqual(expected, actual) {
					t.Errorf("For test #%d ...", testNumber)
					t.Errorf("\tEXPECTED REMAINING TOKENS: %#v", expected)
					t.Errorf("\tACTUAL REMAINING TOKENS:   %#v", actual)
					t.Errorf("\t\ttokens: %#v", test.Tokens)
					t.Errorf("\t\texpected key: %q", test.ExpectedKey)
					t.Errorf("\t\texpected value: %q", test.ExpectedValue)
					continue
				}

				if expected, actual := 1, len(mapStringInterface); expected != actual {
					t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
					continue
				}

				{
					_, found := mapStringInterface[test.ExpectedKey]
					if expected := test.ExpectedKey; !found {
						t.Errorf("For test #%d, expected key %q in map, but was not found.", testNumber, expected)
						continue
					}
				}

				{
					actualValue := mapStringInterface[test.ExpectedKey]
					if expected, actual := test.ExpectedValue, actualValue; expected != actual {
						t.Errorf("For test #%d ...", testNumber)
						t.Errorf("\tEXPECTED: %q", expected)
						t.Errorf("\tACTUAL:   %q", actual)
						continue
					}
				}
			}

		}

	}
}

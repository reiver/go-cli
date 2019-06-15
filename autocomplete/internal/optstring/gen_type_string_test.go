package optstring

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */


import (
	"fmt"

	"testing"
)

func TestTypeString(t *testing.T) {
	tests := []struct {
		Datum    Type
		Expected string
	}{
		{
			Datum: Some("apple"),
			Expected:   "apple",
		},
		{
			Datum: Some("BANANA"),
			Expected:   "BANANA",
		},
		{
			Datum: Some("Cherry"),
			Expected:   "Cherry",
		},
		{
			Datum: Some("dATE"),
			Expected:   "dATE",
		},



		{
			Datum: Some("Hello world!"),
			Expected:   "Hello world!",
		},



		{
			Datum: Some("-5"),
			Expected:   "-5",
		},
		{
			Datum: Some("-4"),
			Expected:   "-4",
		},
		{
			Datum: Some("-3"),
			Expected:   "-3",
		},
		{
			Datum: Some("-2"),
			Expected:   "-2",
		},
		{
			Datum: Some("-1"),
			Expected:   "-1",
		},
		{
			Datum: Some("0"),
			Expected:   "0",
		},
		{
			Datum: Some("1"),
			Expected:   "1",
		},
		{
			Datum: Some("2"),
			Expected:   "2",
		},
		{
			Datum: Some("3"),
			Expected:   "3",
		},
		{
			Datum: Some("4"),
			Expected:   "4",
		},
		{
			Datum: Some("5"),
			Expected:   "5",
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			s := fmt.Sprintf("%d", x)

			test := struct {
				Datum    Type
				Expected string
			}{
				Datum: Some(s),
				Expected:   s,
			}
			tests = append(tests, test)
		}

		{
			y := -x

			s := fmt.Sprintf("%d", y)

			test := struct {
				Datum    Type
				Expected string
			}{
				Datum: Some(s),
				Expected:   s,
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {
		datum := test.Datum

		actual, err := datum.String()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v", testNumber, expected, actual)
			continue
		}
	}
}

func TestTypeStringNone(t *testing.T) {

	datum := None()

	_, err := datum.String()
	if nil == err {
		t.Errorf("Expected an error, but did not actually get one: %v", err)
		return
	}

	if expected, actual := errNone, err; expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v", expected, expected, actual, actual)
		return
	}
}

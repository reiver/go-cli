package optstring

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */


import (
	"testing"
)

func TestTypeValue(t *testing.T) {
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
			Datum: Some("1"),
			Expected:   "1",
		},
		{
			Datum: Some("12"),
			Expected:   "12",
		},
		{
			Datum: Some("-5"),
			Expected:   "-5",
		},
	}


	for testNumber, test := range tests {
		actual, err := test.Datum.Value()
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

func TestTypeValueNone(t *testing.T) {
	datum := None()

	_, err := datum.Value()
	if nil == err {
		t.Errorf("Expected an error, but did not actually get one: (%T) %v", err, err)
		return
	}

}

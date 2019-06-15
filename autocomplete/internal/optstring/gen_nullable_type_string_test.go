package optstring

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */


import (
	"fmt"
	"math"

	"testing"
)

func TestNullableTypeString(t *testing.T) {
	tests := []struct {
		Datum    NullableType
		Expected string
	}{
		{
			Datum: SomeNullable("apple"),
			Expected:           "apple",
		},
		{
			Datum: SomeNullable("BANANA"),
			Expected:           "BANANA",
		},
		{
			Datum: SomeNullable("Cherry"),
			Expected:           "Cherry",
		},
		{
			Datum: SomeNullable("dATE"),
			Expected:           "dATE",
		},



		{
			Datum: SomeNullable("Hello world!"),
			Expected:           "Hello world!",
		},



		{
			Datum: SomeNullable(fmt.Sprintf("%d", math.MinInt64)),
			Expected:           fmt.Sprintf("%d", math.MinInt64),
		},
		{
			Datum: SomeNullable(fmt.Sprintf("%d", -5)),
			Expected:           fmt.Sprintf("%d", -5),
		},
		{
			Datum: SomeNullable(fmt.Sprintf("%d", -4)),
			Expected:           fmt.Sprintf("%d", -4),
		},
		{
			Datum: SomeNullable(fmt.Sprintf("%d", -3)),
			Expected:           fmt.Sprintf("%d", -3),
		},
		{
			Datum: SomeNullable(fmt.Sprintf("%d", -2)),
			Expected:           fmt.Sprintf("%d", -2),
		},
		{
			Datum: SomeNullable(fmt.Sprintf("%d", -1)),
			Expected:           fmt.Sprintf("%d", -1),
		},
		{
			Datum: SomeNullable(fmt.Sprintf("%d", 0)),
			Expected:           fmt.Sprintf("%d", 0),
		},
		{
			Datum: SomeNullable(fmt.Sprintf("%d", 1)),
			Expected:           fmt.Sprintf("%d", 1),
		},
		{
			Datum: SomeNullable(fmt.Sprintf("%d", 2)),
			Expected:           fmt.Sprintf("%d", 2),
		},
		{
			Datum: SomeNullable(fmt.Sprintf("%d", 3)),
			Expected:           fmt.Sprintf("%d", 3),
		},
		{
			Datum: SomeNullable(fmt.Sprintf("%d", 4)),
			Expected:           fmt.Sprintf("%d", 4),
		},
		{
			Datum: SomeNullable(fmt.Sprintf("%d", 5)),
			Expected:           fmt.Sprintf("%d", 5),
		},
		{
			Datum: SomeNullable(fmt.Sprintf("%d", math.MaxInt64)),
			Expected:           fmt.Sprintf("%d", math.MaxInt64),
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			s := fmt.Sprintf("%d", x)

			test := struct {
				Datum    NullableType
				Expected string
			}{
				Datum: SomeNullable(s),
				Expected:           s,
			}
			tests = append(tests, test)
		}

		{
			y := -x

			s := fmt.Sprintf("%d", y)

			test := struct {
				Datum    NullableType
				Expected string
			}{
				Datum: SomeNullable(s),
				Expected:           s,
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

func TestNullableTypeStringNoneNullable(t *testing.T) {

	datum := NoneNullable()

	_, err := datum.String()
	if nil == err {
		t.Errorf("Expected an error, but did not actually get one: %v", err)
		return
	}

	if expected, actual := errNoneNullable, err; expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v", expected, expected, actual, actual)
		return
	}
}

func TestNullableTypeStringNull(t *testing.T) {

	datum := Null()

	_, err := datum.String()
	if nil == err {
		t.Errorf("Expected an error, but did not actually get one: %v", err)
		return
	}

	if expected, actual := errNull, err; expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v", expected, expected, actual, actual)
		return
	}
}

package optint64

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */


import (
	"math"

	"testing"
)

func TestNullableTypeInt64(t *testing.T) {
	tests := []struct {
		Datum    NullableType
		Expected int64
	}{
		{
			Datum: SomeNullable(math.MinInt64),
			Expected:           math.MinInt64,
		},
		{
			Datum: SomeNullable(-5),
			Expected:           -5,
		},
		{
			Datum: SomeNullable(-4),
			Expected:           -4,
		},
		{
			Datum: SomeNullable(-3),
			Expected:           -3,
		},
		{
			Datum: SomeNullable(-2),
			Expected:           -2,
		},
		{
			Datum: SomeNullable(-1),
			Expected:           -1,
		},
		{
			Datum: SomeNullable(0),
			Expected:           0,
		},
		{
			Datum: SomeNullable(1),
			Expected:           1,
		},
		{
			Datum: SomeNullable(2),
			Expected:           2,
		},
		{
			Datum: SomeNullable(3),
			Expected:           3,
		},
		{
			Datum: SomeNullable(4),
			Expected:           4,
		},
		{
			Datum: SomeNullable(5),
			Expected:           5,
		},
		{
			Datum: SomeNullable(math.MaxInt64),
			Expected:           math.MaxInt64,
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			test := struct {
				Datum    NullableType
				Expected int64
			}{
				Datum: SomeNullable(x),
				Expected:           x,
			}
			tests = append(tests, test)
		}

		{
			y := -x

			test := struct {
				Datum    NullableType
				Expected int64
			}{
				Datum: SomeNullable(y),
				Expected:           y,
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {
		datum := test.Datum

		actual, err := datum.Int64()
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

func TestNullableTypeInt64NoneNullable(t *testing.T) {

	datum := NoneNullable()

	_, err := datum.Int64()
	if nil == err {
		t.Errorf("Expected an error, but did not actually get one: %v", err)
		return
	}

	if expected, actual := errNoneNullable, err; expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v", expected, expected, actual, actual)
		return
	}
}

func TestNullableTypeInt64Null(t *testing.T) {

	datum := Null()

	_, err := datum.Int64()
	if nil == err {
		t.Errorf("Expected an error, but did not actually get one: %v", err)
		return
	}

	if expected, actual := errNull, err; expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v", expected, expected, actual, actual)
		return
	}
}

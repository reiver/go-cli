package optint64

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */


import (
	"math"

	"testing"
)

func TestTypeInt64(t *testing.T) {
	tests := []struct {
		Datum    Type
		Expected int64
	}{
		{
			Datum:    Some(math.MinInt64),
			Expected:      math.MinInt64,
		},
		{
			Datum:    Some(-5),
			Expected:      -5,
		},
		{
			Datum:    Some(-4),
			Expected:      -4,
		},
		{
			Datum:    Some(-3),
			Expected:      -3,
		},
		{
			Datum:    Some(-2),
			Expected:      -2,
		},
		{
			Datum:    Some(-1),
			Expected:      -1,
		},
		{
			Datum:    Some(0),
			Expected:      0,
		},
		{
			Datum:    Some(1),
			Expected:      1,
		},
		{
			Datum:    Some(2),
			Expected:      2,
		},
		{
			Datum:    Some(3),
			Expected:      3,
		},
		{
			Datum:    Some(4),
			Expected:      4,
		},
		{
			Datum:    Some(5),
			Expected:      5,
		},
		{
			Datum:    Some(math.MaxInt64),
			Expected:      math.MaxInt64,
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			test := struct {
				Datum    Type
				Expected int64
			}{
				Datum:    Some(x),
				Expected: x,
			}
			tests = append(tests, test)
		}

		{
			y := -x

			test := struct {
				Datum    Type
				Expected int64
			}{
				Datum:    Some(y),
				Expected: y,
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

func TestTypeInt64None(t *testing.T) {

	datum := None()

	_, err := datum.Int64()
	if nil == err {
		t.Errorf("Expected an error, but did not actually get one: %v", err)
		return
	}

	if expected, actual := errNone, err; expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v", expected, expected, actual, actual)
		return
	}
}

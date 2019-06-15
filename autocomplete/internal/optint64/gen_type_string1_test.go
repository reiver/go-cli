package optint64

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */


import (
	"fmt"
	"math"
	"testing"
)

func TestTypeString(t *testing.T) {
	tests := []struct{
		Datum    Type
		Expected string
	}{
		{
			Datum:                                Some(math.MinInt64),
			Expected: fmt.Sprintf("optint64.Some(%d)", math.MinInt64),
		},
		{
			Datum                               : Some(-5),
			Expected: fmt.Sprintf("optint64.Some(%d)", -5),
		},
		{
			Datum                               : Some(-4),
			Expected: fmt.Sprintf("optint64.Some(%d)", -4),
		},
		{
			Datum                               : Some(-3),
			Expected: fmt.Sprintf("optint64.Some(%d)", -3),
		},
		{
			Datum                               : Some(-2),
			Expected: fmt.Sprintf("optint64.Some(%d)", -2),
		},
		{
			Datum                               : Some(-1),
			Expected: fmt.Sprintf("optint64.Some(%d)", -1),
		},
		{
			Datum                               : Some(0),
			Expected: fmt.Sprintf("optint64.Some(%d)", 0),
		},
		{
			Datum                               : Some(1),
			Expected: fmt.Sprintf("optint64.Some(%d)", 1),
		},
		{
			Datum                               : Some(2),
			Expected: fmt.Sprintf("optint64.Some(%d)", 2),
		},
		{
			Datum                               : Some(3),
			Expected: fmt.Sprintf("optint64.Some(%d)", 3),
		},
		{
			Datum                               : Some(4),
			Expected: fmt.Sprintf("optint64.Some(%d)", 4),
		},
		{
			Datum                               : Some(5),
			Expected: fmt.Sprintf("optint64.Some(%d)", 5),
		},
		{
			Datum:                                Some(math.MaxInt64),
			Expected: fmt.Sprintf("optint64.Some(%d)", math.MaxInt64),
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			test := struct {
				Datum Type
				Expected string
			}{
				Datum:                                Some(x),
				Expected: fmt.Sprintf("optint64.Some(%d)", x),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			test := struct {
				Datum Type
				Expected string
			}{
				Datum:                                Some(y),
				Expected: fmt.Sprintf("optint64.Some(%d)", y),
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {
		datum := test.Datum

		if expected, actual := test.Expected, datum.String(); expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v", testNumber, expected, actual)
			continue
		}
	}
}

func TestTypeStringNone(t *testing.T) {

	datum := None()

	if expected, actual := "optint64.None()", datum.String(); expected != actual {
		t.Errorf("Expected %q, but actually got %q", expected, actual)
		return
	}
}

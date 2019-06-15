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

func TestNullableTypeString(t *testing.T) {
	tests := []struct{
		Datum    NullableType
		Expected string
	}{
		{
			Datum:                                SomeNullable(math.MinInt64),
			Expected: fmt.Sprintf("optint64.SomeNullable(%d)", math.MinInt64),
		},
		{
			Datum                               : SomeNullable(-5),
			Expected: fmt.Sprintf("optint64.SomeNullable(%d)", -5),
		},
		{
			Datum                               : SomeNullable(-4),
			Expected: fmt.Sprintf("optint64.SomeNullable(%d)", -4),
		},
		{
			Datum                               : SomeNullable(-3),
			Expected: fmt.Sprintf("optint64.SomeNullable(%d)", -3),
		},
		{
			Datum                               : SomeNullable(-2),
			Expected: fmt.Sprintf("optint64.SomeNullable(%d)", -2),
		},
		{
			Datum                               : SomeNullable(-1),
			Expected: fmt.Sprintf("optint64.SomeNullable(%d)", -1),
		},
		{
			Datum                               : SomeNullable(0),
			Expected: fmt.Sprintf("optint64.SomeNullable(%d)", 0),
		},
		{
			Datum                               : SomeNullable(1),
			Expected: fmt.Sprintf("optint64.SomeNullable(%d)", 1),
		},
		{
			Datum                               : SomeNullable(2),
			Expected: fmt.Sprintf("optint64.SomeNullable(%d)", 2),
		},
		{
			Datum                               : SomeNullable(3),
			Expected: fmt.Sprintf("optint64.SomeNullable(%d)", 3),
		},
		{
			Datum                               : SomeNullable(4),
			Expected: fmt.Sprintf("optint64.SomeNullable(%d)", 4),
		},
		{
			Datum                               : SomeNullable(5),
			Expected: fmt.Sprintf("optint64.SomeNullable(%d)", 5),
		},
		{
			Datum:                                SomeNullable(math.MaxInt64),
			Expected: fmt.Sprintf("optint64.SomeNullable(%d)", math.MaxInt64),
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			test := struct {
				Datum NullableType
				Expected string
			}{
				Datum:                                SomeNullable(x),
				Expected: fmt.Sprintf("optint64.SomeNullable(%d)", x),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			test := struct {
				Datum NullableType
				Expected string
			}{
				Datum:                                SomeNullable(y),
				Expected: fmt.Sprintf("optint64.SomeNullable(%d)", y),
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

func TestNullableTypeStringNoneNullable(t *testing.T) {

	datum := NoneNullable()

	if expected, actual := "optint64.NoneNullable()", datum.String(); expected != actual {
		t.Errorf("Expected %q, but actually got %q", expected, actual)
		return
	}
}

func TestNullableTypeStringNull(t *testing.T) {

	datum := Null()

	if expected, actual := "optint64.Null()", datum.String(); expected != actual {
		t.Errorf("Expected %q, but actually got %q", expected, actual)
		return
	}
}

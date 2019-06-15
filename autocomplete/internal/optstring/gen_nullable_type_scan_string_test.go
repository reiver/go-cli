package optstring

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */


import (
	"fmt"
	"math"
	"time"

	"testing"
)

func TestNullableTypeScan(t *testing.T) {
	tests := []struct {
		Datum    interface{}
		Expected NullableType
	}{
		{
			Datum:    nil,
			Expected: Null(),
		},



		{
			Datum:    None(),
			Expected: NoneNullable(),
		},



		{
			Datum:            Some("apple"),
			Expected: SomeNullable("apple"),
		},
		{
			Datum:            Some("BANANA"),
			Expected: SomeNullable("BANANA"),
		},
		{
			Datum:            Some("Cherry"),
			Expected: SomeNullable("Cherry"),
		},
		{
			Datum:            Some("dATE"),
			Expected: SomeNullable("dATE"),
		},



		{
			Datum:    SomeNullable("apple"),
			Expected: SomeNullable("apple"),
		},
		{
			Datum:    SomeNullable("BANANA"),
			Expected: SomeNullable("BANANA"),
		},
		{
			Datum:    SomeNullable("Cherry"),
			Expected: SomeNullable("Cherry"),
		},
		{
			Datum:    SomeNullable("dATE"),
			Expected: SomeNullable("dATE"),
		},



		{
			Datum:                 "apple",
			Expected: SomeNullable("apple"),
		},
		{
			Datum:                 "BANANA",
			Expected: SomeNullable("BANANA"),
		},
		{
			Datum:                 "Cherry",
			Expected: SomeNullable("Cherry"),
		},
		{
			Datum:                 "dATE",
			Expected: SomeNullable("dATE"),
		},



		{
			Datum:                 fmt.Sprintf("%d", math.MinInt64),
			Expected: SomeNullable(fmt.Sprintf("%d", math.MinInt64)),
		},
		{
			Datum:                 fmt.Sprintf("%d", -5),
			Expected: SomeNullable(fmt.Sprintf("%d", -5)),
		},
		{
			Datum:                 fmt.Sprintf("%d", -4),
			Expected: SomeNullable(fmt.Sprintf("%d", -4)),
		},
		{
			Datum:                 fmt.Sprintf("%d", -3),
			Expected: SomeNullable(fmt.Sprintf("%d", -3)),
		},
		{
			Datum:                 fmt.Sprintf("%d", -2),
			Expected: SomeNullable(fmt.Sprintf("%d", -2)),
		},
		{
			Datum:                 fmt.Sprintf("%d", -1),
			Expected: SomeNullable(fmt.Sprintf("%d", -1)),
		},
		{
			Datum:                 fmt.Sprintf("%d", 0),
			Expected: SomeNullable(fmt.Sprintf("%d", 0)),
		},
		{
			Datum:                 fmt.Sprintf("%d", 1),
			Expected: SomeNullable(fmt.Sprintf("%d", 1)),
		},
		{
			Datum:                 fmt.Sprintf("%d", 2),
			Expected: SomeNullable(fmt.Sprintf("%d", 2)),
		},
		{
			Datum:                 fmt.Sprintf("%d", 3),
			Expected: SomeNullable(fmt.Sprintf("%d", 3)),
		},
		{
			Datum:                 fmt.Sprintf("%d", 4),
			Expected: SomeNullable(fmt.Sprintf("%d", 4)),
		},
		{
			Datum:                 fmt.Sprintf("%d", 5),
			Expected: SomeNullable(fmt.Sprintf("%d", 5)),
		},
		{
			Datum:                 fmt.Sprintf("%d", math.MaxInt64),
			Expected: SomeNullable(fmt.Sprintf("%d", math.MaxInt64)),
		},



		{
			Datum:          []byte(fmt.Sprintf("%d", math.MinInt64)),
			Expected: SomeNullable(fmt.Sprintf("%d", math.MinInt64)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", -5)),
			Expected: SomeNullable(fmt.Sprintf("%d", -5)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", -4)),
			Expected: SomeNullable(fmt.Sprintf("%d", -4)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", -3)),
			Expected: SomeNullable(fmt.Sprintf("%d", -3)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", -2)),
			Expected: SomeNullable(fmt.Sprintf("%d", -2)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", -1)),
			Expected: SomeNullable(fmt.Sprintf("%d", -1)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", 0)),
			Expected: SomeNullable(fmt.Sprintf("%d", 0)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", 1)),
			Expected: SomeNullable(fmt.Sprintf("%d", 1)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", 2)),
			Expected: SomeNullable(fmt.Sprintf("%d", 2)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", 3)),
			Expected: SomeNullable(fmt.Sprintf("%d", 3)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", 4)),
			Expected: SomeNullable(fmt.Sprintf("%d", 4)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", 5)),
			Expected: SomeNullable(fmt.Sprintf("%d", 5)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", math.MaxInt64)),
			Expected: SomeNullable(fmt.Sprintf("%d", math.MaxInt64)),
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			s := fmt.Sprintf("%d", x)

			test := struct {
				Datum    interface{}
				Expected NullableType
			}{
				Datum:                 s,
				Expected: SomeNullable(s),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			s := fmt.Sprintf("%d", y)

			test := struct {
				Datum    interface{}
				Expected NullableType
			}{
				Datum:                 s,
				Expected: SomeNullable(s),
			}
			tests = append(tests, test)
		}

		{
			s := fmt.Sprintf("%d", x)
			b := []byte(s)

			test := struct {
				Datum    interface{}
				Expected NullableType
			}{
				Datum:                 b,
				Expected: SomeNullable(s),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			s := fmt.Sprintf("%d", y)
			b := []byte(s)

			test := struct {
				Datum    interface{}
				Expected NullableType
			}{
				Datum:                 b,
				Expected: SomeNullable(s),
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {
		var datum NullableType

		if err := datum.Scan(test.Datum); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, datum; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v", testNumber, expected, actual)
			continue
		}
	}
}

func TestNullableTypeScanFail(t *testing.T) {
	tests := []struct {
		Datum interface{}
	}{
		{
			Datum: false,
		},
		{
			Datum: true,
		},



		{
			Datum: int64(-5),
		},
		{
			Datum: int64(-4),
		},
		{
			Datum: int64(-3),
		},
		{
			Datum: int64(-2),
		},
		{
			Datum: int64(-1),
		},
		{
			Datum: int64(0),
		},
		{
			Datum: int64(1),
		},
		{
			Datum: int64(2),
		},
		{
			Datum: int64(3),
		},
		{
			Datum: int64(4),
		},
		{
			Datum: int64(5),
		},



		{
			Datum: float64(-1.1),
		},
		{
			Datum: float64(-1.0),
		},
		{
			Datum: float64(-0.1),
		},
		{
			Datum: float64(0.0),
		},
		{
			Datum: float64(0.1),
		},
		{
			Datum: float64(1.0),
		},
		{
			Datum: float64(1.1),
		},



		{
			Datum: time.Now(),
		},
	}


	for testNumber, test := range tests {
		var datum NullableType

		if err := datum.Scan(test.Datum); nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one: (%T) %v", testNumber, err, err)
			continue
		}
	}
}

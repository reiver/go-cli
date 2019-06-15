package optstring

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */


import (
	"fmt"
	"time"

	"testing"
)

func TestTypeScan(t *testing.T) {
	tests := []struct {
		Datum    interface{}
		Expected Type
	}{
		{
			Datum:    Some("apple"),
			Expected: Some("apple"),
		},
		{
			Datum:    Some("BANANA"),
			Expected: Some("BANANA"),
		},
		{
			Datum:    Some("Cherry"),
			Expected: Some("Cherry"),
		},
		{
			Datum:    Some("dATE"),
			Expected: Some("dATE"),
		},



		{
			Datum:         "apple",
			Expected: Some("apple"),
		},
		{
			Datum:         "BANANA",
			Expected: Some("BANANA"),
		},
		{
			Datum:         "Cherry",
			Expected: Some("Cherry"),
		},
		{
			Datum:         "dATE",
			Expected: Some("dATE"),
		},



		{
			Datum:  []byte("apple"),
			Expected: Some("apple"),
		},
		{
			Datum:  []byte("BANANA"),
			Expected: Some("BANANA"),
		},
		{
			Datum:  []byte("Cherry"),
			Expected: Some("Cherry"),
		},
		{
			Datum:  []byte("dATE"),
			Expected: Some("dATE"),
		},



		{
			Datum:    Some("Hello world!"),
			Expected: Some("Hello world!"),
		},



		{
			Datum:         "Hello world!",
			Expected: Some("Hello world!"),
		},



		{
			Datum:  []byte("Hello world!"),
			Expected: Some("Hello world!"),
		},



		{
			Datum:         "-5",
			Expected: Some("-5"),
		},
		{
			Datum:         "-4",
			Expected: Some("-4"),
		},
		{
			Datum:         "-3",
			Expected: Some("-3"),
		},
		{
			Datum:         "-3",
			Expected: Some("-3"),
		},
		{
			Datum:         "-2",
			Expected: Some("-2"),
		},
		{
			Datum:         "-1",
			Expected: Some("-1"),
		},
		{
			Datum:         "0",
			Expected: Some("0"),
		},
		{
			Datum:         "1",
			Expected: Some("1"),
		},
		{
			Datum:         "2",
			Expected: Some("2"),
		},
		{
			Datum:         "3",
			Expected: Some("3"),
		},
		{
			Datum:         "4",
			Expected: Some("4"),
		},
		{
			Datum:         "5",
			Expected: Some("5"),
		},



		{
			Datum:  []byte("-5"),
			Expected: Some("-5"),
		},
		{
			Datum:  []byte("-4"),
			Expected: Some("-4"),
		},
		{
			Datum:  []byte("-3"),
			Expected: Some("-3"),
		},
		{
			Datum:  []byte("-2"),
			Expected: Some("-2"),
		},
		{
			Datum:  []byte("-1"),
			Expected: Some("-1"),
		},
		{
			Datum:  []byte("0"),
			Expected: Some("0"),
		},
		{
			Datum:  []byte("1"),
			Expected: Some("1"),
		},
		{
			Datum:  []byte("2"),
			Expected: Some("2"),
		},
		{
			Datum:  []byte("3"),
			Expected: Some("3"),
		},
		{
			Datum:  []byte("4"),
			Expected: Some("4"),
		},
		{
			Datum:  []byte("5"),
			Expected: Some("5"),
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			s := fmt.Sprintf("%d", x)

			test := struct {
				Datum    interface{}
				Expected Type
			}{
				Datum:         s,
				Expected: Some(s),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			s := fmt.Sprintf("%d", y)

			test := struct {
				Datum    interface{}
				Expected Type
			}{
				Datum:         s,
				Expected: Some(s),
			}
			tests = append(tests, test)
		}

		{
			s := fmt.Sprintf("%d", x)
			b := []byte(s)

			test := struct {
				Datum    interface{}
				Expected Type
			}{
				Datum:         b,
				Expected: Some(s),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			s := fmt.Sprintf("%d", y)
			b := []byte(s)

			test := struct {
				Datum    interface{}
				Expected Type
			}{
				Datum:         b,
				Expected: Some(s),
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {
		var datum Type

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

func TestTypeScanFail(t *testing.T) {
	tests := []struct {
		Datum interface{}
	}{
		{
			Datum: nil,
		},



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
		var datum Type

		if err := datum.Scan(test.Datum); nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one: (%T) %v", testNumber, err, err)
			continue
		}
	}
}

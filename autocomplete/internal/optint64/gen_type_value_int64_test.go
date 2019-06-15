package optint64

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */


import (
	"math"
	"testing"
)

func TestTypeValue(t *testing.T) {
	tests := []struct {
		Datum    Type
		Expected int64
	}{
		{
			Datum:     Some(1),
			Expected: int64(1),
		},
		{
			Datum:     Some(12),
			Expected: int64(12),
		},
		{
			Datum:     Some(-5),
			Expected: int64(-5),
		},
		{
			Datum:     Some(math.MaxInt64),
			Expected: int64(math.MaxInt64),
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

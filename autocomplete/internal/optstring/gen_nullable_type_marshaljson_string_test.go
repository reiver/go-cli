package optstring

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */


import (
	"encoding/json"
	"fmt"
	"time"

	"testing"
)

func TestNullableTypeMarshalJSONStruct(t *testing.T) {
	type MyThing struct {
		ID          NullableType `json:"id"`
		EMail       string       `json:"email"`
		WhenCreated time.Time    `json:"when_created"`
		ParentID    NullableType `json:"parent_id"`
	}

	var datum MyThing

	datum.ID = SomeNullable("some-ID")
	datum.EMail = "joeblow@example.net"
	{
		tm, err := time.Parse(time.RFC3339, "2017-04-12T13:02:28Z")
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}

		datum.WhenCreated = tm
	}
	datum.ParentID = SomeNullable("another-ID")

	actualBytes, err := json.Marshal(datum)
	if nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	{
		expected := `{` +
			`"id":"some-ID",` +
			`"email":"joeblow@example.net",` +
			`"when_created":"2017-04-12T13:02:28Z",` +
			`"parent_id":"another-ID"` +
			`}`

		if actual := string(actualBytes); expected != actual {
			t.Errorf("EXPECTED: %q", expected)
			t.Errorf("ACTUAL:   %q", actual)
			return
		}
	}
}

func TestNullableTypeMarshalJSON(t *testing.T) {
	tests := []struct {
		Datum    NullableType
		Expected string
	}{
		{
			Datum: SomeNullable("apple"),
			Expected:    `"apple"`,
		},
		{
			Datum: SomeNullable("BANANA"),
			Expected:    `"BANANA"`,
		},
		{
			Datum: SomeNullable("Cherry"),
			Expected:    `"Cherry"`,
		},
		{
			Datum: SomeNullable("dATE"),
			Expected:    `"dATE"`,
		},



		{
			Datum: SomeNullable("Hello world!"),
			Expected:    `"Hello world!"`,
		},



		{
			Datum: SomeNullable("5"),
			Expected:    `"5"`,
		},
		{
			Datum: SomeNullable("-4"),
			Expected:    `"-4"`,
		},
		{
			Datum: SomeNullable("-3"),
			Expected:    `"-3"`,
		},
		{
			Datum: SomeNullable("-2"),
			Expected:    `"-2"`,
		},
		{
			Datum: SomeNullable("-1"),
			Expected:    `"-1"`,
		},
		{
			Datum: SomeNullable("0"),
			Expected:    `"0"`,
		},
		{
			Datum: SomeNullable("1"),
			Expected:    `"1"`,
		},
		{
			Datum: SomeNullable("2"),
			Expected:    `"2"`,
		},
		{
			Datum: SomeNullable("3"),
			Expected:    `"3"`,
		},
		{
			Datum: SomeNullable("4"),
			Expected:    `"4"`,
		},
		{
			Datum: SomeNullable("5"),
			Expected:    `"5"`,
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			test := struct {
				Datum    NullableType
				Expected string
			}{
				Datum:     SomeNullable(fmt.Sprintf("%d", x)),
				Expected: fmt.Sprintf(`"%d"`, x),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			test := struct {
				Datum    NullableType
				Expected string
			}{
				Datum:     SomeNullable(fmt.Sprintf("%d", y)),
				Expected: fmt.Sprintf(`"%d"`, y),
			}
			tests = append(tests, test)
		}
	}

	for testNumber, test := range tests {

		actualBytes, err := json.Marshal(test.Datum)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, string(actualBytes); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q", testNumber, expected, actual)
			continue
		}
	}
}

func TestNullableTypeMarshalJSONNoneNullable(t *testing.T) {
	datum := NoneNullable()

	_, err := datum.MarshalJSON()
	if nil == err {
		t.Errorf("Expected an error, but not actually get one: (%T) %v", err, err)
		return
	}
}

func TestNullableTypeMarshalJSONull(t *testing.T) {
	datum := Null()

	actualBytes, err := datum.MarshalJSON()
	if nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	if expected, actual := "null", string(actualBytes); expected != actual {
		t.Errorf("Expected %q, but actually got %q", expected, actual)
		return
	}
}

package cliflags_test

import (
	"github.com/reiver/go-cli/flags"

	"testing"
)

func TestKeyValuesStorer(t *testing.T) {

	var x cliflags.Storer = new(cliflags.KeyValues) // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == x {
		t.Errorf("This should not happen.")
		return
	}
}

func TestKeyValuesDoubleStoreError(t *testing.T) {

	var keyvalues cliflags.KeyValues

	var expectedKey   string = "when"
	var expectedValue string = "now"

	{
		var key   string = expectedKey
		var value string = expectedValue

		if err := keyvalues.Store(key, value); nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %q", err, err)
			return
		}
	}

	{
		var key   string = "n"
		var value string = "5"

		if err := keyvalues.Store(key, value); nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %q", err, err)
			return
		}
	}

	{
		var key   string = "when"
		var value string = "later"

		if err := keyvalues.Store(key, value); nil == err {
			t.Errorf("Expected an error, but did not actually geg one: #%v", err)
			return
		}
	}

	{
		if expected, actual := 2, keyvalues.Len(); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		{
			actualValue := keyvalues.Load(expectedKey)

			if expected, actual := cliflags.SomeValue(expectedValue), actualValue; expected != actual {
				t.Errorf("Expected value %#v, but actually got value %#v.", expected, actual)
				return
			}
		}
	}
}

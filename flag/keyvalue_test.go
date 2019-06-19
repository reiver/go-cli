package cliflag_test

import (
	"github.com/reiver/go-cli/flag"

	"testing"
)

func TestKeyValueStorer(t *testing.T) {

	var x cliflag.Storer = new(cliflag.KeyValue) // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == x {
		t.Errorf("This should not happen.")
		return
	}
}

func TestKeyValueNotLoaded(t *testing.T) {

	var keyvalue cliflag.KeyValue

	key, value, err := keyvalue.Unwrap()
	if nil == err {
		t.Errorf("Expect an error, but did not actually get one: %#v", err)
		return
	}

	if expected, actual := "", key; expected != actual {
		t.Errorf("Expected key to be %q, but key was actually %q", expected, actual)
		return
	}

	if expected, actual := "", value; expected != actual {
		t.Errorf("Expected value to be %q, but value was actually %q", expected, actual)
		return
	}
}

func TestKeyValueUnwrap(t *testing.T) {

	var keyvalue cliflag.KeyValue

	var expectedKey   string = "when"
	var expectedValue string = "now"

	{
		var key   string = expectedKey
		var value string = expectedValue

		if err := keyvalue.Store(key, value); nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %q", err, err)
			return
		}
	}

	{
		key, value, err := keyvalue.Unwrap()
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %q", err, err)
			return
		}

		if expected, actual := expectedKey, key; expected != actual {
			t.Errorf("Expected key %q, but actually got key %q.", expected, actual)
			return
		}

		if expected, actual := expectedValue, value; expected != actual {
			t.Errorf("Expected value %q, but actually got value %q.", expected, actual)
			return
		}
	}

}

func TestKeyValueDoubleStoreError(t *testing.T) {

	var keyvalue cliflag.KeyValue

	var expectedKey   string = "when"
	var expectedValue string = "now"

	{
		var key   string = expectedKey
		var value string = expectedValue

		if err := keyvalue.Store(key, value); nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %q", err, err)
			return
		}
	}

	{
		var key   string = "n"
		var value string = "5"

		if err := keyvalue.Store(key, value); nil == err {
			t.Errorf("Expected an error, but did not actually geg one: #%v", err)
			return
		}
	}

	{
		var key   string = "when"
		var value string = "later"

		if err := keyvalue.Store(key, value); nil == err {
			t.Errorf("Expected an error, but did not actually geg one: #%v", err)
			return
		}
	}

	{
		key, value, err := keyvalue.Unwrap()
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %q", err, err)
			return
		}

		if expected, actual := expectedKey, key; expected != actual {
			t.Errorf("Expected key %q, but actually got key %q.", expected, actual)
			return
		}

		if expected, actual := expectedValue, value; expected != actual {
			t.Errorf("Expected value %q, but actually got value %q.", expected, actual)
			return
		}
	}
}

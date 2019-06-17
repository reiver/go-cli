package cli

import (
	"reflect"

	"testing"
)

func TestUnserialize(t *testing.T) {

	tests := []struct{
		Serialized string
		Expected []string
	}{
		{
			Serialized: "",
			Expected: []string{},
		},



		{
			Serialized: "apple",
			Expected: []string{"apple"},
		},
		{
			Serialized: "apple\u001Fbanana",
			Expected: []string{"apple","banana"},
		},
		{
			Serialized: "apple\u001Fbanana\u001Fcherry",
			Expected: []string{"apple","banana","cherry"},
		},



		{
			Serialized: "🙂\u001F😲😟\u001F👽👾👻",
			Expected: []string{"🙂","😲😟","👽👾👻"},
		},
	}

	for testNumber, test := range tests {

		actual := unserialize(test.Serialized)

		if expected := test.Expected; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d...", testNumber)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}

	}
}

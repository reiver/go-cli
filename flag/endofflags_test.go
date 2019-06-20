package cliflag

import (
	"testing"
)

func TestInternalEndOfFlags(t *testing.T) {

	var datum EndOfFlags = internalEndOfFlags{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == datum {
		t.Errorf("This should never happen.")
	}
}

func TestEndOfFlags(t *testing.T) {
	tests := []struct{
		Token string
		Expected string
	}{
		{
			Token:           "",
			Expected: "cli: end of flags: \"\" is not a flag",
		},



		{
			Token:           "apple",
			Expected: "cli: end of flags: \"apple\" is not a flag",
		},
		{
			Token:           "banana",
			Expected: "cli: end of flags: \"banana\" is not a flag",
		},
		{
			Token:           "cherry",
			Expected: "cli: end of flags: \"cherry\" is not a flag",
		},



		{
			Token:           "👽👾👻",
			Expected: "cli: end of flags: \"👽👾👻\" is not a flag",
		},
	}

	for testNumber, test := range tests {

		var endOfFlags EndOfFlags = internalEndOfFlags{token: test.Token}

		{
			var err error = endOfFlags

			if expected, actual := test.Expected, err.Error(); expected != actual {
				t.Errorf("For test #%d ...", testNumber)
				t.Errorf("\tEXPECTED: %#v", expected)
				t.Errorf("\tACTUAL:   %#v", actual)
				continue
			}
		}
	}
}

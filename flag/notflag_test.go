package cliflag

import (
	"testing"
)

func TestInternalNotFlag(t *testing.T) {

	var datum NotFlag = internalNotFlag{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == datum {
		t.Errorf("This should never happen.")
	}
}

func TestNotFlag(t *testing.T) {
	tests := []struct{
		Token string
		Expected string
	}{
		{
			Token:           "",
			Expected: "cli: \"\" is not a flag",
		},



		{
			Token:           "apple",
			Expected: "cli: \"apple\" is not a flag",
		},
		{
			Token:           "banana",
			Expected: "cli: \"banana\" is not a flag",
		},
		{
			Token:           "cherry",
			Expected: "cli: \"cherry\" is not a flag",
		},



		{
			Token:           "👽👾👻",
			Expected: "cli: \"👽👾👻\" is not a flag",
		},
	}

	for testNumber, test := range tests {

		var notFlag NotFlag = internalNotFlag{token: test.Token}

		{
			var err error = notFlag

			if expected, actual := test.Expected, err.Error(); expected != actual {
				t.Errorf("For test #%d ...", testNumber)
				t.Errorf("\tEXPECTED: %#v", expected)
				t.Errorf("\tACTUAL:   %#v", actual)
				continue
			}
		}
	}
}

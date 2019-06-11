package cli

import (
	"bytes"

	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct{
		String   string
		Expected string
	}{
		{ // 0
			String:   "",
			Expected: "",
		},




		{ // 1
			String:   "apple",
			Expected: "apple",
		},
		{ // 2
			String:   "apple banana",
			Expected: "apple banana",
		},
		{ // 3
			String:   "apple banana cherry",
			Expected: "apple banana cherry",
		},



		{ // 4
			String:   "hello world",
			Expected: "hello world",
		},



		{ // 5
			String:   "خُدا حافِظ",
			Expected: "خُدا حافِظ",
		},



		{ // 6
			String:   "🐸🐒",
			Expected: "🐸🐒",
		},
		{ // 7
			String:   "🍉🍑🍌",
			Expected: "🍉🍑🍌",
		},
		{ // 8
			String:   "👽👻👾😉",
			Expected: "👽👻👾😉",
		},



		{ // 9
			String:   "\x1Bapple",
			Expected: "\x1B\x1Bapple",
		},
		{ // 10
			String:   "a\x1Bpple",
			Expected: "a\x1B\x1Bpple",
		},
		{ // 11
			String:   "ap\x1Bple",
			Expected: "ap\x1B\x1Bple",
		},
		{ // 12
			String:   "app\x1Ble",
			Expected: "app\x1B\x1Ble",
		},
		{ // 13
			String:   "appl\x1Be",
			Expected: "appl\x1B\x1Be",
		},
		{ // 14
			String:   "apple\x1B",
			Expected: "apple\x1B\x1B",
		},



		{ // 15
			String:   "\x1B\x1Bapple",
			Expected: "\x1B\x1B\x1B\x1Bapple",
		},
		{ // 16
			String:   "a\x1B\x1Bpple",
			Expected: "a\x1B\x1B\x1B\x1Bpple",
		},
		{ // 17
			String:   "ap\x1B\x1Bple",
			Expected: "ap\x1B\x1B\x1B\x1Bple",
		},
		{ // 18
			String:   "app\x1B\x1Ble",
			Expected: "app\x1B\x1B\x1B\x1Ble",
		},
		{ // 19
			String:   "appl\x1B\x1Be",
			Expected: "appl\x1B\x1B\x1B\x1Be",
		},
		{ // 20
			String:   "apple\x1B\x1B",
			Expected: "apple\x1B\x1B\x1B\x1B",
		},



		{ // 21
			String:   "\x1B\x1B\x1Bapple",
			Expected: "\x1B\x1B\x1B\x1B\x1B\x1Bapple",
		},
		{ // 22
			String:   "a\x1B\x1B\x1Bpple",
			Expected: "a\x1B\x1B\x1B\x1B\x1B\x1Bpple",
		},
		{ // 23
			String:   "ap\x1B\x1B\x1Bple",
			Expected: "ap\x1B\x1B\x1B\x1B\x1B\x1Bple",
		},
		{ // 24
			String:   "app\x1B\x1B\x1Ble",
			Expected: "app\x1B\x1B\x1B\x1B\x1B\x1Ble",
		},
		{ // 25
			String:   "appl\x1B\x1B\x1Be",
			Expected: "appl\x1B\x1B\x1B\x1B\x1B\x1Be",
		},
		{ // 26
			String:   "apple\x1B\x1B\x1B",
			Expected: "apple\x1B\x1B\x1B\x1B\x1B\x1B",
		},



		{ // 27
			String:   "\x1B\x1B\x1B\x1Bapple",
			Expected: "\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1Bapple",
		},
		{ // 28
			String:   "a\x1B\x1B\x1B\x1Bpple",
			Expected: "a\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1Bpple",
		},
		{ // 29
			String:   "ap\x1B\x1B\x1B\x1Bple",
			Expected: "ap\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1Bple",
		},
		{ // 30
			String:   "app\x1B\x1B\x1B\x1Ble",
			Expected: "app\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1Ble",
		},
		{ // 31
			String:   "appl\x1B\x1B\x1B\x1Be",
			Expected: "appl\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1Be",
		},
		{ // 32
			String:   "apple\x1B\x1B\x1B\x1B",
			Expected: "apple\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B",
		},



		{ // 33
			String:   "\x1B\x1Bbanana",
			Expected: "\x1B\x1B\x1B\x1Bbanana",
		},
		{ // 34
			String:   "\x1Bb\x1Banana",
			Expected: "\x1B\x1Bb\x1B\x1Banana",
		},
		{ // 35
			String:   "\x1Bba\x1Bnana",
			Expected: "\x1B\x1Bba\x1B\x1Bnana",
		},
		{ // 36
			String:   "\x1Bban\x1Bana",
			Expected: "\x1B\x1Bban\x1B\x1Bana",
		},
		{ // 37
			String:   "\x1Bbana\x1Bna",
			Expected: "\x1B\x1Bbana\x1B\x1Bna",
		},
		{ // 38
			String:   "\x1Bbanan\x1Ba",
			Expected: "\x1B\x1Bbanan\x1B\x1Ba",
		},
		{ // 39
			String:   "\x1Bbanana\x1B",
			Expected: "\x1B\x1Bbanana\x1B\x1B",
		},



		{ // 40
			String:   "b\x1B\x1Banana",
			Expected: "b\x1B\x1B\x1B\x1Banana",
		},
		{ // 41
			String:   "b\x1Ba\x1Bnana",
			Expected: "b\x1B\x1Ba\x1B\x1Bnana",
		},
		{ // 42
			String:   "b\x1Ban\x1Bana",
			Expected: "b\x1B\x1Ban\x1B\x1Bana",
		},
		{ // 43
			String:   "b\x1Bana\x1Bna",
			Expected: "b\x1B\x1Bana\x1B\x1Bna",
		},
		{ // 44
			String:   "b\x1Banan\x1Ba",
			Expected: "b\x1B\x1Banan\x1B\x1Ba",
		},
		{ // 45
			String:   "b\x1Banana\x1B",
			Expected: "b\x1B\x1Banana\x1B\x1B",
		},



		{ // 46
			String:   "ba\x1B\x1Bnana",
			Expected: "ba\x1B\x1B\x1B\x1Bnana",
		},
		{ // 47
			String:   "ba\x1Bn\x1Bana",
			Expected: "ba\x1B\x1Bn\x1B\x1Bana",
		},
		{ // 48
			String:   "ba\x1Bna\x1Bna",
			Expected: "ba\x1B\x1Bna\x1B\x1Bna",
		},
		{ // 49
			String:   "ba\x1Bnan\x1Ba",
			Expected: "ba\x1B\x1Bnan\x1B\x1Ba",
		},
		{ // 50
			String:   "ba\x1Bnana\x1B",
			Expected: "ba\x1B\x1Bnana\x1B\x1B",
		},



		{ // 51
			String:   "ban\x1B\x1Bana",
			Expected: "ban\x1B\x1B\x1B\x1Bana",
		},
		{ // 52
			String:   "ban\x1Ba\x1Bna",
			Expected: "ban\x1B\x1Ba\x1B\x1Bna",
		},
		{ // 53
			String:   "ban\x1Ban\x1Ba",
			Expected: "ban\x1B\x1Ban\x1B\x1Ba",
		},
		{ // 54
			String:   "ban\x1Bana\x1B",
			Expected: "ban\x1B\x1Bana\x1B\x1B",
		},



		{ // 55
			String:   "bana\x1B\x1Bna",
			Expected: "bana\x1B\x1B\x1B\x1Bna",
		},
		{ // 56
			String:   "bana\x1Bn\x1Ba",
			Expected: "bana\x1B\x1Bn\x1B\x1Ba",
		},
		{ // 57
			String:   "bana\x1Bna\x1B",
			Expected: "bana\x1B\x1Bna\x1B\x1B",
		},



		{ // 58
			String:   "banan\x1B\x1Ba",
			Expected: "banan\x1B\x1B\x1B\x1Ba",
		},
		{ // 59
			String:   "banan\x1Ba\x1B",
			Expected: "banan\x1B\x1Ba\x1B\x1B",
		},



		{ // 60
			String:   "banana\x1B\x1B",
			Expected: "banana\x1B\x1B\x1B\x1B",
		},



		{ // 61
			String:   "c\x1Bhe\x1B\x1Brr\x1Fy",
			Expected: "c\x1B\x1Bhe\x1B\x1B\x1B\x1Brr\x1B\x1Fy",
		},



		{ // 62
			String:   "🍉\x1B🍑🍌,👽\x1F👻👾😉",
			Expected: "🍉\x1B\x1B🍑🍌,👽\x1B\x1F👻👾😉",
		},



		{ // 63
			String:   "🍐\x1F🥝\x1F🍎 🍔🍟\x1F\x1F🍕🌭\t👿👿\x1F abc\r\n",
			Expected: "🍐\x1B\x1F🥝\x1B\x1F🍎 🍔🍟\x1B\x1F\x1B\x1F🍕🌭\t👿👿\x1B\x1F abc\r\n",
		},
	}

	for testNumber, test := range tests {
		var buffer bytes.Buffer

		if err := encode(&buffer, test.String); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one (%T) %q", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For test #%d...", testNumber)
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}
	}
}

package clitest_test

import (
	"github.com/reiver/go-cli/test"

	"testing"
)

func TestWriteCloserCloseNilReceiver(t *testing.T) {

	var wc *clitest.WriteCloser = nil

	err := wc.Close()

	if expected, actual := "Nil Receiver", err.Error(); expected != actual {
		t.Errorf("Expected an error %q, but actually got %q.", expected, actual)
		return
	}
}

func TestWriteCloserWriteNilReceiver(t *testing.T) {

	var wc *clitest.WriteCloser = nil

	var b [12]byte = [12]byte{'H','e','l','l','o',' ','w','o','r','l','d','!'}
	n, err := wc.Write(b[:])

	if expected, actual := "Nil Receiver", err.Error(); expected != actual {
		t.Errorf("Expected an error %q, but actually got %q.", expected, actual)
		return
	}

	if expected, actual := 0, n; expected != actual {
		t.Errorf("Expected %d, but actually got %d.", expected, actual)
		return
	}
}

func TestWriteCloserClose(t *testing.T) {

	var target clitest.WriteCloser

	{
		var b [13]byte = [...]byte{'G','o','o','d', ' ', 'm','o','r','n','i','n','g', '!'}

		n, err := target.Write(b[:])
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %q", err, err)
			return
		}
		if expected, actual := len(b), n; expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}
	}

	if err := target.Close(); nil != err{
		t.Errorf("Did not expect an error, but actually got one: (%T) %q", err, err)
		return
	}

	{
		var b [9]byte = [...]byte{'G','o','o','d', ' ', 'b','y','e', '!'}

		n, err := target.Write(b[:])
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one: %#v", err)
			return
		}
		if expected, actual := "Closed", err.Error(); expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
		if expected, actual := 0, n; expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}
	}
}

func TestWriteCloser(t *testing.T) {

	tests := []struct{
		Writes [][]byte
		Expected string
	}{
		{
			Writes: [][]byte{},
			Expected: "",
		},



		{
			Writes: [][]byte{
				[]byte("apple"),
			},
			Expected: "apple",
		},
		{
			Writes: [][]byte{
				[]byte("apple"),
				[]byte("banana"),
			},
			Expected: "applebanana",
		},
		{
			Writes: [][]byte{
				[]byte("apple"),
				[]byte("banana"),
				[]byte("cherry"),
			},
			Expected: "applebananacherry",
		},



		{
			Writes: [][]byte{
				[]byte("Hello"),
				[]byte(" "),
				[]byte("world"),
				[]byte("!"),
			},
			Expected: "Hello world!",
		},
	}

	for testNumber, test := range tests {

		var target clitest.WriteCloser

		for writesNumber, b := range test.Writes {

			{
				n, err := target.Write(b)
				if nil != err {
					t.Errorf("For test #%d and writes #%d, did not expect an error, but actually got one: (%T) %q", testNumber, writesNumber, err, err)
					continue
				}
				if expected, actual := len(b), n; expected != actual {
					t.Errorf("For test #%d and writes #%d, expected %d, but actually got %d.", testNumber, writesNumber, expected, actual)
					continue
				}
			}
		}

		if expected, actual := test.Expected, target.String(); expected != actual {
			t.Errorf("For test #%d...", testNumber)
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}
	}
}
